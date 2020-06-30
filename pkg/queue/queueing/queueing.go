package queueing

import (
	"context"
	"errors"
	"github.com/L1ghtman2k/ScoreTrak/pkg/exec"
	"github.com/L1ghtman2k/ScoreTrak/pkg/exec/resolver"
	"github.com/L1ghtman2k/ScoreTrak/pkg/logger"
	"math/rand"
	"strconv"
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
	ID             uint64
	Group          string
	Name           string
	ReturningTopic string
}

type QCheck struct {
	Service QService
	RoundID uint64
	Passed  bool
	Log     string
	Err     string
}

type IndexedQueue struct {
	Q *QCheck
	I int
}

//func WaitTimeout(wg *sync.WaitGroup, deadline time.Time) bool {
//	c := make(chan struct{})
//	go func() {
//		wg.Wait()
//		close(c)
//	}()
//	select {
//	case <-c:
//		return false
//	case <-time.After(time.Until(deadline)):
//		return true
//	}
//} //https://gist.github.com/r4um/c1ab51b8757fc2d75d30320933cdbdf6

func TopicFromServiceRound(ser *QService, roundID uint64) string {
	if roundID == 0 {
		var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
		return "test_" + strconv.FormatUint(ser.ID, 10) + strconv.Itoa(seededRand.Int()) + "_ack"
	}
	return strconv.FormatUint(roundID, 10) + "_ack"
}

func CommonExecute(sd *ScoringData, execDeadline time.Time, l logger.LogInfoFormat) QCheck {
	if time.Now().After(sd.Deadline) {
		return QCheck{Service: sd.Service, Passed: false, Log: "", Err: "The check arrived late to the worker", RoundID: sd.RoundID}
	}
	executable := resolver.ExecutableByName(sd.Service.Name)
	exec.UpdateExecutableProperties(executable, sd.Properties)
	ctx := context.Background()
	ctx, cancel := context.WithDeadline(ctx, execDeadline)
	defer cancel()

	e := exec.NewExec(ctx, sd.Host, executable, l)
	type checkRet struct {
		passed bool
		log    string
		err    error
	}
	cq := make(chan checkRet)
	go func() {
		passed, log, err := e.Execute()
		cq <- checkRet{passed: passed, log: log, err: err}
	}()
	select {
	case res := <-cq:
		var errstr string
		if res.err != nil {
			errstr = res.err.Error()
		}
		return QCheck{Service: sd.Service, Passed: res.passed, Log: res.log, Err: errstr, RoundID: sd.RoundID}
	case <-time.After(time.Until(execDeadline.Add(time.Second))):
		panic(errors.New("check timed out, which should not have happened. this is most likely a bug. Please check logs for more info"))
	}
}

type RoundTookTooLongToExecute struct {
	Msg string
}

func (e *RoundTookTooLongToExecute) Error() string { return e.Msg }
