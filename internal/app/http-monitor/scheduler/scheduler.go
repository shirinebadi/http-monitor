package scheduler

import (
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

	con := nats.New(s.Cfg)
	nats := nats.Nats{Cfg: s.Cfg, Con: con}

	urls := make([]model.Status, 0)

	go func() {
		for {

			deliveredJob := <-s.Jobs

			log.Print("New Url Added:\n ", deliveredJob)

			urls = append(urls, deliveredJob)

			for i, url := range urls {

				if counter == s.Cfg.Common.Period {

					nats.Publish(&url)

					if i == len(urls)-1 {
						counter = 1
					}

				}

			}

			time.Sleep(1 * time.Second)
			counter++
		}
	}()
}
