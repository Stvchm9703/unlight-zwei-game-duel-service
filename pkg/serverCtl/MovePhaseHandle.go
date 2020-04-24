package serverCtl

import (
	pb "ULZGameDuelService/proto"
	"fmt"
	"log"
	"sync"
	// cm "ULZGameDuelService/pkg/common"
	// "context"
	// "log"
	// "sync"
	// "time"
	// "github.com/gogo/status"
	// "google.golang.org/grpc/codes"
)

/**
 * movePhaseHandle :
 * 		for handle the determine_move_phase:proxy logic
 */
func (this *ULZGameDuelServiceBackend) determineMovePhaseHandle(
	gameSet *pb.GameDataSet,
	stateMod *pb.PhaseSnapMod,
	effectMod *pb.EffectNodeSnapMod,
	moveMod *pb.MovePhaseSnapMod,
) {
	// go to request the move result
	// SECTION: skill-calculation

	domain := pb.PlayerSide_HOST
	var tmpval int32
	distance := gameSet.Range
	switch {
	case isMoveFowBack(moveMod.HostOpt) && !isMoveFowBack(moveMod.DuelOpt):
		domain = pb.PlayerSide_HOST
	case !isMoveFowBack(moveMod.HostOpt) && !isMoveFowBack(moveMod.DuelOpt):
		domain = pb.PlayerSide_HOST
	case !isMoveFowBack(moveMod.HostOpt) && isMoveFowBack(moveMod.DuelOpt):
		domain = pb.PlayerSide_DUELER
	case isMoveFowBack(moveMod.HostOpt) && isMoveFowBack(moveMod.DuelOpt):
		_effectCheckForMovePhase(gameSet, effectMod, moveMod)
		fmt.Printf("movMod:%#v", moveMod)
		if moveMod.HostVal > moveMod.DuelVal {
			domain = pb.PlayerSide_HOST
			tmpval = moveMod.HostVal - moveMod.DuelVal
		} else if moveMod.DuelVal > moveMod.HostVal {
			domain = pb.PlayerSide_DUELER
			tmpval = moveMod.HostVal - moveMod.DuelVal
		} else {
			domain = pb.PlayerSide_HOST
		}
	}
	switch {
	case domain == pb.PlayerSide_HOST && moveMod.HostOpt == pb.MovePhaseOpt_FORWARD:
		fallthrough
	case domain == pb.PlayerSide_DUELER && moveMod.DuelOpt == pb.MovePhaseOpt_FORWARD:
		switch {
		case tmpval > 2:
			fallthrough
		case tmpval == 1 && (distance == pb.RangeType_MIDDLE || distance == pb.RangeType_SHORT):
			distance = pb.RangeType_SHORT
		case tmpval == 1 && distance == pb.RangeType_LONG:
			distance = pb.RangeType_MIDDLE
		}
		break

	case domain == pb.PlayerSide_HOST && moveMod.HostOpt == pb.MovePhaseOpt_BACKWARD:
		fallthrough
	case domain == pb.PlayerSide_DUELER && moveMod.DuelOpt == pb.MovePhaseOpt_BACKWARD:
		switch {
		case tmpval > 2:
			fallthrough
		case tmpval == 1 && (distance == pb.RangeType_MIDDLE || distance == pb.RangeType_LONG):
			distance = pb.RangeType_LONG
		case tmpval == 1 && distance == pb.RangeType_SHORT:
			distance = pb.RangeType_MIDDLE
		}
	}

	gameSet.Range = distance
	gameSet.FirstAttack = domain
	gameSet.CurrPhase = domain

	if moveMod.HostOpt == pb.MovePhaseOpt_STAY {
		gameSet.HostCardDeck[gameSet.HostCurrCardKey].HpInst++
	}
	if moveMod.DuelOpt == pb.MovePhaseOpt_STAY {
		gameSet.DuelCardDeck[gameSet.DuelCurrCardKey].HpInst++
	}

	fmt.Printf("Range:%s, FirstAttack: %s", distance.String(), domain.String())

	//
	// this.executeEffectNode(gameSet, stateMod, effectMod)
	cleanEffectResult(
		pb.EventHookPhase_determine_move_phase,
		pb.EventHookType_Proxy,
		effectMod,
	)

	// this.proxyHandle(gameSet *pb.GameDataSet, phaseMod *pb.PhaseSnapMod, effectMod *pb.EffectNodeSnapMod, snapMod ...interface{})

	go func() {
		wkbox := this.searchAliveClient()
		if _, err := (wkbox).SetPara(gameSet.RoomKey, gameSet); err != nil {
			log.Println(err)
		}
		wkbox.Preserve(false)
	}()
	// real act
	go func() {
		wkbox := this.searchAliveClient()
		if _, err := (wkbox).SetPara(gameSet.RoomKey+moveMod.RdsKeyName(), moveMod); err != nil {
			log.Println(err)
		}
		wkbox.Preserve(false)
	}()

	go func() {
		mbox := this.searchAliveClient()
		if _, err := (mbox).SetPara(gameSet.RoomKey+effectMod.RdsKeyName(), effectMod); err != nil {
			log.Println(err)
		}
		mbox.Preserve(false)
	}()
	// send ok message

}
func _effectCheckForMovePhase(
	gmSet *pb.GameDataSet,
	eflist *pb.EffectNodeSnapMod,
	snapMove *pb.MovePhaseSnapMod,
) {
	fmt.Println("Start Filter effect")
	/** TODO : concat the Effect node that for move-phase calculating
	 */
	// =======================================================================
	twg := sync.WaitGroup{}
	twg.Add(2)
	var hostCurrentEF, duelCurrentEF pb.EffectResult
	var hostEF, duelEF []*pb.EffectResult
	go func() {
		hostEF = pb.NodeFilter(eflist.PendingEf, func(v *pb.EffectResult) bool {
			return v.TriggerTime.EventPhase == pb.EventHookPhase_determine_move_phase &&
				v.TriggerTime.HookType == pb.EventHookType_Proxy &&
				v.TarSide == pb.PlayerSide_HOST &&
				v.TarCard == gmSet.HostCurrCardKey
		})
		for _, v := range hostEF {
			if v.EfOption == pb.EffectOption_Hard_Instance_Change {
				hostCurrentEF.EfOption = pb.EffectOption_Hard_Instance_Change
			}
			if v.DisableMove {
				hostCurrentEF.DisableMove = v.DisableMove
			}
			if v.DisableChange {
				hostCurrentEF.DisableChange = v.DisableChange
			}
			if v.BindingFunc != "" {
				hostCurrentEF.BindingFunc += ";" + v.BindingFunc
			}
		}
		twg.Done()
	}()

	go func() {
		duelEF = pb.NodeFilter(eflist.PendingEf, func(v *pb.EffectResult) bool {
			return v.TriggerTime.EventPhase == pb.EventHookPhase_determine_move_phase &&
				v.TriggerTime.HookType == pb.EventHookType_Proxy &&
				v.TarSide == pb.PlayerSide_DUELER &&
				v.TarCard == gmSet.HostCurrCardKey
		})
		for _, v := range duelEF {
			if v.EfOption == pb.EffectOption_Hard_Instance_Change {
				duelCurrentEF.EfOption = pb.EffectOption_Hard_Instance_Change
				duelCurrentEF.Mp = v.Mp
			}
			if v.DisableMove {
				duelCurrentEF.DisableMove = v.DisableMove
			}
			if v.DisableChange {
				duelCurrentEF.DisableChange = v.DisableChange
			}
			if v.BindingFunc != "" {
				duelCurrentEF.BindingFunc += ";" + v.BindingFunc
			}
		}

		twg.Done()
	}()
	twg.Wait()
	// =======================================================================
	if hostCurrentEF.DisableMove {
		snapMove.HostVal = 0
	} else if hostCurrentEF.EfOption == pb.EffectOption_Hard_Instance_Change {
		snapMove.HostVal = hostCurrentEF.Mp
	}

	if duelCurrentEF.DisableMove {
		snapMove.DuelVal = 0
	} else if duelCurrentEF.EfOption == pb.EffectOption_Hard_Instance_Change {
		snapMove.DuelVal = duelCurrentEF.Mp
	}

	// =======================================================================
}

func isMoveFowBack(opt pb.MovePhaseOpt) bool {
	return opt == pb.MovePhaseOpt_BACKWARD || opt == pb.MovePhaseOpt_FORWARD
}
