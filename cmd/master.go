package cmd

import (
	"github.com/ScoreTrak/ScoreTrak/pkg/auth/authfx"
	"github.com/ScoreTrak/ScoreTrak/pkg/queue/queuefx"
	"log"

	"github.com/ScoreTrak/ScoreTrak/pkg/config/configfx"
	"github.com/ScoreTrak/ScoreTrak/pkg/server/serverfx"
	"github.com/ScoreTrak/ScoreTrak/pkg/storage/storagefx"
	"github.com/ScoreTrak/ScoreTrak/pkg/telemetry/telemetryfx"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

// masterCmd represents the master command
var masterCmd = &cobra.Command{
	Use:   "master",
	Short: "master runs the grpc server and runner if in single-node mode",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("starting master server")

		app := fx.New(
			// Create configs
			configfx.Module,

			// Telemetry
			telemetryfx.Module,

			// Create database components
			storagefx.Module,

			// Add auth components
			authfx.Module,

			// Create server
			serverfx.Module,

			// Create queueing components
			queuefx.Module,
			// fx.Provide(queue.NewMasterStreamPubSub),
			// fx.Provide(queue.NewWorkerQueue),

			// Create platform module. Responsible for creating works in docker, docker swarm, and kubernetes
			// fx.Provide(platform.NewPlatform),

			// Create scheduler
			// schedulerfx.Module,
		)

		app.Run()
	},
}

func init() {
	rootCmd.AddCommand(masterCmd)
}
