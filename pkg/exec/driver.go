package exec

import (
	"errors"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"reflect"
	"time"
)

type Executable interface {
	Validate() error
	Execute(e Exec) (passed bool, log string, err error)
}

type Exec struct {
	Timeout    time.Time
	Host       string
	executable Executable
}

func NewExec(t time.Time, h string, e Executable) *Exec {
	return &Exec{Timeout: t, Host: h, executable: e}
}

func (e Exec) Execute() (passed bool, log string, err error) {
	oldTimeout := e.Timeout
	e.Timeout = e.Timeout.Add(-time.Second)
	completed := make(chan bool, 1)
	defer close(completed)
	go func() {
		passed, log, err = e.executable.Execute(e)
		completed <- true
	}()
	select {
	case <-completed:
		break
	case <-time.After(time.Until(oldTimeout)):
		return false, "", errors.New("check took too long to execute")
	}
	return e.executable.Execute(e)
}
func (e Exec) Validate() error {
	return e.executable.Validate()
}

func UpdateExecutableProperties(v Executable, p map[string]string) {
	rv := reflect.ValueOf(v).Elem()
	for key, val := range p {
		rf := rv.FieldByName(key)
		rf.SetString(val)
	}
}
