package services

import (
	"errors"
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"

	"github.com/ScoreTrak/ScoreTrak/pkg/exec"
	"github.com/masterzen/winrm"
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

var ErrWinrmRequiresUsernameAndPassword = errors.New("winrm check_service needs username, and password")

func (w *Winrm) Validate() error {
	if w.Password != "" && w.Username != "" {
		return nil
	}
	return ErrWinrmRequiresUsernameAndPassword
}

func (w *Winrm) Execute(e exec.Exec) (passed bool, logOutput string, err error) {
	isHTTPS := IsSecure(w.Scheme)
	if w.Port == "" {
		if isHTTPS {
			w.Port = "5986"
		} else {
			w.Port = "5985"
		}
	}
	i, err := strconv.Atoi(w.Port)
	if err != nil {
		return false, "", fmt.Errorf("unable to convert port number to integer: %w", err)
	}
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(e.Host, w.Port), time.Until(e.Deadline())/3)
	if err != nil {
		return false, "", fmt.Errorf("port was not open on a remote host: %w", err)
	}
	_ = conn.Close()
	endpoint := winrm.NewEndpoint(e.Host, i, isHTTPS, true, nil, nil, nil, time.Until(e.Deadline()))
	params := winrm.DefaultParameters
	params.Dial = (&net.Dialer{
		Timeout: time.Until(e.Deadline()),
	}).Dial
	if strings.ToLower(w.ClientType) == "ntlm" {
		params.TransportDecorator = func() winrm.Transporter { return &winrm.ClientNTLM{} }
	}
	client, err := winrm.NewClientWithParameters(endpoint, w.Username, w.Password, params)
	if err != nil {
		return false, "", fmt.Errorf("unable to initialize winrm client: %w", err)
	}
	procStdout, procStderr, returnCode, err := client.RunWithString(w.Command, "")
	if err != nil {
		return false, "", fmt.Errorf("unable to execute provided command: %w", err)
	}
	if returnCode != 0 {
		return false, "", fmt.Errorf("%w: %s", ErrNonZeroReturn, procStderr)
	}
	if w.ExpectedOutput != "" && !strings.Contains(procStdout, strings.ToLower(w.ExpectedOutput)) {
		return false, "", fmt.Errorf("%w. Output Received: %s", ErrDidNotMatchExpectedOutput, procStdout)
	}
	return true, Success, nil
}

var ErrNonZeroReturn = errors.New("process returned a non-zero code")

// endpoint := winrm.NewEndpoint(e.Host, i, isHttps, true, nil, nil, nil, time.Until(e.Deadline()))
// // params := winrm.NewParameters(strconv.Itoa(int(time.Until(e.Deadline()).Seconds()))+"S", "en-US", 153600)
// client, err := winrm.NewClient(endpoint, w.Username, w.Password)
// if err != nil {
//	return false, "Unable to initialize winrm client", err
// }
