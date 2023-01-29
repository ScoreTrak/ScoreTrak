package grpc

import (
	"context"
	"errors"
	"fmt"
	"github.com/ScoreTrak/ScoreTrak/pkg/auth"
	"github.com/ScoreTrak/ScoreTrak/pkg/config"
	"github.com/ScoreTrak/ScoreTrak/pkg/server"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
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

func NewGrpcServer(staticConfig config.StaticConfig, logger *zap.Logger, authInterceptor *auth.Interceptor) (*grpc.Server, error) {

	var s *grpc.Server

	var serverOptions []grpc.ServerOption

	var ErrProdCertMissing = errors.New("production requires certfile, and keyfile")

	tls_enabled := staticConfig.Server.TLS.CertFile != "" || staticConfig.Server.TLS.KeyFile != ""

	if tls_enabled {
		creds, err := credentials.NewClientTLSFromFile(staticConfig.Server.TLS.CertFile, staticConfig.Server.TLS.KeyFile)
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
	s = grpc.NewServer(serverOptions...)

	// Reflection
	if !staticConfig.Prod {
		reflection.Register(s)
	}

	return s, nil
}

func InitGrpcServer(lc fx.Lifecycle, config server.Config, server *grpc.Server) {
	address := fmt.Sprintf("%s:%s", config.Address, config.Port)

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
