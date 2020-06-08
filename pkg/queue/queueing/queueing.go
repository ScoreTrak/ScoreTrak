package queueing

import "time"

type ScoringData struct {
	Service    QService
	Properties []QProperty
	Timeout    time.Duration
	Host       string
}

type QService struct {
	ID   uint64
	Name string
}

type QProperty struct {
	Key   string
	Value string
}

type QCheck struct {
	Service QService
	Passed  bool
	Log     string
	Err     error
}
