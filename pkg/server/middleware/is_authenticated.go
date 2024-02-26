package middleware

import (
	"context"
	"github.com/justinas/alice"
	"github.com/ory/kratos-client-go"
	"github.com/scoretrak/scoretrak/internal/entities"
	"github.com/scoretrak/scoretrak/pkg/auth/user"
	"github.com/scoretrak/scoretrak/pkg/config"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"net/http"
)

func NewIsAuthenticatedConstructor(um *IsAuthenticatedMiddleware) alice.Constructor {
	return um.Handler
}

// This middleware supplies user information
type IsAuthenticatedMiddleware struct {
	ory      *client.APIClient
	logger   *otelzap.SugaredLogger
	dbClient *entities.Client
}

func NewIsAuthenticatedMiddleware(cfg *config.Config, oryClient *client.APIClient, logger *otelzap.SugaredLogger) *IsAuthenticatedMiddleware {
	return &IsAuthenticatedMiddleware{ory: oryClient, logger: logger}
}

func (m *IsAuthenticatedMiddleware) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		ctx := context.Background()

		_, ok := user.FromContext(ctx)
		if !ok {
			http.Error(rw, "Unauthenticated", http.StatusUnauthorized)
		} else {
			next.ServeHTTP(rw, r)
		}
	})
}
