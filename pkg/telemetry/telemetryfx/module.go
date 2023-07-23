package telemetryfx

import (
	"github.com/ScoreTrak/ScoreTrak/pkg/telemetry"
	"go.uber.org/fx"
)

var Module = fx.Options(
	// OTEL Resource
	fx.Provide(
		telemetry.NewResource,
	),

	// Logging
	fx.Provide(
		telemetry.NewLogger,
	),

	// Custom Loggers
	fx.Provide(
		telemetry.NewWatermillLogger,
		telemetry.NewCronLogger,
	),
	fx.WithLogger(telemetry.NewFxEventLogger),

	// Tracing
	fx.Provide(
		telemetry.NewTracerProvider,
		telemetry.NewOtlpTraceGrpcExporter,
		//telemetry.NewOtlpTraceHttpExporter,
	),

	// Metrics
	fx.Provide(
	//telemetry.NewMeterProvider,
	//telemetry.NewOtlpMetricGrpcExporter,
	//telemetry.NewOtlpMetricHttpExporter,
	//telemetry.NewPrometheusExporter,
	),
)
