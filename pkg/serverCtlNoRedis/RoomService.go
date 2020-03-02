package serverCtlNoRedis

import (
	// _ "ULZGameDuelService"
	cm "ULZGameDuelService/pkg/common"
	cf "ULZGameDuelService/pkg/config"
	pb "ULZGameDuelService/proto"
	"errors"
	"log"
	"sync"
	// ants "github.com/panjf2000/ants/v2"
)

var _ pb.RoomServiceServer = (*ULZGameDuelServiceBackend)(nil)

// Remark: the framework make consider "instant" request
//
type ULZGameDuelServiceBackend struct {
	// pb.ULZGameDuelServiceServer
	mu       *sync.Mutex
	CoreKey  string
	Roomlist map[string]*RoomMgr
}

// New : Create new backend
func New(conf *cf.ConfTmp) *ULZGameDuelServiceBackend {
	ck := "RSCore" + cm.HashText(conf.APIServer.IP)

	g := ULZGameDuelServiceBackend{
		CoreKey:  ck,
		mu:       &sync.Mutex{},
		Roomlist: make(map[string]*RoomMgr),
	}
	return &g
}

func (this *ULZGameDuelServiceBackend) Shutdown() {
	log.Println("in shtdown proc")
	/// TODO: send closing msg to all client
	for _, v := range this.Roomlist {
		log.Println("Server OS.sigKill")
		// v.BroadCast("RmSvrMgr", cm.MsgSystShutdown(v.Room.Key))
		v.ClearAll()
	}
	// this.CloseDB()
	log.Println("endof shutdown proc:", this.CoreKey)
}

// PrintReqLog

// ======================================================================================================
// RoomMgr : Room Manager
type RoomMgr struct {
	pb.Room
	clientConn map[string]*pb.RoomService_ServerBroadcastServer
}

// ----------------------------------------------------------------------------------------------------
// roommgr.clientConn

func (rm *ULZGameDuelServiceBackend) GetStream(roomKey *string, userId *string) *pb.RoomService_ServerBroadcastServer {
	a, ok := rm.Roomlist[*roomKey]
	b, ok := a.clientConn[*userId]
	if ok {
		return b
	}
	return nil
}

func (rm *ULZGameDuelServiceBackend) AddStream(roomKey *string, userId *string, stream *pb.RoomService_ServerBroadcastServer) (bool, error) {
	// _, ok := rm.bc_stream[user_id]

	if rm.Roomlist == nil {
		log.Println("help this")
	} else {
		log.Println("this-roomlist : ", rm.Roomlist)
	}

	a, ok := rm.Roomlist[*roomKey]
	if !ok {
		return false, errors.New("ROOM_NOT_EXIST")
	}

	_, ok = a.clientConn[*userId]
	if ok {
		return false, errors.New("USER_EXIST")
	}
	a.clientConn[*userId] = stream
	return true, nil
}

func (rm *ULZGameDuelServiceBackend) DelStream(roomKey *string, userId *string) (bool, error) {
	log.Println("Del Stream:", userId)
	a, ok := rm.Roomlist[*roomKey]
	if !ok {
		return false, errors.New("ROOM_NOT_EXIST")
	}
	if a.clientConn[*userId] != nil {
		*(a.clientConn[*userId]) = nil
		delete(a.clientConn, *userId)
		return true, nil
	}
	return false, errors.New("StreamNotExist")
}

func (rm *ULZGameDuelServiceBackend) BroadCast(roomkey *string, from *string, message *pb.RoomMsg) error {
	log.Println("BS!", message)
	log.Println(rm.Roomlist[*roomkey])
	if rm.Roomlist == nil {
		log.Println("help this")
	} else {
		log.Println("this-roomlist : ", rm.Roomlist)
	}
	rmb, ok := rm.Roomlist[*roomkey]
	if !ok {
		log.Println("room not exist")
		return errors.New("ROOM_NOT_EXIST")
	}
	for k, v := range rmb.clientConn {
		if k != *from {
			(*v).Send(message)
		}
	}
	return nil
}

// ---------------------------------------------------------------------------------------------
// RoomStreamBox Controlling

func (rm *RoomMgr) ClearAll() {
	log.Println("ClearAll Proc")
	for _, vc := range rm.clientConn {
		(*vc).Send(cm.MsgSystShutdown(&rm.Room.Key))
	}
	for k := range rm.clientConn {
		*(rm.clientConn[k]) = nil
		delete(rm.clientConn, k)
	}
	return
}
