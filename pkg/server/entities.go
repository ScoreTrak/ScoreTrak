package server

import (
	"github.com/ScoreTrak/ScoreTrak/internal/entities"
	"github.com/ScoreTrak/ScoreTrak/internal/entities/ogent"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
)

func NewEntityServer(dbClient *entities.Client, mp metric.MeterProvider, tp trace.TracerProvider) (*ogent.Server, error) {
	//crs := cors.New(cors.Options{
	//	AllowedOrigins:   []string{"http://localhost:5173"},
	//	AllowedMethods:   nil,
	//	AllowedHeaders:   nil,
	//	ExposedHeaders:   nil,
	//	MaxAge:           0,
	//	AllowCredentials: true,
	//	//AllowPrivateNetwork:  false,
	//	//OptionsPassthrough:   false,
	//	//OptionsSuccessStatus: 0,
	//	Debug: true,
	//})
	osrv, err := ogent.NewServer(
		ogent.NewOgentHandler(dbClient),
		ogent.WithTracerProvider(tp),
		ogent.WithMeterProvider(mp),
	)
	if err != nil {
		return nil, err
	}
	return osrv, nil
}
