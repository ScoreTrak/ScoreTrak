package telemetry

import (
	"context"
	"github.com/ScoreTrak/ScoreTrak/pkg/config"
	"github.com/ScoreTrak/ScoreTrak/pkg/version"
	"github.com/spf13/viper"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"
	"go.uber.org/fx"
	"log"
)

var ServiceName = "scoretrak"

func NewJaegerTracerProvider(exp *jaeger.Exporter, resource *resource.Resource) *trace.TracerProvider {
	tp := trace.NewTracerProvider(
		trace.WithBatcher(exp),
		trace.WithResource(resource),
	)

	return tp
}

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

func NewJaegerExporter() *jaeger.Exporter {
	exp, _ := jaeger.New(jaeger.WithCollectorEndpoint(
		jaeger.WithEndpoint(viper.GetString("open-telemetry.jaeger.endpoint")),
		//jaeger.WithUsername(viper.GetString("open-telemetry.jaeger.username")),
		//jaeger.WithPassword(viper.GetString("open-telemetry.jaeger.password")),
	))

	return exp
}

func NewOtlpGrpcExporter() *otlptrace.Exporter {
	exp, _ := otlptracegrpc.New(
		context.Background(),
		otlptracegrpc.WithInsecure())
	return exp
}

func NewResource(config config.StaticConfig) *resource.Resource {
	resourceWithAttributes := resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceNameKey.String(ServiceName),
		attribute.Bool("production", config.Prod),
		attribute.String("version", version.Version),
	)

	return resourceWithAttributes
}
