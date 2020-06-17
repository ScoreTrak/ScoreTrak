package services

import (
	"ScoreTrak/pkg/exec"
	"errors"
	"golang.org/x/crypto/ssh"
	"time"
)

type SSH struct {
	Username       string `json:"username"`
	Password       string `json:"password"`
	Port           string `json:"port"`
	Command        string `json:"command"`
	ExpectedOutput string `json:"expected_output"`
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
	client, err := ssh.Dial("tcp", e.Host, sshConfig)
	if err != nil {
		return false, "Unable to dial the remote host", err
	}
	session, err := client.NewSession()
	if err != nil {
		return false, "Unable to establish the session", err
	}
	defer session.Close()
	out, err := session.CombinedOutput("mkdir test")
	if err != nil {
		return false, "Unable to execute the command", err
	}
	if s.ExpectedOutput != "" && string(out) != s.ExpectedOutput {
		return false, "The output of the command did not match Expected Output", nil //TODO: Make a more meaningful output
	}
	return true, "Success!", nil
}
