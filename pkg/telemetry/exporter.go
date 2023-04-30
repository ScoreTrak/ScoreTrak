package telemetry

import (
	"context"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetrichttp"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/sdk/metric"
)

func NewOtlpTraceGrpcExporter() (*otlptrace.Exporter, error) {
	exp, err := otlptracegrpc.New(
		context.Background(),
	)
	if err != nil {
		return nil, err
	}
	return exp, nil
}

func NewOtlpTraceHttpExporter() (*otlptrace.Exporter, error) {
	exp, err := otlptracehttp.New(
		context.Background(),
	)
	if err != nil {
		return nil, err
	}
	return exp, nil
}

func NewOtlpMetricGrpcExporter() (metric.Exporter, error) {
	exp, err := otlpmetricgrpc.New(
		context.Background(),
	)
	if err != nil {
		return nil, err
	}
	return exp, nil
}

func NewOtlpMetricHttpExporter() (metric.Exporter, error) {
	exp, err := otlpmetrichttp.New(
		context.Background(),
	)
	if err != nil {
		return nil, err
	}
	return exp, nil
}

func NewPrometheusExporter() (*prometheus.Exporter, error) {
	exp, err := prometheus.New()
	if err != nil {
		return nil, err
	}
	return exp, nil
}
