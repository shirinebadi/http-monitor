package nats

import (
	"log"

	"github.com/nats-io/nats.go"
	"github.com/shirinebadi/http-monitor/internal/app/http-monitor/config"
)

func New(n config.Config) *nats.Conn {
	print(n.Nats.Host)
	nc, err := nats.Connect(n.Nats.Host)
	if err != nil {
		log.Fatal(err)
	}

	return nc
}
