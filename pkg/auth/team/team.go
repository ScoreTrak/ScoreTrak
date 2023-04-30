package team

import (
	"context"
)

const COMPETITION_CONTEXT_KEY = "team"

func NewContext(ctx context.Context, c string) context.Context {
	return context.WithValue(ctx, COMPETITION_CONTEXT_KEY, c)
}

func FromContext(ctx context.Context) (string, bool) {
	i, ok := ctx.Value(COMPETITION_CONTEXT_KEY).(string)
	return i, ok
}
