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
func (this *ULZGameDuelServiceBackend) EventPhaseConfirm(ctx context.Context, req *pb.GDPhaseConfirmReq) (*pb.Empty, error) {

	return nil, status.Error(codes.Unimplemented, "EVENT_PHASE_CONFIRM")
}
func (this *ULZGameDuelServiceBackend) EventPhaseResult(ctx context.Context, req *pb.GDGetInfoReq) (*pb.GDPhaseConfirmResp, error) {
	// get phase-mod
	cm.PrintReqLog(ctx, "Event-Phase-Result", req)
	start := time.Now()
	this.mu.Lock()
	defer func() {
		// wkbox.Preserve(false)
		this.mu.Unlock()
		elapsed := time.Since(start)
		log.Printf("Event-Phase-Result took %s", elapsed)
	}()

	if req.IsWatcher {
		return nil, status.Error(codes.InvalidArgument, "ERR_PLAYER")
	}

	var snapPhase pb.PhaseSnapMod
	snapPhasekey := req.RoomKey + snapPhase.RdsKeyName()
	wkbox := this.searchAliveClient()
	if _, err := (wkbox).GetPara(snapPhasekey, &snapPhase); err != nil {
		log.Println(err)
		return nil, status.Error(codes.NotFound, err.Error())
	}

	if req.CurrentPhase != snapPhase.EventPhase {
		wkbox.Preserve(false)
		return nil, status.Error(codes.InvalidArgument, "ERR_PHASE")
	}

	if snapPhase.HookType != pb.EventHookType_Proxy {
		wkbox.Preserve(false)
		return nil, status.Error(codes.InvalidArgument, "ERR_PHASE")
	}

	if (req.Side == pb.PlayerSide_HOST && snapPhase.IsHostReady) ||
		(req.Side == pb.PlayerSide_DUELER && snapPhase.IsDuelReady) {
		wkbox.Preserve(false)
		return nil, status.Error(codes.AlreadyExists, "ALREADY_READY")
	}

	if req.Side == pb.PlayerSide_HOST {
		snapPhase.IsHostReady = true
	} else if req.Side == pb.PlayerSide_DUELER {
		snapPhase.IsDuelReady = true
	}

	if _, err := wkbox.SetPara(snapPhasekey, snapPhase); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if snapPhase.IsDuelReady && snapPhase.IsHostReady {
		go func() {
			wg := sync.WaitGroup{}
			wg.Add(2)
			var gameSet pb.GameDataSet
			var gameSetKey = req.RoomKey
			go func() {
				wkbox := this.searchAliveClient()
				if _, err := (wkbox).GetPara(gameSetKey, &gameSet); err != nil {
					log.Println(err)
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
				}
				wkbox.Preserve(false)
				wg.Done()
			}()
			wg.Wait()
			log.Printf("Rm:%s, EvtPh:%s, Ready", req.RoomKey, req.CurrentPhase.String())
			this.moveNextPhase(&gameSet, &snapPhase, &effectMod)
		}()
	}
	return &pb.GDPhaseConfirmResp{
		RoomKey: req.RoomKey,
	}, nil
	// wkbox.Preserve(false)
}

