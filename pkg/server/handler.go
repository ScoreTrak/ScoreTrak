package server

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"github.com/ScoreTrak/ScoreTrak/pkg/entities"
	"github.com/ScoreTrak/ScoreTrak/pkg/entities/ogent"
	"github.com/ScoreTrak/ScoreTrak/pkg/entities/round"
)

type Handler struct {
	*ogent.OgentHandler
	dbClient *entities.Client
}

func NewHandler(dbClient *entities.Client) *Handler {
	return &Handler{
		OgentHandler: ogent.NewOgentHandler(dbClient),
		dbClient:     dbClient,
	}
}

func (h Handler) ReadRoundLatest(ctx context.Context) (ogent.ReadRoundLatestRes, error) {
	rnd, err := h.dbClient.Round.Query().Limit(1).Where(round.StatusEQ(round.StatusFinished)).Order(round.ByRoundNumber(sql.OrderDesc())).First(ctx)
	if err != nil {
		return nil, err
	}
	if rnd == nil {
		return &ogent.ReadRoundLatestNoContent{}, nil
	}
	return ogent.NewRoundRead(rnd), nil
}

func (h Handler) ListRoundChecksLatest(ctx context.Context) (ogent.ListRoundChecksLatestRes, error) {
	rnd, err := h.dbClient.Round.Query().Limit(1).Where(round.StatusEQ(round.StatusFinished)).Order(round.ByRoundNumber(sql.OrderDesc())).WithChecks().First(ctx)
	if err != nil {
		return nil, err
	}

	list := ogent.NewRoundChecksLists(rnd.Edges.Checks)
	res := ogent.ListRoundChecksLatestOKApplicationJSON(list)
	return &res, nil

}
