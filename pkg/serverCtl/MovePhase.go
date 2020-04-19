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

func (this *ULZGameDuelServiceBackend) MovePhaseConfirm(ctx context.Context, req *pb.GDMoveConfirmReq) (*pb.Empty, error) {
	cm.PrintReqLog(ctx, "Move-Phase-Confirm", req)
	start := time.Now()
	this.mu.Lock()
	wkbox := this.searchAliveClient()
	defer func() {
		wkbox.Preserve(false)
		this.mu.Unlock()
		elapsed := time.Since(start)
		log.Printf("Move-Phase-Confirm took %s", elapsed)
	}()
	// get data -----
	wg := sync.WaitGroup{}
	wg.Add(2)
	// MovePhaseSnapMod
	var snapMove pb.MovePhaseSnapMod
	var snapMovekey = req.RoomKey + snapMove.RdsKeyName()
	errch := make(chan error)
	go func() {
		if _, err := (wkbox).GetPara(&snapMovekey, &snapMove); err != nil {
			log.Println(err)
			errch <- status.Errorf(codes.NotFound, err.Error())
		}
		wg.Done()
	}()
	// PhaseSnapMod
	var snapPhase pb.PhaseSnapMod
	snapPhasekey := req.RoomKey + snapPhase.RdsKeyName()
	go func() {
		if _, err := (wkbox).GetPara(&snapPhasekey, &snapMove); err != nil {
			log.Println(err)
			errch <- status.Errorf(codes.NotFound, err.Error())
		}
		wg.Done()
	}()

	// var eventNode pb.EventSnapMod

	wg.Wait()

	if err := <-errch; err != nil {
		return nil, err
	}
	// ======= -----

	if snapPhase.EventPhase != pb.EventHookPhase_move_card_drop_phase {
		return nil, status.Error(codes.FailedPrecondition, "NOT_IN_MOVE_PHASE")
	}

	if (req.Side == pb.PlayerSide_HOST && snapPhase.IsHostReady) ||
		(req.Side == pb.PlayerSide_DUELER && snapPhase.IsDuelReady) {
		return nil, status.Error(codes.AlreadyExists, "ALREADY_READY")
	}

	if req.Side == pb.PlayerSide_HOST {
		snapMove.HostVal = req.Point
		snapMove.HostOpt = req.MoveOpt
		snapMove.HostTrigSkl = req.TriggerSkl
		snapMove.HostCard = req.UpdateCard
		// snapPhase
		snapPhase.IsHostReady = true

	} else if req.Side == pb.PlayerSide_DUELER {
		snapMove.DuelVal = req.Point
		snapMove.DuelOpt = req.MoveOpt
		snapMove.DuelTrigSkl = req.TriggerSkl
		snapMove.DuelCard = req.UpdateCard
		// snapPhase
		snapPhase.IsDuelReady = true
	}
	// do snap-mod update
	// wg.Add(1)

	wg.Add(2)
	go func() {
		if _, err := (wkbox).SetPara(&snapMovekey, snapMove); err != nil {
			log.Println(err)
			errch <- status.Errorf(codes.Internal, err.Error())
		}
		wg.Done()
	}()
	go func() {
		if _, err := (wkbox).SetPara(&snapPhasekey, snapPhase); err != nil {
			log.Println(err)
			errch <- status.Errorf(codes.Internal, err.Error())
		}
		wg.Done()
	}()
	wg.Wait()
	if err := <-errch; err != nil {
		return nil, err
	}

	// broadcast ok ?
	// side := req.Side.String()
	go this.BroadCast(&pb.GDBroadcastResp{
		RoomKey:      req.RoomKey,
		Msg:          req.Side.String() + "_READY",
		Command:      pb.CastCmd_GET_MOVE_PHASE_RESULT,
		CurrentPhase: pb.EventHookPhase_move_card_drop_phase,
		PhaseHook:    pb.EventHookType_Proxy,
		Side:         req.Side,
		InstanceSet:  nil,
	})

	if snapPhase.IsHostReady && snapPhase.IsDuelReady {
		go this.BroadCast(&pb.GDBroadcastResp{
			RoomKey:      req.RoomKey,
			Msg:          "Both Ready",
			Command:      pb.CastCmd_GET_MOVE_PHASE_RESULT,
			CurrentPhase: pb.EventHookPhase_move_card_drop_phase,
			PhaseHook:    pb.EventHookType_After,
			InstanceSet:  nil,
		})
		// go store before move next

		// both-ready; then move-next phase
		go func() {
			mbox := this.searchAliveClient()
			var gameDt pb.GameDataSet
			mbox.GetPara(&req.RoomKey, &gameDt)
			gameDt.HookType = pb.EventHookType_After
			mbox.Preserve(false)
			this.moveNextPhase(&gameDt, &snapPhase)
		}()
	}
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
// remark it should from determine-move-
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
		snapMovekey := req.RoomKey + snapMove.RdsKeyName()
		if _, err := (wkbox).GetPara(&snapMovekey, &snapMove); err != nil {
			log.Println(err)
			errch <- status.Errorf(codes.NotFound, err.Error())
		}
		wg.Done()
	}()

	go func() {
		snapMovekey := req.RoomKey + movePhase.RdsKeyName()
		wkbox1 := this.searchAliveClient()
		if _, err := (wkbox1).GetPara(&snapMovekey, &movePhase); err != nil {
			log.Println(err)
			errch <- status.Errorf(codes.NotFound, err.Error())
		}
		wg.Done()
	}()

	wg.Wait()

	go func() {
		snapMovekey := req.RoomKey + PhaseMod.RdsKeyName()
		wkbox1 := this.searchAliveClient()
		if _, err := (wkbox1).GetPara(&snapMovekey, &PhaseMod); err != nil {
			log.Println(err)
			errch <- status.Errorf(codes.NotFound, err.Error())
		}
		if !req.IsWatcher && req.Side == pb.PlayerSide_HOST {
			PhaseMod.IsHostReady = true
		} else if !req.IsWatcher && req.Side == pb.PlayerSide_DUELER {
			PhaseMod.IsDuelReady = true
		}
		go (wkbox1).SetPara(&snapMovekey, PhaseMod)
		if PhaseMod.IsHostReady && PhaseMod.IsDuelReady {
			go this.BroadCast(&pb.GDBroadcastResp{
				RoomKey:      req.RoomKey,
				Msg:          fmt.Sprintf("MV_PHASE:ACK_Both_SideResolve:"),
				Command:      pb.CastCmd_GET_MOVE_PHASE_RESULT,
				CurrentPhase: pb.EventHookPhase_attack_card_drop_phase,
				PhaseHook:    pb.EventHookType_Proxy,
			})

			var gamDT pb.GameDataSet
			wkbox1.GetPara(&req.RoomKey, &gamDT)
			// suppose (gamDT.EventPhase == PhaseMod.EventPhase) === pb.determine
			// gamDT.EventPhase = PhaseMod.EventPhase
			gamDT.HookType = pb.EventHookType_After
			go this.moveNextPhase(&gamDT, &PhaseMod)
		}

	}()

	return &snapMove, nil
}
