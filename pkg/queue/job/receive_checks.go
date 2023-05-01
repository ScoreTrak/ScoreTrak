package job

import (
	"context"
	"encoding/json"
	"github.com/ScoreTrak/ScoreTrak/internal/entities"
	"github.com/golang-queue/queue/core"
	"log"
)

const (
	QUEUE_NAME_RECEIVE_CHECKS = "receive_checks"
)

type ReceiveChecksJob struct {
	core.QueuedMessage
	Check entities.Check
}

func (j *ReceiveChecksJob) Bytes() []byte {
	b, err := json.Marshal(j)
	if err != nil {
		panic(err)
	}
	return b
}

func ReceiveChecks(ctx context.Context, m core.QueuedMessage) error {
	v, ok := m.(*ReceiveChecksJob)
	if !ok {
		if err := json.Unmarshal(m.Bytes(), &v); err != nil {
			return err
		}
	}

	log.Println(v.Check)
	return nil
}
