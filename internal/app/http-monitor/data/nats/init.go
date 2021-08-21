package nats

import (
	"log"

	"github.com/nats-io/nats.go"
	"github.com/shirinebadi/http-monitor/internal/app/http-monitor/config"
)

func New(n config.Config) nats.JetStreamContext {
	nc, _ := nats.Connect(n.Nats.Host)
	js, _ := nc.JetStream()
	stream, err := js.StreamInfo(n.Nats.Name)
	if err != nil {
		log.Println(err)
	}
	if stream == nil {
		log.Printf("creating stream %q and subjects %q", n.Nats.Name, n.Nats.Topic)
		_, err = js.AddStream(&nats.StreamConfig{
			Name:     n.Nats.Name,
			Subjects: []string{n.Nats.Topic},
		})
		if err != nil {
			log.Fatal(err)
			return nil
		}

	}
	return js
}
