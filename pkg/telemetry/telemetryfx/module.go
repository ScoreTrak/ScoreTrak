package telemetryfx

import (
	"github.com/ScoreTrak/ScoreTrak/pkg/telemetry"
	"go.uber.org/fx"
)

var Module = fx.Options(
	// OTEL Resource
	fx.Provide(
		telemetry.NewInstance,
		telemetry.NewResource,
	),

	// Logging
	fx.Provide(telemetry.NewLogger),

	// Tracing
	fx.Provide(
		telemetry.NewTracerProvider,
		telemetry.NewOtlpTraceGrpcExporter,
		//telemetry.NewOtlpTraceHttpExporter,
	),

	// Metrics
	fx.Provide(
		telemetry.NewMeterProvider,
		//telemetry.NewOtlpMetricGrpcExporter,
		//telemetry.NewOtlpMetricHttpExporter,
		telemetry.NewPrometheusExporter,
	),
)
