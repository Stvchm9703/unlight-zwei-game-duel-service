package serverCtl

import (
	cm "ULZGameDuelService/pkg/common"
	pb "ULZGameDuelService/proto"
	"context"
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
	incomePhase := (snapMod.FirstAttack == snapMod.CurrPhase)
	if isAttack && !incomePhase || isDefence && incomePhase {
		return nil, status.Error(codes.Internal, "AD_PHASE:InvaildPhase")
	}

	if isAttack {
		// snapMod.HostCard
		return nil, status.Error(codes.Unimplemented, "AD_PHASE_CONFIRM")
	} else if isDefence {

	} else {

	}
	return nil, status.Error(codes.Unimplemented, "AD_PHASE_CONFIRM")
}

func (this *ULZGameDuelServiceBackend) ADPhaseResult(context.Context, *pb.GDGetInfoReq) (*pb.GDADResultResp, error) {
	return nil, status.Error(codes.Unimplemented, "AD_PHASE_RESULT")
}
func (this *ULZGameDuelServiceBackend) ADPhaseDiceResult(context.Context, *pb.GDGetInfoReq) (*pb.GDADDiceResult, error) {
	return nil, status.Error(codes.Unimplemented, "AD_PHASE_DICE_RESULT")
}
