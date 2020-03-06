package serverCtl

import (
	pb "ULZGameDuelService/proto"
	"context"
	"fmt"
	"log"
	"sort"
	"sync"

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

func (this *ULZGameDuelServiceBackend) moveNextPhase(gameDS *pb.GameDataSet, shiftPhase pb.EventHookPhase, shiftType pb.EventCardType) {
	switch gameDS.HookType {
	case pb.EventHookType_Before, pb.EventHookType_After:
		this.phaseTrigEf(gameDS, shiftPhase, shiftType)
		break
	case pb.EventHookType_Proxy:
		break
	}
	// upshift the phase
}

// phaseTrigEf : general phase trigger effect
// it only handle Instance_Change / direct-dmg
// NOTE not available for atk/def, move phase calculation
func (this *ULZGameDuelServiceBackend) phaseTrigEf(gameDS *pb.GameDataSet, shiftPhase pb.EventHookPhase, shiftType pb.EventCardType) {
	var efResult pb.EffectNodeSnapMod
	var efResList []*pb.EffectResult
	taskHandler := "phaseTrigEf"
	wkbox := this.searchAliveClient()
	searchKey := gameDS.RoomKey + ":"
	if gameDS.EffectCounter != nil {
		efResList = gameDS.EffectCounter
	} else {
		if _, err := wkbox.GetPara(&searchKey, efResult); err != nil {
			log.Println(err)
		}
	}
	if len(efResList) == 0 {
		return
	}

	tarEf := nodeFilter(efResList, func(v *pb.EffectResult) bool {
		return (v.TriggerTime.EventPhase == gameDS.EventPhase) &&
			(v.TriggerTime.HookType == gameDS.HookType)
	})
	if len(tarEf) == 0 {
		return
	}
	sort.Slice(tarEf, func(i, j int) bool {
		return tarEf[i].TriggerTime.SubCount < tarEf[i].TriggerTime.SubCount
	})

	DirectDmg := nodeFilter(tarEf, func(v *pb.EffectResult) bool {
		return (v.EfOption == pb.EffectOption_Instance_Change)
	})

	// return be4 run loop
	if len(DirectDmg) == 0 {
		return
	}

	wg := sync.WaitGroup{}

	// Status release
	// 	 Damage part first
	wg.Add(3)
	go func() {
		bcMsg := pb.GDBroadcastResp{
			RoomKey:      gameDS.RoomKey,
			Msg:          fmt.Sprintf("Damage from effect to player"),
			Command:      pb.CastCmd_GET_INSTANCE_CARD,
			CurrentPhase: gameDS.EventPhase,
			PhaseHook:    gameDS.HookType,
			EffectTrig:   DirectDmg,
		}
		this.BroadCast(&gameDS.RoomKey, &taskHandler, &bcMsg)
		wg.Done()
	}()
	go func() {
		for _, v := range DirectDmg {
			if v.TarSide == pb.PlayerSide_HOST {
				// hp change
				pointCalcute(&gameDS.HostCardDeck[v.TarCard].HpInst, &gameDS.HostCardDeck[v.TarCard].HpOrig, &v.Hp)
				// ap change
				pointCalcute(&gameDS.HostCardDeck[v.TarCard].ApInst, &gameDS.HostCardDeck[v.TarCard].ApOrig, &v.Ap)
				// dp change
				pointCalcute(&gameDS.HostCardDeck[v.TarCard].DpInst, &gameDS.HostCardDeck[v.TarCard].DpOrig, &v.Dp)

			} else {
				// hp change
				pointCalcute(&gameDS.DuelCardDeck[v.TarCard].HpInst, &gameDS.DuelCardDeck[v.TarCard].HpOrig, &v.Hp)
				// ap change
				pointCalcute(&gameDS.DuelCardDeck[v.TarCard].ApInst, &gameDS.DuelCardDeck[v.TarCard].ApOrig, &v.Ap)
				// dp change
				pointCalcute(&gameDS.DuelCardDeck[v.TarCard].DpInst, &gameDS.DuelCardDeck[v.TarCard].DpOrig, &v.Dp)
			}
			fmt.Println(v)
		}
		wg.Done()
	}()
	wg.Wait()

	hostFixEf := nodeFilter(tarEf, func(v *pb.EffectResult) bool {
		return (v.EfOption == pb.EffectOption_Status_FixValue) &&
			(v.TarCard == gameDS.HostCurrCardKey) &&
			(v.TarSide == pb.PlayerSide_HOST)
	})
	hostFixFin := pb.EffectResult{}
	if len(hostFixEf) > 0 {
		for _, v := range hostFixEf {
			if v.Hp > hostFixFin.Hp {
				hostFixFin.Hp = v.Hp
			}
			if v.Ap > hostFixFin.Ap {
				hostFixFin.Ap = v.Ap
			}
			if v.Dp > hostFixFin.Dp {
				hostFixFin.Dp = v.Dp
			}
		}
	}

	duelFixEf := nodeFilter(tarEf, func(v *pb.EffectResult) bool {
		return (v.EfOption == pb.EffectOption_Status_FixValue) &&
			(v.TarCard == gameDS.DuelCurrCardKey) &&
			(v.TarSide == pb.PlayerSide_DUELER)
	})

	duelFixFin := pb.EffectResult{}
	if len(hostFixEf) > 0 {
		for _, v := range hostFixEf {
			if v.Hp > duelFixFin.Hp {
				duelFixFin.Hp = v.Hp
			}
			if v.Ap > duelFixFin.Ap {
				duelFixFin.Ap = v.Ap
			}
			if v.Dp > duelFixFin.Dp {
				duelFixFin.Dp = v.Dp
			}
		}
	}

	fmt.Printf("hostFix : %#v \n", hostFixEf)
	fmt.Printf("duelFix : %#v \n", duelFixEf)

	return
}

// EffectResult sorting
func nodeFilter(vs []*pb.EffectResult, f func(*pb.EffectResult) bool) []*pb.EffectResult {
	vsf := make([]*pb.EffectResult, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
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
