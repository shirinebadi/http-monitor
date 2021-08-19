package nats

import (
	"log"

	"github.com/nats-io/nats.go"
	"github.com/shirinebadi/http-monitor/internal/app/http-monitor/config"
	"github.com/shirinebadi/http-monitor/internal/app/http-monitor/model"
)

type Nats struct {
	Cfg config.Config
	Con *nats.EncodedConn
}

func (n *Nats) Publish(s model.Status) {

	err := n.Con.Publish(n.Cfg.Nats.Topic, s)
	if err != nil {
		log.Fatal(err)
	}
}

func (n *Nats) Subscribe() model.Status {

	newReq := model.Status{}

	if _, err := n.Con.Subscribe(n.Cfg.Nats.Topic, func(s model.Status) {

		log.Print(s.ID, " Delivered to Worker")
		newReq = s
	}); err != nil {

		log.Fatal(err)
	}

	return newReq

}
