package server

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/shirinebadi/http-monitor/internal/app/http-monitor/config"
	db "github.com/shirinebadi/http-monitor/internal/app/http-monitor/data/db/server"
	"github.com/shirinebadi/http-monitor/internal/app/http-monitor/handler"
	"github.com/spf13/cobra"
)

func main(cfg config.Config) {

	myDB, err := db.Init()
	if err != nil {
		log.Fatal("failed to setup db: %s", err.Error())
	}

	e := echo.New()

	userI := db.Mydb{DB: myDB}

	token := handler.Token{Cfg: cfg}

	urlI := db.Mydb{DB: myDB}

	user := handler.UserHandler{UserI: &userI, Token: token}
	url := handler.UrlHandler{RequestI: &userI, UrlI: &urlI, Token: token}

	e.POST("/register", user.Register)
	e.POST("/login", user.Login)
	e.POST("/request", url.Send)

	address := cfg.Server.Address

	if err := e.Start(address); err != nil {
		log.Fatal(err)
	}

}

func Register(root *cobra.Command, cfg config.Config) {
	runServer := &cobra.Command{
		Use:   "server",
		Short: "server for container scheduling",
		Run: func(cmd *cobra.Command, args []string) {
			main(cfg)
		},
	}

	root.AddCommand(runServer)
}
