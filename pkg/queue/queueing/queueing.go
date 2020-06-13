package queueing

import "time"

type ScoringData struct {
	Service    QService
	Properties map[string]string
	Timeout    time.Duration
	Host       string
}

type QService struct {
	ID    uint64
	Group string
	Name  string
}

type QCheck struct {
	Service QService
	Passed  bool
	Log     string
	Err     error
}
