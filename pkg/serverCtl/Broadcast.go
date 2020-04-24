package serverCtl

import (
	cm "ULZGameDuelService/pkg/common"
	ws "ULZGameDuelService/pkg/websocket"
	pb "ULZGameDuelService/proto"
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gogo/status"
	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc/codes"
	// "google."
	// "time"
)

func (this *ULZGameDuelServiceBackend) ServerBroadcast(rReq *pb.GDGetInfoReq, stream pb.GameDuelService_ServerBroadcastServer) error {
	cm.PrintReqLog(nil, "server-broadcast", rReq)
	return status.Error(codes.Internal, "UseNAtsConn:")
}

func (this *ULZGameDuelServiceBackend) SendMessage(ctx context.Context, msg *pb.GDBroadcastResp) (*pb.Empty, error) {
	cm.PrintReqLog(nil, "server-broadcast:msg", msg)
	this.BroadCast(msg)
	return &pb.Empty{}, nil
}

func (rsb *ULZGameDuelServiceBackend) BroadCast(cp *pb.GDBroadcastResp) error {
	// rsb.castServer.Broadcast(cp)
	msgpt, err := proto.Marshal(cp)
	if err != nil {
		return err
	}
	rsb.natscli.Publish(cp.RoomKey, msgpt)
	return nil
}

// func (rsb *ULZGameDuelServiceBackend) RunWebSocketServer(config cf.CfAPIServer) error {
// 	hub := ws.NewHub()
// 	go hub.Run()
// 	router := gin.New()
// 	router.GET("/:roomId", Wrapfunc(rsb, hub))
// 	rsb.castServer = hub
// 	return router.Run(config.IP + ":" + strconv.Itoa(config.PollingPort))
// }

// wraper to gin handler
func Wrapfunc(rsb *ULZGameDuelServiceBackend, hub *ws.SocketHub) gin.HandlerFunc {
	return func(c *gin.Context) {
		serveWs(rsb, hub, c)
	}
}

// serveWs handles websocket requests from the peer.
func serveWs(rsb *ULZGameDuelServiceBackend, hub *ws.SocketHub, c *gin.Context) {
	conn, err := ws.Upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(c.Param("roomId"))
	reqKey := c.Param("roomId")

	var tmp pb.GameDataSet
	wkbox := rsb.searchAliveClient()
	if _, err := wkbox.GetPara(reqKey, &tmp); err != nil {
		c.AbortWithStatus(412)
	}
	client := ws.NewClient(reqKey, hub, conn)
	go client.WritePump()
	go client.ReadPump()
}
