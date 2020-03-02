package serverCtl

import (
	pb "ULZGameDuelService/proto"
	"context"

	"github.com/gogo/status"
	"google.golang.org/grpc/codes"
	// Static files
	// _ "ULZGameDuelService/statik"
)

func (this *ULZGameDuelServiceBackend) MovePhaseConfirm(context.Context, *pb.GDMoveConfirmReq) (*pb.Empty, error) {
	// d := this.
	return nil, status.Error(codes.Unimplemented, "MOVE_PHASE_CONFIRM")
}

func (this *ULZGameDuelServiceBackend) MovePhaseResult(context.Context, *pb.GDGetInfoReq) (*pb.GDMoveConfirmResp, error) {
	return nil, status.Error(codes.Unimplemented, "MOVE_PHASE_RESULT")
}
