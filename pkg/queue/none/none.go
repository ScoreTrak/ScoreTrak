package none

import (
	"ScoreTrak/pkg/exec"
	"ScoreTrak/pkg/exec/resolver"
	"ScoreTrak/pkg/logger"
	"ScoreTrak/pkg/queue/queueing"
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

type None struct {
	l logger.LogInfoFormat
}

func (n None) Send(sds []*queueing.ScoringData) ([]*queueing.QCheck, error, error) {
	wg := &sync.WaitGroup{}
	wg.Add(len(sds))
	ret := make([]*queueing.QCheck, len(sds))
	for i, sd := range sds {
		go func(sd *queueing.ScoringData, i int) {
			defer wg.Done()
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
					n.l.Error(err)
					qc := queueing.QCheck{Service: sd.Service, Passed: false, Log: "Encountered an unexpected error during the check. This is most likely a bug", Err: err.Error(), RoundID: sd.RoundID}
					ret[i] = &qc
					return
				}
			}()
			executable := resolver.ExecutableByName(sd.Service.Name)
			exec.UpdateExecutableProperties(executable, sd.Properties)
			ctx := context.Background()
			execDeadline := sd.Deadline.Add(-time.Second * 2)
			ctx, cancel := context.WithDeadline(ctx, execDeadline)
			defer cancel()
			e := exec.NewExec(ctx, sd.Host, executable, n.l)
			fmt.Println(fmt.Sprintf("Executing a check for service ID %d for round %d", sd.Service.ID, sd.RoundID))
			wge := sync.WaitGroup{}
			wge.Add(1)
			var (
				passed bool
				log    string
				err    error
			)
			go func() {
				passed, log, err = e.Execute()
				wge.Done()
			}()
			if queueing.WaitTimeout(&wge, execDeadline.Add(time.Second)) {
				panic(errors.New("check timed out, which should not have happened. this is most likely a bug. Please check logs for more info"))
			}

			var errstr string
			if err != nil {
				errstr = err.Error()
			}
			qc := queueing.QCheck{Service: sd.Service, Passed: passed, Log: log, Err: errstr, RoundID: sd.RoundID}
			ret[i] = &qc
		}(sd, i)
	}
	if queueing.WaitTimeout(wg, sds[0].Deadline) {
		return nil, nil, errors.New("round took too long to score. this might be due to many reasons like a worker going down, or the number of rounds being too big for one master")
	}
	return ret, nil, nil
}

func (n None) Receive() {
	panic(errors.New("you should not call Receive when queue is none"))
}

func (n None) Acknowledge(q queueing.QCheck) {
	panic(errors.New("you should not call Acknowledge when queue is none"))
}

func NewNoneQueue(l logger.LogInfoFormat) (*None, error) {
	return &None{l}, nil
}
