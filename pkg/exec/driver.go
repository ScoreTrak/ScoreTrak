package exec

import (
	"context"
	"errors"
	"github.com/ScoreTrak/ScoreTrak/pkg/logger"
	"reflect"
	"time"
)

type Executable interface {
	Validate() error
	Execute(e Exec) (passed bool, log string, err error)
}

type Exec struct {
	Context    context.Context
	Host       string
	executable Executable
	Logger     logger.LogInfoFormat
}

func NewExec(ctx context.Context, h string, e Executable, l logger.LogInfoFormat) *Exec {
	return &Exec{Context: ctx, Host: h, executable: e, Logger: l}
}

func (e Exec) Execute() (passed bool, log string, err error) {
	err = e.Validate()
	if err != nil {
		return false, "Check did not pass parameter validation", err
	}
	if time.Now().After(e.Deadline()) {
		return false, "Unable to start the check", errors.New("deadline passed to a check wasn't set, or was negative. This is most likely a bug, or a misconfiguration")
	}
	return e.executable.Execute(e)
}
func (e Exec) Validate() error {
	return e.executable.Validate()
}

func (e Exec) Deadline() time.Time {
	deadline, ok := e.Context.Deadline()
	if !ok {
		panic(errors.New("deadline was not set"))
	}
	return deadline
}

func UpdateExecutableProperties(v Executable, p map[string]string) {
	rv := reflect.ValueOf(v).Elem()
	for key, val := range p {
		rf := rv.FieldByName(key)
		rf.SetString(val)
	}
}
