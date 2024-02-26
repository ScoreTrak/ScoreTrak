package models

import (
	"time"
)

type ElaspedModel struct {
	StartedAt  time.Time `json:"started_at"`
	FinishedAt time.Time `json:"finished_at"`
}
