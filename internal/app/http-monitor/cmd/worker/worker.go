package worker

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/shirinebadi/http-monitor/internal/app/http-monitor/config"
	db "github.com/shirinebadi/http-monitor/internal/app/http-monitor/data/db/server"
	"github.com/shirinebadi/http-monitor/internal/app/http-monitor/data/nats"
	"github.com/spf13/cobra"
)

func main(cfg config.Config) {
	con := nats.New(cfg)
	nats := nats.Nats{Cfg: cfg, Con: con}

	myDB, err := db.Init()
	if err != nil {
		log.Fatal("failed to setup db: %s", err.Error())
	}

	for {
		url := nats.Subscribe()
		fmt.Print("url is: ", url)

		dbI := db.Mydb{DB: myDB}

		address, err := dbI.SearchUrl(url.Url)

		if err != nil {
			time.Sleep(10 * time.Second)
			continue
		}

		fmt.Print("hi", address.Body)

		time.Sleep(10 * time.Second)

		for i := 0; i < 10; i++ {
			go func() {
				resp, err := http.Get(address.Body)
				if err != nil {
					log.Print(err)
				}

				url.StatusCode = append(url.StatusCode, int32(resp.StatusCode))

				dbI.Update(&url)
			}()
		}

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
