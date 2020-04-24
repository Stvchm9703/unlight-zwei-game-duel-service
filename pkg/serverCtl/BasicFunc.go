package serverCtl

import (
	cm "ULZGameDuelService/pkg/common"
	pb "ULZGameDuelService/proto"
	"context"
	"log"
	"sync"
	"time"

	"github.com/gogo/status"
	"google.golang.org/grpc/codes"
	// Static files
	// _ "ULZGameDuelService/statik"
)

func (this *ULZGameDuelServiceBackend) CreateGame(ctx context.Context, req *pb.GDCreateReq) (*pb.GameDataSet, error) {
	cm.PrintReqLog(ctx, "Create-Game", req)
	start := time.Now()
	this.mu.Lock()
	// wkbox := this.searchAliveClient()
	defer func() {
		this.mu.Unlock()
		elapsed := time.Since(start)
		log.Printf("Create-Game took %s", elapsed)
	}()
	wkbox := this.searchAliveClient()
	l, err := (wkbox).ListRem(&req.RoomKey)
	if err != nil {
		log.Println(err)
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	if len(*l) == 1 {
		// room exist =1
		var returner pb.GameDataSet
		if _, err := (wkbox).GetPara(req.RoomKey, &returner); err != nil {
			log.Println(err)
			return nil, status.Errorf(codes.NotFound, err.Error())
		}
		return &returner, status.Error(codes.AlreadyExists, "create-game,the room exist 1")
		wkbox.Preserve(false)
	}
	gameSetKey := req.RoomKey
	new_gameset := pb.GameDataSet{
		// by request
		RoomKey:         req.RoomKey,
		HostId:          req.HostId,
		DuelId:          req.DuelerId,
		Nvn:             req.Nvn,
		HostCardDeck:    req.HostCardDeck,
		DuelCardDeck:    req.DuelCardDeck,
		GameTurn:        0,
		HostCurrCardKey: 0,
		DuelCurrCardKey: 0,
		Range:           pb.RangeType_MIDDLE,
		EventPhase:      pb.EventHookPhase_gameset_start,
		HookType:        pb.EventHookType_Before,
		FirstAttack:     pb.PlayerSide_HOST,
		CurrPhase:       0,
		EffectCounter:   nil,
	}

	// new_gameset.HostEvent
	// return nil, status.Error(codes.Unimplemented, "CREATE_GAME")
	// Set Para
	if _, err := wkbox.SetPara(gameSetKey, new_gameset); err != nil {
		log.Println(err)
		wkbox.Preserve(false)
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	wg := sync.WaitGroup{}
	errCh := make(chan error)
	// event-card-control
	new_gameset.HostEventCardDeck = genCardSet(150, 0)
	new_gameset.DuelEventCardDeck = genCardSet(150, 0)
	wg.Add(6)
	// Host-Event-Card-Deck
	go func() {
		mwkbox := this.searchAliveClient()
		var set pb.EventCardListSet
		set.Set = new_gameset.HostEventCardDeck
		if _, err := mwkbox.SetPara(req.RoomKey+set.RdsKeyName(pb.PlayerSide_HOST), set); err != nil {
			log.Println(err)
			errCh <- status.Errorf(codes.Internal, err.Error())
		}
		mwkbox.Preserve(false)
		wg.Done()
	}()
	// Duel-Event-Card-Deck
	go func() {
		mwkbox := this.searchAliveClient()
		var set pb.EventCardListSet
		set.Set = new_gameset.DuelEventCardDeck
		if _, err := mwkbox.SetPara(req.RoomKey+set.RdsKeyName(pb.PlayerSide_DUELER), set); err != nil {
			panic(err)
			errCh <- status.Errorf(codes.Internal, err.Error())
		}
		mwkbox.Preserve(false)
		wg.Done()
	}()
	// PhaseSnapMod
	go func() {
		mwkbox := this.searchAliveClient()
		phase_inst := pb.PhaseSnapMod{
			Turns:       1,
			EventPhase:  pb.EventHookPhase_gameset_start,
			HookType:    pb.EventHookType_Proxy,
			IsHostReady: false,
			IsDuelReady: false,
		}

		if _, err := mwkbox.SetPara(req.RoomKey+phase_inst.RdsKeyName(), phase_inst); err != nil {
			log.Println(err)
			errCh <- status.Errorf(codes.Internal, err.Error())
		}
		mwkbox.Preserve(false)
		wg.Done()
	}()
	// move-instance
	go func() {
		mwkbox := this.searchAliveClient()
		move_instance := pb.MovePhaseSnapMod{
			Turns: 0,
		}
		if _, err := mwkbox.SetPara(req.RoomKey+move_instance.RdsKeyName(), move_instance); err != nil {
			log.Println(err)
			errCh <- status.Errorf(codes.Internal, err.Error())
		}
		mwkbox.Preserve(false)
		wg.Done()
	}()
	// ad-phase-instance
	go func() {
		mwkbox := this.searchAliveClient()
		ad_instance := pb.ADPhaseSnapMod{
			Turns:       0,
			FirstAttack: 0,
			EventPhase:  pb.EventHookPhase_start_turn_phase,
		}
		if _, err := mwkbox.SetPara(req.RoomKey+ad_instance.RdsKeyName(), ad_instance); err != nil {
			log.Println(err)
			errCh <- status.Errorf(codes.Internal, err.Error())
		}
		mwkbox.Preserve(false)
		wg.Done()
	}()
	//EffectNodeMod
	go func() {
		mwkbox := this.searchAliveClient()
		ef_instance := pb.EffectNodeSnapMod{
			Turns:     0,
			PendingEf: nil,
		}

		if _, err := mwkbox.SetPara(req.RoomKey+ef_instance.RdsKeyName(), ef_instance); err != nil {
			log.Println(err)
			errCh <- status.Errorf(codes.Internal, err.Error())
			// return nil, status.Errorf(codes.Internal, err.Error())
		}
		mwkbox.Preserve(false)
		wg.Done()
	}()
	wg.Wait()
	if errRes := <-errCh; errRes != nil {
		return nil, errRes
	}
	wkbox.Preserve(false)
	return &new_gameset, nil
}

func (this *ULZGameDuelServiceBackend) GetGameData(ctx context.Context, req *pb.GDGetInfoReq) (*pb.GameDataSet, error) {
	cm.PrintReqLog(ctx, "get-room-info", req)

	start := time.Now()
	this.mu.Lock()
	wkbox := this.searchAliveClient()
	defer func() {
		this.mu.Unlock()
		elapsed := time.Since(start)
		log.Printf("Get-Room took %s", elapsed)
		(wkbox).Preserve(false)
	}()

	var tmp pb.GameDataSet
	var eflist pb.EffectNodeSnapMod
	wg := sync.WaitGroup{}
	errCh := make(chan error)
	wg.Add(2)
	go func() {
		if _, err := wkbox.GetPara(req.RoomKey, &tmp); err != nil {
			log.Fatalln(err)
			errCh <- status.Errorf(codes.Internal, err.Error())
		}
		wg.Done()
	}()
	go func() {
		if _, err := wkbox.GetPara(req.RoomKey+eflist.RdsKeyName(), &eflist); err != nil {
			log.Fatalln(err)
			errCh <- status.Errorf(codes.Internal, err.Error())
		}
		wg.Done()
	}()
	if errRes := <-errCh; errRes != nil {
		return nil, errRes
	}
	tmp.EffectCounter = eflist.PendingEf
	// return nil, status.Error(codes.Unimplemented, "CREATE_GAME")
	return &tmp, nil
}

func (this *ULZGameDuelServiceBackend) QuitGame(context.Context, *pb.GDCreateReq) (*pb.Empty, error) {
	return nil, status.Error(codes.Unimplemented, "CREATE_GAME")

}
