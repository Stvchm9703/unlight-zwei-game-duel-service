package serverCtlNoRedis

import (
	// cm "ULZGameDuelService/pkg/common"
	cm "ULZGameDuelService/pkg/common"
	pb "ULZGameDuelService/proto"

	// "context"
	"log"
	"time"
)

// GetRoomList :
func (b *ULZGameDuelServiceBackend) GetRoomList(req *pb.RoomCreateReq, stream pb.RoomService_GetRoomListServer) error {
	start := time.Now()
	cm.PrintReqLog(nil, "get-room-list", req)

	defer func() {
		elapsed := time.Since(start)
		log.Printf("Get-Room-List took %s", elapsed)
	}()

	var tmp []*string

	for v := range b.Roomlist {
		if b.Roomlist[v].Room.CharCardNvn == req.CharCardNvn &&
			(req.CostLimitMax != 0 && req.CostLimitMax == b.Roomlist[v].Room.CostLimitMax) &&
			(req.CostLimitMax != 0 && req.CostLimitMax == b.Roomlist[v].Room.CostLimitMax) &&
			(req.CharCardLimitMax != nil && req.CharCardLimitMax == b.Roomlist[v].Room.CharCardLimitMax) &&
			(req.CharCardLimitMin != nil && req.CharCardLimitMin == b.Roomlist[v].Room.CharCardLimitMin) {

			tmp = append(tmp, &b.Roomlist[v].Room.Key)
			stream.Send(&b.Roomlist[v].Room)
		}
	}
	for v := range b.Roomlist {
		if b.Roomlist[v].Room.CharCardNvn == req.CharCardNvn &&
			(req.CostLimitMax != 0 && req.CostLimitMax == b.Roomlist[v].Room.CostLimitMax) &&
			(req.CostLimitMax != 0 && req.CostLimitMax == b.Roomlist[v].Room.CostLimitMax) {
			rtmp := false
			for k := range tmp {
				if *tmp[k] == b.Roomlist[v].Room.Key {
					rtmp = true
				}
			}
			if !rtmp {
				tmp = append(tmp, &b.Roomlist[v].Room.Key)
				stream.Send(&b.Roomlist[v].Room)
			}
		}
	}
	for v := range b.Roomlist {
		if b.Roomlist[v].Room.CharCardNvn == req.CharCardNvn {
			rtmp := false
			for k := range tmp {
				if *tmp[k] == b.Roomlist[v].Room.Key {
					rtmp = true
				}
			}
			if !rtmp {
				tmp = append(tmp, &b.Roomlist[v].Room.Key)
				stream.Send(&b.Roomlist[v].Room)
			}
		}
	}

	return nil
}
