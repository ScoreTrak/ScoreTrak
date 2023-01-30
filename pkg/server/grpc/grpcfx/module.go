package grpcfx

import (
	"github.com/ScoreTrak/ScoreTrak/pkg/handler/handlerfx"
	sgrpc "github.com/ScoreTrak/ScoreTrak/pkg/server/grpc"
	"go.uber.org/fx"
	"google.golang.org/grpc"
)

var Module = fx.Options(
	fx.Provide(
		sgrpc.NewGrpcServer,
		func(server *grpc.Server) grpc.ServiceRegistrar { return server },
	),
	handlerfx.GrpcModule,
	fx.Invoke(sgrpc.InitGrpcServer),
)
