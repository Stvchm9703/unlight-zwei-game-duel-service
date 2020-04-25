package serverCtl

import (
	cm "ULZGameDuelService/pkg/common"
	pb "ULZGameDuelService/proto"
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	// "github.com/docker/docker/builder/builder-next/adapters/snapshot"
	"github.com/gogo/status"
	"google.golang.org/grpc/codes"
	// Static files
	// _ "ULZGameDuelService/statik"
)

/**
 * Move-Phase-Confirm : (rpc)
 * 		request handle in move-card-drop-phase:proxy
 */
func (this *ULZGameDuelServiceBackend) MovePhaseConfirm(ctx context.Context, req *pb.GDMoveConfirmReq) (*pb.Empty, error) {
	cm.PrintReqLog(ctx, "Move-Phase-Confirm", req)
	start := time.Now()
	this.mu.Lock()
	defer func() {
		// wkbox.Preserve(false)
		this.mu.Unlock()
		elapsed := time.Since(start)
		log.Printf("Move-Phase-Confirm took %s", elapsed)
	}()
	// get data -----
	wg := sync.WaitGroup{}
	// ======================================================================
	wg.Add(4)
	errch := make(chan error)

	var gameSet pb.GameDataSet
	var gameSetKey = req.RoomKey
	go func() {
		wkbox := this.searchAliveClient()
		if _, err := (wkbox).GetPara(gameSetKey, &gameSet); err != nil {
			log.Println(err)
			errch <- err
		}
		wkbox.Preserve(false)
		wg.Done()
	}()
	// MovePhaseSnapMod
	var snapMove pb.MovePhaseSnapMod
	var snapMovekey = req.RoomKey + snapMove.RdsKeyName()
	go func() {
		wkbox := this.searchAliveClient()
		if _, err := (wkbox).GetPara(snapMovekey, &snapMove); err != nil {
			log.Println(err)
			errch <- err
		}
		wkbox.Preserve(false)
		wg.Done()
	}()
	// PhaseSnapMod
	var snapPhase pb.PhaseSnapMod
	snapPhasekey := req.RoomKey + snapPhase.RdsKeyName()
	go func() {
		wkbox := this.searchAliveClient()
		if _, err := (wkbox).GetPara(snapPhasekey, &snapMove); err != nil {
			log.Println(err)
			errch <- err
		}
		wkbox.Preserve(false)
		wg.Done()
	}()
	// effect-result
	var effectMod pb.EffectNodeSnapMod
	snapEfKey := req.RoomKey + effectMod.RdsKeyName()
	go func() {
		wkbox := this.searchAliveClient()
		if _, err := (wkbox).GetPara(snapEfKey, &effectMod); err != nil {
			log.Println(err)
			errch <- err
		}
		wkbox.Preserve(false)
		wg.Done()
	}()
	wg.Wait()
	if err := <-errch; err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	// ======================================================================

	// ======= -----

	if snapPhase.EventPhase != pb.EventHookPhase_move_card_drop_phase {
		return nil, status.Error(codes.FailedPrecondition, "NOT_IN_MOVE_PHASE")
	}

	if (req.Side == pb.PlayerSide_HOST && snapPhase.IsHostReady) ||
		(req.Side == pb.PlayerSide_DUELER && snapPhase.IsDuelReady) {
		return nil, status.Error(codes.AlreadyExists, "ALREADY_READY")
	}

	// ======================================================================
	// @Skill-Calculate
	var Eff []*pb.EffectResult
	var Val *int32
	var err error

	DisableSkill := pb.NodeFilter(effectMod.PendingEf, func(v *pb.EffectResult) bool {
	})

	Val, Eff, err = this.skillClient.SkillCalculateWrap(
		req.UpdateCard,
		req.TriggerSkl,
		pb.EventCardType_MOVE,
	)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	// ======================================================================
	var currentAtkKey int32
	if req.Side == pb.PlayerSide_HOST {
		snapMove.HostOpt = req.MoveOpt
		snapMove.HostTrigSkl = req.TriggerSkl
		snapMove.HostCard = req.UpdateCard
		snapMove.HostVal = *Val
		// snapPhase
		snapPhase.IsHostReady = true
		currentAtkKey = gameSet.HostCurrCardKey

	} else if req.Side == pb.PlayerSide_DUELER {
		snapMove.DuelOpt = req.MoveOpt
		snapMove.DuelTrigSkl = req.TriggerSkl
		snapMove.DuelCard = req.UpdateCard
		// snapPhase
		snapMove.DuelVal = *Val
		snapPhase.IsDuelReady = true
		currentAtkKey = gameSet.DuelCurrCardKey
	}

	effectMod.PendingEf = append(effectMod.PendingEf, Eff...)

	// add instant mp
	var addEff []*pb.EffectResult
	var addVal int32
	wg.Add(2)
	go func() {
		addEff = pb.NodeFilter(effectMod.PendingEf, func(v *pb.EffectResult) bool {
			return (v.TriggerTime.EventPhase == pb.EventHookPhase_move_card_drop_phase &&
				v.TriggerTime.HookType == pb.EventHookType_Proxy &&
				v.EfOption == pb.EffectOption_Status_Addition &&
				v.TarSide == req.Side &&
				v.TarCard == currentAtkKey)
		})
		for _, v := range addEff {
			addVal += v.Mp
		}
		wg.Done()
	}()
	var fixEff []*pb.EffectResult
	var fixVal int32
	go func() {
		fixEff = pb.NodeFilter(effectMod.PendingEf, func(v *pb.EffectResult) bool {
			return (v.TriggerTime.EventPhase == pb.EventHookPhase_determine_move_phase &&
				v.TriggerTime.HookType == pb.EventHookType_Proxy &&
				v.EfOption == pb.EffectOption_Status_FixValue &&
				v.TarSide == req.Side &&
				v.TarCard == currentAtkKey)
		})
		for _, v := range fixEff {
			if v.Ap > fixVal {
				fixVal = v.Mp
			}
		}
		wg.Done()
	}()
	wg.Wait()
	if len(fixEff) > 0 {
		if req.Side == pb.PlayerSide_HOST {
			snapMove.HostVal = fixVal
		} else if req.Side == pb.PlayerSide_DUELER {
			snapMove.DuelVal = fixVal
		}
	} else {
		if req.Side == pb.PlayerSide_HOST {
			snapMove.HostVal += addVal
		} else if req.Side == pb.PlayerSide_DUELER {
			snapMove.DuelVal += addVal
		}
	}

	// effectMod

	// do snap-mod update
	// wg.Add(1)

	// ======================================================================
	wg.Add(3)
	go func() {
		wkbox := this.searchAliveClient()
		if _, err := (wkbox).SetPara(snapMovekey, snapMove); err != nil {
			log.Println(err)
			errch <- err
		}
		wkbox.Preserve(false)
		wg.Done()
	}()
	go func() {
		wkbox := this.searchAliveClient()
		if _, err := (wkbox).SetPara(snapPhasekey, snapPhase); err != nil {
			log.Println(err)
			errch <- err
		}
		wkbox.Preserve(false)
		wg.Done()
	}()
	go func() {
		wkbox := this.searchAliveClient()
		if _, err := (wkbox).SetPara(snapEfKey, effectMod); err != nil {
			log.Println(err)
			errch <- err
		}
		wkbox.Preserve(false)
		wg.Done()
	}()
	wg.Wait()
	if err := <-errch; err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	// ======================================================================

	// broadcast ok ?
	// side := req.Side.String()
	go this.BroadCast(&pb.GDBroadcastResp{
		RoomKey:      req.RoomKey,
		Msg:          fmt.Sprintf("MOV_PHASE:Val_Ready:%v", req.Side.String(), *Val+addVal),
		Command:      pb.CastCmd_GET_MOVE_PHASE_RESULT,
		CurrentPhase: pb.EventHookPhase_move_card_drop_phase,
		PhaseHook:    pb.EventHookType_Proxy,
		Side:         req.Side,
		InstanceSet:  nil,
	})
	go func() {
		if snapPhase.IsHostReady && snapPhase.IsDuelReady {
			this.BroadCast(&pb.GDBroadcastResp{
				RoomKey:      req.RoomKey,
				Msg:          "Both Ready",
				Command:      pb.CastCmd_GET_MOVE_PHASE_RESULT,
				CurrentPhase: pb.EventHookPhase_move_card_drop_phase,
				PhaseHook:    pb.EventHookType_Proxy,
				InstanceSet:  nil,
			})
			// go store before move next
			// both-ready; then move-next phase
			mbox := this.searchAliveClient()
			var gameDt pb.GameDataSet
			mbox.GetPara(req.RoomKey, &gameDt)
			// gameDt.HookType = pb.EventHookType_Proxy
			mbox.Preserve(false)
			this.moveNextPhase(
				&gameDt,
				&snapPhase,
				&effectMod,
				&snapMove,
			)
		}
	}()
	return &pb.Empty{}, nil

}

