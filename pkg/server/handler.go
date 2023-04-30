package server

import (
	"github.com/ScoreTrak/ScoreTrak/internal/entities"
	"github.com/ScoreTrak/ScoreTrak/internal/entities/ogent"
)

type handler struct {
	*ogent.OgentHandler
	dbClient *entities.Client
}

func NewHandler(dbClient *entities.Client) *handler {
	return &handler{
		OgentHandler: ogent.NewOgentHandler(dbClient),
		dbClient:     dbClient,
	}
}
