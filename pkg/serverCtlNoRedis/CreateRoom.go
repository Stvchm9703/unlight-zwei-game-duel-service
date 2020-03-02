package serverCtlNoRedis

import (
	cm "ULZGameDuelService/pkg/common"
	pb "ULZGameDuelService/proto"
	"context"
	"log"
	"time"
)

// CreateRoom :

func (this *ULZGameDuelServiceBackend) CreateRoom(ctx context.Context, req *pb.RoomCreateReq) (*pb.Room, error) {
	cm.PrintReqLog(ctx, "create-room", req)
	start := time.Now()
	this.mu.Lock()
	defer func() {
		this.mu.Unlock()
		elapsed := time.Since(start)
		log.Printf("Create-Room took %s", elapsed)
	}()
	// for loop it
	tmptime := time.Now().String() + req.Host.GetId()
	var f = ""
	for {
		f = cm.HashText(tmptime)
		_, xist := this.Roomlist["Rm"+f]
		if !xist {
			break
		}
	}
	log.Println(req.Host)
	if req.CharCardNvn == 0 {
		req.CharCardNvn = 1
	}
	if req.CostLimitMax == 0 {
		req.CostLimitMax = 200
	}
	if req.CharCardLimitMax == nil {
		req.CharCardLimitMax = &pb.RmCharCardInfo{
			Cost: 40,
		}
	}
	rmTmp := pb.Room{
		Key:              "Rm" + f,
		Id:               f[0:5],
		Host:             req.Host,
		Dueler:           nil,
		Status:           pb.RoomStatus_ON_WAIT,
		Turns:            0,
		CostLimitMax:     req.CostLimitMax,
		CostLimitMin:     req.CostLimitMin,
		CharCardLimitMax: req.CharCardLimitMax,
		CharCardLimitMin: req.CharCardLimitMin,
		CharCardNvn:      req.CharCardNvn,
	}
	rmTmp1 := RoomMgr{
		Room:       rmTmp,
		clientConn: make(map[string]*pb.RoomService_ServerBroadcastServer),
	}
	// !FIXME
	this.Roomlist["Rm"+f] = &rmTmp1
	return &rmTmp, nil
}
