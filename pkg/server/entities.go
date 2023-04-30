package server

import (
	"github.com/ScoreTrak/ScoreTrak/internal/entities/ogent"
)

func NewEntityServer(h *handler) (*ogent.Server, error) {
	osrv, err := ogent.NewServer(
		h,
		//ogent.WithTracerProvider(*tp),
		//ogent.WithMeterProvider(*mp),
	)
	if err != nil {
		return nil, err
	}
	return osrv, nil
}
