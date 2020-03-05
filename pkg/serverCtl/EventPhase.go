package serverCtl

import (
	pb "ULZGameDuelService/proto"
	"context"

	"github.com/gogo/status"
	"google.golang.org/grpc/codes"
	// Static files
	// _ "ULZGameDuelService/statik"
)

/**
 * 		go-routine {
 * 			---- determine_battle_point_phase
 * 			move next phase [determine_battle_point_phase:before] {
 * 				if ef-node > 0
 * 					do exec event-node
 * 						event-phase
 * 						? go send ACK [EventResult][]
 * 			}
 * 			move next phase [determine_battle_point_phase:proxy] {
 * 				1. dice-roll from sub-client
 * 				2. store dice-roll first-result
 * 			}
 * 			move next phase [determine_battle_point_phase:after] {
 * 				if ef-node > 0
 * 					do exec event-node
 * 						event-phase
 * 						? go send ACK [EventResult][]
 * 					store dice-roll final-result
 * 			}
 *			----- battle_result_phase
 * 			move next phase [battle_result_phase : before] {
 * 				if ef-node [battle-result-phase] > 0
 * 					do exec event-node
 * 			}
 *			move next phase [battle_result_phase : proxy] {
 * 				go broadcast request player [ADPhaseDiceResult]
 * 			}
 * 			move next phase [battle_result_phase : after] {
 * 				if ef-node [battle-result-phase] > 0
 * 					do exec event-node
 * 			}
 *
 * 			----- damage_phase -----
 *			move next phase [damage_phase : before] {
 * 				if ef-node [battle-result-phase] > 0
 * 					do exec event-node
 * 			}
 *			move next phase [damage_phase : proxy] {
 * 				update gameDataSet
 * 				? send damage ?
 *			}
 *			move next phase [damage_phase : after] {
 * 				if ef-node [battle-result-phase] > 0
 * 					do exec event-node
 *			}
 *
 * 			----- dead_chara_change_phase -----
 *			move next phase [dead_chara_change_phase : before] {
 * 				if ef-node [battle-result-phase] > 0
 * 					do exec event-node
 * 			}
 *			move next phase [dead_chara_change_phase : proxy] {
 * 				update gameDataSet
 * 				? send damage ?
 *			}
 *			move next phase [dead_chara_change_phase : after] {
 * 				if ef-node [battle-result-phase] > 0
 * 					do exec event-node
 *			}
 * 			----- determine_dead_chara_change_phase -----
 *			move next phase [determine_dead_chara_change_phase : before] {
 * 				if ef-node [battle-result-phase] > 0
 * 					do exec event-node
 * 			}
 *			move next phase [determine_dead_chara_change_phase : proxy] {
 * 				update gameDataSet
 * 				? send damage ?
 *			}
 *			move next phase [determine_dead_chara_change_phase : after] {
 * 				if ef-node [battle-result-phase] > 0
 * 					do exec event-node
 *			}
 * 		}
 *
 */
func (this *ULZGameDuelServiceBackend) EventPhaseConfirm(context.Context, *pb.GDPhaseConfirmReq) (*pb.Empty, error) {
	return nil, status.Error(codes.Unimplemented, "EVENT_PHASE_CONFIRM")
}
func (this *ULZGameDuelServiceBackend) EventPhaseResult(context.Context, *pb.GDGetInfoReq) (*pb.GDPhaseConfirmResp, error) {
	return nil, status.Error(codes.Unimplemented, "EVENT_PHASE_RESULT")
}

func (this *ULZGameDuelServiceBackend) eventPhaseHandle(gameDS *pb.GameDataSet, shiftPhase pb.EventHookPhase, shiftHook pb.EventHookType) {

}
