package server

import (
	"github.com/ScoreTrak/ScoreTrak/internal/entities"
	"github.com/ScoreTrak/ScoreTrak/internal/entities/ogent"
)

func NewEntityServer(dbClient *entities.Client) (*ogent.Server, error) {
	osrv, err := ogent.NewServer(ogent.NewOgentHandler(dbClient))
	if err != nil {
		return nil, err
	}
	return osrv, nil
}
