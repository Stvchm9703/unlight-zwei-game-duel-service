package serverCtl

import (
	pb "ULZGameDuelService/proto"
	"context"
	"log"
	"sort"

	"github.com/gogo/status"
	"github.com/jinzhu/copier"
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
//
func (this *ULZGameDuelServiceBackend) phaseTrigEf(gameDS *pb.GameDataSet, shiftPhase pb.EventHookPhase, shiftType pb.EventCardType) {
	var efResult pb.EffectNodeSnapMod
	var efResList []*pb.EffectResult
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

	gameDSTmp := pb.GameDataSet{}
	copier.Copy(&gameDSTmp, gameDS)

	// return be4 run loop
	if len(DirectDmg) == 0 {
		return
	}

	for _, v := range DirectDmg {
		if v.TarSide == pb.PlayerSide_HOST {
			gameDSTmp.HostCardDeck[v.TarCard].HpInst += v.Hp
			gameDSTmp.HostCardDeck[v.TarCard].ApInst += v.Ap
			gameDSTmp.HostCardDeck[v.TarCard].DpInst += v.Dp
		} else {
			gameDSTmp.DuelCardDeck[v.TarCard].HpInst += v.Hp
			gameDSTmp.DuelCardDeck[v.TarCard].ApInst += v.Ap
			gameDSTmp.DuelCardDeck[v.TarCard].DpInst += v.Dp
		}
	}

	//Hp handle:
	// if the value higher than orig, reset asd orig
	// if the value lower than 0, set as 0 ,
	// 		if it is current => raise dead-flag
	//

	// for k,v:= {

	// }
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
