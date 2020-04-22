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
	// ======================================================================
	wg.Add(1)
	var val *int32
	var eff []*pb.EffectResult
	go func() {
		var err error
		val, eff, err = this.skillClient.SkillCalculateWrap(
			snapMod.AttackCard,
			snapMod.AttackTrigSkl,
			pb.EventCardType_ATTACK,
		)
		if err != nil {
			errch <- err
			return
		}
		wg.Done()
	}()
	wg.Wait()
	if errt := <-errch; errt != nil {
		fmt.Printf("attack-phase-handle::%v", errt)
		return false, errt
	}
	// ======================================================================
	// ADD instant node value
	snapMod.AttackVal = *val
	snapMod.IsProcessed = true
	effectMod.PendingEf = append(effectMod.PendingEf, eff...)
	var currentAtkKey int32
	if stateMod.CurrAttack == pb.PlayerSide_HOST {
		currentAtkKey = gameSet.HostCurrCardKey
	} else if stateMod.CurrAttack == pb.PlayerSide_DUELER {
		currentAtkKey = gameSet.DuelCurrCardKey
	}

	// dfEff := effectMod.PendingEf
	var addEff []*pb.EffectResult
	var FixEff []*pb.EffectResult
	wg.Add(2)
	go func() {
		addEff = pb.NodeFilter(effectMod.PendingEf, func(v *pb.EffectResult) bool {
			return (v.TriggerTime.EventPhase == pb.EventHookPhase_attack_card_drop_phase &&
				v.TriggerTime.HookType == pb.EventHookType_Proxy &&
				v.EfOption == pb.EffectOption_Status_Addition &&
				v.TarSide == stateMod.CurrAttack &&
				v.TarCard == currentAtkKey)
		})
		wg.Done()
	}()

	// Hard Set Value
	go func() {
		FixEff = pb.NodeFilter(effectMod.PendingEf, func(v *pb.EffectResult) bool {
			return (v.TriggerTime.EventPhase == pb.EventHookPhase_attack_card_drop_phase &&
				v.TriggerTime.HookType == pb.EventHookType_Proxy &&
				v.EfOption == pb.EffectOption_Status_FixValue &&
				v.TarSide == stateMod.CurrAttack &&
				v.TarCard == currentAtkKey)
		})
		wg.Done()
	}()
	wg.Wait()
	if len(FixEff) > 0 {
		tmpAtk := int32(0)
		for _, v := range FixEff {
			if v.Ap > tmpAtk {
				tmpAtk = v.Ap
			}
		}
		snapMod.AttackVal = tmpAtk
	} else {
		for _, v := range addEff {
			snapMod.AttackVal += v.Ap
		}
	}
	// ======================================================================

	errch = make(chan error)
	wg.Add(2)
	go func() {
		wkbox := this.searchAliveClient()
		var snapModkey = gameSet.RoomKey + snapMod.RdsKeyName()
		if _, err := (wkbox).SetPara(&snapModkey, snapMod); err != nil {
			log.Println(err)
			errch <- err
		}
		wg.Done()
	}()
	go func() {
		wkbox := this.searchAliveClient()
		var snapModkey = gameSet.RoomKey + snapMod.RdsKeyName()
		if _, err := (wkbox).SetPara(&snapModkey, stateMod); err != nil {
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
