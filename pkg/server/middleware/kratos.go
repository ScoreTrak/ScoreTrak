package middleware

import (
	"context"
	"errors"
	"fmt"
	"github.com/ScoreTrak/ScoreTrak/pkg/auth/session"
	"github.com/ScoreTrak/ScoreTrak/pkg/config"
	"github.com/ogen-go/ogen/middleware"
	"github.com/ory/client-go"
	"net/http"
)

var ORY_KRATOS_MISSING_COOKE_ERROR = errors.New("No session found in cooke")

type KratosMiddleware struct {
	ory        *client.APIClient
	cookieName string
}

func NewKratosMiddleware(cfg *config.Config, oryClient *client.APIClient) *KratosMiddleware {
	return &KratosMiddleware{ory: oryClient, cookieName: fmt.Sprintf(`ory_session_%s`, cfg.Auth.Ory.Slug)}
}

func (k *KratosMiddleware) Middleware(req middleware.Request, next middleware.Next) (middleware.Response, error) {
	//cookie, ok := req.Params.Cookie(k.cookieName)
	//req.Params.Cookie("")
	//if !ok {
	//}
	//if string(cookie) == nil {
	//}
}

func (k *KratosMiddleware) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		s, err := k.validateSession(r)
		if err != nil {
			return
		}
		if !*s.Active {
			return
		}
		ctx := session.NewContext(context.Background(), s)
		next.ServeHTTP(rw, r.WithContext(ctx))
	})
}

func (k *KratosMiddleware) validateSession(r *http.Request) (*client.Session, error) {
	cookie, err := r.Cookie(k.cookieName)
	if err != nil {
		return nil, err
	}
	if cookie == nil {
		return nil, ORY_KRATOS_MISSING_COOKE_ERROR
	}
	resp, _, err := k.ory.FrontendApi.ToSession(context.Background()).Cookie(cookie.String()).Execute()
	if err != nil {
		return nil, err
	}
	return resp, nil
}
