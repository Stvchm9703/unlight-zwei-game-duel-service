package serverCtl

import (
	pb "ULZGameDuelService/proto"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
	// Static files
	// _ "ULZGameDuelService/statik"
)

// Handle Proxy phase
func (this *ULZGameDuelServiceBackend) proxyHandle(
	gameSet *pb.GameDataSet,
	phaseMod *pb.PhaseSnapMod,
	effectMod *pb.EffectNodeSnapMod,
	snapMod ...interface{},
) {
	switch phaseMod.EventPhase {
	case pb.EventHookPhase_gameset_start:
		this.gamesetStart(gameSet)
		break
	case pb.EventHookPhase_start_turn_phase:
		this.startTurnPhase(gameSet)
		break
	case pb.EventHookPhase_refill_action_card_phase:
		this.refillActionCard(gameSet)
		break
	// case pb.EventHookPhase_move_card_drop_phase:
	// 		this.moveCardDropPhase(gameSet)
	// 		break
	case pb.EventHookPhase_determine_move_phase:
		snapMovMod, _ := snapMod[0].(*pb.MovePhaseSnapMod)
		this.determineMovePhaseHandle(gameSet, phaseMod, effectMod, snapMovMod)
		break
	case pb.EventHookPhase_finish_move_phase:
		this.finishMovePhase(gameSet)
		break
	case pb.EventHookPhase_chara_change_phase:
		// this.charaChangePhase(gameSet)
		break
	case pb.EventHookPhase_determine_chara_change_phase:
		//
	// case pb.EventHookPhase_attack_card_drop_phase:
	// 		snapADMod, _ := snapMod[0].(*pb.ADPhaseSnapMod)
	// 		this.attackPhaseHandle(gameSet, snapADMod, phaseMod, effectMod)
	// 		break
	// case pb.EventHookPhase_defence_card_drop_phase:
	// 		snapADMod, _ := snapMod[0].(*pb.ADPhaseSnapMod)
	// 		this.defencePhaseHandle(gameSet, snapADMod, phaseMod, effectMod)
	// 		break
	case pb.EventHookPhase_determine_battle_point_phase:

	case pb.EventHookPhase_battle_result_phase:
		snapADMod, _ := snapMod[0].(*pb.ADPhaseSnapMod)
		this.battlePhaseHandle(gameSet, phaseMod, effectMod, snapADMod)

	case pb.EventHookPhase_damage_phase:

	case pb.EventHookPhase_dead_chara_change_phase:

	case pb.EventHookPhase_determine_dead_chara_change_phase:

	case pb.EventHookPhase_change_initiative_phase:

	case pb.EventHookPhase_finish_turn_phase:

	case pb.EventHookPhase_gameset_end:
	}
	fmt.Println("Hello world")

	return
}

func (this *ULZGameDuelServiceBackend) gamesetStart(gameSet *pb.GameDataSet) {
	log.Printf("game-start: key: %s\n", gameSet.RoomKey)
	go this.BroadCast(&pb.GDBroadcastResp{
		RoomKey:      gameSet.RoomKey,
		Msg:          fmt.Sprintf("GameStart:%v", gameSet.GameTurn),
		Command:      pb.CastCmd_INSTANCE_STATUS_CHANGE,
		CurrentPhase: gameSet.EventPhase,
		PhaseHook:    gameSet.HookType,
	})
}

func (this *ULZGameDuelServiceBackend) startTurnPhase(gameSet *pb.GameDataSet) {
	log.Printf("start-turn-phase: turn %v", gameSet.GameTurn)
	gameSet.GameTurn++
	this.BroadCast(&pb.GDBroadcastResp{
		RoomKey:      gameSet.RoomKey,
		Msg:          fmt.Sprintf("TurnStart:%v", gameSet.GameTurn),
		Command:      pb.CastCmd_INSTANCE_STATUS_CHANGE,
		CurrentPhase: gameSet.EventPhase,
		PhaseHook:    gameSet.HookType,
	})

}

