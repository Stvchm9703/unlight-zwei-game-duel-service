package serverCtl

import (
	pb "ULZGameDuelService/proto"
	"context"
	// Static files
	// _ "ULZGameDuelService/statik"
)

func (this *ULZGameDuelServiceBackend) ServerBroadcast(*pb.GDGetInfoReq, pb.GameDuelService_ServerBroadcastServer) error

func (this *ULZGameDuelServiceBackend) CreateGame(context.Context, *pb.GDCreateReq) (*pb.GameDataSet, error)

func (this *ULZGameDuelServiceBackend) GetGameData(context.Context, *pb.GDGetInfoReq) (*pb.GameDataSet, error)

func (this *ULZGameDuelServiceBackend) QuitGame(context.Context, *pb.GDCreateReq) (*pb.Empty, error)
