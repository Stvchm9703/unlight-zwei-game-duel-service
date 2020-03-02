package serverCtl

import (
	pb "ULZGameDuelService/proto"
	"context"

	"github.com/gogo/status"
	"google.golang.org/grpc/codes"
	// Static files
	// _ "ULZGameDuelService/statik"
)

func (this *ULZGameDuelServiceBackend) InstSetEventCard(context.Context, *pb.GDInstanceDT) (*pb.Empty, error) {
	return nil, status.Error(codes.Unimplemented, "EVENT_PHASE_CONFIRM")
}
func (this *ULZGameDuelServiceBackend) InstGetEventCard(context.Context, *pb.GDGetInfoReq) (*pb.GDInstanceDT, error) {
	return nil, status.Error(codes.Unimplemented, "EVENT_PHASE_RESULT")

}
