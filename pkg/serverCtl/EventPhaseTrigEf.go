package serverCtl

import (
	pb "ULZGameDuelService/proto"
	"fmt"
	"sort"
	"sync"
)

// phaseTrigEf : general phase trigger effect-node
// it only handle Instance_Change / direct-dmg
// NOTE not available for atk/def, move phase calculation
func (this *ULZGameDuelServiceBackend) phaseTrigEf(
	gameDS *pb.GameDataSet,
	phaseMod *pb.PhaseSnapMod,
	effectMod *pb.EffectNodeSnapMod,
) {
	efResList := effectMod.PendingEf
	if len(efResList) == 0 {
		return
	}

	tarEf := pb.NodeFilter(efResList, func(v *pb.EffectResult) bool {
		return (v.TriggerTime.EventPhase == phaseMod.EventPhase) &&
			(v.TriggerTime.HookType == phaseMod.HookType)
	})
	if len(tarEf) == 0 {
		return
	}

	sort.Slice(tarEf, func(i, j int) bool {
		return tarEf[i].TriggerTime.SubCount < tarEf[i].TriggerTime.SubCount
	})

	DirectDmg := pb.NodeFilter(tarEf, func(v *pb.EffectResult) bool {
		return (v.EfOption == pb.EffectOption_Instance_Change)
	})

	HardDmg := pb.NodeFilter(tarEf, func(v *pb.EffectResult) bool {
		return (v.EfOption == pb.EffectOption_Hard_Instance_Change)
	})

	// return be4 run loop
	if len(DirectDmg) == 0 {
		return
	}

	wg := sync.WaitGroup{}
	var waitForClean []*pb.EffectResult

	// Status release
	// 	 Damage part first
	wg.Add(3)
	go func() {
		bcMsg := pb.GDBroadcastResp{
			RoomKey:      gameDS.RoomKey,
			Msg:          fmt.Sprintf("Damage from effect to player"),
			Command:      pb.CastCmd_INSTANCE_STATUS_CHANGE,
			CurrentPhase: gameDS.EventPhase,
			PhaseHook:    gameDS.HookType,
			EffectTrig:   append(DirectDmg, HardDmg...),
		}
		this.BroadCast(&bcMsg)
		wg.Done()
	}()
	go func() {
		for _, v := range DirectDmg {
			if v.TarSide == pb.PlayerSide_HOST && v.TarCard == gameDS.HostCurrCardKey {
				// hp change
				pointCalcute(&gameDS.HostCardDeck[v.TarCard].HpInst, &gameDS.HostCardDeck[v.TarCard].HpOrig, &v.Hp)
				// ap change
				pointCalcute(&gameDS.HostCardDeck[v.TarCard].ApInst, &gameDS.HostCardDeck[v.TarCard].ApOrig, &v.Ap)
				// dp change
				pointCalcute(&gameDS.HostCardDeck[v.TarCard].DpInst, &gameDS.HostCardDeck[v.TarCard].DpOrig, &v.Dp)

			} else if v.TarCard == gameDS.DuelCurrCardKey {
				// hp change
				pointCalcute(&gameDS.DuelCardDeck[v.TarCard].HpInst, &gameDS.DuelCardDeck[v.TarCard].HpOrig, &v.Hp)
				// ap change
				pointCalcute(&gameDS.DuelCardDeck[v.TarCard].ApInst, &gameDS.DuelCardDeck[v.TarCard].ApOrig, &v.Ap)
				// dp change
				pointCalcute(&gameDS.DuelCardDeck[v.TarCard].DpInst, &gameDS.DuelCardDeck[v.TarCard].DpOrig, &v.Dp)
			}
			fmt.Println(v)
		}
		waitForClean = append(waitForClean, DirectDmg...)
		wg.Done()
	}()
	go func() {
		for _, v := range HardDmg {
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
		waitForClean = append(waitForClean, HardDmg...)
		wg.Done()
	}()
	wg.Wait()

	wg.Add(2)
	// Status_FixValue :
	// Hard life
	var hostFixFin, duelFixFin pb.EffectResult
	go func() {
		hostFixEf := pb.NodeFilter(tarEf, func(v *pb.EffectResult) bool {
			return (v.EfOption == pb.EffectOption_Status_FixValue) &&
				(v.TarCard == gameDS.HostCurrCardKey) &&
				(v.TarSide == pb.PlayerSide_HOST)
		})
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
			waitForClean = append(waitForClean, hostFixEf...)
		}
		wg.Done()
	}()
	go func() {
		duelFixEf := pb.NodeFilter(tarEf, func(v *pb.EffectResult) bool {
			return (v.EfOption == pb.EffectOption_Status_FixValue) &&
				(v.TarCard == gameDS.DuelCurrCardKey) &&
				(v.TarSide == pb.PlayerSide_DUELER)
		})
		if len(duelFixEf) > 0 {
			for _, v := range duelFixEf {
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
			waitForClean = append(waitForClean, duelFixEf...)
		}
		wg.Done()
	}()
	wg.Wait()

	wg.Add(2)
	go func() {
		bcMsg := pb.GDBroadcastResp{
			RoomKey:      gameDS.RoomKey,
			Msg:          fmt.Sprintf("Fix Effect from effect to player"),
			Command:      pb.CastCmd_INSTANCE_STATUS_CHANGE,
			CurrentPhase: gameDS.EventPhase,
			PhaseHook:    gameDS.HookType,
			EffectTrig: []*pb.EffectResult{
				&hostFixFin, &duelFixFin,
			},
		}
		this.BroadCast(&bcMsg)
	}()
	// go wkbox.SetPara(&gameDS.RoomKey, gameDS)
	// clean the effect by one cd
	fmt.Printf("%#v", waitForClean)
	pb.CleanAfterExec(efResList, waitForClean)

	go func() {
		this.BroadCast(&pb.GDBroadcastResp{
			RoomKey:      gameDS.RoomKey,
			Msg:          fmt.Sprintf("EndOfEventPhase,WaitForConfirm"),
			Command:      pb.CastCmd_GET_EFFECT_RESULT,
			CurrentPhase: gameDS.EventPhase,
			PhaseHook:    gameDS.HookType,
		})
	}()
	fmt.Println(effectMod)
	// save data
	wkbox := this.searchAliveClient()
	wg.Add(2)
	go func() {
		wkbox.SetPara(gameDS.RoomKey, gameDS)
		wg.Done()
	}()
	go func() {
		wkbox.SetPara(gameDS.RoomKey+effectMod.RdsKeyName(), effectMod)
		wg.Done()
	}()
	wg.Wait()
	wkbox.Preserve(false)
	return
}
