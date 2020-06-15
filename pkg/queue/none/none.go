package none

import (
	"ScoreTrak/pkg/exec"
	"ScoreTrak/pkg/exec/resolver"
	"ScoreTrak/pkg/queue/queueing"
	"errors"
	"fmt"
	"sync"
	"time"
)

type None struct{}

func (n None) Send(sds []*queueing.ScoringData) []*queueing.QCheck {
	wg := &sync.WaitGroup{}
	wg.Add(len(sds))
	ret := make([]*queueing.QCheck, len(sds))
	for i, sd := range sds {
		go func(sd *queueing.ScoringData, i int) {
			defer wg.Done()
			executable := resolver.ExecutableByName(sd.Service.Name)
			exec.UpdateExecutableProperties(executable, sd.Properties)
			e := exec.NewExec(sd.Timeout.Add(-time.Second), sd.Host, executable)
			fmt.Println(fmt.Sprintf("Executing a check for service ID %d for round %d", sd.Service.ID, sd.RoundID))
			err := e.Validate()
			if err != nil {
				qc := queueing.QCheck{Service: sd.Service, Passed: false, Log: "Check did not pass parameter validation", Err: err.Error(), RoundID: sd.RoundID}
				ret[i] = &qc
				return
			}
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

func NewNoneQueue() (*None, error) {
	return &None{}, nil
}