func (this *ULZGameDuelServiceBackend) moveNextPhase(
	gameDS *pb.GameDataSet,
	phaseMod *pb.PhaseSnapMod,
	effectMod *pb.EffectNodeSnapMod,
	snapMod ...interface{},
) {
	/**
	 * run the current phase:type
	 * e.g. : phase_a:before => phase_a:proxy
	 * 		: phase_a:after => phase_b:before
	 *
	 * # Proxy may have two-side confirm, therefore it may stop at proxy state
	 */
	this.executeEffectNode(gameDS, phaseMod, effectMod)
	wg := sync.WaitGroup{}

	// =======================================================================
	wg.Add(3)
	go func() {
		wkbox := this.searchAliveClient()
		wkbox.SetPara(gameDS.RoomKey, gameDS)
		wkbox.Preserve(false)
		wg.Done()
	}()
	go func() {
		wkbox := this.searchAliveClient()
		wkbox.SetPara(gameDS.RoomKey+phaseMod.RdsKeyName(), phaseMod)
		wkbox.Preserve(false)
		wg.Done()
	}()
	go func() {
		wkbox := this.searchAliveClient()
		wkbox.SetPara(gameDS.RoomKey+effectMod.RdsKeyName(), effectMod)
		wkbox.Preserve(false)
		wg.Done()
	}()
	wg.Wait()

	if b, _ := this.checkDeadFlaging(gameDS, phaseMod, effectMod); b {
		this.proxyHandle(gameDS, phaseMod, effectMod, snapMod)
		return
	}
	// =======================================================================

	nextPhase, nextType := upnextEventPhase(phaseMod.EventPhase, phaseMod.HookType)

	if phaseMod.FirstAttack == phaseMod.CurrAttack &&
		phaseMod.EventPhase == pb.EventHookPhase_damage_phase &&
		phaseMod.HookType == pb.EventHookType_After {
		gameDS.EventPhase = pb.EventHookPhase_change_initiative_phase
		phaseMod.EventPhase = pb.EventHookPhase_change_initiative_phase
		gameDS.HookType = pb.EventHookType_Before
		phaseMod.HookType = pb.EventHookType_Before
	} else {
		gameDS.EventPhase = nextPhase
		gameDS.HookType = nextType
		phaseMod.EventPhase = nextPhase
		phaseMod.HookType = nextType
	}

	// after shift
	if phaseMod.HookType == pb.EventHookType_Proxy {
		this.proxyHandle(gameDS, phaseMod, effectMod, snapMod)
		return
	}

	// upshift the phase

	// non-proxy: move next again
	this.moveNextPhase(gameDS, phaseMod, effectMod, snapMod)

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

func (this *ULZGameDuelServiceBackend) checkDeadFlaging(
	gameSet *pb.GameDataSet,
	phaseMod *pb.PhaseSnapMod,
	effectMod *pb.EffectNodeSnapMod,
) (bool, error) {
	fmt.Printf("current phase: %#v,\tcurrent hook: %#v ;\n", gameSet.EventPhase, gameSet.HookType)
	// fmt.Printf("target phase: %#v ,\ttarget hook: %#v; \n", shiftPhase, shiftType)
	// assigner := "upshiftPhaseHandler"
	//=================================================================
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
	wg.Add(2)
	go func() {
		if isHostDead && len(gameSet.HostCardDeck) > hostAllDead {
			this.BroadCast(&pb.GDBroadcastResp{
				RoomKey: gameSet.RoomKey,
				Msg:     fmt.Sprintf("HostCharIsDead,ChangeChar"),
				Command: pb.CastCmd_INSTANCE_STATUS_CHANGE,
			})
		} else if isHostDead && len(gameSet.HostCardDeck) == hostAllDead {
			this.BroadCast(&pb.GDBroadcastResp{
				RoomKey: gameSet.RoomKey,
				Msg:     fmt.Sprintf("HostCharIsDead"),
				Command: pb.CastCmd_INSTANCE_STATUS_CHANGE,
			})
			EndGameFlag = true
		}
		wg.Done()
	}()
	go func() {
		if isDuelDead && len(gameSet.DuelCardDeck) > duelAllDead {
			this.BroadCast(&pb.GDBroadcastResp{
				RoomKey: gameSet.RoomKey,
				Msg:     fmt.Sprintf("DuelCharIsDead,ChangeChar"),
				Command: pb.CastCmd_INSTANCE_STATUS_CHANGE,
			})
		} else if isDuelDead && len(gameSet.DuelCardDeck) == duelAllDead {
			this.BroadCast(&pb.GDBroadcastResp{
				RoomKey: gameSet.RoomKey,
				Msg:     fmt.Sprintf("DuelCharIsDead"),
				Command: pb.CastCmd_INSTANCE_STATUS_CHANGE,
			})
			EndGameFlag = true
		}
		wg.Done()
	}()
	wg.Wait()
	// wg.Add(1)
	if EndGameFlag {
		gameSet.EventPhase = pb.EventHookPhase_gameset_end
		gameSet.HookType = pb.EventHookType_Proxy
		phaseMod.EventPhase = pb.EventHookPhase_gameset_end
		phaseMod.HookType = pb.EventHookType_Proxy
	}
	if ChangeFlag {
		gameSet.EventPhase = pb.EventHookPhase_dead_chara_change_phase
		gameSet.HookType = pb.EventHookType_Proxy
		phaseMod.EventPhase = pb.EventHookPhase_dead_chara_change_phase
		phaseMod.HookType = pb.EventHookType_Proxy
	}
	//=================================================================

	// --------------------------------------------------
	// start shift-next
	return EndGameFlag || ChangeFlag, nil
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

func (this *ULZGameDuelServiceBackend) executeEffectNode(
	gameSet *pb.GameDataSet,
	stateMod *pb.PhaseSnapMod,
	effectMod *pb.EffectNodeSnapMod,
) (*pb.GameDataSet, error) {
	nodelist := pb.NodeFilter(effectMod.PendingEf, func(v *pb.EffectResult) bool {
		return v.TriggerTime.EventPhase == stateMod.EventPhase &&
			v.TriggerTime.HookType == stateMod.HookType
	})
	gameSet1 := gameSet
	var err error
	if len(nodelist) > 0 {
		nextPhase, nextType := upnextEventPhase(stateMod.EventPhase, stateMod.HookType)
		gameSet1, err = this.skillClient.EffectCalculateWrap(
			gameSet.RoomKey,
			&pb.EffectTiming{
				EventPhase: stateMod.EventPhase,
				HookType:   stateMod.HookType,
			},
			&pb.EffectTiming{
				EventPhase: nextPhase,
				HookType:   nextType,
			},
			gameSet,
		)
		cleanEffectResult(stateMod.EventPhase, stateMod.HookType, effectMod)
	}
	fmt.Printf("updated gameSet %#v \n err?: %v\n", gameSet1, err)
	return gameSet1, err
}

func upnextEventPhase(
	inPhase pb.EventHookPhase, inType pb.EventHookType,
) (outPhase pb.EventHookPhase, outType pb.EventHookType) {
	if inType == pb.EventHookType_Proxy {
		outPhase = inPhase
		outType = pb.EventHookType_After
	} else if inType == pb.EventHookType_Before {
		outPhase = inPhase
		outType = pb.EventHookType_Proxy
	} else if inType == pb.EventHookType_After {
		outType = pb.EventHookType_Before
		switch inPhase {
		case pb.EventHookPhase_start_turn_phase:
			outPhase = pb.EventHookPhase_refill_action_card_phase
			break
		case pb.EventHookPhase_refill_action_card_phase:
			outPhase = pb.EventHookPhase_move_card_drop_phase
			break
		case pb.EventHookPhase_move_card_drop_phase:
			outPhase = pb.EventHookPhase_determine_move_phase
			break
		case pb.EventHookPhase_determine_move_phase:
			outPhase = pb.EventHookPhase_finish_move_phase
			break
		case pb.EventHookPhase_finish_move_phase:
			outPhase = pb.EventHookPhase_attack_card_drop_phase
			break
		case pb.EventHookPhase_attack_card_drop_phase:
			outPhase = pb.EventHookPhase_defence_card_drop_phase
			break
		case pb.EventHookPhase_defence_card_drop_phase:
			outPhase = pb.EventHookPhase_determine_battle_point_phase
			break
		case pb.EventHookPhase_determine_battle_point_phase:
			outPhase = pb.EventHookPhase_battle_result_phase
			break
		case pb.EventHookPhase_battle_result_phase:
			outPhase = pb.EventHookPhase_damage_phase
			break
		case pb.EventHookPhase_dead_chara_change_phase:
			outPhase = pb.EventHookPhase_determine_dead_chara_change_phase
			break
		case pb.EventHookPhase_determine_dead_chara_change_phase:
			outPhase = pb.EventHookPhase_change_initiative_phase
			break
		case pb.EventHookPhase_change_initiative_phase:
			outPhase = pb.EventHookPhase_attack_card_drop_phase
			break

		default:
			outPhase = inPhase + 1
		}
	}
	return
}

func cleanEffectResult(phase pb.EventHookPhase, ehType pb.EventHookType, ens *pb.EffectNodeSnapMod) {
	var removeKey []int
	pef := ens.PendingEf
	for k, v := range ens.PendingEf {
		if v.EndTime.EventPhase == phase && v.EndTime.HookType == ehType {
			v.RemainCd--
		}
		if v.RemainCd == 0 {
			removeKey = append(removeKey, k)
		}
	}
	for _, v := range removeKey {
		pef = append(pef[:v], pef[v+1:]...)
	}
	ens.PendingEf = pef
}
