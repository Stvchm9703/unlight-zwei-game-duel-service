package natsCli 

import (
	stan "github.com/nats-io/stan.go"
)

type NatsCli struct {
	coreCli string
	nats stan.NatsCli
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

func NewNatsCli(key string) (*NatsCli, error){
	sc,_ := stan.Connect("", "")
	return &NatsCli, nil
}