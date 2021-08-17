package server

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/shirinebadi/http-monitor/internal/app/http-monitor/config"
	db "github.com/shirinebadi/http-monitor/internal/app/http-monitor/db/server"
	"github.com/shirinebadi/http-monitor/internal/app/http-monitor/handler"
	"github.com/spf13/cobra"
)

func main(cfg config.Config) {

	myDB, err := db.Init()
	if err != nil {
		log.Fatal("failed to setup db: %s", err.Error())
	}

	userI := db.Mydb{DB: myDB}
	e := echo.New()
	user := handler.UserHandler{UserI: &userI}

	e.POST("/register", user.Register)
	e.POST("/login", user.Login)

	address := cfg.Server.Address

	if err := e.Start(address); err != nil {
		log.Fatal(err)
	}

}

func Register(root *cobra.Command, cfg config.Config) {
	print("bye")
	runServer := &cobra.Command{
		Use:   "server",
		Short: "server for container scheduling",
		Run: func(cmd *cobra.Command, args []string) {
			main(cfg)
		},
	}

	root.AddCommand(runServer)
}
