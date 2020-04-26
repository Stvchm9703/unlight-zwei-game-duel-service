package serverCtl

import (
	pb "ULZGameDuelService/proto"
	"encoding/json"
	"fmt"
	"log"
	"strings"
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
	twg := sync.WaitGroup{}
	twg.Add(2)
	var hostCurrentEF, duelCurrentEF pb.EffectResult
	var hostEF, duelEF []*pb.EffectResult
	var hostNoHeal, hostFocHeal, hostPois1, hostPosi2 bool
	var duelNoHeal, duelFocHeal, duelPois1, duelPosi2 bool
	var hostEfFlag, duelEfFlag []string
	go func() {
		hostEF = pb.NodeFilter(effectMod.PendingEf, func(v *pb.EffectResult) bool {
			return v.TriggerTime.EventPhase == pb.EventHookPhase_determine_move_phase &&
				v.TriggerTime.HookType == pb.EventHookType_Proxy &&
				v.TarSide == pb.PlayerSide_HOST &&
				v.TarCard == gameSet.HostCurrCardKey
		})
		tmpFuncMap := make(map[string]string)
		for _, v := range hostEF {
			switch {
			case v.DisableChange:
				hostCurrentEF.DisableChange = v.DisableChange
				fallthrough
			case v.DisableMove:
				hostCurrentEF.DisableMove = v.DisableMove
				fallthrough
			case v.BindingFunc != "":
				tmpFuncMap[fmt.Sprint(v.SkillId)] = v.BindingFunc
				fallthrough

			case v.StatusId == 1:
				hostPois1 = true
				hostEfFlag = append(hostEfFlag, "poison")
			case v.StatusId == 2:
				hostPosi2 = true
				hostEfFlag = append(hostEfFlag, "poison2")
			case v.StatusId == 16:
				hostFocHeal = true
				hostEfFlag = append(hostEfFlag, "regene")
			case v.StatusId == 27:
				hostNoHeal = true
				hostEfFlag = append(hostEfFlag, "dark")
			}
			rs, _ := json.Marshal(tmpFuncMap)
			v.BindingFunc = string(rs)
		}
		twg.Done()
	}()

	go func() {
		duelEF = pb.NodeFilter(effectMod.PendingEf, func(v *pb.EffectResult) bool {
			return v.TriggerTime.EventPhase == pb.EventHookPhase_determine_move_phase &&
				v.TriggerTime.HookType == pb.EventHookType_Proxy &&
				v.TarSide == pb.PlayerSide_DUELER &&
				v.TarCard == gameSet.DuelCurrCardKey
		})
		tmpFuncMap := make(map[string]string)
		for _, v := range duelEF {
			switch {
			case v.DisableChange:
				hostCurrentEF.DisableChange = v.DisableChange
				fallthrough
			case v.DisableMove:
				hostCurrentEF.DisableMove = v.DisableMove
				fallthrough
			case v.BindingFunc != "":
				tmpFuncMap[fmt.Sprint(v.SkillId)] = v.BindingFunc
				fallthrough
			case v.StatusId == 1:
				duelPois1 = true
				duelEfFlag = append(duelEfFlag, "poison")

			case v.StatusId == 2:
				duelPosi2 = true
				duelEfFlag = append(duelEfFlag, "poison2")

			case v.StatusId == 16:
				duelFocHeal = true
				duelEfFlag = append(duelEfFlag, "regene")

			case v.StatusId == 27:
				duelNoHeal = true
				duelEfFlag = append(duelEfFlag, "dark")

			}
			rs, _ := json.Marshal(tmpFuncMap)
			v.BindingFunc = string(rs)
		}
		twg.Done()
	}()
	twg.Wait()
	// =======================================================================
	if hostCurrentEF.DisableMove {
		moveMod.HostVal = 0
	} else if hostCurrentEF.EfOption == pb.EffectOption_Hard_Instance_Change {
		moveMod.HostVal = hostCurrentEF.Mp
	}

	if duelCurrentEF.DisableMove {
		moveMod.DuelVal = 0
	} else if duelCurrentEF.EfOption == pb.EffectOption_Hard_Instance_Change {
		moveMod.DuelVal = duelCurrentEF.Mp
	}

	// =======================================================================

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

	/**
	 * proxy-recover => status-posion => status-regene
	 */

	//  proxy-recover
	if moveMod.HostOpt == pb.MovePhaseOpt_STAY && !hostNoHeal {
		gameSet.HostCardDeck[gameSet.HostCurrCardKey].HpInst++
	}

	if moveMod.DuelOpt == pb.MovePhaseOpt_STAY && !duelNoHeal {
		gameSet.DuelCardDeck[gameSet.DuelCurrCardKey].HpInst++
	}

	// status-posion
	if hostPois1 {
		gameSet.HostCardDeck[gameSet.HostCurrCardKey].HpInst--
	}
	if duelPois1 {
		gameSet.DuelCardDeck[gameSet.DuelCurrCardKey].HpInst--
	}

	// status-posion2
	if hostPosi2 {
		gameSet.HostCardDeck[gameSet.HostCurrCardKey].HpInst -= 2
	}
	if duelPosi2 {
		gameSet.DuelCardDeck[gameSet.DuelCurrCardKey].HpInst -= 2
	}

	// status-regene
	if hostFocHeal {
		gameSet.HostCardDeck[gameSet.HostCurrCardKey].HpInst++
	}
	if duelFocHeal {
		gameSet.DuelCardDeck[gameSet.DuelCurrCardKey].HpInst++
	}

	fmt.Printf("Range:%s;FirstAttack:%s\n", distance.String(), domain.String())
	fmt.Printf("host:%s", strings.Join(hostEfFlag, ","))
	fmt.Printf("duel:%s", strings.Join(duelEfFlag, ","))
	//
	// this.executeEffectNode(gameSet, stateMod, effectMod)

	// cleanEffectResult(
	// 	pb.EventHookPhase_determine_move_phase,
	// 	pb.EventHookType_Proxy,
	// 	effectMod,
	// )

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

	this.BroadCast(&pb.GDBroadcastResp{
		RoomKey: gameSet.RoomKey,
		Msg: fmt.Sprintf(
			"Range:%s;FA:%s;MovEf:Self:%s;Duel:%s;",
			distance.String(), domain.String(),
			strings.Join(hostEfFlag, ","),
			strings.Join(duelEfFlag, ",")),
		Command:      pb.CastCmd_INSTANCE_STATUS_CHANGE,
		CurrentPhase: pb.EventHookPhase_determine_move_phase,
		PhaseHook:    pb.EventHookType_Proxy,
	})

}

