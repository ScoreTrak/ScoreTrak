package none

import (
	"ScoreTrak/pkg/exec"
	"ScoreTrak/pkg/exec/resolver"
	"ScoreTrak/pkg/logger"
	"ScoreTrak/pkg/queue/queueing"
	"errors"
	"fmt"
	"sync"
	"time"
)

type None struct {
	l logger.LogInfoFormat
}

func (n None) Send(sds []*queueing.ScoringData) []*queueing.QCheck {
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
			e := exec.NewExec(sd.Timeout.Add(-time.Second), sd.Host, executable)
			fmt.Println(fmt.Sprintf("Executing a check for service ID %d for round %d", sd.Service.ID, sd.RoundID))
			passed, log, err := e.Execute()
			var errstr string
			if err != nil {
				errstr = err.Error()
			}
			qc := queueing.QCheck{Service: sd.Service, Passed: passed, Log: log, Err: errstr, RoundID: sd.RoundID}
			ret[i] = &qc
		}(sd, i)
	}
	wg.Wait()
	return ret
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
