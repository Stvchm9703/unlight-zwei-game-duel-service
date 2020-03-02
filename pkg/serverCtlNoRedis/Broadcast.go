package serverCtlNoRedis

import (
	cm "ULZGameDuelService/pkg/common"
	pb "ULZGameDuelService/proto"
	"context"
	"log"

	"github.com/gogo/status"
	"google.golang.org/grpc/codes"
)

func (this *ULZGameDuelServiceBackend) ServerBroadcast(rReq *pb.RoomReq, stream pb.RoomService_ServerBroadcastServer) error {
	// log.Println("\nServer Broadcast Connect\n methods: ServerBroadcast")
	cm.PrintReqLog(nil, "server-broadcast", rReq)
	_, err := this.AddStream(&rReq.Key, &rReq.User.Id, &stream)
	if err != nil {
		return status.Error(codes.NotFound, err.Error())
	}

	go func() {
		<-stream.Context().Done()
		log.Println("close done")
		_, err := this.DelStream(&rReq.Key, &rReq.User.Id)
		if err != nil {
			log.Println(err)
		}
		this.BroadCast(&rReq.Key, &rReq.User.Id,
			cm.MsgUserQuitRoom(&rReq.Key, &rReq.User.Id, &rReq.User.Name))
	}()
	for {
	}
}

func (this *ULZGameDuelServiceBackend) SendMessage(ctx context.Context, msg *pb.RoomMsg) (*pb.Empty, error) {
	cm.PrintReqLog(nil, "server-broadcast:msg", msg)

	this.BroadCast(&msg.Key, &msg.FromId, msg)
	return nil, nil
}
