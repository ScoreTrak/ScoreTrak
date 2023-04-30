package middleware

import (
	"github.com/justinas/alice"
)

func NewMiddlewareChain(constructors ...alice.Constructor) alice.Chain {
	return alice.New(constructors...)
}
