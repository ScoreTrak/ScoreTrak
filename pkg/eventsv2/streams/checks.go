package streams

import (
	"context"
	"fmt"
	"github.com/nats-io/nats.go/jetstream"
	"time"
)

const (
	STREAM_CHECKS = "CHECKS"
)

type ChecksStream struct {
	stream jetstream.Stream
}

func NewCheckStream(ctx context.Context, js jetstream.JetStream) (*ChecksStream, error) {
	jsctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	streamConfig := jetstream.StreamConfig{
		Name:        STREAM_CHECKS,
		Description: "score checks",
		Subjects: []string{
			fmt.Sprintf("%s.>", STREAM_CHECKS),
		},
	}

	stream, err := js.CreateOrUpdateStream(jsctx, streamConfig)
	if err != nil {
		return nil, err
	}

	return &ChecksStream{stream: stream}, nil
}
