package websocket

import (
	"fmt"

	pb "ULZGameDuelService/proto"

	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
)

type SocketConn struct {
	// The websocket SocketConn.
	ws *websocket.Conn
	// Buffered channel of outbound messages.
	send chan []byte
}

type SocketHub struct {
	// Registered clients.
	clients map[*SocketClient]bool
	// Inbound messages from the clients.
	broadcast chan []byte
	// server-side cell-update
	roomMsg chan *pb.GDBroadcastResp
	// Register requests from the clients.
	register chan *SocketClient
	// Unregister requests from clients.
	unregister chan *SocketClient
}

func NewHub() *SocketHub {
	return &SocketHub{
		broadcast:  make(chan []byte),
		register:   make(chan *SocketClient),
		unregister: make(chan *SocketClient),
		roomMsg:    make(chan *pb.GDBroadcastResp),
		clients:    make(map[*SocketClient]bool),
	}
}

// // wraper to gin handler
// func Wrapfunc(rsb *ULZRoomServiceBackend, hub *SocketHub) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		serveWs(rsb, hub, c)
// 	}
// }

func (h *SocketHub) Run() {
	for {
		select {
		case client := <-h.register:
			fmt.Printf("\nclient regi :%#v \n", client)
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}

		case cellResp := <-h.roomMsg:
			fmt.Printf("[WS:CellCast] Receive Cast\n\t- Message : %#v\n", cellResp)
			msgpt, _ := proto.Marshal(cellResp)
			for cli := range h.clients {
				if cli.roomKey == cellResp.RoomKey {
					fmt.Println("have one :", cli.roomKey)
					select {
					case cli.send <- msgpt:
					default:
						close(cli.send)
						delete(h.clients, cli)
					}
				}
			}
		}
	}
}

func (h *SocketHub) Broadcast(msg *pb.GDBroadcastResp) {
	fmt.Printf("[rsb:Broadcast]\n\t%#v\n", msg)
	h.roomMsg <- msg
}
