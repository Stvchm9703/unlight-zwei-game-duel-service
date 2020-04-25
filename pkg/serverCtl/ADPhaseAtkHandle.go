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

func (this *ULZGameDuelServiceBackend) attackPhaseHandle(
	gameSet *pb.GameDataSet,
	snapMod *pb.ADPhaseSnapMod,
	stateMod *pb.PhaseSnapMod,
	effectMod *pb.EffectNodeSnapMod,
) (bool, error) {
	// do effect calculate
	// SECTION: skill-calculation
	wg := sync.WaitGroup{}
	errch := make(chan error)

	tmpAtk := int32(0)
	isDisable := false

	var currentAtkKey int32
	if stateMod.CurrAttack == pb.PlayerSide_HOST {
		currentAtkKey = gameSet.HostCurrCardKey
	} else if stateMod.CurrAttack == pb.PlayerSide_DUELER {
		currentAtkKey = gameSet.DuelCurrCardKey
	}
	// ======================================================================
	var val *int32
	var eff []*pb.EffectResult
	var err error

	disableSkill := pb.NodeFilter(effectMod.PendingEf, func(v *pb.EffectResult) bool {
		return (v.TriggerTime.EventPhase == pb.EventHookPhase_attack_card_drop_phase &&
			v.TriggerTime.HookType == pb.EventHookType_Proxy &&
			v.TarSide == stateMod.CurrAttack &&
			v.TarCard == currentAtkKey &&
			v.StatusId == 10)
	})

	if len(disableSkill) == 0 {
		val, eff, err = this.skillClient.SkillCalculateWrap(
			snapMod.AttackCard,
			snapMod.AttackTrigSkl,
			pb.EventCardType_ATTACK,
		)
		if err != nil {
			return false, err
		}
	} else {
		val = &snapMod.AttackVal
	}

	// ======================================================================
	// ADD instant node value
	snapMod.AttackVal = *val
	snapMod.IsProcessed = true
	effectMod.PendingEf = append(effectMod.PendingEf, eff...)
	// dfEff := effectMod.PendingEf
	addEff := pb.NodeFilter(effectMod.PendingEf, func(v *pb.EffectResult) bool {
		return (v.TriggerTime.EventPhase == pb.EventHookPhase_attack_card_drop_phase &&
			v.TriggerTime.HookType == pb.EventHookType_Proxy &&
			v.TarSide == stateMod.CurrAttack &&
			v.TarCard == currentAtkKey)
	})
	for _, v := range addEff {
		switch {
		case v.StatusId == 4, v.StatusId == 19, v.StatusId == 21:
			tmpAtk += v.Ap
		case v.StatusId == 5:
			tmpAtk -= v.Ap
		case v.StatusId == 9:
			isDisable = true
		case v.StatusId == 23:
			if v.Ap == 1 {
				tmpAtk++
			} else if v.Ap == 2 {
				tmpAtk += 2
			} else if v.Ap > 2 {
				tmpAtk += 5
			}
		}
	}

	if isDisable {
		snapMod.AttackVal = 0
	} else {
		snapMod.AttackVal += tmpAtk
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
		if _, err := (wkbox).SetPara(gameSet.RoomKey+snapMod.RdsKeyName(), stateMod); err != nil {
			log.Println(err)
			errch <- err
		}
		wg.Done()
	}()
	wg.Wait()
	if err := <-errch; err != nil {
		return false, err
	}
	// send ok message
	go this.BroadCast(&pb.GDBroadcastResp{
		RoomKey:      gameSet.RoomKey,
		Msg:          fmt.Sprintf("AD_PHASE:ATK_RESULT:%v", snapMod.AttackVal),
		Command:      pb.CastCmd_GET_ATK_PHASE_RESULT,
		Side:         snapMod.CurrAttacker,
		CurrentPhase: pb.EventHookPhase_attack_card_drop_phase,
		PhaseHook:    pb.EventHookType_Proxy,
	})
	return true, nil
}
