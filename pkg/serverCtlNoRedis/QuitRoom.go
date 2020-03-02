package serverCtlNoRedis

import (
	cm "ULZGameDuelService/pkg/common"
	pb "ULZGameDuelService/proto"
	"context"
	"log"
	"time"

	// types "github.com/gogo/protobuf/types"
	"github.com/gogo/status"
	"google.golang.org/grpc/codes"
)

func (b *ULZGameDuelServiceBackend) QuitRoom(ctx context.Context, req *pb.RoomReq) (*pb.Empty, error) {
	start := time.Now()
	b.mu.Lock()
	cm.PrintReqLog(ctx, "quit-room", req)
	defer func() {
		b.mu.Unlock()
		elapsed := time.Since(start)
		log.Printf("Quit-Room took %s", elapsed)
	}()

	tmp, ok := b.Roomlist[req.Key]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "ROOM_NOT_EXIST")
	}

	_, err := b.DelStream(&req.Key, &req.User.Id)
	if err != nil {
		log.Println(err)
	}
	// broadcast to room
	b.BroadCast(&req.Key, &b.CoreKey,
		cm.MsgUserQuitRoom(&req.Key, &req.User.Id, &req.User.Name))

	//
	// edit room
	if tmp.Room.Host.Id == req.User.Id {
		// remove room stream
		b.BroadCast(&req.Key, &b.CoreKey,
			cm.MsgHostQuitRoom(&tmp.Room.Key, &req.User.Id))
		tmp.ClearAll()
		b.Roomlist[req.Key] = nil
		delete(b.Roomlist, req.Key)

	} else if tmp.Room.Dueler.Id == req.User.Id {
		tmp.Room.Dueler = nil
		tmp.Room.Status = pb.RoomStatus_ON_WAIT
	}

	// return nil, errors.New("NotImplement")
	return nil, nil

}
