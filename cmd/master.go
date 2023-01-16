package cmd

import (
	"github.com/ScoreTrak/ScoreTrak/pkg/auth"
	"github.com/ScoreTrak/ScoreTrak/pkg/platform"
	"github.com/ScoreTrak/ScoreTrak/pkg/policy"
	"github.com/ScoreTrak/ScoreTrak/pkg/policy/policyclient"
	"github.com/ScoreTrak/ScoreTrak/pkg/queue"
	"github.com/ScoreTrak/ScoreTrak/pkg/report/reportclient"
	"github.com/ScoreTrak/ScoreTrak/pkg/runner"
	"github.com/ScoreTrak/ScoreTrak/pkg/server/grpc/grpcfx"
	"github.com/ScoreTrak/ScoreTrak/pkg/storage/seed"
	"github.com/ScoreTrak/ScoreTrak/pkg/storage/storagefx"
	"github.com/ScoreTrak/ScoreTrak/pkg/telemetry/telemetryfx"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
	"log"
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
				NewStaticConfig,
				NewDynamicConfig,
				NewStorageConfig,
				NewQueueConfig,
				NewPlatformConfig,
				NewMasterQueueConfig,
				NewJWTConfig,
				NewServerConfig,
			),

			// Observability
			telemetryfx.Module,

			// Create database components
			storagefx.Module,
			fx.Invoke(seed.SeedDB),

			// Create queueing components
			fx.Provide(queue.NewMasterStreamPubSub),
			fx.Provide(queue.NewWorkerQueue),

			// Create policy and report clients
			fx.Provide(
				policy.NewPolicy,
				policyclient.NewPolicyClient,
				reportclient.NewReportClient,
			),

			// Create platform module. Responsible for creating works in docker, docker swarm, and kubernetes
			fx.Provide(platform.NewPlatform),

			// Create auth components
			fx.Provide(
				auth.NewJWTManager,
				auth.NewAuthInterceptor,
			),

			// Create grpc server
			grpcfx.Module,

			// Create connect server
			//connectfx.Module,

			// Create runner
			fx.Provide(runner.NewRunner),

			// Register lifecycle hooks for the server, runner, policy/report client
			fx.Invoke(policyclient.InitPolicyClient, reportclient.InitReportClient),
			fx.Invoke(runner.InitRunner),
		)

		app.Run()
	},
}

func init() {
	rootCmd.AddCommand(masterCmd)
}
