package handler

import (
	"context"
	"github.com/scoretrak/scoretrak/internal/api-stub"
	"github.com/scoretrak/scoretrak/internal/entities"
)

func (h *Handler) GetCompetition(ctx context.Context, params api_stub.GetCompetitionParams) (*api_stub.Competition, error) {
	competition, err := h.dbClient.Competition.Get(ctx, params.ID)
	if err != nil {
		return nil, err
	}

	var c api_stub.Competition
	err = convertEntityToResponse(competition, c.UnmarshalJSON)
	if err != nil {
		return nil, err
	}

	return &c, nil
}

func (h *Handler) CreateCompetition(ctx context.Context, req *api_stub.CreateCompetitionRequest) (*api_stub.Competition, error) {
	var ce entities.Competition
	err := ConvertStruct(req, &ce)
	if err != nil {
		return nil, err
	}

	newCompetition, err := h.dbClient.Competition.Create().
		SetName(ce.Name).
		SetDisplayName(ce.DisplayName).
		SetNillableViewableToPublic(ce.ViewableToPublic).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	var c api_stub.Competition
	err = convertEntityToResponse(newCompetition, c.UnmarshalJSON)
	if err != nil {
		return nil, err
	}

	return &c, nil
}

func (h *Handler) UpdateCompetition(ctx context.Context, req *api_stub.UpdateCompetitionRequest, params api_stub.UpdateCompetitionParams) (*api_stub.Competition, error) {
	competition := entities.Competition{ID: params.ID}
	err := ConvertStruct(req, &competition)
	if err != nil {
		return nil, err
	}

	updatedCompetition, err := h.dbClient.Competition.UpdateOne(&competition).Save(ctx)
	if err != nil {
		return nil, err
	}

	var c api_stub.Competition
	err = convertEntityToResponse(updatedCompetition, c.UnmarshalJSON)
	if err != nil {
		return nil, err
	}

	return &c, nil
}
