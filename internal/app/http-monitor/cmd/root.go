package cmd

import (
	"github.com/shirinebadi/http-monitor/internal/app/http-monitor/cmd/scheduler"
	"github.com/shirinebadi/http-monitor/internal/app/http-monitor/cmd/server"
	"github.com/shirinebadi/http-monitor/internal/app/http-monitor/cmd/worker"
	"github.com/shirinebadi/http-monitor/internal/app/http-monitor/config"
	"github.com/spf13/cobra"
)

func NewRootCommand() *cobra.Command {
	var root = &cobra.Command{
		Use: "http-monitor",
	}
	cfg := config.Init()

	server.Register(root, cfg)
	scheduler.Register(root, cfg)
	worker.Register(root, cfg)

	return root
}
