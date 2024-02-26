package executors

//import (
//	"context"
//	"encoding/json"
//	"errors"
//	"fmt"
//	"github.com/scoretrak/scoretrak/pkg/scorer"
//	"net"
//	"strconv"
//	"strings"
//	"time"
//
//	"github.com/masterzen/winrm"
//)
//
//type WinrmProperties struct {
//	Username       string
//	Password       string
//	Host           string
//	Port           string
//	Command        string
//	ExpectedOutput string
//	Scheme         string
//	ClientType     string
//}
//
//var ErrWinrmRequiresUsernameAndPassword = errors.New("winrm check_service needs username, and password")
//
////func (w *Winrm) Validate() error {
////	if w.Password != "" && w.Username != "" {
////		return nil
////	}
////	return ErrWinrmRequiresUsernameAndPassword
////}
//
//func ScoreWinrm(ctx context.Context, ow *outcome_writer.OutcomeWriter, properties []byte) {
//	winrmproperties := &WinrmProperties{}
//	err := json.Unmarshal(properties, &winrmproperties)
//	if err != nil {
//		return &scorer.Outcome{
//			Passed: false,
//			Error:  fmt.Errorf("unable to unmarshall properties: %v", err),
//		}
//	}
//
//	isHTTPS := IsSecure(winrmproperties.Scheme)
//	if winrmproperties.Port == "" {
//		if isHTTPS {
//			winrmproperties.Port = "5986"
//		} else {
//			winrmproperties.Port = "5985"
//		}
//	}
//	i, err := strconv.Atoi(winrmproperties.Port)
//	if err != nil {
//		return &scorer.Outcome{
//			Passed: false,
//			Error:  fmt.Errorf("unable to convert port number to integer: %w", err),
//		}
//	}
//
//	timeout, _ := ctx.Deadline()
//
//	err = tcpPortDial(net.JoinHostPort(winrmproperties.Host, winrmproperties.Port), time.Until(timeout)/3)
//	if err != nil {
//		return &scorer.Outcome{
//			Passed: false,
//			Error:  err,
//		}
//	}
//
//	endpoint := winrm.NewEndpoint(winrmproperties.Host, i, isHTTPS, true, nil, nil, nil, time.Until(timeout))
//	params := winrm.DefaultParameters
//	params.Dial = (&net.Dialer{
//		Timeout: time.Until(timeout),
//	}).Dial
//	if strings.ToLower(winrmproperties.ClientType) == "ntlm" {
//		params.TransportDecorator = func() winrm.Transporter { return &winrm.ClientNTLM{} }
//	}
//	client, err := winrm.NewClientWithParameters(endpoint, winrmproperties.Username, winrmproperties.Password, params)
//	if err != nil {
//		return &scorer.Outcome{
//			Passed: false,
//			Error:  fmt.Errorf("unable to initialize winrm client: %w", err),
//		}
//	}
//	procStdout, procStderr, returnCode, err := client.RunWithString(winrmproperties.Command, "")
//	if err != nil {
//		return &scorer.Outcome{
//			Passed: false,
//			Error:  fmt.Errorf("unable to execute provided command: %w", err),
//		}
//	}
//	if returnCode != 0 {
//		return &scorer.Outcome{
//			Passed: false,
//			Error:  fmt.Errorf("%w: %s", ErrNonZeroReturn, procStderr),
//		}
//	}
//	if winrmproperties.ExpectedOutput != "" && !strings.Contains(procStdout, strings.ToLower(winrmproperties.ExpectedOutput)) {
//		return &scorer.Outcome{
//			Passed: false,
//			Error:  fmt.Errorf("%w. Output Received: %s", ErrDidNotMatchExpectedOutput, procStdout),
//		}
//	}
//	return &scorer.Outcome{
//		Passed: true,
//		Error:  err,
//	}
//}
//
//var ErrNonZeroReturn = errors.New("process returned a non-zero code")
//
//// endpoint := winrm.NewEndpoint(e.HostAddress, i, isHttps, true, nil, nil, nil, time.Until(e.Deadline()))
//// // params := winrm.NewParameters(strconv.Itoa(int(time.Until(e.Deadline()).Seconds()))+"S", "en-US", 153600)
//// client, err := winrm.NewClient(endpoint, w.Username, w.Password)
//// if err != nil {
////	return false, "Unable to initialize winrm client", err
//// }
