package serverCtl

import (
	pb "ULZGameDuelService/proto"
	"context"

	"github.com/gogo/status"
	"google.golang.org/grpc/codes"
	// Static files
	// _ "ULZGameDuelService/statik"
)

func (this *ULZGameDuelServiceBackend) DrawPhaseConfirm(context.Context, *pb.GDGetInfoReq) (*pb.Empty, error) {
	return nil, status.Error(codes.Unimplemented, "DRAW_PHASE_CONFIRM")
}
