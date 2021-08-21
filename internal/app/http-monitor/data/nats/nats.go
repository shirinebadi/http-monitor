package nats

import (
	"encoding/json"
	"log"

	"github.com/nats-io/nats.go"
	"github.com/shirinebadi/http-monitor/internal/app/http-monitor/config"
	"github.com/shirinebadi/http-monitor/internal/app/http-monitor/model"
)

type Nats struct {
	Cfg  config.Config
	Cn   nats.JetStreamContext
	Jobs chan *model.Status
}

func (n *Nats) Publish(s *model.Status) {
	StatusJson, _ := json.Marshal(&s)
	_, err := n.Cn.Publish(n.Cfg.Nats.Topic, StatusJson)
	if err != nil {
		log.Fatal(err)
	}
}

func (n *Nats) Subscribe() {

	if _, err := n.Cn.Subscribe(n.Cfg.Nats.Topic, func(msg *nats.Msg) {
		var s model.Status
		msg.Ack()

		err := json.Unmarshal(msg.Data, &s)
		if err != nil {
			log.Fatal("Error in nats subscribe: ", err)
		}

		log.Print(s.ID, " Delivered to Worker")

		n.Jobs <- &s

	}, nats.Durable("monitor"), nats.ManualAck()); err != nil {
		log.Fatal(err)
	}

}
