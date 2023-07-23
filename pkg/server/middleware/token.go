package middleware

import (
	"context"
	"github.com/ScoreTrak/ScoreTrak/pkg/auth/user"
	"github.com/ScoreTrak/ScoreTrak/pkg/config"
	"github.com/ScoreTrak/ScoreTrak/pkg/entities"
	"github.com/justinas/alice"
	"github.com/ory/kratos-client-go"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"net/http"
)

func NewAPITokenConstructor(um *APITokenMiddleware) alice.Constructor {
	return um.Handler
}

type APITokenMiddleware struct {
	ory        *client.APIClient
	logger     *otelzap.SugaredLogger
	dbClient   *entities.Client
	cookieName string
}

func NewAPITokenMiddleware(cfg *config.Config, oryClient *client.APIClient, logger *otelzap.SugaredLogger) *APITokenMiddleware {
	return &APITokenMiddleware{ory: oryClient, cookieName: cfg.Auth.Ory.CookieName, logger: logger}
}

func (m *APITokenMiddleware) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		ctx := context.Background()

		s, err := m.validateSession(r)
		if err != nil {
			m.logger.Ctx(ctx).Errorw(err.Error())
			next.ServeHTTP(rw, r.WithContext(ctx))
			return
		}
		if !s.GetActive() {
			m.logger.Ctx(ctx).Infow("no user logged in")
			next.ServeHTTP(rw, r.WithContext(ctx))
			return
		}
		i, err := m.getIdentity(ctx, s)
		ctx = user.NewContext(context.Background(), i)
		next.ServeHTTP(rw, r.WithContext(ctx))
	})
}

func (m *APITokenMiddleware) validateSession(r *http.Request) (*client.Session, error) {
	cookie, err := r.Cookie(m.cookieName)
	if err != nil {
		return nil, err
	}
	if cookie == nil {
		return nil, ORY_KRATOS_MISSING_COOKE_ERROR
	}
	resp, _, err := m.ory.FrontendApi.ToSession(context.Background()).Cookie(cookie.String()).Execute()
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *APITokenMiddleware) getIdentity(ctx context.Context, s *client.Session) (*client.Identity, error) {
	i, _, err := m.ory.IdentityApi.GetIdentity(ctx, s.Identity.Id).Execute()
	if err != nil {
		return nil, err
	}
	return i, nil
}
