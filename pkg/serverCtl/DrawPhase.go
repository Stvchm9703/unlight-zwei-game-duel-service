package serverCtl

import (
	cm "ULZGameDuelService/pkg/common"
	pb "ULZGameDuelService/proto"
	"context"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"

	"github.com/gogo/status"
	"google.golang.org/grpc/codes"
	// Static files
	// _ "ULZGameDuelService/statik"
)

func (this *ULZGameDuelServiceBackend) DrawPhaseConfirm(ctx context.Context, req *pb.GDGetInfoReq) (*pb.Empty, error) {
	cm.PrintReqLog(ctx, "Draw-Phase-Confirm", req)
	start := time.Now()
	this.mu.Lock()
	defer func() {
		this.mu.Unlock()
		elapsed := time.Since(start)
		log.Printf("Draw-Phase-Confirm took %s", elapsed)
	}()
	if req.IsWatcher {
		return nil, cm.StatusErrorNonPlayer()
	}
	// change
	wkbox := this.searchAliveClient()
	var returner pb.PhaseSnapMod
	if _, err := (wkbox).GetPara(req.RoomKey+returner.RdsKeyName(), &returner); err != nil {
		log.Println(err)
		wkbox.Preserve(false)
		return nil, status.Errorf(codes.NotFound, err.Error())
	}

	if returner.EventPhase != req.CurrentPhase ||
		returner.EventPhase != pb.EventHookPhase_refill_action_card_phase ||
		req.CurrentPhase != pb.EventHookPhase_refill_action_card_phase {
		return nil, status.Error(codes.FailedPrecondition, "NOT_IN_DRAW_PHASE")
	}

	if (req.Side == pb.PlayerSide_HOST && returner.IsHostReady) ||
		(req.Side == pb.PlayerSide_DUELER && returner.IsDuelReady) {
		return nil, status.Error(codes.AlreadyExists, "ALREADY_READY")
	}

	if req.Side == pb.PlayerSide_HOST {
		returner.IsHostReady = true
	} else if req.Side == pb.PlayerSide_DUELER {
		returner.IsDuelReady = true
	}
	//  #change
	// broadcast ok ?
	go this.BroadCast(&pb.GDBroadcastResp{
		RoomKey:      req.RoomKey,
		Msg:          req.Side.String() + "_READY",
		Command:      pb.CastCmd_GET_DRAW_PHASE_RESULT,
		CurrentPhase: pb.EventHookPhase_refill_action_card_phase,
		PhaseHook:    pb.EventHookType_Proxy,
		Side:         req.Side,
		InstanceSet:  nil,
	})
	if _, err := wkbox.SetPara(req.RoomKey+returner.RdsKeyName(), returner); err != nil {
		return nil, err
	}
	wkbox.Preserve(false)
	log.Println("ACK-msg")
	if returner.IsHostReady && returner.IsDuelReady {
		go func() {
			// broadcast for ready next phase
			this.BroadCast(&pb.GDBroadcastResp{
				RoomKey:      req.RoomKey,
				Msg:          "Both_Ready",
				Command:      pb.CastCmd_GET_DRAW_PHASE_RESULT,
				CurrentPhase: pb.EventHookPhase_refill_action_card_phase,
				PhaseHook:    pb.EventHookType_Proxy,
				Side:         0,
				InstanceSet:  nil,
			})

			returner.EventPhase = pb.EventHookPhase_refill_action_card_phase
			returner.HookType = pb.EventHookType_Proxy
			// check event hook
			var gmset pb.GameDataSet
			mbox := this.searchAliveClient()
			if _, err := mbox.GetPara(req.RoomKey, &gmset); err != nil {
				log.Println(err)
			}
			var efMod pb.EffectNodeSnapMod
			if _, err := mbox.GetPara(req.RoomKey+efMod.RdsKeyName(), &efMod); err != nil {
				log.Println(err)
			}
			mbox.Preserve(false)
			this.moveNextPhase(&gmset, &returner, &efMod)
		}()
	}
	return &pb.Empty{}, nil
	// return nil, status.Error(codes.Unimplemented, "DRAW_PHASE_CONFIRM")
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
		if _, err := (wkbox).GetPara(key, &hostECInStore); err != nil {
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
		if _, err := (wkbox).SetPara(key, duelECInStore); err != nil {
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
	log.Printf("Rm %s: End of Draw Card proxy, wait for draw-confirm", gameSet.RoomKey)
	/**
	 * since the effect-node will done by executeEffectNode in move-next-phase
	 */
}
