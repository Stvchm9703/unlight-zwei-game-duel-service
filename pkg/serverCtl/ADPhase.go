package serverCtl

import (
	cm "ULZGameDuelService/pkg/common"
	pb "ULZGameDuelService/proto"
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/gogo/status"
	"google.golang.org/grpc/codes"
	// Static files
	// _ "ULZGameDuelService/statik"
)

/**
 *
 * 	[atk-drop-phase:proxy] {
 * 		do store mod
 * 		if trig-skill > 0
 * 			request skill-calcu for calc
 * 			| -> result val store as VAl
 * 			| -> skill effect push the effect-node
 * 			|	do store mod
 *
 * 		move next phase [atk-drop-phase:after] {
 * 			if ef-node > 0
 * 				do exec event-phase
 * 			go send ACK [EventResult]
 * 		}
 * 		go send ACK [ADPhaseResult]
 *
 * 		go-routine move next phase [def-drop-phase:before] {
 * 			if ef-node > 0
 * 				do exec event-phase
 * 			go send ACK [EventResult]
 * 			? go clock (5min) ?
 * 				| -> is-alive?
 * 		}
 * 	}
 * 	|
 * 	[def-drop-phase:proxy] {
 * 		do store mod
 * 		if trig-skill > 0
 * 			request skill-calcu for calc
 * 			| -> result val store as VAl
 * 			| -> skill effect push the effect-node
 * 			|	do store mod
 *
 * 		move next phase [def-drop-phase:after] {
 * 			if ef-node > 0
 * 				do exec event-phase
 * 			go send ACK [EventResult]
 * 		}
 * 		go send ACK [ADPhaseResult]
 *
 *		go-routine move nextphase [determine_battle_point_phase] {
 * 			if
 *      }
 *
 */

func (this *ULZGameDuelServiceBackend) ADPhaseConfirm(ctx context.Context, req *pb.GDADConfirmReq) (*pb.Empty, error) {
	cm.PrintReqLog(ctx, "AtkDef-Phase-Confirm", req)
	start := time.Now()
	this.mu.Lock()
	wkbox := this.searchAliveClient()
	defer func() {
		wkbox.Preserve(false)
		this.mu.Unlock()
		elapsed := time.Since(start)
		log.Printf("AtkDef-Phase-Confirm took %s", elapsed)
	}()

	// get data in routine
	wg := sync.WaitGroup{}
	wg.Add(3)
	errCh := make(chan error)
	var returner pb.GameDataSet
	go func() {
		if _, err := (wkbox).GetPara(&req.RoomKey, &returner); err != nil {
			log.Println(err)
			errCh <- status.Errorf(codes.NotFound, err.Error())
		}
		wg.Done()
	}()
	var snapMod pb.ADPhaseSnapMod
	snapModkey := req.RoomKey + ":ADPhMod"
	go func() {
		if _, err := (wkbox).GetPara(&snapModkey, &snapMod); err != nil {
			log.Println(err)
			errCh <- status.Errorf(codes.NotFound, err.Error())
		}
		wg.Done()
	}()
	var phaseInst pb.PhaseSnapMod
	go func() {
		tmpKey := req.RoomKey + ":PhaseState"
		if _, err := wkbox.GetPara(&tmpKey, &phaseInst); err != nil {
			log.Println(err)
			errCh <- status.Errorf(codes.Internal, err.Error())
		}
		wg.Done()
	}()

	// check grep data error
	if errRes := <-errCh; errRes != nil {
		return nil, errRes
	}

	//  attack phase
	if req.CurrentPhase != phaseInst.EventPhase {
		return nil, status.Error(codes.InvalidArgument, "AD_PHASE:InvaildPhase")
	}
	// FirstAttack = Host & CurrPhase = Host  -> Host is First-Attack
	isAttack := (phaseInst.EventPhase == pb.EventHookPhase_attack_card_drop_phase)
	isDefence := (phaseInst.EventPhase == pb.EventHookPhase_defence_card_drop_phase)
	// snapMod.FirstAttack == snapMod.CurrPhase -> attack
	// snapMod.FirstAttack != snapMod.CurrPhase -> defence
	// incomePhase := (snapMod.FirstAttack == snapMod.CurrAttacker)

	if isAttack {
		// snapMod.HostCard
		snapMod.AttackCard = req.UpdateCard
		snapMod.AttackTrigSkl = req.TriggerSkl
		if _, err := (wkbox).SetPara(&snapModkey, snapMod); err != nil {
			log.Println(err)
			return nil, status.Errorf(codes.Internal, err.Error())
		}
		go this.attackPhaseHandle(&req.RoomKey, &snapMod, &phaseInst)

		// this return notice the sender the process is ongoing
		// sender need to wait broadcast to move next phase
		return &pb.Empty{}, nil
	} else if isDefence {
		snapMod.DefenceCard = req.UpdateCard
		snapMod.DefenceTrigSkl = req.TriggerSkl
		if _, err := (wkbox).SetPara(&snapModkey, snapMod); err != nil {
			log.Println(err)
			return nil, status.Errorf(codes.Internal, err.Error())
		}
		go this.defencePhaseHandle(&req.RoomKey, &snapMod, &phaseInst)

		// this return notice the sender the process is ongoing
		// sender need to wait broadcast to move next phase
		return &pb.Empty{}, nil

	}
	return nil, status.Error(codes.Unimplemented, "AD_PHASE_CONFIRM")
}

