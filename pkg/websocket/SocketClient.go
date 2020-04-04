package websocket

import (
	"log"
	"time"

	"github.com/gorilla/websocket"
)

// SocketClient is a middleman between the websocket SocketConn and the hub.
type SocketClient struct {
	hub     *SocketHub
	conn    *websocket.Conn
	roomKey string
	send    chan []byte
}

func NewClient(RoomKey string, Hub *SocketHub, Connection *websocket.Conn) *SocketClient {
	cli := &SocketClient{
		roomKey: RoomKey,
		hub:     Hub,
		conn:    Connection,
		send:    make(chan []byte, 256),
	}
	cli.hub.register <- cli
	return cli
}

// readPump pumps messages from the websocket SocketConn to the hub.
//
// The application runs readPump in a per-SocketConn goroutine. The application
// ensures that there is at most one reader on a SocketConn by executing all
// reads from this goroutine.
func (c *SocketClient) ReadPump() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()
	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error {
		c.conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})
	for {
		// Note: Igore client send out
		_, _, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		// message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		// c.hub.broadcast <- message
	}
}

// writePump pumps messages from the hub to the websocket SocketConn.
//
// A goroutine running writePump is started for each SocketConn. The
// application ensures that there is at most one writer to a SocketConn by
// executing all writes from this goroutine.
func (c *SocketClient) WritePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(websocket.BinaryMessage)
			if err != nil {
				return
			}
			w.Write(message)

			// Add queued chat messages to the current websocket message.
			n := len(c.send)
			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write(<-c.send)
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}
