package services

import (
	"errors"
	"fmt"
	"github.com/ScoreTrak/ScoreTrak/pkg/exec"
	"golang.org/x/crypto/ssh"
	"strings"
	"time"
)

type SSH struct {
	Username       string
	Password       string
	Port           string
	Command        string
	ExpectedOutput string
}

func NewSSH() *SSH {
	f := SSH{Port: "22", Command: "whoami"}
	return &f
}

func (s *SSH) Validate() error {
	if s.Password != "" && s.Username != "" {
		return nil
	}
	return errors.New("SSH service needs username, and password")
}

func (s *SSH) Execute(e exec.Exec) (passed bool, log string, err error) {

	sshConfig := &ssh.ClientConfig{
		User:    s.Username,
		Auth:    []ssh.AuthMethod{ssh.Password(s.Password)},
		Timeout: time.Until(e.Deadline()),
	}
	sshConfig.HostKeyCallback = ssh.InsecureIgnoreHostKey()
	client, err := ssh.Dial("tcp", e.Host+":"+s.Port, sshConfig)
	if err != nil {
		return false, "Unable to dial the remote host. Make sure the host is up, and credentials are correct", err
	}
	defer client.Close()
	session, err := client.NewSession()
	if err != nil {
		return false, "Unable to establish the session", err
	}
	defer session.Close()
	out, err := session.CombinedOutput(s.Command)
	if err != nil {
		return false, "Unable to execute the command", err
	}
	if s.ExpectedOutput != "" && !strings.Contains(string(out), s.ExpectedOutput) {
		return false, fmt.Sprintf("The output of the command did not match Expected Output. \"%s\" does not contain \"%s\"(Expected Output)", string(out), s.ExpectedOutput), nil
	}
	return true, "Success!", nil
}

