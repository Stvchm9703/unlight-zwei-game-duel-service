package serverCtl

import (
	pb "ULZGameDuelService/proto"
	"context"

	"github.com/gogo/status"
	"google.golang.org/grpc/codes"
	// Static files
	// _ "ULZGameDuelService/statik"
)

func (this *ULZGameDuelServiceBackend) ADPhaseConfirm(context.Context, *pb.GDADConfirmReq) (*pb.Empty, error) {
	return nil, status.Error(codes.Unimplemented, "AD_PHASE_CONFIRM")
}
func (this *ULZGameDuelServiceBackend) ADPhaseResult(context.Context, *pb.GDGetInfoReq) (*pb.GDADResultResp, error) {
	return nil, status.Error(codes.Unimplemented, "AD_PHASE_RESULT")
}
func (this *ULZGameDuelServiceBackend) ADPhaseDiceResult(context.Context, *pb.GDGetInfoReq) (*pb.GDADDiceResult, error) {
	return nil, status.Error(codes.Unimplemented, "AD_PHASE_DICE_RESULT")
}
