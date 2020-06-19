package services

import (
	"ScoreTrak/pkg/exec"
	"errors"
	"github.com/masterzen/winrm"
	"net"
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
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(e.Host, w.Port), time.Until(e.Deadline())/3)
	if err != nil {
		return false, "Port was not open on a remote host", err
	}
	conn.Close()
	endpoint := winrm.NewEndpoint(e.Host, i, isHttps, true, nil, nil, nil, time.Until(e.Deadline()))
	//params := winrm.NewParameters(strconv.Itoa(int(time.Until(e.Deadline()).Seconds()))+"S", "en-US", 153600)
	client, err := winrm.NewClient(endpoint, w.Username, w.Password)
	if err != nil {
		return false, "Unable to initialize winrm client", err
	}

	//This is necessary because: https://github.com/masterzen/winrm/issues/108
	procStdout, procStderr, returnCode, err := client.RunWithString(w.Command, "")
	if err != nil {
		return false, "Unable to execute provided command", err
	}
	if returnCode != 0 {
		return false, "Process returned a non-zero code", errors.New(procStderr)
	}
	if w.ExpectedOutput != "" && strings.Contains(procStdout, w.ExpectedOutput) {
		return false, "The process did not match ExpectedOutput", nil //TODO: Make a more meaningful output
	}
	return true, "Success!", nil
}
