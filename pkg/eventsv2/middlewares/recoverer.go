package middlewares

import (
	"fmt"
	"github.com/nats-io/nats.go/jetstream"
	"github.com/pkg/errors"
	"log"
	"runtime/debug"
)

// RecoveredPanicError holds the recovered panic's error along with the stacktrace.
type RecoveredPanicError struct {
	V          interface{}
	Stacktrace string
}

func (p RecoveredPanicError) Error() string {
	return fmt.Sprintf("panic occurred: %#v, stacktrace: \n%s", p.V, p.Stacktrace)
}

// Recoverer recovers from any panic in the handler and appends RecoveredPanicError with the stacktrace
// to any error returned from the handler.
func Recoverer(m jetstream.MessageHandler) jetstream.MessageHandler {
	return func(m jetstream.Msg) {
		var err error
		panicked := true

		defer func() {
			if r := recover(); r != nil || panicked {
				err = errors.WithStack(RecoveredPanicError{V: r, Stacktrace: string(debug.Stack())})
			}
		}()

		panicked = false
		log.Println(err)
	}
}