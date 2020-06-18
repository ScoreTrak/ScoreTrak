package queueing

import (
	"sync"
	"time"
)

type ScoringData struct {
	Service    QService
	Properties map[string]string
	Deadline   time.Time
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

func WaitTimeout(wg *sync.WaitGroup, deadline time.Time) bool {
	c := make(chan struct{})
	go func() {
		defer close(c)
		wg.Wait()
	}()
	select {
	case <-c:
		return false
	case <-time.After(time.Until(deadline)):
		return true
	}
} //https://gist.github.com/r4um/c1ab51b8757fc2d75d30320933cdbdf6
