package serverCtl

import (
	pb "ULZGameDuelService/proto"
	"context"

	"github.com/gogo/status"
	"google.golang.org/grpc/codes"
	// Static files
	// _ "ULZGameDuelService/statik"
)

func (this *ULZGameDuelServiceBackend) ChangePhaseConfirm(context.Context, *pb.GDChangeConfirmReq) (*pb.Empty, error) {
	return nil, status.Error(codes.Unimplemented, "CHANGE_PHASE_CONFIRM")
}
func (this *ULZGameDuelServiceBackend) ChangePhaseResult(context.Context, *pb.GDGetInfoReq) (*pb.Empty, error) {
	return nil, status.Error(codes.Unimplemented, "CHANGE_PHASE_RESULT")

}
