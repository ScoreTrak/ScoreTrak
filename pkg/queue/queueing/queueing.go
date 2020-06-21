package queueing

import (
	"ScoreTrak/pkg/exec"
	"ScoreTrak/pkg/exec/resolver"
	"ScoreTrak/pkg/logger"
	"context"
	"errors"
	"math/rand"
	"strconv"
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
	wg := sync.WaitGroup{}
	wg.Add(1)
	var (
		passed bool
		log    string
		err    error
	)
	go func() {
		passed, log, err = e.Execute()
		wg.Done()
	}()
	if WaitTimeout(&wg, execDeadline.Add(time.Second)) {
		panic(errors.New("check timed out, which should not have happened. this is most likely a bug. Please check logs for more info"))
	}
	var errstr string
	if err != nil {
		errstr = err.Error()
	}
	if time.Now().After(sd.Deadline) {
		l.Error("Service scored late. Please fix the implementation of the following service: ", sd)
	}
	return QCheck{Service: sd.Service, Passed: passed, Log: log, Err: errstr, RoundID: sd.RoundID}
}
