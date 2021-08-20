package worker

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/shirinebadi/http-monitor/internal/app/http-monitor/config"
	db "github.com/shirinebadi/http-monitor/internal/app/http-monitor/data/db"
	"github.com/shirinebadi/http-monitor/internal/app/http-monitor/data/nats"
	"github.com/shirinebadi/http-monitor/internal/app/http-monitor/model"
	"github.com/spf13/cobra"
)

func main(cfg config.Config) {
	ch := make(chan *model.Status)
	con := nats.New(cfg)
	nats := nats.Nats{Cfg: cfg, Con: con, Jobs: ch}

	myDB, err := db.Init()
	if err != nil {
		log.Fatal("failed to setup db: %s", err.Error())
	}

	go func() {
		nats.Subscribe()
		time.Sleep(20 * time.Second)
	}()

	go func() {
		for {

			for u := range ch {

				dbI := db.Mydb{DB: myDB}

				address, err := dbI.SearchUrl(u.Url)

				if err != nil {
					time.Sleep(10 * time.Second)
					continue
				}

				time.Sleep(10 * time.Second)

				for i := 0; i < 10; i++ {
					go func() {
						resp, err := http.Get(address.Body)
						if err != nil {
							log.Println(err)
						} else {
							u.StatusCode = append(u.StatusCode, int32(resp.StatusCode))
							fmt.Println(u)

							dbI.Update(u)
						}
					}()
				}
			}

		}

	}()
	for {
	}
}

func Register(root *cobra.Command, cfg config.Config) {
	runWorker := &cobra.Command{
		Use:   "worker",
		Short: "worker for http monitor",
		Run: func(cmd *cobra.Command, args []string) {
			main(cfg)
		},
	}

	root.AddCommand(runWorker)
}