func (this *ULZGameDuelServiceBackend) ADPhaseResult(ctx context.Context, req *pb.GDGetInfoReq) (*pb.GDADResultResp, error) {
	cm.PrintReqLog(ctx, "AtkDef-Phase-Result", req)
	start := time.Now()
	this.mu.Lock()
	wkbox := this.searchAliveClient()
	defer func() {
		wkbox.Preserve(false)
		this.mu.Unlock()
		elapsed := time.Since(start)
		log.Printf("AtkDef-Phase-Result took %s", elapsed)
	}()

	var snapMod pb.ADPhaseSnapMod
	snapModkey := req.RoomKey + snapMod.RdsKeyName()

	var stateMod pb.PhaseSnapMod
	stateModkey := req.RoomKey + stateMod.RdsKeyName()

	errCh := make(chan error)
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		if _, err := (wkbox).GetPara(&snapModkey, &snapMod); err != nil {
			log.Println(err)
			errCh <- status.Errorf(codes.Internal, err.Error())
		}
		wg.Done()
	}()
	go func() {
		if _, err := (wkbox).GetPara(&stateModkey, &stateMod); err != nil {
			log.Println(err)
			errCh <- status.Errorf(codes.Internal, err.Error())
		}
		wg.Done()
	}()
	wg.Wait()
	if err := <-errCh; err != nil {
		return nil, err
	}

	// return nil, status.Error(codes.Unimplemented, "AD_PHASE_RESULT")
	side := snapMod.CurrAttacker
	var pt int32 = 0
	if stateMod.EventPhase != req.CurrentPhase {
		return nil, status.Error(codes.InvalidArgument, "AD_PHASE:InvaildPhase")
	}

	if snapMod.EventPhase == pb.EventHookPhase_attack_card_drop_phase {
		side = snapMod.CurrAttacker
		// Attack-pt
		pt = snapMod.AttackVal
	} else {
		if snapMod.CurrAttacker == pb.PlayerSide_HOST {
			side = pb.PlayerSide_DUELER
		} else {
			side = pb.PlayerSide_HOST
		}
		// Deference-pt
		pt = snapMod.DefenceVal
	}

	go func() {
		if req.Side == pb.PlayerSide_HOST && !req.IsWatcher {
			stateMod.IsHostReady = true
		} else if req.Side == pb.PlayerSide_DUELER && !req.IsWatcher {
			stateMod.IsDuelReady = true
		}
		go (wkbox).SetPara(&stateModkey, stateMod)

		if stateMod.IsHostReady && stateMod.IsDuelReady {
			go this.BroadCast(&pb.GDBroadcastResp{
				RoomKey:      req.RoomKey,
				Msg:          fmt.Sprintf("AD_PHASE:ACK_Both_SideResolve:"),
				Command:      pb.CastCmd_GET_MOVE_PHASE_RESULT,
				CurrentPhase: pb.EventHookPhase_attack_card_drop_phase,
				PhaseHook:    pb.EventHookType_Proxy,
			})
			wkbox1 := this.searchAliveClient()
			var gamDT pb.GameDataSet
			wkbox1.GetPara(&req.RoomKey, &gamDT)
			gamDT.HookType = pb.EventHookType_After
			go this.moveNextPhase(&gamDT, &stateMod)
		}
	}()

	return &pb.GDADResultResp{
		RoomKey:      req.RoomKey,
		Side:         side,
		CurrentPhase: snapMod.EventPhase,
		Point:        pt,
	}, nil

}
func (this *ULZGameDuelServiceBackend) ADPhaseDiceResult(context.Context, *pb.GDGetInfoReq) (*pb.GDADDiceResult, error) {
	return nil, status.Error(codes.Unimplemented, "AD_PHASE_DICE_RESULT")
}
