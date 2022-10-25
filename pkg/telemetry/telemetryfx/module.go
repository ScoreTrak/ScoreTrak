package telemetryfx

import (
	"github.com/ScoreTrak/ScoreTrak/pkg/telemetry"
	"go.uber.org/fx"
)

var OtlpModule = fx.Module("otlp",
	fx.Provide(
		telemetry.NewOtlpGrpcExporter,
		telemetry.NewOtlpGrpcTracerProvider,
	),
)

var JaegerModule = fx.Module("jaeger",
	fx.Provide(
		telemetry.NewJaegerExporter,
		telemetry.NewJaegerTracerProvider,
	),
)

var Module = fx.Options(
	// Logging
	fx.Provide(telemetry.NewLogger),
	// Tracing
	fx.Provide(telemetry.NewResource),

	//     Otlp
	OtlpModule,
	//     Jaeger
	//JaegerModule,

	// Set Global Tracer
	fx.Invoke(telemetry.RegisterTracerProvider),
)
