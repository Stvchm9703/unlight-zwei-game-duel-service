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

var (
	InstSetECMsg     *pb.GDBroadcastResp = &pb.GDBroadcastResp{}
	InstMsgStoreTime int                 = 15000
)

func (this *ULZGameDuelServiceBackend) InstSetEventCard(ctx context.Context, req *pb.GDInstanceDT) (*pb.Empty, error) {
	cm.PrintReqLog(ctx, "Inst-Set-EventCard", req)
	start := time.Now()
	this.mu.Lock()
	wkbox := this.searchAliveClient()
	defer func() {
		wkbox.Preserve(false)
		this.mu.Unlock()
		elapsed := time.Since(start)
		log.Printf("Inst-Set-EventCard took %s", elapsed)
	}()
	sidetmp := req.Side.String()
	instKey := req.RoomKey + "_instMsg@" + req.Side.String()
	wkbox.SetParaWTO(&instKey, req, InstMsgStoreTime)

	var tmpSet pb.GameDataSet
	if _, err := wkbox.GetPara(&req.RoomKey, &tmpSet); err != nil {
		log.Fatalln(err)
		return nil, status.Errorf(codes.NotFound, err.Error())
	}

	err := this.BroadCast(&req.RoomKey, &sidetmp, InstSetECMsg)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return nil, status.Error(codes.Unimplemented, "EVENT_PHASE_CONFIRM")
}
func (this *ULZGameDuelServiceBackend) InstGetEventCard(ctx context.Context, req *pb.GDGetInfoReq) (*pb.GDInstanceDT, error) {
	return nil, status.Error(codes.Unimplemented, "EVENT_PHASE_RESULT")

}
