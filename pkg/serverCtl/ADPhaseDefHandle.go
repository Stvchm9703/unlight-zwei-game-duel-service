package serverCtl

import (
	pb "ULZGameDuelService/proto"
	// svcpb "ULZGameDuelService/pkg/scriptRunner/proto"
	"fmt"
	"log"
	"sync"
	// Static files
	// _ "ULZGameDuelService/statik"
)

func (this *ULZGameDuelServiceBackend) defencePhaseHandle(
	gameSet *pb.GameDataSet,
	snapMod *pb.ADPhaseSnapMod,
	stateMod *pb.PhaseSnapMod,
	effectMod *pb.EffectNodeSnapMod,
) (bool, error) {
	// do effect calculate
	// SECTION: skill-calculation
	wg := sync.WaitGroup{}
	errch := make(chan error)

	tmpDef := int32(0)
	isDisable := false

	var currentDefKey int32
	var side pb.PlayerSide
	if snapMod.CurrAttacker == pb.PlayerSide_HOST {
		currentDefKey = gameSet.DuelCurrCardKey
		side = pb.PlayerSide_DUELER
	} else if snapMod.CurrAttacker == pb.PlayerSide_DUELER {
		currentDefKey = gameSet.HostCurrCardKey
		side = pb.PlayerSide_HOST
	}
	// ======================================================================
	wg.Add(1)
	var val *int32
	var eff []*pb.EffectResult
	var err error

	disableSkill := pb.NodeFilter(effectMod.PendingEf, func(v *pb.EffectResult) bool {
		return (v.TriggerTime.EventPhase == pb.EventHookPhase_attack_card_drop_phase &&
			v.TriggerTime.HookType == pb.EventHookType_Proxy &&
			v.TarSide == stateMod.CurrAttack &&
			v.TarCard == currentDefKey &&
			v.StatusId == 10)
	})

	if len(disableSkill) == 0 {
		val, eff, err = this.skillClient.SkillCalculateWrap(
			snapMod.DefenceCard,
			snapMod.DefenceTrigSkl,
			pb.EventCardType_DEFENCE,
		)
		if err != nil {
			fmt.Printf("defence-phase-handle::%v", err)
			return false, err
		}
	} else {
		val = &snapMod.DefenceVal
	}

	// ======================================================================
	// do update
	snapMod.DefenceVal = *val
	snapMod.IsProcessed = true
	effectMod.PendingEf = append(effectMod.PendingEf, eff...)

	addEff := pb.NodeFilter(effectMod.PendingEf, func(v *pb.EffectResult) bool {
		return (v.TriggerTime.EventPhase == pb.EventHookPhase_attack_card_drop_phase &&
			v.TriggerTime.HookType == pb.EventHookType_Proxy &&
			v.EfOption == pb.EffectOption_Status_Addition &&
			v.TarSide != stateMod.CurrAttack &&
			v.TarCard == currentDefKey)
	})

	for _, v := range addEff {
		switch {
		case v.StatusId == 6, v.StatusId == 19, v.StatusId == 21:
			tmpDef += v.Dp
		case v.StatusId == 7:
			tmpDef -= v.Dp
		case v.StatusId == 9:
			isDisable = true
		case v.StatusId == 23:
			if v.Dp == 1 {
				tmpDef++
			} else if v.Dp == 2 {
				tmpDef += 2
			} else if v.Dp > 2 {
				tmpDef += 5
			}
		}
	}
	if isDisable {
		snapMod.DefenceVal = 0
	} else {
		snapMod.DefenceVal += tmpDef
	}
	// ======================================================================
	errch = make(chan error)
	wg.Add(2)
	go func() {
		wkbox := this.searchAliveClient()
		if _, err := (wkbox).SetPara(gameSet.RoomKey+snapMod.RdsKeyName(), snapMod); err != nil {
			log.Println(err)
			errch <- err
		}
		wg.Done()
	}()

	go func() {
		wkbox := this.searchAliveClient()
		if _, err := (wkbox).SetPara(gameSet.RoomKey+effectMod.RdsKeyName(), *effectMod); err != nil {
			log.Println(err)
			errch <- err
		}
		wg.Done()
	}()
	wg.Wait()
	if err := <-errch; err != nil {
		return false, err
	}
	// ======================================================================

	// send ok message

	go this.BroadCast(&pb.GDBroadcastResp{
		RoomKey:      gameSet.RoomKey,
		Msg:          fmt.Sprintf("AD_PHASE:DEF_RESULT:%v", snapMod.DefenceVal),
		Command:      pb.CastCmd_GET_DEF_PHASE_RESULT,
		Side:         side,
		CurrentPhase: pb.EventHookPhase_defence_card_drop_phase,
		PhaseHook:    pb.EventHookType_Proxy,
	})
	return true, nil
}
