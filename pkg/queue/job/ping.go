package job

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/golang-queue/queue/core"
	"log"
)

type PingJob struct {
	core.QueuedMessage
	Message string
}

func (j *PingJob) Bytes() []byte {
	b, err := json.Marshal(j)
	if err != nil {
		panic(err)
	}
	return b
}

func Ping(ctx context.Context, m core.QueuedMessage) error {
	v, ok := m.(*PingJob)
	if !ok {
		if err := json.Unmarshal(m.Bytes(), &v); err != nil {
			return err
		}
	}

	if v.Message == "ping" {
		log.Println("pong")
		return nil
	}

	return errors.New("not ping")
}
