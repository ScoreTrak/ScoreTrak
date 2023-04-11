package telemetryfx

import (
	"github.com/ScoreTrak/ScoreTrak/pkg/telemetry"
	"go.uber.org/fx"
)

var LoggingModule = fx.Options(
	// Logging
	fx.Provide(telemetry.NewLogger),
)

var OTELModule = fx.Options(
	// Resource
	fx.Provide(telemetry.NewResource),

	// Tracing
	fx.Provide(
		telemetry.NewTracerProvider,
	),

	// Metrics
	fx.Provide(
		telemetry.NewMeterProvider,
	),

	// Exporters
	fx.Provide(
		telemetry.NewOtlpGrpcExporter,
	),

	// Set Global Tracer
	fx.Invoke(telemetry.RegisterTracerProvider),
)

var Module = fx.Options(
	// Open Telemetry
	OTELModule,

	// Logging
	LoggingModule,
)
