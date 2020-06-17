package services

import (
	"ScoreTrak/pkg/exec"
	"errors"
	"github.com/masterzen/winrm"
	"strconv"
	"time"
)

type Winrm struct {
	Username       string `json:"username"`
	Password       string `json:"password"`
	Port           string `json:"port"`
	Command        string `json:"command"`
	ExpectedOutput string `json:"expected_output"`
	Scheme         string `json:"scheme"`
}

func NewWinrm() *Winrm {
	f := Winrm{Command: "whoami", Scheme: "http"}
	return &f
}

func (w *Winrm) Validate() error {
	if w.Password != "" && w.Username != "" {
		return nil
	}
	return errors.New("winrm service needs username, and password")
}

func (w *Winrm) Execute(e exec.Exec) (passed bool, log string, err error) {
	var isHttps bool
	isHttps = exec.IsSecure(w.Scheme)
	if w.Port == "" {
		if isHttps {
			w.Port = "5986"
		} else {
			w.Port = "5985"
		}
	}
	i, err := strconv.Atoi(w.Port)
	if err != nil {
		return false, "Unable to convert port number to integer", err
	}
	endpoint := winrm.NewEndpoint(e.Host, i, isHttps, true, nil, nil, nil, time.Until(e.Deadline()))
	client, err := winrm.NewClient(endpoint, w.Username, w.Password)
	if err != nil {
		return false, "Unable to initialize winrm client", err
	}
	procStdout, procStderr, returnCode, err := client.RunWithString(w.Command, "")
	if err != nil {
		return false, "Unable to execute provided command", err
	}
	if returnCode != 0 {
		return false, "Process returned a non-zero code", errors.New(procStderr)
	}
	if w.ExpectedOutput != "" && procStdout != w.ExpectedOutput {
		return false, "The process did not match ExpectedOutput", nil //TODO: Make a more meaningful output
	}
	return true, "Success!", nil
}
