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
	// ======================================================================
	wg.Add(1)
	var val *int32
	var eff []*pb.EffectResult
	go func() {
		var err error
		val, eff, err = this.skillClient.SkillCalculateWrap(
			snapMod.DefenceCard,
			snapMod.DefenceTrigSkl,
			pb.EventCardType_DEFENCE,
		)
		if err != nil {
			errch <- err
			return
		}
		wg.Done()
	}()
	wg.Wait()
	if errt := <-errch; errt != nil {
		fmt.Printf("defence-phase-handle::%v", errt)
		return false, errt
	}
	// ======================================================================
	// do update
	snapMod.DefenceVal = *val
	snapMod.IsProcessed = true
	effectMod.PendingEf = append(effectMod.PendingEf, eff...)
	side := snapMod.CurrAttacker
	var currentDefKey int32
	if side == pb.PlayerSide_HOST {
		side = pb.PlayerSide_DUELER
		currentDefKey = gameSet.DuelCurrCardKey
	} else {
		side = pb.PlayerSide_HOST
		currentDefKey = gameSet.HostCurrCardKey
	}
	var addEff []*pb.EffectResult
	var FixEff []*pb.EffectResult
	wg.Add(2)
	go func() {
		addEff = pb.NodeFilter(effectMod.PendingEf, func(v *pb.EffectResult) bool {
			return (v.TriggerTime.EventPhase == pb.EventHookPhase_attack_card_drop_phase &&
				v.TriggerTime.HookType == pb.EventHookType_Proxy &&
				v.EfOption == pb.EffectOption_Status_Addition &&
				v.TarSide == side &&
				v.TarCard == currentDefKey)
		})
		wg.Done()
	}()

	// Hard Set Value
	go func() {
		FixEff = pb.NodeFilter(effectMod.PendingEf, func(v *pb.EffectResult) bool {
			return (v.TriggerTime.EventPhase == pb.EventHookPhase_attack_card_drop_phase &&
				v.TriggerTime.HookType == pb.EventHookType_Proxy &&
				v.EfOption == pb.EffectOption_Status_FixValue &&
				v.TarSide == side &&
				v.TarCard == currentDefKey)
		})
		wg.Done()
	}()
	wg.Wait()
	if len(FixEff) > 0 {
		tmpAtk := int32(0)
		for _, v := range FixEff {
			if v.Dp > tmpAtk {
				tmpAtk = v.Dp
			}
		}
		snapMod.DefenceVal = tmpAtk
	} else {
		for _, v := range addEff {
			snapMod.DefenceVal += v.Dp
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
		var snapModkey = gameSet.RoomKey + effectMod.RdsKeyName()
		if _, err := (wkbox).SetPara(&snapModkey, *effectMod); err != nil {
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
