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
	// Tracing
	fx.Provide(telemetry.NewResource),

	// Exporters
	fx.Provide(
		telemetry.NewOtlpGrpcExporter,
		telemetry.NewOtlpGrpcTracerProvider,
	),

	// Set Global Tracer
	fx.Invoke(telemetry.RegisterTracerProvider),
)

var Module = fx.Options(
	// Logging
	LoggingModule,

	// Open Telemetry
	//OTELModule,
)
