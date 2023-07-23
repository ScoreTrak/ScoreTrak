package services

import (
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/ScoreTrak/ScoreTrak/pkg/exec"
	"golang.org/x/crypto/ssh"
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

var ErrSSHRequiresUsernameAndPassword = errors.New("ssh check_service needs username, and password")

func (s *SSH) Validate() error {
	if s.Password != "" && s.Username != "" {
		return nil
	}
	return ErrSSHRequiresUsernameAndPassword
}

func (s *SSH) Execute(e exec.Exec) (passed bool, logOutput string, err error) {
	sshConfig := &ssh.ClientConfig{
		User:    s.Username,
		Auth:    []ssh.AuthMethod{ssh.Password(s.Password)},
		Timeout: time.Until(e.Deadline()),
	}
	sshConfig.HostKeyCallback = ssh.InsecureIgnoreHostKey() //nolint:gosec
	client, err := ssh.Dial("tcp", e.HostAddress+":"+s.Port, sshConfig)
	if err != nil {
		return false, "", fmt.Errorf("unable to dial the remote host. Make sure the host is up, and credentials are correct: %w", err)
	}
	defer func(client *ssh.Client) {
		err := client.Close()
		if err != nil {
			log.Println(fmt.Errorf("unable to close ssh client: %w", err))
		}
	}(client)
	session, err := client.NewSession()
	if err != nil {
		return false, "", fmt.Errorf("unable to establish the session: %w", err)
	}
	out, err := session.CombinedOutput(s.Command)
	if err != nil {
		return false, "", fmt.Errorf("unable to execute the command: %w", err)
	}
	if s.ExpectedOutput != "" && !strings.Contains(string(out), s.ExpectedOutput) {
		return false, "", fmt.Errorf("%w. Output Received: %s", ErrDidNotMatchExpectedOutput, string(out))
	}
	return true, Success, nil
}
