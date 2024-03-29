package exec

import (
	"context"
	"errors"
	"fmt"
	"log"

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
}

func NewExec(ctx context.Context, host string, e Executable) *Exec {
	return &Exec{Context: ctx, Host: host, executable: e}
}

var ErrCheckDeadlineExceeded = errors.New("unable to start the check: deadline passed to a check wasn't set, or was negative. This is most likely a misconfiguration(round_duration too small)")

// Execute first validates the given check, and if validation passes it executes a given check.
func (e Exec) Execute() (passed bool, log string, err error) {
	err = e.Validate()
	if err != nil {
		return false, "", fmt.Errorf("check did not pass parameter validation: %w", err)
	}
	if time.Now().After(e.Deadline()) {
		return false, "", ErrCheckDeadlineExceeded
	}
	return e.executable.Execute(e)
}

// Validate validates the service parameters
func (e Exec) Validate() error {
	return e.executable.Validate()
}

// Deadline calculates deadline of a service
func (e Exec) Deadline() time.Time {
	deadline, ok := e.Context.Deadline()
	if !ok {
		log.Panicln("deadline was not set")
	}
	return deadline
}

var ErrUnknownPanic = errors.New("unknown panic")
var ErrPanic = errors.New("panic")

// UpdateExecutableProperties sets the properties of a given check by extracting them from map(string) => string, and setting them via reflection.
func UpdateExecutableProperties(v Executable, p map[string]string) (err error) {
	defer func() {
		if x := recover(); x != nil {
			switch x := x.(type) {
			case string:
				err = fmt.Errorf("%w: %s", ErrPanic, x)
			case error:
				err = x
			default:
				err = ErrUnknownPanic
			}
		}
	}()

	rv := reflect.ValueOf(v).Elem()
	for key, val := range p {
		if val != "" { // Eliminate unnecessary default value
			rf := rv.FieldByName(key)
			rf.SetString(val)
		}
	}
	return nil
}
