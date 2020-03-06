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
	sort.Slice(tarEf, func(i, j int) bool {
		return tarEf[i].TriggerTime.SubCount < tarEf[i].TriggerTime.SubCount
	})
	FixEf := nodeFilter(tarEf, func(v *pb.EffectResult) bool {
		return (v.EfOption == pb.EffectOption_Status_FixValue)
	})

	DirectDmg := nodeFilter(tarEf, func(v *pb.EffectResult) bool {
		return (v.EfOption == pb.EffectOption_Instance_Change)
	})

	// return be4 run loop
	if len(DirectDmg) == 0 {
		return
	}
	wg := sync.WaitGroup{}
	for _, v := range DirectDmg {
		bcMsg := pb.GDBroadcastResp{
			RoomKey:      gameDS.RoomKey,
			Msg:          fmt.Sprintf("Damage from ", v.AssignFrom),
			Command:      pb.CastCmd_GET_INSTANCE_CARD,
			CurrentPhase: gameDS.EventPhase,
			PhaseHook:    gameDS.HookType,
		}
		wg.Add(1)
		if v.TarSide == pb.PlayerSide_HOST {
			// hp change
			gameDS.HostCardDeck[v.TarCard].HpInst += v.Hp
			if gameDS.HostCardDeck[v.TarCard].HpInst > gameDS.HostCardDeck[v.TarCard].HpOrig {
				gameDS.HostCardDeck[v.TarCard].HpInst = gameDS.HostCardDeck[v.TarCard].HpOrig
			}
			if gameDS.HostCardDeck[v.TarCard].HpInst < 0 {
				gameDS.HostCardDeck[v.TarCard].HpInst = 0
			}

			// ap change
			gameDS.HostCardDeck[v.TarCard].ApInst += v.Ap
			if gameDS.HostCardDeck[v.TarCard].ApInst > gameDS.HostCardDeck[v.TarCard].ApOrig {
				gameDS.HostCardDeck[v.TarCard].ApInst = gameDS.HostCardDeck[v.TarCard].ApOrig
			}
			if gameDS.HostCardDeck[v.TarCard].ApInst < 0 {
				gameDS.HostCardDeck[v.TarCard].ApInst = 0
			}

			// dp change
			gameDS.HostCardDeck[v.TarCard].DpInst += v.Dp
			if gameDS.HostCardDeck[v.TarCard].DpInst > gameDS.HostCardDeck[v.TarCard].DpOrig {
				gameDS.HostCardDeck[v.TarCard].DpInst = gameDS.HostCardDeck[v.TarCard].DpOrig
			}
			if gameDS.HostCardDeck[v.TarCard].DpInst < 0 {
				gameDS.HostCardDeck[v.TarCard].DpInst = 0
			}

		} else {
			// hp change
			gameDS.DuelCardDeck[v.TarCard].HpInst += v.Hp
			if gameDS.DuelCardDeck[v.TarCard].HpInst > gameDS.DuelCardDeck[v.TarCard].HpOrig {
				gameDS.DuelCardDeck[v.TarCard].HpInst = gameDS.DuelCardDeck[v.TarCard].HpOrig
			}
			if gameDS.DuelCardDeck[v.TarCard].HpInst < 0 {
				gameDS.DuelCardDeck[v.TarCard].HpInst = 0
			}

			// ap change
			gameDS.DuelCardDeck[v.TarCard].ApInst += v.Ap
			if gameDS.DuelCardDeck[v.TarCard].ApInst > gameDS.DuelCardDeck[v.TarCard].ApOrig {
				gameDS.DuelCardDeck[v.TarCard].ApInst = gameDS.DuelCardDeck[v.TarCard].ApOrig
			}
			if gameDS.DuelCardDeck[v.TarCard].ApInst < 0 {
				gameDS.DuelCardDeck[v.TarCard].ApInst = 0
			}

			// dp change
			gameDS.DuelCardDeck[v.TarCard].DpInst += v.Dp
			if gameDS.DuelCardDeck[v.TarCard].DpInst > gameDS.DuelCardDeck[v.TarCard].DpOrig {
				gameDS.DuelCardDeck[v.TarCard].DpInst = gameDS.DuelCardDeck[v.TarCard].DpOrig
			}
			if gameDS.DuelCardDeck[v.TarCard].DpInst < 0 {
				gameDS.DuelCardDeck[v.TarCard].DpInst = 0
			}
		}
		go func() {
			this.BroadCast(&gameDS.RoomKey, &taskHandler, &bcMsg)
			wg.Done()
		}()
	}

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
