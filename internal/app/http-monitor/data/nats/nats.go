package nats

import (
	"log"

	"github.com/nats-io/nats.go"
	"github.com/shirinebadi/http-monitor/internal/app/http-monitor/config"
	"github.com/shirinebadi/http-monitor/internal/app/http-monitor/model"
)

type Nats struct {
	Cfg config.Config
	Con *nats.Conn
}

type NatsI interface {
	Publish(u model.Url)
}

func (n *Nats) Publish(u model.Url) {
	ec, err := nats.NewEncodedConn(n.Con, nats.GOB_ENCODER)
	if err != nil {
		log.Fatal(err)
	}

	err = ec.Publish(n.Cfg.Nats.Topic, u)
	if err != nil {
		log.Fatal(err)
	}
}

func (n *Nats) Subscribe() {
	c, err := nats.NewEncodedConn(n.Con, nats.GOB_ENCODER)
	if err != nil {
		log.Fatal(err)
	}

	defer c.Close()

}
