package telemetry

import (
	"context"
	"github.com/ScoreTrak/ScoreTrak/pkg/config"
	"github.com/ScoreTrak/ScoreTrak/pkg/version"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"
	"go.uber.org/fx"
	"log"
)

var ServiceName = "scoretrak"

func NewOtlpGrpcTracerProvider(exp *otlptrace.Exporter, resource *resource.Resource) *trace.TracerProvider {
	tp := trace.NewTracerProvider(
		trace.WithSampler(trace.AlwaysSample()),
		trace.WithBatcher(exp),
		trace.WithResource(resource),
	)

	return tp
}

func RegisterTracerProvider(lc fx.Lifecycle, tp *trace.TracerProvider) {
	otel.SetTracerProvider(tp)

	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			log.Println("Stopping tracer provider")
			err := tp.Shutdown(ctx)
			if err != nil {
				return err
			}
			return nil
		},
	})
}

func NewOtlpGrpcExporter() *otlptrace.Exporter {
	exp, _ := otlptracegrpc.New(
		context.Background(),
		otlptracegrpc.WithInsecure())
	return exp
}

func NewResource(config config.StaticConfig) *resource.Resource {
	r, _ := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(ServiceName),
			semconv.ServiceVersionKey.String(version.Version),
			attribute.Bool("production", config.Prod),
		),
	)

	return r
}
