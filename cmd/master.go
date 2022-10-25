package cmd

import (
	"context"
	"errors"
	"fmt"
	"github.com/ScoreTrak/ScoreTrak/pkg/auth"
	"github.com/ScoreTrak/ScoreTrak/pkg/config"
	"github.com/ScoreTrak/ScoreTrak/pkg/handler/handlerfx"
	"github.com/ScoreTrak/ScoreTrak/pkg/platform"
	"github.com/ScoreTrak/ScoreTrak/pkg/policy"
	"github.com/ScoreTrak/ScoreTrak/pkg/policy/policyclient"
	"github.com/ScoreTrak/ScoreTrak/pkg/queue"
	"github.com/ScoreTrak/ScoreTrak/pkg/report"
	"github.com/ScoreTrak/ScoreTrak/pkg/report/reportclient"
	"github.com/ScoreTrak/ScoreTrak/pkg/runner"
	"github.com/ScoreTrak/ScoreTrak/pkg/storage/seed"
	"github.com/ScoreTrak/ScoreTrak/pkg/storage/storagefx"
	"github.com/ScoreTrak/ScoreTrak/pkg/telemetry/telemetryfx"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/spf13/cobra"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
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

			// Create starter objects
			fx.Supply(&policy.Policy{ID: 1}),
			fx.Supply(&report.Report{ID: 1, Cache: "{}"}),

			// Create policy and report clients
			fx.Provide(
				policyclient.NewPolicyClient,
				reportclient.NewReportClient,
			),

			// Create stuff
			fx.Provide(
				platform.NewPlatform,
				auth.NewJWTManager,
				auth.NewAuthInterceptor,
			),

			// Create server components
			fx.Provide(NewGrpcServer),
			handlerfx.Module,

			// Runner
			fx.Provide(runner.NewRunner),

			// Register Lifecycle hooks for the server, runner, policy/report client
			fx.Invoke(InitGrpcServer),
			fx.Invoke(runner.InitRunner),
			fx.Invoke(policyclient.InitPolicyClient, reportclient.InitReportClient),
		)

		app.Run()
	},
}

func init() {
	rootCmd.AddCommand(masterCmd)
}

func NewGrpcServer(staticConfig config.StaticConfig, logger *zap.Logger, authInterceptor *auth.Interceptor) (*grpc.Server, error) {

	var server *grpc.Server

	var serverOptions []grpc.ServerOption

	var ErrProdCertMissing = errors.New("production requires certfile, and keyfile")
	if staticConfig.CertFile != "" && staticConfig.KeyFile != "" {
		creds, err := credentials.NewClientTLSFromFile(staticConfig.CertFile, staticConfig.KeyFile)
		if err != nil {
			return nil, err
		}
		serverOptions = append(serverOptions, grpc.Creds(creds))
	} else if staticConfig.Prod {
		return nil, ErrProdCertMissing
	}

	var unaryServerInterceptors []grpc.UnaryServerInterceptor
	var streamServerInterceptors []grpc.StreamServerInterceptor

	// Logging
	{
		logOpts := []grpc_zap.Option{
			grpc_zap.WithLevels(grpc_zap.DefaultCodeToLevel),
		}
		grpc_zap.ReplaceGrpcLoggerV2(logger)
		unaryServerInterceptors = append(unaryServerInterceptors, []grpc.UnaryServerInterceptor{
			grpc_ctxtags.UnaryServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
			grpc_zap.UnaryServerInterceptor(logger, logOpts...),
		}...)

		streamServerInterceptors = append(streamServerInterceptors, []grpc.StreamServerInterceptor{
			grpc_ctxtags.StreamServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
			grpc_zap.StreamServerInterceptor(logger, logOpts...),
		}...)
	}

	// Recovery
	{
		recoveryOpts := []grpc_recovery.Option{
			grpc_recovery.WithRecoveryHandler(func(p interface{}) (err error) {
				return status.Errorf(codes.Unknown, "panic triggered: %v", p)
			}),
		}
		if staticConfig.Prod {
			unaryServerInterceptors = append(unaryServerInterceptors, grpc_recovery.UnaryServerInterceptor(recoveryOpts...))
			streamServerInterceptors = append(streamServerInterceptors, grpc_recovery.StreamServerInterceptor(recoveryOpts...))
		}
	}

	// Auth
	{
		unaryServerInterceptors = append(unaryServerInterceptors, authInterceptor.Unary())
		streamServerInterceptors = append(streamServerInterceptors, authInterceptor.Stream())
	}

	// Observability
	{
		unaryServerInterceptors = append(unaryServerInterceptors, otelgrpc.UnaryServerInterceptor())
		streamServerInterceptors = append(streamServerInterceptors, otelgrpc.StreamServerInterceptor())
	}

	// Chaining
	{
		serverOptions = append(serverOptions, grpc_middleware.WithUnaryServerChain(unaryServerInterceptors...))
		serverOptions = append(serverOptions, grpc_middleware.WithStreamServerChain(streamServerInterceptors...))
	}

	// New GRPC Server
	server = grpc.NewServer(serverOptions...)

	// Reflection
	if !staticConfig.Prod {
		reflection.Register(server)
	}

	return server, nil
}

func InitGrpcServer(lc fx.Lifecycle, staticConfig config.StaticConfig, server *grpc.Server) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", staticConfig.Port))
	if err != nil {
		log.Panicf("Failed to listen: %v", err)
	}

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
