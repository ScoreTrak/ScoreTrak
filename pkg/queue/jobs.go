package queue

import (
	"encoding/json"
	"github.com/ScoreTrak/ScoreTrak/internal/entities"
)

type ScoreServiceJob struct {
	Service    entities.Service
	Properties entities.Properties
}

func (j *ScoreServiceJob) Bytes() []byte {
	b, err := json.Marshal(j)
	if err != nil {
		panic(err)
	}
	return b
}
