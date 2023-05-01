package job

import (
	"context"
	"encoding/json"
	"github.com/ScoreTrak/ScoreTrak/internal/entities"
	"github.com/golang-queue/queue/core"
	"log"
)

type ScoreHostServiceJob struct {
	core.QueuedMessage
	CompetitionID string
	Host          entities.Host
	HostService   entities.HostService
	Properties    []entities.Property
}

func (j *ScoreHostServiceJob) Bytes() []byte {
	b, err := json.Marshal(j)
	if err != nil {
		panic(err)
	}
	return b
}

func ScoreHostService(ctx context.Context, m core.QueuedMessage) error {
	v, ok := m.(*ScoreHostServiceJob)
	if !ok {
		if err := json.Unmarshal(m.Bytes(), &v); err != nil {
			return err
		}
	}

	log.Printf("Comp %s, Host %s, HostService %s", v.CompetitionID, v.Host.ID, v.HostService.ID)
	return nil
}
