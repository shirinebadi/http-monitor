package worker

import (
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
	cn := nats.New(cfg)
	nats := nats.Nats{Cfg: cfg, Cn: cn, Jobs: ch}

	myDB, err := db.Init()
	if err != nil {
		log.Fatal("failed to setup db: ", err.Error())
	}

	go func() {
		nats.Subscribe()
	}()

	go func() {

		dbI := db.Mydb{DB: myDB}

		for {

			u := <-ch
			*u, _ = dbI.GetFirst(u.ID)

			address, err := dbI.SearchUrl(u.Url)

			if err != nil {
				time.Sleep(10 * time.Second)
				continue
			}

			for i := 0; i < 10; i++ {
				go func() {
					resp, err := http.Get(address.Body)
					if err != nil {
						u.StatusCode = append(u.StatusCode, int32(500))
					} else {
						u.StatusCode = append(u.StatusCode, int32(resp.StatusCode))
					}

					if err := dbI.Update(u); err != nil {
						log.Print(err)
					}

				}()
			}
		}

	}()
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
