package services

import (
	"errors"
	"fmt"
	"github.com/ScoreTrak/ScoreTrak/pkg/exec"
	"github.com/masterzen/winrm"
	"net"
	"strconv"
	"strings"
	"time"
)

type Winrm struct {
	Username       string
	Password       string
	Port           string
	Command        string
	ExpectedOutput string
	Scheme         string
	ClientType     string
}

func NewWinrm() *Winrm {
	return &Winrm{Command: "whoami", Scheme: "http", ClientType: "NTLM"}
}

func (w *Winrm) Validate() error {
	if w.Password != "" && w.Username != "" {
		return nil
	}
	return errors.New("winrm check_service needs username, and password")
}

func (w *Winrm) Execute(e exec.Exec) (passed bool, log string, err error) {
	isHttps := exec.IsSecure(w.Scheme)
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
	params := winrm.DefaultParameters
	params.Dial = (&net.Dialer{
		Timeout: time.Until(e.Deadline()),
	}).Dial
	if strings.ToLower(w.ClientType) == "ntlm" {
		params.TransportDecorator = func() winrm.Transporter { return &winrm.ClientNTLM{} }
	}
	client, err := winrm.NewClientWithParameters(endpoint, w.Username, w.Password, params)
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
	if w.ExpectedOutput != "" && !strings.Contains(procStdout, strings.ToLower(w.ExpectedOutput)) {
		return false, fmt.Sprintf("The output of the command did not match Expected Output. Output Received: %s", procStdout), nil
	}
	return true, "Success!", nil
}

//endpoint := winrm.NewEndpoint(e.Host, i, isHttps, true, nil, nil, nil, time.Until(e.Deadline()))
////params := winrm.NewParameters(strconv.Itoa(int(time.Until(e.Deadline()).Seconds()))+"S", "en-US", 153600)
//client, err := winrm.NewClient(endpoint, w.Username, w.Password)
//if err != nil {
//	return false, "Unable to initialize winrm client", err
//}
