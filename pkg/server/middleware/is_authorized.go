package middleware

import (
	"github.com/scoretrak/scoretrak/internal/entities"
	"github.com/justinas/alice"
	"github.com/ory/kratos-client-go"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"net/http"
)

func NewIsAuthorizedConstructor(um *IsAuthorizedMiddleware) alice.Constructor {
	return um.Handler
}

// This middleware supplies user information
type IsAuthorizedMiddleware struct {
	ory      *client.APIClient
	logger   *otelzap.SugaredLogger
	dbClient *entities.Client
}

func NewIsAuthorizedMiddleware(oryClient *client.APIClient, logger *otelzap.SugaredLogger) *IsAuthorizedMiddleware {
	return &IsAuthorizedMiddleware{ory: oryClient, logger: logger}
}

func (m *IsAuthorizedMiddleware) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		//ctx := context.Background()

		// Get user
		//_, ok := user.FromContext(ctx)
		// Get role
		// Get list of permission given to the user
		// Run auth check ughhh

	})
}
