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
