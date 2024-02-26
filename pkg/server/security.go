package server

import (
	"context"
	"fmt"
	api_stub "github.com/scoretrak/scoretrak/internal/api-stub"
	"github.com/scoretrak/scoretrak/internal/entities"
	"github.com/scoretrak/scoretrak/internal/entities/apitoken"
	"time"
)

type ApiTokenSecurityHandler struct {
	api_stub.SecurityHandler
	dbClient *entities.Client
}

func NewApiTokenSecurityHandler(dbClient *entities.Client) *ApiTokenSecurityHandler {
	return &ApiTokenSecurityHandler{
		dbClient: dbClient,
	}
}

func (a *ApiTokenSecurityHandler) HandleApiToken(ctx context.Context, operationName string, t api_stub.ApiToken) (context.Context, error) {
	apiToken, err := a.dbClient.ApiToken.Query().Where(apitoken.Token(t.Token)).First(ctx)
	if err != nil {
		return ctx, err
	}

	// Ensure that the api token has not expired
	if apiToken.ExpiredAt.Before(time.Now()) {
		return ctx, fmt.Errorf("the api token has expired")
	}

	return ctx, nil
}
