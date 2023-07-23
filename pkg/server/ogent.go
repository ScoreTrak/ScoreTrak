package server

import (
	"github.com/ScoreTrak/ScoreTrak/pkg/entities/ogent"
)

func NewOgentServer(h *Handler) (*ogent.Server, error) {
	osrv, err := ogent.NewServer(
		h,
	)
	if err != nil {
		return nil, err
	}
	return osrv, nil
}
