package cmd

import (
	"context"
	"fmt"
	"github.com/ScoreTrak/ScoreTrak/pkg/auth"
	"github.com/ScoreTrak/ScoreTrak/pkg/config"
	"github.com/ScoreTrak/ScoreTrak/pkg/handler/handlerfx"
	"github.com/ScoreTrak/ScoreTrak/pkg/platform"
	"github.com/ScoreTrak/ScoreTrak/pkg/policy"
	"github.com/ScoreTrak/ScoreTrak/pkg/policy/policyclient"
	"github.com/ScoreTrak/ScoreTrak/pkg/queue"
	"github.com/ScoreTrak/ScoreTrak/pkg/report/reportclient"
	"github.com/ScoreTrak/ScoreTrak/pkg/runner"
	"github.com/ScoreTrak/ScoreTrak/pkg/storage/seed"
	"github.com/ScoreTrak/ScoreTrak/pkg/storage/storagefx"
	"github.com/ScoreTrak/ScoreTrak/pkg/telemetry/telemetryfx"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

// masterCmd represents the master command
var masterCmd = &cobra.Command{
	Use:   "master",
	Short: "master runs the grpc server and runner if in single-node mode",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("starting master server")

		app := fx.New(
			// Create configs
			fx.Provide(NewStaticConfig, NewDynamicConfig, NewStorageConfig, NewQueueConfig, NewPlatformConfig, NewMasterQueueConfig, NewJWTConfig),

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

			// Create platform
			fx.Provide(platform.NewPlatform),

			// Create auth deps
			fx.Provide(auth.NewJWTManager, auth.NewAuthInterceptor),

			// Create grpc server
			fx.Provide(),
			handlerfx.GrpcModule,

			// Create connect server

			// Create runner
			fx.Provide(runner.NewRunner),

			// Register Lifecycle hooks for the server, runner, policy/report client
			fx.Invoke(policyclient.InitPolicyClient, reportclient.InitReportClient),
			fx.Invoke(runner.InitRunner),
			fx.Invoke(InitGrpcServer),
		)

		app.Run()
	},
}

func init() {
	rootCmd.AddCommand(masterCmd)
}

func InitGrpcServer(lc fx.Lifecycle, staticConfig config.StaticConfig, server *grpc.Server) {
	address := fmt.Sprintf("%s:%s", staticConfig.Server.Address, staticConfig.Server.Port)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Panicf("Failed to listen: %v", err)
	}
	log.Printf("Created Listener at %s", address)

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			log.Println("Starting grpc server")
			go func() {
				err := server.Serve(lis)
				if err != nil {
					time.Sleep(time.Second)
					log.Panicf("%v", err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Println("Stopping grpc server")
			err := lis.Close()
			if err != nil {
				return err
			}
			server.Stop()
			return nil
		},
	})
}