func (this *ULZGameDuelServiceBackend) refillActionCard(
	gameSet *pb.GameDataSet,
	phaseMod *pb.PhaseSnapMod,
	effectMod *pb.EffectNodeSnapMod,
) {
	log.Printf("refill-action-card-phase: turn %v", gameSet.RoomKey)
	wg := sync.WaitGroup{}
	wg.Add(2)
	var hostECInStore, duelECInStore pb.EventCardListSet
	var hostDrawSet, duelDrawSet []*pb.ECShortHand
	hostNoMoreCardFlag, duelNoMoreCardFlag := false, false
	errch := make(chan error)
	go func() {
		wkbox := this.searchAliveClient()
		key := gameSet.RoomKey + hostECInStore.RdsKeyName(pb.PlayerSide_HOST)
		if _, err := (wkbox).GetPara(&key, &hostECInStore); err != nil {
			errch <- err
		}
		inDeck := pb.EventCardFilter(hostECInStore.Set, func(v *pb.EventCard) bool {
			return v.Position == pb.EventCardPos_BLOCK
		})
		inHand := pb.EventCardFilter(hostECInStore.Set, func(v *pb.EventCard) bool {
			return v.Position == pb.EventCardPos_INSIDE
		})
		// destoryed := pb.EventCardFilter(hostECToHands.Set, func(v *pb.EventCard) bool {
		// 	return v.Position == pb.EventCardPos_DESTROY
		// })
		numberToDraw := 9 - len(inHand)

		if numberToDraw > len(inDeck) {
			hostNoMoreCardFlag = true
			numberToDraw = len(inDeck)
		}

		var tmpSet []int32
		rand.Seed(int64(time.Now().UnixNano()))
		stoper := 0
		for stoper < numberToDraw {
			// rand
			tmpNum := rand.Intn(len(inDeck))
			isExist := false
			for _, v := range tmpSet {
				if v == inDeck[tmpNum].Id {
					isExist = true
				}
			}
			if !isExist {
				tmpSet = append(tmpSet, inDeck[tmpNum].Id)
				stoper++
			}
		}
		for _, v := range tmpSet {
			for _, vv := range hostECInStore.Set {
				if vv.Id == v {
					vv.Position = pb.EventCardPos_INSIDE
					hostDrawSet = append(hostDrawSet, vv.ToECShostHand())
				}
			}
		}

		if _, err := (wkbox).SetPara(&key, hostECInStore); err != nil {
			errch <- err
		}
		wg.Done()
	}()
	go func() {
		wkbox := this.searchAliveClient()
		key := gameSet.RoomKey + duelECInStore.RdsKeyName(pb.PlayerSide_DUELER)
		if _, err := (wkbox).GetPara(&key, &duelECInStore); err != nil {
			errch <- err
		}
		inDeck := pb.EventCardFilter(duelECInStore.Set, func(v *pb.EventCard) bool {
			return v.Position == pb.EventCardPos_BLOCK
		})
		inHand := pb.EventCardFilter(duelECInStore.Set, func(v *pb.EventCard) bool {
			return v.Position == pb.EventCardPos_INSIDE
		})
		// destoryed := pb.EventCardFilter(hostECToHands.Set, func(v *pb.EventCard) bool {
		// 	return v.Position == pb.EventCardPos_DESTROY
		// })
		numberToDraw := 9 - len(inHand)

		if numberToDraw > len(inDeck) {
			duelNoMoreCardFlag = true
			numberToDraw = len(inDeck)
		}

		var tmpSet []int32
		rand.Seed(int64(time.Now().UnixNano()))
		stoper := 0
		for stoper < numberToDraw {
			// rand
			tmpNum := rand.Intn(len(inDeck))
			isExist := false
			for _, v := range tmpSet {
				if v == inDeck[tmpNum].Id {
					isExist = true
				}
			}
			if !isExist {
				tmpSet = append(tmpSet, inDeck[tmpNum].Id)
				stoper++
			}
		}
		for _, v := range tmpSet {
			for _, vv := range duelECInStore.Set {
				if vv.Id == v {
					vv.Position = pb.EventCardPos_INSIDE
					duelDrawSet = append(duelDrawSet, vv.ToECShostHand())
				}
			}
		}
		if _, err := (wkbox).SetPara(&key, duelECInStore); err != nil {
			errch <- err
		}
		wg.Done()
	}()
	wg.Wait()

	if hostNoMoreCardFlag {
		this.BroadCast(&pb.GDBroadcastResp{
			RoomKey:      gameSet.RoomKey,
			Msg:          fmt.Sprintf("HostDrawCard"),
			Command:      pb.CastCmd_GET_DRAW_PHASE_RESULT,
			CurrentPhase: gameSet.EventPhase,
			PhaseHook:    gameSet.HookType,
			InstanceSet:  hostDrawSet,
		})

	} else {
		this.BroadCast(&pb.GDBroadcastResp{
			RoomKey:      gameSet.RoomKey,
			Msg:          fmt.Sprintf("HostDrawCard:NoMoreCard"),
			Command:      pb.CastCmd_GET_DRAW_PHASE_RESULT,
			CurrentPhase: gameSet.EventPhase,
			PhaseHook:    gameSet.HookType,
			InstanceSet:  hostDrawSet,
		})
	}
	if duelNoMoreCardFlag {
		this.BroadCast(&pb.GDBroadcastResp{
			RoomKey:      gameSet.RoomKey,
			Msg:          fmt.Sprintf("DuelDrawCard"),
			Command:      pb.CastCmd_GET_DRAW_PHASE_RESULT,
			CurrentPhase: gameSet.EventPhase,
			PhaseHook:    gameSet.HookType,
			InstanceSet:  duelDrawSet,
		})
	} else {
		this.BroadCast(&pb.GDBroadcastResp{
			RoomKey:      gameSet.RoomKey,
			Msg:          fmt.Sprintf("DuelDrawCard:NoMoreCard"),
			Command:      pb.CastCmd_GET_DRAW_PHASE_RESULT,
			CurrentPhase: gameSet.EventPhase,
			PhaseHook:    gameSet.HookType,
			InstanceSet:  duelDrawSet,
		})
	}
	log.Println("End of Draw Card proxy")

}
func (this *ULZGameDuelServiceBackend) determineMovePhase(gameSet *pb.GameDataSet) {

}

func (this *ULZGameDuelServiceBackend) finishMovePhase(gameSet *pb.GameDataSet) {}

func (this *ULZGameDuelServiceBackend) determineBattlePointPhase(gameSet *pb.GameDataSet) {}

func (this *ULZGameDuelServiceBackend) damageResultPhase(gameSet *pb.GameDataSet) {}

func (this *ULZGameDuelServiceBackend) deadCharaChangePhase(gameSet *pb.GameDataSet) {}

func (this *ULZGameDuelServiceBackend) determineDeadCharaChangePhase(gameSet *pb.GameDataSet) {}

func (this *ULZGameDuelServiceBackend) changeInitiativePhase(gameSet *pb.GameDataSet) {}

func (this *ULZGameDuelServiceBackend) finishTurnPhase(gameSet *pb.GameDataSet) {}
