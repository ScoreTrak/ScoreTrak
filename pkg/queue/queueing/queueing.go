package queueing

import (
	"context"
	"errors"
	"github.com/L1ghtman2k/ScoreTrak/pkg/exec"
	"github.com/L1ghtman2k/ScoreTrak/pkg/exec/resolver"
	"github.com/L1ghtman2k/ScoreTrak/pkg/logger"
	"github.com/gofrs/uuid"
	"math/rand"
	"strconv"
	"time"
)

type ScoringData struct {
	Service    QService
	Properties map[string]string
	Deadline   time.Time
	Host       string
	RoundID    uint
}

type QService struct {
	ID             uuid.UUID
	Group          string
	Name           string
	ReturningTopic string
}

type QCheck struct {
	Service QService
	RoundID uint
	Passed  bool
	Log     string
	Err     string
}

type IndexedQueue struct {
	Q *QCheck
	I int
}

type Config struct {
	Use   string `default:"none"`
	Kafka struct {
	}
	NSQ struct {
		NSQD struct {
			Port string `default:"4150"`
			Host string `default:"nsqd"`
		}
		IgnoreAllScoresIfWorkerFails bool   `default:"true"`
		Topic                        string `default:"default"`
		MaxInFlight                  int    `default:"200"`
		ConcurrentHandlers           int    `default:"200"`
		NSQLookupd                   struct {
			Hosts []string `default:"[\"nsqlookupd\"]"`
			Port  string   `default:"4161"`
		}
	}
}

func TopicFromServiceRound(ser *QService, roundID uint) string {
	if roundID == 0 {
		var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
		return "test_" + ser.ID.String() + strconv.Itoa(seededRand.Int()) + "_ack"
	}
	return ser.ID.String() + "_ack"
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
