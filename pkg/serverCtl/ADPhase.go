package serverCtl

import (
	cm "ULZGameDuelService/pkg/common"
	pb "ULZGameDuelService/proto"
	"context"
	"fmt"
	"log"
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

	var returner pb.GameDataSet
	if _, err := (wkbox).GetPara(&req.RoomKey, &returner); err != nil {
		log.Println(err)
		return nil, status.Errorf(codes.NotFound, err.Error())
	}

	var snapModkey = req.RoomKey + ":ADPhMod"
	var snapMod pb.ADPhaseSnapMod
	if _, err := (wkbox).GetPara(&snapModkey, &snapMod); err != nil {
		log.Println(err)
		// return nil, status.Errorf(codes.NotFound, err.Error())
	}

	//  attack phase
	if req.CurrentPhase != snapMod.EventPhase {
		return nil, status.Error(codes.InvalidArgument, "AD_PHASE:InvaildPhase")
	}
	// FirstAttack = Host & CurrPhase = Host  -> Host is First-Attack
	isAttack := (snapMod.EventPhase == pb.EventHookPhase_attack_card_drop_phase)
	isDefence := (snapMod.EventPhase == pb.EventHookPhase_defence_card_drop_phase)
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
		go this.attackPhaseHandle(&returner, &snapMod)

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
		go this.defencePhaseHandle(&returner, &snapMod)

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

	var snapModkey = req.RoomKey + ":ADPhMod"
	var snapMod pb.ADPhaseSnapMod
	if _, err := (wkbox).GetPara(&snapModkey, &snapMod); err != nil {
		log.Println(err)
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	// return nil, status.Error(codes.Unimplemented, "AD_PHASE_RESULT")
	side := snapMod.CurrAttacker
	var pt int32 = 0
	if snapMod.EventPhase == pb.EventHookPhase_attack_card_drop_phase {
		side = snapMod.CurrAttacker
		pt = snapMod.AttackVal
	} else {
		if snapMod.CurrAttacker == pb.PlayerSide_HOST {
			side = pb.PlayerSide_DUELER
		} else {
			side = pb.PlayerSide_HOST
		}
		pt = snapMod.DefenceVal
	}

	if snapMod.EventPhase == pb.EventHookPhase_attack_card_drop_phase &&
		req.Side == snapMod.CurrAttacker {
		// go shiftNext(&req)
		// snapMod.
	}
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

func (this *ULZGameDuelServiceBackend) attackPhaseHandle(gameDS *pb.GameDataSet, snapMod *pb.ADPhaseSnapMod) {
	// do effect calculate
	result := 0
	// do update
	wkbox := this.searchAliveClient()
	snapMod.AttackVal = int32(result)
	snapMod.IsProcessed = true
	var snapModkey = gameDS.RoomKey + ":ADPhMod"
	if _, err := (wkbox).SetPara(&snapModkey, snapMod); err != nil {
		log.Println(err)
	}
	// send ok message
	go this.BroadCast(&gameDS.RoomKey, &snapModkey, &pb.GDBroadcastResp{
		RoomKey:      gameDS.RoomKey,
		Msg:          fmt.Sprintf("AD_PHASE:ATK_RESULT:", result),
		Command:      pb.CastCmd_GET_ATK_PHASE_RESULT,
		Side:         snapMod.CurrAttacker,
		CurrentPhase: pb.EventHookPhase_attack_card_drop_phase,
		PhaseHook:    pb.EventHookType_Proxy,
	})

}

func (this *ULZGameDuelServiceBackend) defencePhaseHandle(gameDS *pb.GameDataSet, snapMod *pb.ADPhaseSnapMod) {
	// do effect calculate
	result := 0
	// do update
	wkbox := this.searchAliveClient()
	snapMod.AttackVal = int32(result)
	var snapModkey = gameDS.RoomKey + ":ADPhMod"
	if _, err := (wkbox).SetPara(&snapModkey, snapMod); err != nil {
		log.Println(err)
	}
	// send ok message
	side := snapMod.CurrAttacker
	if side == 0 {
		side = 1
	} else {
		side = 0
	}
	go this.BroadCast(&gameDS.RoomKey, &snapModkey, &pb.GDBroadcastResp{
		RoomKey:      gameDS.RoomKey,
		Msg:          fmt.Sprintf("AD_PHASE:DEF_RESULT:", result),
		Command:      pb.CastCmd_GET_DEF_PHASE_RESULT,
		Side:         side,
		CurrentPhase: pb.EventHookPhase_defence_card_drop_phase,
		PhaseHook:    pb.EventHookType_Proxy,
	})
}
