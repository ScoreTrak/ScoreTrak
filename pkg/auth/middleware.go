package auth

import (
	"context"
	"errors"
	"github.com/ory/client-go"
	"net/http"
)

const ORY_KRATOS_COOKIE_NAME = "X-SESSION-TOKEN"
const USER_ID_CONTEXT_KEY = "user_id"

var ORY_KRATOS_MISSING_COOKE_ERROR = errors.New("No session found in cooke")

type kratosMiddleware struct {
	ory *client.APIClient
}

func NewKratosMiddleware(oryClient *client.APIClient) *kratosMiddleware {
	return &kratosMiddleware{ory: oryClient}
}

func (k *kratosMiddleware) Session(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		session, err := k.validateSession(r)
		if err != nil {
			return
		}
		if !*session.Active {
			return
		}
		ctx := context.WithValue(r.Context(), USER_ID_CONTEXT_KEY, "")
		next.ServeHTTP(rw, r.WithContext(ctx))
	})
}

func (k *kratosMiddleware) validateSession(r *http.Request) (*client.Session, error) {
	cookie, err := r.Cookie(ORY_KRATOS_COOKIE_NAME)
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
