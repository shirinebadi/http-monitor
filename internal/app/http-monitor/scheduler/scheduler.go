package scheduler

import (
	"fmt"
	"log"
	"time"

	"github.com/shirinebadi/http-monitor/internal/app/http-monitor/config"
	"github.com/shirinebadi/http-monitor/internal/app/http-monitor/data/nats"
	"github.com/shirinebadi/http-monitor/internal/app/http-monitor/model"
)

type Scheduler struct {
	Cfg  config.Config
	Jobs chan model.Status
}

func (s *Scheduler) Run() {
	counter := 1

	cn := nats.New(s.Cfg)
	nats := nats.Nats{Cfg: s.Cfg, Cn: cn}

	urls := make([]model.Status, 0)

	go func() {
		for {
			deliveredJob := <-s.Jobs

			log.Print("New Url Added:\n ", deliveredJob)

			urls = append(urls, deliveredJob)

		}
	}()

	go func() {
		for {

			if counter > s.Cfg.Common.Period {
				counter = 1
			}

			for i, url := range urls {
				fmt.Println(counter, " ", i)

				if counter == s.Cfg.Common.Period {

					nats.Publish(&url)

				}

			}

			time.Sleep(1 * time.Second)
			counter++
		}
	}()
}
