package serverCtl

import (
	pb "ULZGameDuelService/proto"
	"fmt"
	// Static files
	// _ "ULZGameDuelService/statik"
)

// Handle Proxy phase
func (this *ULZGameDuelServiceBackend) proxyHandle(gameSet *pb.GameDataSet) {
	switch gameSet.EventPhase {
	case pb.EventHookPhase_gameset_start:
		gamesetStart(gameSet)
		break
	case pb.EventHookPhase_start_turn_phase:
		startTurnPhase(gameSet)
		break
	case pb.EventHookPhase_determine_move_phase:
		// refillActionCard(gameSet)
		break
	case pb.EventHookPhase_finish_move_phase:

	case pb.EventHookPhase_chara_change_phase:

	case pb.EventHookPhase_determine_chara_change_phase:

	case pb.EventHookPhase_determine_battle_point_phase:

	case pb.EventHookPhase_battle_result_phase:

	case pb.EventHookPhase_damage_phase:

	case pb.EventHookPhase_dead_chara_change_phase:

	case pb.EventHookPhase_determine_dead_chara_change_phase:

	case pb.EventHookPhase_change_initiative_phase:

	case pb.EventHookPhase_finish_turn_phase:

	case pb.EventHookPhase_gameset_end:
	}
	fmt.Println("Hello world")
	go func() {
		wkbox := this.searchAliveClient()
		if _, err := wkbox.SetPara(&gameSet.RoomKey, gameSet); err != nil {
			fmt.Println(err)
		}
		wkbox.Preserve(false)
	}()
	return
}

func gamesetStart(gameSet *pb.GameDataSet) {

}

func startTurnPhase(gameSet *pb.GameDataSet) {

}

func determineMovePhase(gameSet *pb.GameDataSet) {

}

func finishMovePhase(gameSet *pb.GameDataSet) {}

func determineBattlePointPhase(gameSet *pb.GameDataSet) {}

func battleResultPhase(gameSet *pb.GameDataSet) {}

func damageResultPhase(gameSet *pb.GameDataSet) {}

func deadCharaChangePhase(gameSet *pb.GameDataSet) {}

func determineDeadCharaChangePhase(gameSet *pb.GameDataSet) {}

func changeInitiativePhase(gameSet *pb.GameDataSet) {}

func finishTurnPhase(gameSet *pb.GameDataSet) {}
