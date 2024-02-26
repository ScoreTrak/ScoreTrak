package streams

import (
	"context"
	"fmt"
	"github.com/nats-io/nats.go/jetstream"
	"time"
)

const (
	STREAM_PRINT = "PRINT"
)

type PrintStream struct {
	stream jetstream.Stream
}

func NewPrintStream(ctx context.Context, js jetstream.JetStream) (*PrintStream, error) {
	jsctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	streamConfig := jetstream.StreamConfig{
		Name:        STREAM_PRINT,
		Description: "print messages",
		Subjects: []string{
			fmt.Sprintf("%s.>", STREAM_PRINT),
		},
	}

	stream, err := js.CreateOrUpdateStream(jsctx, streamConfig)
	if err != nil {
		return nil, err
	}

	return &PrintStream{stream: stream}, nil
}
