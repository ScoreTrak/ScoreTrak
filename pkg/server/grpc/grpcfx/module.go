package grpcfx

import (
	stgrpc "github.com/ScoreTrak/ScoreTrak/pkg/server/grpc"
	"go.uber.org/fx"
	"google.golang.org/grpc"
)

var Module = fx.Options(
	fx.Provide(
		stgrpc.NewGrpcServer,
		func(server *grpc.Server) grpc.ServiceRegistrar { return server },
	),
	fx.Invoke(stgrpc.InitGrpcServer),
)
