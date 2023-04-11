package cmd

import (
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
			fx.Provide(
				fx.Annotate(
					func() string {
						return cfgFile
					},
					fx.ResultTags(`name:"cfgFile"`),
				),
			),
			configfx.Module,

			// Telemetry
			telemetryfx.Module,

			// Create database components
			storagefx.Module,

			// Create queueing components
			// fx.Provide(queue.NewMasterStreamPubSub),
			// fx.Provide(queue.NewWorkerQueue),

			// Create policy and report clients
			// fx.Provide(
			// 	policy.NewPolicy,
			// 	policyclient.NewPolicyClient,
			// 	reportclient.NewReportClient,
			// ),

			// Create platform module. Responsible for creating works in docker, docker swarm, and kubernetes
			// fx.Provide(platform.NewPlatform),

			// Create auth components
			// authfx.Module,

			// Create server
			serverfx.Module,

			// Create scheduler
			// schedulerfx.Module,

			// Register lifecycle hooks for the runner, policy/report client
			// fx.Invoke(policyclient.InitPolicyClient, reportclient.InitReportClient),
		)

		app.Run()
	},
}

func init() {
	rootCmd.AddCommand(masterCmd)
}
