package telemetry

import (
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
)

func NewMeterProvider(r *resource.Resource, exp *metric.Exporter) *metric.MeterProvider {
	mp := metric.NewMeterProvider(
		metric.WithResource(r),
		metric.WithReader(
			metric.NewPeriodicReader(*exp),
		),
	)

	return mp
}
