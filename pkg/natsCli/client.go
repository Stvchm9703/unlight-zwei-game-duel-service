package natscli

import (
	nats "github.com/nats-io/nats.go"
	// pb "ULZGameDuelService/proto"
)

type NatsCli struct {
	coreCli  string
	natsConn *nats.Conn
}

/**
 *
 *
 *
 * sc, _ := stan.Connect(clusterID, clientID)

// Simple Synchronous Publisher
sc.Publish("foo", []byte("Hello World")) // does not return until an ack has been received from NATS Streaming

// Simple Async Subscriber
sub, _ := sc.Subscribe("foo", func(m *stan.Msg) {
    fmt.Printf("Received a message: %s\n", string(m.Data))
})

// Unsubscribe
sub.Unsubscribe()

// Close connection
sc.Close()
*/

func NewNatsCli(address string, clientId string) (*NatsCli, error) {
	nc, _ := nats.Connect(address)
	return &NatsCli{
		coreCli:  clientId,
		natsConn: nc,
	}, nil
}

func (nc *NatsCli) Broadcast(roomKey string, msg []byte) {
	nc.natsConn.Publish(roomKey, msg)
}

func (nc *NatsCli) Close() {
	nc.natsConn.Close()
}
