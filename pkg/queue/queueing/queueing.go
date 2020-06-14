package queueing

import "time"

type ScoringData struct {
	Service    QService
	Properties map[string]string
	Timeout    time.Time
	Host       string
	RoundID    uint64
}

type QService struct {
	ID    uint64
	Group string
	Name  string
}

type QCheck struct {
	Service QService
	RoundID uint64
	Passed  bool
	Log     string
	Err     string
}
