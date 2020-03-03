package serverCtl

import (
	// _ "ULZGameDuelService"
	cm "ULZGameDuelService/pkg/common"
	cf "ULZGameDuelService/pkg/config"
	rd "ULZGameDuelService/pkg/store/redis"
	pb "ULZGameDuelService/proto"
	"errors"
	"fmt"
	"log"
	"strconv"
	"sync"
	"time"
	// ants "github.com/panjf2000/ants/v2"
)

var _ pb.GameDuelServiceServer = (*ULZGameDuelServiceBackend)(nil)

// Remark: the framework make consider "instant" request
//
type ULZGameDuelServiceBackend struct {
	// pb.ULZGameDuelServiceServer
	mu         *sync.Mutex
	CoreKey    string
	redhdlr    []*rd.RdsCliBox
	roomStream map[string](*RoomStreamBox)
}
type RoomStreamBox struct {
	clientConn map[string]*pb.GameDuelService_ServerBroadcastServer
}

// New : Create new backend
func New(conf *cf.ConfTmp) *ULZGameDuelServiceBackend {
	ck := "RSCore" + cm.HashText(conf.APIServer.IP)
	rdfl := []*rd.RdsCliBox{}
	for i := 0; i < conf.CacheDb.WorkerNode; i++ {
		rdf := rd.New(ck, "wKU"+cm.HashText("num"+strconv.Itoa(i)))

		if cm.Mode == "prod" || cm.Mode == "Debug" {
			rdf.MarshalMethods = "proto"
		}

		if _, err := rdf.Connect(&conf.CacheDb); err == nil {
			rdfl = append(rdfl, rdf)
		}
	}

	g := ULZGameDuelServiceBackend{
		CoreKey:    ck,
		mu:         &sync.Mutex{},
		redhdlr:    rdfl,
		roomStream: make(map[string](*RoomStreamBox)),
	}
	// g.InitDB(&conf.Database)
	return &g
}

func (this *ULZGameDuelServiceBackend) Shutdown() {
	/// TODO: send closing msg to all client
	for _, v := range this.roomStream {
		log.Println("Server OS.sigKill")
		v.ClearAll()
	}
	log.Println("in shutdown proc")
	for _, v := range this.redhdlr {
		if _, err := v.CleanRem(); err != nil {
			log.Println(err)
		}
		if _, e := v.Disconn(); e != nil {
			log.Println(e)
		}
	}
	// this.CloseDB()
	log.Println("endof shutdown proc:", this.CoreKey)
}

// PrintReqLog

// ----------------------------------------------------------------------------------------------------
//

func (rm *ULZGameDuelServiceBackend) GetStream(roomKey *string, userId *string) *pb.GameDuelService_ServerBroadcastServer {
	a, ok := rm.roomStream[*roomKey]
	b, ok := a.clientConn[*userId]
	if ok {
		return b
	}
	return nil
}

func (rm *ULZGameDuelServiceBackend) AddStream(roomKey *string, userId *string, stream *pb.GameDuelService_ServerBroadcastServer) (bool, error) {
	// _, ok := rm.bc_stream[user_id]
	fmt.Println("RoomService.AddStream")
	fmt.Println(rm.roomStream)
	a, ok := rm.roomStream[*roomKey]
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
	log.Println("Del Stream:")
	a, ok := rm.roomStream[*roomKey]
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

func (rm *ULZGameDuelServiceBackend) BroadCast(roomkey *string, from *string, message *pb.GDBroadcastResp) error {
	log.Println("BS!", message)
	rmb, ok := rm.roomStream[*roomkey]
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

func (rm *RoomStreamBox) ClearAll() {
	log.Println("ClearAll Proc")
	// for _, vc := range rm.clientConn {
	// 	// (*vc).Send(cm.MsgSystShutdown(&rm.key))
	// }
	for k := range rm.clientConn {
		*(rm.clientConn[k]) = nil
		delete(rm.clientConn, k)
	}
	return
}

// -------------------------------------------------------------------------

func (b *ULZGameDuelServiceBackend) searchAliveClient() *rd.RdsCliBox {
	for {
		wk := b.checkAliveClient()
		if wk == nil {
			// log.Println("busy at " + time.Now().String())
			time.Sleep(500)
		} else {
			wk.Preserve(true)
			return wk
		}
	}
}

// checkAliveClient
func (b *ULZGameDuelServiceBackend) checkAliveClient() *rd.RdsCliBox {
	for _, v := range b.redhdlr {
		if !*v.IsRunning() {
			return v
		}
	}
	return nil
}

/// <<<=== Worker Goroutine function
