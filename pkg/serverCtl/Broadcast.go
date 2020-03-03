package serverCtl

import (
	cm "ULZGameDuelService/pkg/common"
	pb "ULZGameDuelService/proto"
	"log"

	"github.com/gogo/status"
	"google.golang.org/grpc/codes"
	// "time"
)

func (this *ULZGameDuelServiceBackend) ServerBroadcast(rReq *pb.GDGetInfoReq, stream pb.GameDuelService_ServerBroadcastServer) error {
	cm.PrintReqLog(nil, "server-broadcast", rReq)
	_, err := this.AddStream(&rReq.RoomKey, &rReq.IncomeUserId, &stream)
	if err != nil {
		return status.Error(codes.NotFound, err.Error())
	}
	go func() {
		<-stream.Context().Done()
		log.Println("close done")
		_, err := this.DelStream(&rReq.RoomKey, &rReq.IncomeUserId)
		if err != nil {
			log.Println(err)
		}
	}()
	for {
	}
}
