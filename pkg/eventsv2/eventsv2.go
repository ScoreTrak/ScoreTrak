package eventsv2

import (
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	"github.com/scoretrak/scoretrak/pkg/config"
)

/*
Nats Subject Mapping

https://docs.nats.io/nats-concepts/subjects

How to subscribe to everything!! CHECKS.>
How to subscribe to a specific team with new checks!! CHECKS.created.<team-ulid>.>
How to subscribe to all completed checks!! CHECKS.>
How to subscribe to all completed checks for a specific team!! CHECKS.<team-ulid>.>

How to publish a check to be scored for a specific team!! CHECKS.<team-ulid>
How to publish a saved check to a specific team!! CHECKS.<team-ulid>

*/

func NewNats(cfg *config.Config) (jetstream.JetStream, error) {
	nc, err := nats.Connect(cfg.Queue.NATS.Url)
	if err != nil {
		return nil, err
	}

	js, err := jetstream.New(nc)
	if err != nil {
		return nil, err
	}

	return js, nil
}
