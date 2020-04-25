package serverCtl

import (
	skc "ULZGameDuelService/pkg/scriptRunner"
	pb "ULZGameDuelService/proto"
	"fmt"
	"sync"
)

/** determineBattlePointPhase: event handle in determine-battle-point-phase
 * 		1. move the event card in evtcrdlst
 */
func (this *ULZGameDuelServiceBackend) determineBattlePointPhase(
	gameSet *pb.GameDataSet,
	stateMod *pb.PhaseSnapMod,
	effectMod *pb.EffectNodeSnapMod,
	adStoreMod *pb.ADPhaseSnapMod,
) {
	var hostECInStore, duelECInStore pb.EventCardListSet
	wg := sync.WaitGroup{}
	wg.Add(2)
	errch := make(chan error)
	go func() {
		wkbox := this.searchAliveClient()
		key := gameSet.RoomKey + hostECInStore.RdsKeyName(pb.PlayerSide_HOST)
		if _, err := (wkbox).GetPara(key, &hostECInStore); err != nil {
			errch <- err
		}
		hostECInStore.ECListMoveTo(pb.EventCardPos_OUTSIDE, pb.EventCardPos_DESTROY)
		if _, err := (wkbox).SetPara(key, hostECInStore); err != nil {
			errch <- err
		}
		wg.Done()
	}()
	go func() {
		wkbox := this.searchAliveClient()
		key := gameSet.RoomKey + duelECInStore.RdsKeyName(pb.PlayerSide_DUELER)
		if _, err := (wkbox).GetPara(key, &duelECInStore); err != nil {
			errch <- err
		}
		duelECInStore.ECListMoveTo(pb.EventCardPos_OUTSIDE, pb.EventCardPos_DESTROY)
		if _, err := (wkbox).SetPara(key, duelECInStore); err != nil {
			errch <- err
		}
		wg.Done()
	}()
	wg.Wait()

	this.BroadCast(&pb.GDBroadcastResp{
		RoomKey:      gameSet.RoomKey,
		Msg:          fmt.Sprintf("DetermineBattlePointPhase"),
		Command:      pb.CastCmd_GET_GAMESET_RESULT,
		CurrentPhase: gameSet.EventPhase,
		PhaseHook:    gameSet.HookType,
	})
}

/**  battlePhaseHandle : event handle in battle-phase
 * 		1. dice the result set
 */
func (this *ULZGameDuelServiceBackend) battlePhaseHandle(
	gameSet *pb.GameDataSet,
	stateMod *pb.PhaseSnapMod,
	effectMod *pb.EffectNodeSnapMod,
	adStoreMod *pb.ADPhaseSnapMod,
) *pb.GDADDiceResult {
	var atkRes, defRes skc.DiceResult
	wg := sync.WaitGroup{}
	wg.Add(2)
	errch := make(chan error)
	go func() {
		var err error
		atkRes, err = this.skillClient.DiceCalculateWrap(
			adStoreMod.AttackVal, 1, effectMod.PendingEf)
		if err != nil {
			errch <- err
		}
		wg.Done()
	}()

	go func() {
		var err error
		defRes, err = this.skillClient.DiceCalculateWrap(
			adStoreMod.DefenceVal, 1, effectMod.PendingEf)
		if err != nil {
			errch <- err
		}
		wg.Done()
	}()

	defSide := adStoreMod.CurrAttacker
	if adStoreMod.CurrAttacker == pb.PlayerSide_HOST {
		defSide = pb.PlayerSide_DUELER
	} else {
		defSide = pb.PlayerSide_HOST
	}
	wg.Wait()

	adDiceResult := pb.GDADDiceResult{
		RoomKey:      gameSet.RoomKey,
		Turns:        stateMod.Turns,
		CurrentPhase: stateMod.EventPhase,
		PhaseAb:      adStoreMod.FirstAttack,
		AtkSide:      adStoreMod.CurrAttacker,
		AtkPoint:     atkRes.ToTotal(),
		AtkSkillId:   pb.GetSkillId(adStoreMod.AttackTrigSkl),
		DefSide:      defSide,
		DefPoint:     defRes.ToTotal(),
		DefSkillId:   pb.GetSkillId(adStoreMod.DefenceTrigSkl),
	}

	wkbox1 := this.searchAliveClient()
	wkbox1.SetPara(gameSet.RoomKey+adDiceResult.RdsKeyName(), adDiceResult)
	wkbox1.Preserve(false)
	go this.BroadCast(&pb.GDBroadcastResp{
		RoomKey:      gameSet.RoomKey,
		Msg:          fmt.Sprintf("DiceResult:Atk:%s;Def:%s", atkRes.ToString(), defRes.ToString()),
		Command:      pb.CastCmd_GET_GAMESET_RESULT,
		CurrentPhase: pb.EventHookPhase_battle_result_phase,
		PhaseHook:    pb.EventHookType_Proxy,
	})
	return &adDiceResult
}

