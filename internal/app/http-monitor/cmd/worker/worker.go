package worker

import (
	"log"
	"net/http"
	"time"

	"github.com/shirinebadi/http-monitor/internal/app/http-monitor/config"
	db "github.com/shirinebadi/http-monitor/internal/app/http-monitor/data/db/server"
	"github.com/shirinebadi/http-monitor/internal/app/http-monitor/data/nats"
)

func main(cfg config.Config) {
	con := nats.New(cfg)
	nats := nats.Nats{Cfg: cfg, Con: con}

	myDB, err := db.Init()
	if err != nil {
		log.Fatal("failed to setup db: %s", err.Error())
	}

	for {
		url, err := nats.Subscribe()

		if err != nil {
			time.Sleep(10 * time.Second)
			continue
		}

		dbI := db.Mydb{DB: myDB}

		address, err := dbI.SearchUrl(url.Url)

		resp, err := http.Get(address.Body)
		if err != nil {
			log.Print(err)
		}

		url.StatusCode = resp.StatusCode

		dbI.Update(&url)

	}

}
