package serverCtl

import (
	pb "ULZGameDuelService/proto"
	"fmt"
	"log"
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
		break
	case pb.EventHookPhase_start_turn_phase:
		this.startTurnPhase(gameSet)
		break
	case pb.EventHookPhase_refill_action_card_phase:
		this.refillActionCard(gameSet)
		break
	case pb.EventHookPhase_move_card_drop_phase:
		this.moveCardDropPhase(gameSet)
		break
	case pb.EventHookPhase_determine_move_phase:
		snapMovMod, _ := snapMod[0].(*pb.MovePhaseSnapMod)
		this.movePhaseHandle(gameSet, phaseMod, effectMod, snapMovMod)
		break
	case pb.EventHookPhase_finish_move_phase:
		this.finishMovePhase(gameSet)
		break
	case pb.EventHookPhase_chara_change_phase:
		// this.charaChangePhase(gameSet)
		break
	case pb.EventHookPhase_determine_chara_change_phase:
		//
	// case pb.EventHookPhase_attack_card_drop_phase:
	// 		snapADMod, _ := snapMod[0].(*pb.ADPhaseSnapMod)
	// 		this.attackPhaseHandle(gameSet, snapADMod, phaseMod, effectMod)
	// 		break
	// case pb.EventHookPhase_defence_card_drop_phase:
	// 		snapADMod, _ := snapMod[0].(*pb.ADPhaseSnapMod)
	// 		this.defencePhaseHandle(gameSet, snapADMod, phaseMod, effectMod)
	// 		break
	case pb.EventHookPhase_determine_battle_point_phase:

	case pb.EventHookPhase_battle_result_phase:
		snapADMod, _ := snapMod[0].(*pb.ADPhaseSnapMod)
		this.battlePhaseHandle(gameSet, phaseMod, effectMod, snapADMod)

	case pb.EventHookPhase_damage_phase:

	case pb.EventHookPhase_dead_chara_change_phase:

	case pb.EventHookPhase_determine_dead_chara_change_phase:

	case pb.EventHookPhase_change_initiative_phase:

	case pb.EventHookPhase_finish_turn_phase:

	case pb.EventHookPhase_gameset_end:
	}
	fmt.Println("Hello world")

	return
}

func (this *ULZGameDuelServiceBackend) gamesetStart(gameSet *pb.GameDataSet) {
	log.Printf("game-start: key: %s\n", gameSet.RoomKey)
}

func (this *ULZGameDuelServiceBackend) startTurnPhase(gameSet *pb.GameDataSet) {
	log.Printf("start-turn-phase: turn %v", gameSet.GameTurn)
	gameSet.GameTurn++
}

func (this *ULZGameDuelServiceBackend) refillActionCard(gameSet *pb.GameDataSet) {
	log.Printf("refill-action-card-phase: turn %v", gameSet.RoomKey)

}
func (this *ULZGameDuelServiceBackend) determineMovePhase(gameSet *pb.GameDataSet) {

}

func (this *ULZGameDuelServiceBackend) finishMovePhase(gameSet *pb.GameDataSet) {}

func (this *ULZGameDuelServiceBackend) determineBattlePointPhase(gameSet *pb.GameDataSet) {}

func (this *ULZGameDuelServiceBackend) damageResultPhase(gameSet *pb.GameDataSet) {}

func (this *ULZGameDuelServiceBackend) deadCharaChangePhase(gameSet *pb.GameDataSet) {}

func (this *ULZGameDuelServiceBackend) determineDeadCharaChangePhase(gameSet *pb.GameDataSet) {}

func (this *ULZGameDuelServiceBackend) changeInitiativePhase(gameSet *pb.GameDataSet) {}

func (this *ULZGameDuelServiceBackend) finishTurnPhase(gameSet *pb.GameDataSet) {}
