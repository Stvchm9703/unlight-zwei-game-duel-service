package serverCtl

import (
	pb "ULZGameDuelService/proto"
	"fmt"
	"log"
	"sync"
	// Static files
	// _ "ULZGameDuelService/statik"
)

// Handle Proxy phase
func (this *ULZGameDuelServiceBackend) proxyHandle(
	gameSet *pb.GameDataSet,
	phaseMod *pb.PhaseSnapMod,
	effectMod *pb.EffectNodeSnapMod,
	snapMod ...interface{},
) {
	switch phaseMod.EventPhase {
	case pb.EventHookPhase_gameset_start:
		this.gamesetStart(gameSet)

	case pb.EventHookPhase_start_turn_phase:
		this.startTurnPhase(gameSet)

	// draw-phase
	case pb.EventHookPhase_refill_action_card_phase:
		this.refillActionCard(gameSet, phaseMod, effectMod)

	// move-phase
	case pb.EventHookPhase_move_card_drop_phase:
	// 		this.moveCardDropPhase(gameSet)
	// 		break
	case pb.EventHookPhase_determine_move_phase:
		snapMovMod, _ := snapMod[0].(*pb.MovePhaseSnapMod)
		this.determineMovePhaseHandle(gameSet, phaseMod, effectMod, snapMovMod)

	case pb.EventHookPhase_finish_move_phase:
		snapMovMod, _ := snapMod[0].(*pb.MovePhaseSnapMod)
		this.finishMovePhase(gameSet, phaseMod, effectMod, snapMovMod)

	// char-change-phase
	case pb.EventHookPhase_chara_change_phase:
		// this.charaChangePhase(gameSet)

	case pb.EventHookPhase_determine_chara_change_phase:
		// change char -> skip

		//	attack-phase
	case pb.EventHookPhase_attack_card_drop_phase:
		log.Println("atk-card-drop-ph wait for event-phase-confirm")
	// 		snapADMod, _ := snapMod[0].(*pb.ADPhaseSnapMod)
	// 		this.attackPhaseHandle(gameSet, snapADMod, phaseMod, effectMod)
	// 		break

	// 	defence-phase
	case pb.EventHookPhase_defence_card_drop_phase:
		log.Println("def-card-drop-ph wait for event-phase-confirm")
	// 		snapADMod, _ := snapMod[0].(*pb.ADPhaseSnapMod)
	// 		this.defencePhaseHandle(gameSet, snapADMod, phaseMod, effectMod)
	// 		break

	case pb.EventHookPhase_determine_battle_point_phase:
		// see also : battlePhase.go
		snapADMod, _ := snapMod[0].(*pb.ADPhaseSnapMod)
		this.determineBattlePointPhase(gameSet, phaseMod, effectMod, snapADMod)

	case pb.EventHookPhase_battle_result_phase:
		// see also : battlePhase.go
		snapADMod, _ := snapMod[0].(*pb.ADPhaseSnapMod)
		snapMod[0] = this.battlePhaseHandle(gameSet, phaseMod, effectMod, snapADMod)

	case pb.EventHookPhase_damage_phase:
		snapResultMod, _ := snapMod[0].(*pb.GDADDiceResult)
		this.damagePhaseHandle(gameSet, phaseMod, effectMod, snapResultMod)
		// see also : battlePhase.go

	case pb.EventHookPhase_dead_chara_change_phase:
		this.deadCharaChangePhase(gameSet)

	case pb.EventHookPhase_determine_dead_chara_change_phase:
		this.determineDeadCharaChangePhase(gameSet)

	case pb.EventHookPhase_change_initiative_phase:
		this.changeInitiativePhase(gameSet, phaseMod)

	case pb.EventHookPhase_finish_turn_phase:
		this.finishTurnPhase(gameSet, phaseMod)

	case pb.EventHookPhase_gameset_end:
	}
	fmt.Println("Hello world")

	return
}

func (this *ULZGameDuelServiceBackend) gamesetStart(gameSet *pb.GameDataSet) {
	log.Printf("game-start: key: %s\n", gameSet.RoomKey)
	go this.BroadCast(&pb.GDBroadcastResp{
		RoomKey:      gameSet.RoomKey,
		Msg:          fmt.Sprintf("GameStart:%v", gameSet.GameTurn),
		Command:      pb.CastCmd_INSTANCE_STATUS_CHANGE,
		CurrentPhase: gameSet.EventPhase,
		PhaseHook:    gameSet.HookType,
	})

}

func (this *ULZGameDuelServiceBackend) startTurnPhase(gameSet *pb.GameDataSet) {
	log.Printf("start-turn-phase: turn %v", gameSet.GameTurn)
	gameSet.GameTurn++
	this.BroadCast(&pb.GDBroadcastResp{
		RoomKey:      gameSet.RoomKey,
		Msg:          fmt.Sprintf("TurnStart:%v", gameSet.GameTurn),
		Command:      pb.CastCmd_INSTANCE_STATUS_CHANGE,
		CurrentPhase: gameSet.EventPhase,
		PhaseHook:    gameSet.HookType,
	})

	wg := sync.WaitGroup{}
	wg.Add(2)
	// move-instance
	go func() {
		mwkbox := this.searchAliveClient()
		move_instance := pb.MovePhaseSnapMod{
			Turns: gameSet.GameTurn,
		}
		if _, err := mwkbox.SetPara(gameSet.RoomKey+move_instance.RdsKeyName(), move_instance); err != nil {
			log.Println(err)
		}
		mwkbox.Preserve(false)
		wg.Done()
	}()
	// ad-phase-instance
	go func() {
		mwkbox := this.searchAliveClient()
		adInstance := pb.ADPhaseSnapMod{
			Turns:       gameSet.GameTurn,
			FirstAttack: 0,
			EventPhase:  pb.EventHookPhase_start_turn_phase,
		}
		if _, err := mwkbox.SetPara(gameSet.RoomKey+adInstance.RdsKeyName(), adInstance); err != nil {
			log.Println(err)
		}
		mwkbox.Preserve(false)
		wg.Done()
	}()
	wg.Wait()
}

func (this *ULZGameDuelServiceBackend) deadCharaChangePhase(gameSet *pb.GameDataSet) {}

func (this *ULZGameDuelServiceBackend) determineDeadCharaChangePhase(gameSet *pb.GameDataSet) {}

func (this *ULZGameDuelServiceBackend) changeInitiativePhase(
	gameSet *pb.GameDataSet,
	phaseMod *pb.PhaseSnapMod,
) {
	if phaseMod.CurrAttack == pb.PlayerSide_HOST {
		phaseMod.CurrAttack = pb.PlayerSide_DUELER
	} else if phaseMod.CurrAttack == pb.PlayerSide_DUELER {
		phaseMod.CurrAttack = pb.PlayerSide_HOST
	}
}

func (this *ULZGameDuelServiceBackend) finishTurnPhase(gameSet *pb.GameDataSet, phaseMod *pb.PhaseSnapMod) {
	//  force reset ready flag

}
