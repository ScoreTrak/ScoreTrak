package services

import (
	"ScoreTrak/pkg/exec"
	"errors"
	"github.com/masterzen/winrm"
	"strconv"
	"strings"
	"time"
)

type Winrm struct {
	Username       string `json:"username"`
	Password       string `json:"password"`
	Port           string `json:"port"`
	Command        string `json:"command"`
	ExpectedOutput string `json:"expected_output"`
	Protocol       string `json:"protocol"`
}

func NewWinrm() *Winrm {
	f := Winrm{Port: "5985", Command: "whoami", Protocol: "http"}
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
	if strings.ToLower(w.Protocol) == "https" {
		isHttps = true
	}
	i, err := strconv.Atoi(w.Port)
	if err != nil {
		return false, "Unable to convert port number to integer", err
	}
	endpoint := winrm.NewEndpoint(e.Host, i, isHttps, true, nil, nil, nil, time.Until(e.Timeout))
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
		return false, "The process did not match ExpectedOutput", nil
	}
	return true, "Success!", nil
}