func isMoveFowBack(opt pb.MovePhaseOpt) bool {
	return opt == pb.MovePhaseOpt_BACKWARD || opt == pb.MovePhaseOpt_FORWARD
}

func (this *ULZGameDuelServiceBackend) finishMovePhase(
	gameSet *pb.GameDataSet,
	phaseMod *pb.PhaseSnapMod,
	effectMod *pb.EffectNodeSnapMod,
	snapMovMod *pb.MovePhaseSnapMod,
) {
	hostHp, duelHp := int32(0), int32(0)
	if snapMovMod.HostOpt == pb.MovePhaseOpt_STAY {
		hostHp = 1
	}
	if snapMovMod.DuelOpt == pb.MovePhaseOpt_STAY {
		duelHp = 1
	}
	mvResult := pb.GDMoveConfirmResp{
		RoomKey:      gameSet.RoomKey,
		ResultRange:  gameSet.Range,
		HostCurrCard: gameSet.HostCurrCardKey,
		DuelCurrCard: gameSet.DuelCurrCardKey,
		HostHp:       hostHp,
		DuelHp:       duelHp,
	}
	var hostECInStore, duelECInStore pb.EventCardListSet
	wg := sync.WaitGroup{}
	wg.Add(4)
	errch := make(chan error)
	go func() {
		key := gameSet.RoomKey + mvResult.RdsKeyName()
		wkbox := this.searchAliveClient()
		wkbox.SetPara(key, mvResult)
		wkbox.Preserve(false)
		wg.Done()
	}()
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
	go func() {
		wkbox := this.searchAliveClient()
		var admod pb.ADPhaseSnapMod
		if _, err := wkbox.GetPara(gameSet.RoomKey+admod.RdsKeyName(), &admod); err != nil {
			errch <- err
		}
		admod.CurrAttacker = gameSet.FirstAttack
		admod.FirstAttack = gameSet.FirstAttack
		if _, err := (wkbox).SetPara(gameSet.RoomKey+admod.RdsKeyName(), admod); err != nil {
			errch <- err
		}
		wg.Done()
	}()
	wg.Wait()
	this.BroadCast(&pb.GDBroadcastResp{
		RoomKey:      gameSet.RoomKey,
		Msg:          fmt.Sprintf("MOVE:MOVE_RESULT:"),
		Command:      pb.CastCmd_GET_MOVE_PHASE_RESULT,
		CurrentPhase: pb.EventHookPhase_finish_move_phase,
		PhaseHook:    pb.EventHookType_Proxy,
	})
	log.Println("end of finish-move-phase")
}