/**
 * after the move-next to do [determine-move-phase]
 * it will request client to send move-phase-result to grep the result
 *
 * 	[move-phase]
 * 	|
 * 	{ both side ready }
 * 		| => ( go exec phase-runner )
 * 				| => [move_card_drop_phase:after] {
 * 				|		1. exec triggered skill
 * 				|			| -> get effect-func
 * 				|			| -> get instance value change
 * 				|		2. order effect-func in [ determine-move-phase:before,after ]
 * 				|		3. run the after-ef-result-node by order
 * 				|		4. return the list of node
 * 				|	}
 * 				|
 * 				|	phase-runner get af-ef-result-node list from exec-service:
 * 				|		-> send effent-phase-request
 * 				|
 * 				|	[determine-move-phase] is going end:{
 * 				|		1. store instance-result
 * 				|	 	2. send move-phase-result
 * 				|	}
 * 				|
 *  			| 	if the gameSet have dead result
 * 				|		if nvn == 1 :
 * 				|			game-end-phase
 * 				|		if nvn == 3 :
 * 				|			change-char-phase
 */

//  MovePhaseResult :
// remark it should from finish_move_phase:proxy
func (this *ULZGameDuelServiceBackend) MovePhaseResult(ctx context.Context, req *pb.GDGetInfoReq) (*pb.GDMoveConfirmResp, error) {
	cm.PrintReqLog(ctx, "Move-Phase-Result", req)
	start := time.Now()
	// this.mu.Lock()
	wkbox := this.searchAliveClient()
	defer func() {
		wkbox.Preserve(false)
		// this.mu.Unlock()
		elapsed := time.Since(start)
		log.Printf("Move-Phase-Result took %s", elapsed)
	}()

	var snapMove pb.GDMoveConfirmResp
	var movePhase pb.MovePhaseSnapMod
	var PhaseMod pb.PhaseSnapMod

	wg := sync.WaitGroup{}
	errch := make(chan error)
	wg.Add(2)
	go func() {
		if _, err := (wkbox).GetPara(req.RoomKey+snapMove.RdsKeyName(), &snapMove); err != nil {
			log.Println(err)
			errch <- status.Errorf(codes.NotFound, err.Error())
		}
		wg.Done()
	}()

	go func() {
		wkbox1 := this.searchAliveClient()
		if _, err := (wkbox1).GetPara(req.RoomKey+movePhase.RdsKeyName(), &movePhase); err != nil {
			log.Println(err)
			errch <- status.Errorf(codes.NotFound, err.Error())
		}
		wg.Done()
	}()
	wg.Wait()
	if err := <-errch; err != nil {
		return nil, status.Errorf(codes.NotFound, err.Error())
	}

	go func() {
		if !req.IsWatcher {

			wkbox1 := this.searchAliveClient()
			if _, err := (wkbox1).GetPara(req.RoomKey+PhaseMod.RdsKeyName(), &PhaseMod); err != nil {
				log.Println(err)
			}
			if req.Side == pb.PlayerSide_HOST {
				PhaseMod.IsHostReady = true
			} else if req.Side == pb.PlayerSide_DUELER {
				PhaseMod.IsDuelReady = true
			}
			(wkbox1).SetPara(req.RoomKey+PhaseMod.RdsKeyName(), PhaseMod)
			if PhaseMod.IsHostReady && PhaseMod.IsDuelReady {
				go this.BroadCast(&pb.GDBroadcastResp{
					RoomKey:      req.RoomKey,
					Msg:          fmt.Sprintf("MV_PHASE:ACK_Both_Side_Resolve:%s,AtkFirst", PhaseMod.FirstAttack.String()),
					Command:      pb.CastCmd_INSTANCE_STATUS_CHANGE,
					CurrentPhase: pb.EventHookPhase_finish_move_phase,
					PhaseHook:    pb.EventHookType_Proxy,
				})

				var gamDT pb.GameDataSet
				wkbox1.GetPara(req.RoomKey, &gamDT)
				var effectMod pb.EffectNodeSnapMod
				wkbox1.GetPara(req.RoomKey+effectMod.RdsKeyName(), effectMod)
				wkbox1.Preserve(false)

				// suppose (gamDT.EventPhase == PhaseMod.EventPhase) === pb.determine
				// gamDT.EventPhase = PhaseMod.EventPhase
				this.moveNextPhase(
					&gamDT,
					&PhaseMod,
					&effectMod)
			}
		}
	}()

	return &snapMove, nil
}
