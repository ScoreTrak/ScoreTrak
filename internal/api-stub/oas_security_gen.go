// Code generated by ogen, DO NOT EDIT.

package api_stub

import (
	"context"
	"net/http"
	"strings"

	"github.com/go-faster/errors"

	"github.com/ogen-go/ogen/ogenerrors"
)

// SecurityHandler is handler for security parameters.
type SecurityHandler interface {
	// HandleApiToken handles ApiToken security.
	HandleApiToken(ctx context.Context, operationName string, t ApiToken) (context.Context, error)
}

func findAuthorization(h http.Header, prefix string) (string, bool) {
	v, ok := h["Authorization"]
	if !ok {
		return "", false
	}
	for _, vv := range v {
		scheme, value, ok := strings.Cut(vv, " ")
		if !ok || !strings.EqualFold(scheme, prefix) {
			continue
		}
		return value, true
	}
	return "", false
}

func (s *Server) securityApiToken(ctx context.Context, operationName string, req *http.Request) (context.Context, bool, error) {
	var t ApiToken
	token, ok := findAuthorization(req.Header, "Bearer")
	if !ok {
		return ctx, false, nil
	}
	t.Token = token
	rctx, err := s.sec.HandleApiToken(ctx, operationName, t)
	if errors.Is(err, ogenerrors.ErrSkipServerSecurity) {
		return nil, false, nil
	} else if err != nil {
		return nil, false, err
	}
	return rctx, true, err
}

// SecuritySource is provider of security values (tokens, passwords, etc.).
type SecuritySource interface {
	// ApiToken provides ApiToken security value.
	ApiToken(ctx context.Context, operationName string) (ApiToken, error)
}

func (s *Client) securityApiToken(ctx context.Context, operationName string, req *http.Request) error {
	t, err := s.sec.ApiToken(ctx, operationName)
	if err != nil {
		return errors.Wrap(err, "security source \"ApiToken\"")
	}
	req.Header.Set("Authorization", "Bearer "+t.Token)
	return nil
}
