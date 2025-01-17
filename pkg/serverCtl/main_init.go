package serverCtl

import (
	// _ "ULZGameDuelService"
	cm "ULZGameDuelService/pkg/common"
	cf "ULZGameDuelService/pkg/config"
	sr "ULZGameDuelService/pkg/scriptRunner"
	rd "ULZGameDuelService/pkg/store/redis"
	pb "ULZGameDuelService/proto"
	"fmt"
	"log"
	"strconv"
	"sync"
	"time"

	nats "github.com/nats-io/nats.go"
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
	natscli     *nats.Conn
	skillClient *sr.SkillEffectSvcClient
	// castServer  *ws.SocketHub
}

func (backend *ULZGameDuelServiceBackend) ServiceName() string {
	return "ULZ.GDSvc"
}

type RoomStreamBox struct {
	clientConn map[string]*pb.GameDuelService_ServerBroadcastServer
}

// New : Create new backend
func New(conf *cf.ConfTmp) *ULZGameDuelServiceBackend {
	ck := "ULZ.GDSvc." + cm.HashText(conf.APIServer.IP+time.Now().String())
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

	sd := nats.Options{
		Url:            fmt.Sprintf("%s://%s:%v", conf.NatsConn.Connector, conf.NatsConn.Host, conf.NatsConn.Port),
		AllowReconnect: true,
		MaxReconnect:   10,
		ReconnectWait:  5 * time.Second,
		Timeout:        1 * time.Second,
	}
	nc, err := sd.Connect()
	if err != nil {
		log.Fatalln(err)
	}
	g := ULZGameDuelServiceBackend{
		CoreKey:     ck,
		mu:          &sync.Mutex{},
		redhdlr:     rdfl,
		natscli:     nc,
		skillClient: skc,
		// castServer:  ws.NewHub(),
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
	this.skillClient.ClientClose()
	this.natscli.Close()

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
			time.Sleep(100)
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
