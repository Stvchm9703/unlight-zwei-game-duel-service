package serverCtl

import (
	cm "ULZGameDuelService/pkg/common"
	pb "ULZGameDuelService/proto"
	"context"
	"log"
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
	wkbox := this.searchAliveClient()
	defer func() {
		wkbox.Preserve(false)
		this.mu.Unlock()
		elapsed := time.Since(start)
		log.Printf("Draw-Phase-Confirm took %s", elapsed)
	}()
	if req.IsWatcher {
		return nil, cm.StatusErrorNonPlayer()
	}
	// change
	var returner pb.PhaseSnapMod
	// go func() {}()
	phaseSnapKey := req.RoomKey + ":PhaseState"
	if _, err := (wkbox).GetPara(&phaseSnapKey, &returner); err != nil {
		log.Println(err)
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
	errch := make(chan error)
	go func() {
		if _, err := wkbox.SetPara(&req.RoomKey, returner); err != nil {
			errch <- status.Error(codes.Internal, err.Error())
		}
	}()
	if err := <-errch; err != nil {
		return nil, err
	}
	log.Println("ACK-msg")
	if returner.IsHostReady && returner.IsDuelReady {
		// broadcast for ready next phase
		returner.EventPhase = pb.EventHookPhase_refill_action_card_phase
		returner.HookType = pb.EventHookType_After
		// check event hook
		go this.BroadCast(&pb.GDBroadcastResp{
			RoomKey:      req.RoomKey,
			Msg:          "Both_Ready",
			Command:      pb.CastCmd_GET_DRAW_PHASE_RESULT,
			CurrentPhase: pb.EventHookPhase_refill_action_card_phase,
			PhaseHook:    pb.EventHookType_Proxy,
			Side:         0,
			InstanceSet:  nil,
		})
		// go this.phaseTrigEf()

		// go func() {
		// 	var gameSet pb.GameDataSet
		// 	if _, err := wkbox.SetPara(&req.RoomKey, &gameSet); err != nil {
		// 		errch <- status.Error(codes.Internal, err.Error())
		// 	} else {
		// 		// this.phaseTrigEf(&gameSet)
		// 	}
		// }()
		// if err := <-errch; err != nil {
		// 	return nil, err
		// }
	}
	return &pb.Empty{}, nil
	// return nil, status.Error(codes.Unimplemented, "DRAW_PHASE_CONFIRM")
}