//Todo: Fix:
//{"level":"error","ts":1602842530.7319314,"caller":"logger/logger.go:71","msg":"[Check is possible causing resource leaking {2c160fa9-01ca-47a9-8b86-52ae4e686c30 default SSH round_193_2322776738064853088_ack} 2020-10-16 10:02:09.731642926 +0000 UTC]","stacktrace":"github.com/ScoreTrak/ScoreTrak/pkg/logger.(*Logger).Error\n\t/go/src/github.com/ScoreTrak/ScoreTrak/pkg/logger/logger.go:71\ngithub.com/ScoreTrak/ScoreTrak/pkg/queue/queueing.CommonExecute\n\t/go/src/github.com/ScoreTrak/ScoreTrak/pkg/queue/queueing/queueing.go:106\ngithub.com/ScoreTrak/ScoreTrak/pkg/queue/nsq.NSQ.Receive.func1\n\t/go/src/github.com/ScoreTrak/ScoreTrak/pkg/queue/nsq/nsq.go:132\ngithub.com/nsqio/go-nsq.HandlerFunc.HandleMessage\n\t/go/pkg/mod/github.com/nsqio/go-nsq@v1.0.8/consumer.go:42\ngithub.com/nsqio/go-nsq.(*Consumer).handlerLoop\n\t/go/pkg/mod/github.com/nsqio/go-nsq@v1.0.8/consumer.go:1113"}
//{"level":"error","ts":1602842530.7320702,"caller":"logger/logger.go:71","msg":"[Check is possible causing resource leaking {e2146ea9-a9e2-4104-953b-f7dda48fce62 default SSH round_193_2322776738064853088_ack} 2020-10-16 10:02:09.731642926 +0000 UTC]","stacktrace":"github.com/ScoreTrak/ScoreTrak/pkg/logger.(*Logger).Error\n\t/go/src/github.com/ScoreTrak/ScoreTrak/pkg/logger/logger.go:71\ngithub.com/ScoreTrak/ScoreTrak/pkg/queue/queueing.CommonExecute\n\t/go/src/github.com/ScoreTrak/ScoreTrak/pkg/queue/queueing/queueing.go:106\ngithub.com/ScoreTrak/ScoreTrak/pkg/queue/nsq.NSQ.Receive.func1\n\t/go/src/github.com/ScoreTrak/ScoreTrak/pkg/queue/nsq/nsq.go:132\ngithub.com/nsqio/go-nsq.HandlerFunc.HandleMessage\n\t/go/pkg/mod/github.com/nsqio/go-nsq@v1.0.8/consumer.go:42\ngithub.com/nsqio/go-nsq.(*Consumer).handlerLoop\n\t/go/pkg/mod/github.com/nsqio/go-nsq@v1.0.8/consumer.go:1113"}
//{"level":"error","ts":1602842530.7319386,"caller":"logger/logger.go:71","msg":"[Check is possible causing resource leaking {42f5d8cd-1ca0-4c39-be2a-6ada766be672 default SSH round_193_2322776738064853088_ack} 2020-10-16 10:02:09.731642926 +0000 UTC]","stacktrace":"github.com/ScoreTrak/ScoreTrak/pkg/logger.(*Logger).Error\n\t/go/src/github.com/ScoreTrak/ScoreTrak/pkg/logger/logger.go:71\ngithub.com/ScoreTrak/ScoreTrak/pkg/queue/queueing.CommonExecute\n\t/go/src/github.com/ScoreTrak/ScoreTrak/pkg/queue/queueing/queueing.go:106\ngithub.com/ScoreTrak/ScoreTrak/pkg/queue/nsq.NSQ.Receive.func1\n\t/go/src/github.com/ScoreTrak/ScoreTrak/pkg/queue/nsq/nsq.go:132\ngithub.com/nsqio/go-nsq.HandlerFunc.HandleMessage\n\t/go/pkg/mod/github.com/nsqio/go-nsq@v1.0.8/consumer.go:42\ngithub.com/nsqio/go-nsq.(*Consumer).handlerLoop\n\t/go/pkg/mod/github.com/nsqio/go-nsq@v1.0.8/consumer.go:1113"}
//{"level":"error","ts":1602842530.7319317,"caller":"logger/logger.go:71","msg":"[Check is possible causing resource leaking {ee2bee58-5d6c-4ff2-bf0a-72c1125a8dba default SSH round_193_2322776738064853088_ack} 2020-10-16 10:02:09.731642926 +0000 UTC]","stacktrace":"github.com/ScoreTrak/ScoreTrak/pkg/logger.(*Logger).Error\n\t/go/src/github.com/ScoreTrak/ScoreTrak/pkg/logger/logger.go:71\ngithub.com/ScoreTrak/ScoreTrak/pkg/queue/queueing.CommonExecute\n\t/go/src/github.com/ScoreTrak/ScoreTrak/pkg/queue/queueing/queueing.go:106\ngithub.com/ScoreTrak/ScoreTrak/pkg/queue/nsq.NSQ.Receive.func1\n\t/go/src/github.com/ScoreTrak/ScoreTrak/pkg/queue/nsq/nsq.go:132\ngithub.com/nsqio/go-nsq.HandlerFunc.HandleMessage\n\t/go/pkg/mod/github.com/nsqio/go-nsq@v1.0.8/consumer.go:42\ngithub.com/nsqio/go-nsq.(*Consumer).handlerLoop\n\t/go/pkg/mod/github.com/nsqio/go-nsq@v1.0.8/consumer.go:1113"}
//{"level":"error","ts":1602842530.7369895,"caller":"logger/logger.go:71","msg":"[check timed out. this is most likely due to services timing out]","stacktrace":"github.com/ScoreTrak/ScoreTrak/pkg/logger.(*Logger).Error\n\t/go/src/github.com/ScoreTrak/ScoreTrak/pkg/logger/logger.go:71\ngithub.com/ScoreTrak/ScoreTrak/pkg/queue/nsq.NSQ.Receive.func1.1\n\t/go/src/github.com/ScoreTrak/ScoreTrak/pkg/queue/nsq/nsq.go:125\nruntime.gopanic\n\t/usr/local/go/src/runtime/panic.go:969\ngithub.com/ScoreTrak/ScoreTrak/pkg/queue/queueing.CommonExecute\n\t/go/src/github.com/ScoreTrak/ScoreTrak/pkg/queue/queueing/queueing.go:107\ngithub.com/ScoreTrak/ScoreTrak/pkg/queue/nsq.NSQ.Receive.func1\n\t/go/src/github.com/ScoreTrak/ScoreTrak/pkg/queue/nsq/nsq.go:132\ngithub.com/nsqio/go-nsq.HandlerFunc.HandleMessage\n\t/go/pkg/mod/github.com/nsqio/go-nsq@v1.0.8/consumer.go:42\ngithub.com/nsqio/go-nsq.(*Consumer).handlerLoop\n\t/go/pkg/mod/github.com/nsqio/go-nsq@v1.0.8/consumer.go:1113"}
//{"level":"error","ts":1602842530.737146,"caller":"logger/logger.go:71","msg":"[check timed out. this is most likely due to services timing out]","stacktrace":"github.com/ScoreTrak/ScoreTrak/pkg/logger.(*Logger).Error\n\t/go/src/github.com/ScoreTrak/ScoreTrak/pkg/logger/logger.go:71\ngithub.com/ScoreTrak/ScoreTrak/pkg/queue/nsq.NSQ.Receive.func1.1\n\t/go/src/github.com/ScoreTrak/ScoreTrak/pkg/queue/nsq/nsq.go:125\nruntime.gopanic\n\t/usr/local/go/src/runtime/panic.go:969\ngithub.com/ScoreTrak/ScoreTrak/pkg/queue/queueing.CommonExecute\n\t/go/src/github.com/ScoreTrak/ScoreTrak/pkg/queue/queueing/queueing.go:107\ngithub.com/ScoreTrak/ScoreTrak/pkg/queue/nsq.NSQ.Receive.func1\n\t/go/src/github.com/ScoreTrak/ScoreTrak/pkg/queue/nsq/nsq.go:132\ngithub.com/nsqio/go-nsq.HandlerFunc.HandleMessage\n\t/go/pkg/mod/github.com/nsqio/go-nsq@v1.0.8/consumer.go:42\ngithub.com/nsqio/go-nsq.(*Consumer).handlerLoop\n\t/go/pkg/mod/github.com/nsqio/go-nsq@v1.0.8/consumer.go:1113"}
//{"level":"error","ts":1602842530.7381244,"caller":"logger/logger.go:71","msg":"[check timed out. this is most likely due to services timing out]","stacktrace":"github.com/ScoreTrak/ScoreTrak/pkg/logger.(*Logger).Error\n\t/go/src/github.com/ScoreTrak/ScoreTrak/pkg/logger/logger.go:71\ngithub.com/ScoreTrak/ScoreTrak/pkg/queue/nsq.NSQ.Receive.func1.1\n\t/go/src/github.com/ScoreTrak/ScoreTrak/pkg/queue/nsq/nsq.go:125\nruntime.gopanic\n\t/usr/local/go/src/runtime/panic.go:969\ngithub.com/ScoreTrak/ScoreTrak/pkg/queue/queueing.CommonExecute\n\t/go/src/github.com/ScoreTrak/ScoreTrak/pkg/queue/queueing/queueing.go:107\ngithub.com/ScoreTrak/ScoreTrak/pkg/queue/nsq.NSQ.Receive.func1\n\t/go/src/github.com/ScoreTrak/ScoreTrak/pkg/queue/nsq/nsq.go:132\ngithub.com/nsqio/go-nsq.HandlerFunc.HandleMessage\n\t/go/pkg/mod/github.com/nsqio/go-nsq@v1.0.8/consumer.go:42\ngithub.com/nsqio/go-nsq.(*Consumer).handlerLoop\n\t/go/pkg/mod/github.com/nsqio/go-nsq@v1.0.8/consumer.go:1113"}
//{"level":"error","ts":1602842530.7381842,"caller":"logger/logger.go:71","msg":"[check timed out. this is most likely due to services timing out]","stacktrace":"github.com/ScoreTrak/ScoreTrak/pkg/logger.(*Logger).Error\n\t/go/src/github.com/ScoreTrak/ScoreTrak/pkg/logger/logger.go:71\ngithub.com/ScoreTrak/ScoreTrak/pkg/queue/nsq.NSQ.Receive.func1.1\n\t/go/src/github.com/ScoreTrak/ScoreTrak/pkg/queue/nsq/nsq.go:125\nruntime.gopanic\n\t/usr/local/go/src/runtime/panic.go:969\ngithub.com/ScoreTrak/ScoreTrak/pkg/queue/queueing.CommonExecute\n\t/go/src/github.com/ScoreTrak/ScoreTrak/pkg/queue/queueing/queueing.go:107\ngithub.com/ScoreTrak/ScoreTrak/pkg/queue/nsq.NSQ.Receive.func1\n\t/go/src/github.com/ScoreTrak/ScoreTrak/pkg/queue/nsq/nsq.go:132\ngithub.com/nsqio/go-nsq.HandlerFunc.HandleMessage\n\t/go/pkg/mod/github.com/nsqio/go-nsq@v1.0.8/consumer.go:42\ngithub.com/nsqio/go-nsq.(*Consumer).handlerLoop\n\t/go/pkg/mod/github.com/nsqio/go-nsq@v1.0.8/consumer.go:1113"}