func (this *ULZGameDuelServiceBackend) damagePhaseHandle(
	gameSet *pb.GameDataSet,
	stateMod *pb.PhaseSnapMod,
	effectMod *pb.EffectNodeSnapMod,
	gdDice *pb.GDADDiceResult,
) {

	dmg := gdDice.AtkPoint - gdDice.DefPoint
	if dmg < 0 {
		dmg = 0
	}

	wg := sync.WaitGroup{}
	wg.Add(2)
	var inAtkTurn, inDefTurn []*pb.EffectResult
	var AtkCard, DefCard int32
	if gdDice.AtkSide == pb.PlayerSide_HOST {
		AtkCard = gameSet.HostCurrCardKey
		DefCard = gameSet.DuelCurrCardKey
	} else {
		AtkCard = gameSet.DuelCurrCardKey
		DefCard = gameSet.HostCurrCardKey
	}
	go func() {
		// Attacker
		inAtkTurn = pb.NodeFilter(effectMod.PendingEf, func(v *pb.EffectResult) bool {
			return v.TriggerTime.EventPhase == pb.EventHookPhase_damage_phase &&
				v.TriggerTime.HookType == pb.EventHookType_Proxy &&
				v.TarSide == gdDice.AtkSide &&
				v.TarCard == AtkCard
		})
		wg.Done()
	}()
	go func() {
		// Defencer
		inDefTurn = pb.NodeFilter(effectMod.PendingEf, func(v *pb.EffectResult) bool {
			return v.TriggerTime.EventPhase == pb.EventHookPhase_damage_phase &&
				v.TriggerTime.HookType == pb.EventHookType_Proxy &&
				v.TarSide != gameSet.CurrPhase &&
				v.TarCard == DefCard
		})

		wg.Done()
	}()

	// status calc

	for _, v := range inAtkTurn {
		switch {
		case v.StatusId == 8:
			dmg = dmg * 2
			fallthrough
		case v.StatusId == 13:
			dmg = dmg/2 + dmg%2
		}
	}
	for _, v := range inDefTurn {
		switch {
		case v.StatusId == 12, v.StatusId == 24:
			dmg = 0
		case v.StatusId == 8:
			dmg = dmg * 2
			fallthrough
		case v.StatusId == 13:
			dmg = dmg/2 + dmg%2
		}
	}

	if gdDice.AtkSide == pb.PlayerSide_HOST {
		gameSet.DuelCardDeck[int(DefCard)].HpInst -= dmg
	} else {
		gameSet.HostCardDeck[int(DefCard)].HpInst -= dmg
	}

	this.BroadCast(&pb.GDBroadcastResp{
		RoomKey:      gameSet.RoomKey,
		Msg:          fmt.Sprintf("AD_PHASE:DMG:%v", dmg),
		Command:      pb.CastCmd_GET_GAMESET_RESULT,
		Side:         gdDice.DefSide,
		CurrentPhase: pb.EventHookPhase_damage_phase,
		PhaseHook:    pb.EventHookType_Proxy,
	})
}
