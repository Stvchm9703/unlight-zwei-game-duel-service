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

type movePhaseSnapMod struct {
	Turn    int
	HostVal int
	DuelVal int
	HostOpt pb.MovePhaseOpt
	DuelOpt pb.MovePhaseOpt
}

func (this *ULZGameDuelServiceBackend) MovePhaseConfirm(ctx context.Context, req *pb.GDMoveConfirmReq) (*pb.Empty, error) {
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

	var returner pb.GameDataSet
	if _, err := (wkbox).GetPara(&req.RoomKey, &returner); err != nil {
		log.Println(err)
		return nil, status.Errorf(codes.NotFound, err.Error())
	}

	if returner.EventPhase != pb.EventHookPhase_move_card_drop_phase {
		return nil, status.Error(codes.FailedPrecondition, "NOT_IN_MOVE_PHASE")
	}

	if (req.Side == pb.PlayerSide_HOST && returner.IsHostReady) ||
		(req.Side == pb.PlayerSide_DUELER && returner.IsDuelReady) {
		return nil, status.Error(codes.AlreadyExists, "ALREADY_READY")
	}

	if req.Side == pb.PlayerSide_HOST {
		returner.IsHostReady = true
		// Do Update

	} else if req.Side == pb.PlayerSide_DUELER {
		returner.IsDuelReady = true
		// Do Update
	}

	if returner.IsHostReady && returner.IsDuelReady {
		// broadcast for ready next phase
		returner.IsHostReady = false
		returner.IsDuelReady = false
		returner.EventPhase = pb.EventHookPhase_determine_move_phase
		returner.HookType = pb.EventHookType_Before
		// check event hook
	}

	if _, err := (wkbox).SetPara(&req.RoomKey, &returner); err != nil {
		log.Println(err)
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &pb.Empty{}, nil
}

func (this *ULZGameDuelServiceBackend) MovePhaseResult(context.Context, *pb.GDGetInfoReq) (*pb.GDMoveConfirmResp, error) {
	return nil, status.Error(codes.Unimplemented, "MOVE_PHASE_RESULT")
}
