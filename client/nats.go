package client

import (
	"log"
	"microservice-case2/common"
	"time"

	nats "github.com/nats-io/go-nats"
)

type NatsInterface interface {
	Publish(topic, msg string) error
	Request(topic, msg string) (*nats.Msg, error)
}

type Nats struct {
	conn *nats.Conn
}

func NewNatsClient(url string) NatsInterface {
	nc, err := nats.Connect(common.NatsURL)
	if err != nil {
		log.Fatalf(`Failed to connect to NATS broker. Error=%s`, err.Error())
	}

	return &Nats{
		conn: nc,
	}
}

func (n *Nats) Publish(topic, msg string) error {
	return n.conn.Publish(topic, []byte(msg))
}

func (n *Nats) Request(topic, msg string) (*nats.Msg, error) {
	return n.conn.Request(topic, []byte(msg), 1*time.Second)
}
