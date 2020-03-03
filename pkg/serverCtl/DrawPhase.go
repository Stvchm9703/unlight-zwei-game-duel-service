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

	var returner pb.GameDataSet
	if _, err := (wkbox).GetPara(&req.RoomKey, &returner); err != nil {
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

	if returner.IsHostReady && returner.IsDuelReady {
		// broadcast for ready next phase
	}
	return &pb.Empty{}, nil
	// return nil, status.Error(codes.Unimplemented, "DRAW_PHASE_CONFIRM")
}
