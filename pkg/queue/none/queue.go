package none

import (
	"errors"
	"fmt"
	"time"

	"github.com/ScoreTrak/ScoreTrak/pkg/queue/queueing"
	"github.com/ScoreTrak/ScoreTrak/pkg/workergroup"
)

type None struct{}

var ErrUnknownPanic = errors.New("unknown panic")
var ErrPanic = errors.New("panic")

func (n None) Send(sds []*queueing.ScoringData) ([]*queueing.QCheck, error, error) {
	ret := make([]*queueing.QCheck, len(sds))
	cq := make(chan queueing.IndexedQueue, 1)
	for i, sd := range sds {
		go func(sd *queueing.ScoringData, i int) {
			defer func() {
				if x := recover(); x != nil {
					var err error
					switch x := x.(type) {
					case string:
						err = fmt.Errorf("%w: %s", ErrPanic, x)
					case error:
						err = x
					default:
						err = ErrUnknownPanic
					}
					cq <- queueing.IndexedQueue{Q: &queueing.QCheck{Service: sd.Service, Passed: false, Log: "Encountered an unexpected error during the check.", Err: err.Error(), RoundID: sd.RoundID}, I: i}
					return
				}
			}()
			qc := queueing.CommonExecute(sd, sd.Deadline.Add(-2*time.Second))
			cq <- queueing.IndexedQueue{Q: &qc, I: i}
		}(sd, i)
	}
	counter := len(sds)
	for {
		select {
		case res := <-cq:
			ret[res.I] = res.Q
			counter--
			if counter == 0 {
				return ret, nil, nil
			}
		case <-time.After(time.Until(sds[0].Deadline)):
			return nil, nil, ErrRoundTookTooLongToScore
		}
	}
}

var ErrRoundTookTooLongToScore = errors.New("round took too long to score. this might be due to many reasons like a worker going down, or the number of rounds being too big for one master")

var ErrMethodNotSupportedForNoneQueue = errors.New("method not supported when queue is none")

func (n None) Receive() {
	panic(ErrMethodNotSupportedForNoneQueue)
}

func (n None) Acknowledge(q queueing.QCheck) {
	panic(ErrMethodNotSupportedForNoneQueue)
}

func (n None) Ping(group *workergroup.WorkerGroup) error {
	return ErrMethodNotSupportedForNoneQueue
}

func NewNoneQueue() (*None, error) {
	return &None{}, nil // Use global Variable
}
