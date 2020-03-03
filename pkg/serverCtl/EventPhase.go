package serverCtl

import (
	pb "ULZGameDuelService/proto"
	"context"

	"github.com/gogo/status"
	"google.golang.org/grpc/codes"
	// Static files
	// _ "ULZGameDuelService/statik"
)

func (this *ULZGameDuelServiceBackend) EventPhaseConfirm(context.Context, *pb.GDPhaseConfirmReq) (*pb.Empty, error) {
	return nil, status.Error(codes.Unimplemented, "EVENT_PHASE_CONFIRM")
}
func (this *ULZGameDuelServiceBackend) EventPhaseResult(context.Context, *pb.GDGetInfoReq) (*pb.GDPhaseConfirmResp, error) {
	return nil, status.Error(codes.Unimplemented, "EVENT_PHASE_RESULT")
}
