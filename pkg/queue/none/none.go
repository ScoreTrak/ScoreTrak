package none

import (
	"errors"

	"github.com/ScoreTrak/ScoreTrak/pkg/queue/queueing"
	"github.com/ScoreTrak/ScoreTrak/pkg/service_group"
	"time"
)

type None struct{}

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
						err = errors.New(x)
					case error:
						err = x
					default:
						err = errors.New("unknown panic")
					}
					cq <- queueing.IndexedQueue{Q: &queueing.QCheck{Service: sd.Service, Passed: false, Log: "Encountered an unexpected error during the check. This is most likely a bug", Err: err.Error(), RoundID: sd.RoundID}, I: i}
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
			return nil, nil, errors.New("round took too long to score. this might be due to many reasons like a worker going down, or the number of rounds being too big for one master")
		}
	}
}

func (n None) Receive() {
	panic(errors.New("you should not call Receive when queue is none"))
}

func (n None) Acknowledge(q queueing.QCheck) {
	panic(errors.New("you should not call Acknowledge when queue is none"))
}

func (n None) Ping(group *service_group.ServiceGroup) error {
	return errors.New("you should not call Ping when queue is none")
}

func NewNoneQueue() (*None, error) {
	return &None{}, nil
}
