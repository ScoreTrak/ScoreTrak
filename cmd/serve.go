/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/ScoreTrak/ScoreTrak/pkg/auth/authfx"
	"github.com/ScoreTrak/ScoreTrak/pkg/config/configfx"
	"github.com/ScoreTrak/ScoreTrak/pkg/server/serverfx"
	"github.com/ScoreTrak/ScoreTrak/pkg/storage/storagefx"
	"github.com/ScoreTrak/ScoreTrak/pkg/telemetry/telemetryfx"
	"go.uber.org/fx"

	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start server",
	Run: func(cmd *cobra.Command, args []string) {
		app := fx.New(
			// Create configs
			configfx.Module,

			// Telemetry
			telemetryfx.Module,

			// Create database components
			storagefx.Module,

			// Add auth components
			authfx.Module,

			// Create queueing components
			// fx.Provide(queue.NewMasterStreamPubSub),
			// fx.Provide(queue.NewWorkerQueue),

			// Create platform module. Responsible for creating works in docker, docker swarm, and kubernetes
			// fx.Provide(platform.NewPlatform),

			// Create server
			serverfx.Module,

			// Create scheduler
			// schedulerfx.Module,
		)

		app.Run()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
