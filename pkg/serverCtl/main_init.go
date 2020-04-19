package serverCtl

import (
	// _ "ULZGameDuelService"
	cm "ULZGameDuelService/pkg/common"
	cf "ULZGameDuelService/pkg/config"
	sr "ULZGameDuelService/pkg/scriptRunner"
	rd "ULZGameDuelService/pkg/store/redis"
	ws "ULZGameDuelService/pkg/websocket"
	pb "ULZGameDuelService/proto"
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
	mu          *sync.Mutex
	CoreKey     string
	redhdlr     []*rd.RdsCliBox
	castServer  *ws.SocketHub
	skillClient *sr.SkillEffectSvcClient
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
	skc := sr.ClientListInit("ScriptRunner", conf.EffectCalcService)
	g := ULZGameDuelServiceBackend{
		CoreKey:     ck,
		mu:          &sync.Mutex{},
		redhdlr:     rdfl,
		castServer:  ws.NewHub(),
		skillClient: skc,
	}
	// g.InitDB(&conf.Database)
	return &g
}

func (this *ULZGameDuelServiceBackend) Shutdown() {
	/// TODO: send closing msg to all client
	// for _, v := range this.roomStream {
	// 	log.Println("Server OS.sigKill")
	// 	v.ClearAll()
	// }

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
//	Skill-client
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
