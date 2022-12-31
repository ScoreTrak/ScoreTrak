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

var Module = fx.Options(
	// Logging
	fx.Provide(telemetry.NewLogger),
	// Tracing
	fx.Provide(telemetry.NewResource),

	// OTLP
	OtlpModule,

	// Set Global Tracer
	fx.Invoke(telemetry.RegisterTracerProvider),
)
