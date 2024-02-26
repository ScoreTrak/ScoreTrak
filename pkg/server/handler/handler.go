package handler

import (
	api_stub "github.com/scoretrak/scoretrak/internal/api-stub"
	"github.com/scoretrak/scoretrak/internal/entities"
)

type Handler struct {
	api_stub.Handler
	dbClient *entities.Client
}

func NewHandler(dbClient *entities.Client) *Handler {
	return &Handler{
		dbClient: dbClient,
	}
}

//func (h *Handler) GetCompetitions(ctx context.Context) (*api_stub.Competitions, error) {
//	competitions, err := h.dbClient.Competition.Query().Limit().Offset().All(ctx)
//	if err != nil {
//		return nil, err
//	}
//
//	var cs api_stub.Competitions
//	err = convertEntityToResponse(competitions, cs.UnmarshalJSON)
//	if err != nil {
//		return nil, err
//	}
//
//	return &cs, nil
//}
