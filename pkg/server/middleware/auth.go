package middleware

import (
	"context"
	"errors"
	"github.com/ScoreTrak/ScoreTrak/internal/entities"
	"github.com/ScoreTrak/ScoreTrak/pkg/auth/user"
	"github.com/ScoreTrak/ScoreTrak/pkg/config"
	"github.com/justinas/alice"
	"github.com/ory/kratos-client-go"
	"net/http"
)

var ORY_KRATOS_MISSING_COOKE_ERROR = errors.New("No session found in cooke")

var COMPETITION_HEADER = "X-Scoretrak-Competition-ID"
var TEAM_HEADER = "X-Scoretrak-Team-ID"

type AuthMiddleware struct {
	ory        *client.APIClient
	dbClient   *entities.Client
	cookieName string
}

func NewAuthMiddleware(cfg *config.Config, oryClient *client.APIClient) *AuthMiddleware {
	return &AuthMiddleware{ory: oryClient, cookieName: cfg.Auth.Ory.CookieName}
}

func NewAuthConstructor(km *AuthMiddleware) alice.Constructor {
	return km.Handler
}

func (a *AuthMiddleware) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		ctx := context.Background()

		s, err := a.validateSession(r)
		if err != nil {
			next.ServeHTTP(rw, r)
			return
		}
		if !s.GetActive() {
			next.ServeHTTP(rw, r)
			return
		}
		i, err := a.getIdentity(ctx, s)
		ctx = user.NewContext(context.Background(), i)
		next.ServeHTTP(rw, r.WithContext(ctx))
	})
}

func (a *AuthMiddleware) validateSession(r *http.Request) (*client.Session, error) {
	cookie, err := r.Cookie(a.cookieName)
	if err != nil {
		return nil, err
	}
	if cookie == nil {
		return nil, ORY_KRATOS_MISSING_COOKE_ERROR
	}
	resp, _, err := a.ory.FrontendApi.ToSession(context.Background()).Cookie(cookie.String()).Execute()
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *AuthMiddleware) getCompetition(r *http.Request) string {
	compId := r.Header.Get(COMPETITION_HEADER)
	return compId
}

func (a *AuthMiddleware) getTeam(r *http.Request) string {
	teamId := r.Header.Get(TEAM_HEADER)
	return teamId
}

func (a *AuthMiddleware) getIdentity(ctx context.Context, s *client.Session) (*client.Identity, error) {
	i, _, err := a.ory.IdentityApi.GetIdentity(ctx, s.Identity.Id).Execute()
	if err != nil {
		return nil, err
	}
	return i, nil
}
