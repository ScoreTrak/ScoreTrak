package session

import (
	"context"
	"github.com/ory/client-go"
)

const CONTEXT_SESSION_KEY = "ory_session"

func NewContext(ctx context.Context, s *client.Session) context.Context {
	return context.WithValue(ctx, CONTEXT_SESSION_KEY, s)
}

func FromContext(ctx context.Context) (*client.Session, bool) {
	s, ok := ctx.Value(CONTEXT_SESSION_KEY).(*client.Session)
	return s, ok
}
