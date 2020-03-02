package serverCtlNoRedis

import (
	cm "ULZGameDuelService/pkg/common"
	pb "ULZGameDuelService/proto"
	"context"
	"log"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// UpdateRoom :
func (b *ULZGameDuelServiceBackend) UpdateRoom(ctx context.Context, req *pb.RoomCreateReq) (*pb.Room, error) {
	// return nil, status.Errorf(codes.Unimplemented, "method DeleteRoom not implemented")
	start := time.Now()
	b.mu.Lock()
	cm.PrintReqLog(ctx, "update-room", req)
	defer func() {
		b.mu.Unlock()
		elapsed := time.Since(start)
		log.Printf("Quit-Room took %s", elapsed)
	}()

	rmg, ok := b.Roomlist[req.Key]
	if !ok {
		return nil, status.Error(codes.NotFound, "ROOM_NOT_FOUND")
	}

	rmg.Room.Password = req.Password
	rmg.Room.CostLimitMax = req.CostLimitMax
	rmg.Room.CostLimitMin = req.CostLimitMin
	rmg.Room.CharCardNvn = req.CharCardNvn
	rmg.Room.CharCardLimitMax = req.CharCardLimitMax
	rmg.Room.CharCardLimitMin = req.CharCardLimitMin
	b.BroadCast(&rmg.Key, &req.Host.Id, cm.MsgHostUpdateRoom(&rmg.Key, &rmg.Password))
	return &rmg.Room, nil
}
