package exec

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"reflect"
	"time"
)

type Executable interface {
	Validate() error
	Execute(e Exec) (passed bool, log string, err error)
}

type Exec struct {
	Timeout    time.Duration
	Host       string
	executable Executable
}

func NewExec(t time.Duration, h string, e Executable) *Exec {
	return &Exec{Timeout: t, Host: h, executable: e}
}

func (e Exec) Execute() (passed bool, log string, err error) {
	//TODO:Terminate based on global timeout logic
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
