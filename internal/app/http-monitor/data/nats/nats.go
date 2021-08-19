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
	Publish(s model.Status)
	Subscribe() (model.Status, error)
}

func (n *Nats) Publish(s model.Status) {
	ec, err := nats.NewEncodedConn(n.Con, nats.GOB_ENCODER)
	if err != nil {
		log.Fatal(err)
	}

	err = ec.Publish(n.Cfg.Nats.Topic, s)
	if err != nil {
		log.Fatal(err)
	}
}

func (n *Nats) Subscribe() (model.Status, error) {
	c, err := nats.NewEncodedConn(n.Con, nats.GOB_ENCODER)
	if err != nil {
		log.Fatal(err)
	}

	newReq := model.Status{}

	defer c.Close()

	if _, err := c.QueueSubscribe(n.Cfg.Nats.Topic, n.Cfg.Nats.Queue, func(s model.Status) {
		log.Print(s.Id, " Delivered to Worker")
		newReq = s
	}); err != nil {
		log.Fatal(err)
	}

	return newReq, err

}
