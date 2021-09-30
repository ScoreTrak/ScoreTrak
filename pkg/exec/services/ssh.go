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

func (s *SSH) Validate() error {
	if s.Password != "" && s.Username != "" {
		return nil
	}
	return errors.New("SSH check_service needs username, and password")
}

func (s *SSH) Execute(e exec.Exec) (passed bool, logOutput string, err error) {
	sshConfig := &ssh.ClientConfig{
		User:    s.Username,
		Auth:    []ssh.AuthMethod{ssh.Password(s.Password)},
		Timeout: time.Until(e.Deadline()),
	}
	sshConfig.HostKeyCallback = ssh.InsecureIgnoreHostKey()
	client, err := ssh.Dial("tcp", e.Host+":"+s.Port, sshConfig)
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
	defer func(session *ssh.Session) {
		err := session.Close()
		if err != nil {
			log.Println(fmt.Errorf("unable to close ssh session: %w", err))
		}
	}(session)
	out, err := session.CombinedOutput(s.Command)
	if err != nil {
		return false, "", fmt.Errorf("unable to execute the command: %w", err)
	}
	if s.ExpectedOutput != "" && !strings.Contains(string(out), s.ExpectedOutput) {
		return false, fmt.Sprintf("The output of the command did not match Expected Output. Output Received: %s", string(out)), nil
	}
	return true, Success, nil
}
