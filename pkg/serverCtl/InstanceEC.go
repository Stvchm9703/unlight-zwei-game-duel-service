package serverCtl

import (
	cm "ULZGameDuelService/pkg/common"
	pb "ULZGameDuelService/proto"
	"context"
	"log"
	"sync"
	"time"

	"github.com/gogo/status"
	"google.golang.org/grpc/codes"
	// Static files
	// _ "ULZGameDuelService/statik"
)

var (
	// InstSetECMsg     *pb.GDBroadcastResp = &pb.GDBroadcastResp{}
	InstMsgStoreTime int = 15000
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
	// sidetmp := req.Side.String()
	instKey := req.RoomKey + "_instMsg@" + req.Side.String()
	go wkbox.SetParaWTO(&instKey, req, InstMsgStoreTime)

	// broadcast have no ddp on redis execution
	// go first
	go this.BroadCast(&pb.GDBroadcastResp{
		RoomKey:      req.RoomKey,
		Msg:          "",
		Command:      pb.CastCmd_GET_INSTANCE_CARD,
		CurrentPhase: req.CurrentPhase,
		InstanceSet:  req.UpdateCard,
	})

	var tmpSet []pb.EventCard
	ky := req.RoomKey
	if req.Side == pb.PlayerSide_HOST {
		ky = req.RoomKey + ":HtEvtCrdDk"
	} else if req.Side == pb.PlayerSide_DUELER {
		ky = req.RoomKey + ":DlEvtCrdDk"
	}
	if _, err := wkbox.GetPara(&ky, &tmpSet); err != nil {
		log.Fatalln(err)
		return nil, status.Errorf(codes.NotFound, err.Error())
	}
	wg := sync.WaitGroup{}

	for _, v := range tmpSet {
		wg.Add(1)
		go func() {
			for _, vk := range req.UpdateCard {
				if vk.CardId == v.Id {
					v.IsInvert = vk.IsInvert
					v.Position = vk.Position
					break
				}
			}
			wg.Done()
		}()
	}
	if _, err := wkbox.SetPara(&ky, tmpSet); err != nil {
		log.Println(err)
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &pb.Empty{}, nil
}
