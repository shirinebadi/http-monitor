package scheduler

import (
	"log"
	"time"

	"github.com/shirinebadi/http-monitor/internal/app/http-monitor/config"
	db "github.com/shirinebadi/http-monitor/internal/app/http-monitor/data/db/server"
	"github.com/shirinebadi/http-monitor/internal/app/http-monitor/data/nats"
	"github.com/shirinebadi/http-monitor/internal/app/http-monitor/model"
	"github.com/spf13/cobra"
)

func main(cfg config.Config) {
	con := nats.New(cfg)
	nats := nats.Nats{Cfg: cfg, Con: con}

	urls := make([]model.Status, 0)
	counter := 1

	myDB, err := db.Init()
	if err != nil {
		log.Fatal("failed to setup db: ", err.Error())
	}

	dbI := db.Mydb{DB: myDB}

	firstId, _ := dbI.GetFirst()

	urls = append(urls, firstId)

	go func() {
		for {

			url, err := dbI.GetRecent(firstId.ID)

			if err != nil {
				time.Sleep(10 * time.Second)
				continue
			}

			firstId = url

			log.Print("Adding to urls: ", url)

			urls = append(urls, url)

		}

	}()

	go func() {
		for {

			for i, url := range urls {

				if counter == cfg.Common.Period {

					log.Print("done: ", url)

					nats.Publish(url)

					if i == len(urls)-1 {
						counter = 1
					}

				}

			}

			time.Sleep(2 * time.Second)

			counter++
		}
	}()

	for {

	}
}

func Register(root *cobra.Command, cfg config.Config) {
	runScheduler := &cobra.Command{
		Use:   "scheduler",
		Short: "scheduler for http monitor",
		Run: func(cmd *cobra.Command, args []string) {
			main(cfg)
		},
	}

	root.AddCommand(runScheduler)
}
