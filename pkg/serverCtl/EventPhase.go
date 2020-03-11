package serverCtl

import (
	pb "ULZGameDuelService/proto"
	"context"
	"fmt"
	"sync"

	"github.com/gogo/status"
	"google.golang.org/grpc/codes"
	// Static files
	// _ "ULZGameDuelService/statik"
)

/**
 * 		go-routine {
 * 			---- start_turn_phase
 * 				[start_turn_phase:before] {
 * 					phaseTrigEf ()
 * 				}
 * 			[start_turn_phase:proxy] {
 * 				1. turn ++
 * 			}
 * 				[start_turn_phase:after] {
 * 					phaseTrigEf ()
 * 				}
 * 			-------------------------------------
 * 			draw-phase
 * 			---- refill_action_card_phase
 *
 * 				[refill_action_card_phase:before] {
 *					phaseTrigEf ()
 * 				}
 *			[refill_action_card_phase:proxy] {
 *				1. wait for [*confirm] request
 * 			}
 *				[refill_action_card_phase:before] {
 *					phaseTrigEf ()
 * 				}
 *
 * 			-------------------------------------
 *			----  determine_battle_point_phase
 *
 * 				[determine_battle_point_phase:before] {
 * 					phaseTrigEf ()
 * 				}
 *
 * 			[determine_battle_point_phase:proxy] {
 * 				1. dice-roll from sub-client
 * 				2. store dice-roll first-result
 * 			}
 * 			 	[determine_battle_point_phase:after] {
 * 					phaseTrigEf ()
 * 				}
 * 			-------------------------------------
 *
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

func (this *ULZGameDuelServiceBackend) moveNextPhase(gameDS *pb.GameDataSet, shiftPhase pb.EventHookPhase, shiftType pb.EventHookType) {
	switch gameDS.HookType {
	case pb.EventHookType_Before, pb.EventHookType_After:
		this.phaseTrigEf(gameDS)
		break
	case pb.EventHookType_Proxy:
		this.proxyHandle(gameDS)
		break
	}
	// upshift the phase
	this.upshiftPhase(gameDS, shiftPhase, shiftType)

}

// point calculate
func pointCalcute(inst *int32, orig *int32, value *int32) {
	*inst += *value
	if *inst > *orig {
		*inst = *orig
	}
	if *inst < 0 {
		*inst = 0
	}
}

func (this *ULZGameDuelServiceBackend) upshiftPhase(gameSet *pb.GameDataSet, shiftPhase pb.EventHookPhase, shiftType pb.EventHookType) (bool, error) {
	fmt.Printf("current phase: %#v,\tcurrent hook: %#v ;\n", gameSet.EventPhase, gameSet.HookType)
	fmt.Printf("target phase: %#v ,\ttarget hook: %#v; \n", shiftPhase, shiftType)
	assigner := "upshiftPhaseHandler"

	// check dead char
	isHostDead := (gameSet.HostCardDeck[gameSet.HostCurrCardKey].HpInst <= 0)
	isDuelDead := (gameSet.DuelCardDeck[gameSet.DuelCurrCardKey].HpInst <= 0)

	hostAllDead := 0
	duelAllDead := 0

	wg := sync.WaitGroup{}
	if gameSet.Nvn > 1 {
		wg.Add(2)
		go func() {
			for k := range gameSet.HostCardDeck {
				if gameSet.HostCardDeck[k].HpInst <= 0 {
					hostAllDead++
				}
			}
			wg.Done()
		}()
		go func() {
			for k := range gameSet.DuelCardDeck {
				if gameSet.DuelCardDeck[k].HpInst <= 0 {
					duelAllDead++
				}
			}
			wg.Done()
		}()
		wg.Wait()
	}

	ChangeFlag := false
	EndGameFlag := false
	wg.Add(1)
	go func() {
		if isHostDead && len(gameSet.HostCardDeck) > hostAllDead {
			this.BroadCast(&gameSet.RoomKey, &assigner, &pb.GDBroadcastResp{
				RoomKey: gameSet.RoomKey,
				Msg:     fmt.Sprintf("HostCharIsDead,ChangeChar"),
				Command: pb.CastCmd_INSTANCE_STATUS_CHANGE,
			})
		} else if isHostDead && len(gameSet.HostCardDeck) == hostAllDead {
			this.BroadCast(&gameSet.RoomKey, &assigner, &pb.GDBroadcastResp{
				RoomKey: gameSet.RoomKey,
				Msg:     fmt.Sprintf("HostCharIsDead"),
				Command: pb.CastCmd_INSTANCE_STATUS_CHANGE,
			})
		}
		wg.Done()
	}()
	go func() {
		if isDuelDead && len(gameSet.DuelCardDeck) > duelAllDead {
			this.BroadCast(&gameSet.RoomKey, &assigner, &pb.GDBroadcastResp{
				RoomKey: gameSet.RoomKey,
				Msg:     fmt.Sprintf("DuelCharIsDead,ChangeChar"),
				Command: pb.CastCmd_INSTANCE_STATUS_CHANGE,
			})
		} else if isDuelDead && len(gameSet.DuelCardDeck) == duelAllDead {
			this.BroadCast(&gameSet.RoomKey, &assigner, &pb.GDBroadcastResp{
				RoomKey: gameSet.RoomKey,
				Msg:     fmt.Sprintf("DuelCharIsDead"),
				Command: pb.CastCmd_INSTANCE_STATUS_CHANGE,
			})
		}
		wg.Done()
	}()
	wg.Wait()
	// wg.Add(1)
	if EndGameFlag {
		gameSet.EventPhase = pb.EventHookPhase_gameset_end
		gameSet.HookType = pb.EventHookType_Before
		return true, nil
	}
	if ChangeFlag {
		gameSet.EventPhase = pb.EventHookPhase_dead_chara_change_phase
		gameSet.HookType = pb.EventHookType_Before
		return true, nil
	}

	// --------------------------------------------------
	// start shift-next
	if gameSet.HookType == pb.EventHookType_Before {
		gameSet.HookType = pb.EventHookType_Proxy
		return false, nil
	} else if gameSet.HookType == pb.EventHookType_After {
		gameSet.HookType = pb.EventHookType_Before
	}

	// switch gameSet.EventPhase {
	// case pb.EventHookPhase_battle_result_phase:
	// 	gameSet.EventPhase = pb.EventHookPhase_finish_turn_phase
	// 	break
	// case pb.EventHookPhase_determine_dead_chara_change_phase:
	// 	break
	// case pb.EventHookPhase_change_initiative_phase:
	// 	gameSet.EventPhase = pb.EventHookPhase_attack_card_drop_phase
	// 	break
	// default:
	// 	gameSet.EventPhase++
	// }

	return false, nil
}

func isManualHandlePhase(in pb.EventHookPhase) bool {
	switch in {
	case pb.EventHookPhase_move_card_drop_phase:
	case pb.EventHookPhase_chara_change_phase:
	case pb.EventHookPhase_attack_card_drop_phase:
	case pb.EventHookPhase_defence_card_drop_phase:
	case pb.EventHookPhase_dead_chara_change_phase:
		return true
	}
	return false
}
