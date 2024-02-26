package executors

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/scoretrak/scoretrak/pkg/scorer/outcome"
	"github.com/scoretrak/scoretrak/pkg/scorer/outcome_writer"
	"log"
	"strings"
	"time"

	"golang.org/x/crypto/ssh"
)

type SSHProperties struct {
	Username       string `json:"username" validate:"required"`
	Password       string `json:"password" validate:"required"`
	Host           string `json:"host"`
	Port           string `json:"port"`
	Command        string `json:"command"`
	ExpectedOutput string `json:"expected_output"`
}

//var ErrSSHRequiresUsernameAndPassword = errors.New("ssh check_service needs username, and password")

func ScoreSSH(ctx context.Context, ow *outcome_writer.OutcomeWriter, properties []byte) {
	sshproperties := &SSHProperties{}
	err := json.Unmarshal(properties, &sshproperties)
	if err != nil {
		ow.SetError(fmt.Errorf("unable to unmarshall properties: %v", err))
		return
	}

	err = validate.Struct(sshproperties)
	if err != nil {
		ow.SetError(fmt.Errorf("validation error: %w", err))
		return
	}

	var sshConfig *ssh.ClientConfig

	timeout, ok := ctx.Deadline()
	if ok {
		sshConfig = &ssh.ClientConfig{
			User:    sshproperties.Username,
			Auth:    []ssh.AuthMethod{ssh.Password(sshproperties.Password)},
			Timeout: time.Until(timeout),
		}
	} else {
		sshConfig = &ssh.ClientConfig{
			User: sshproperties.Username,
			Auth: []ssh.AuthMethod{ssh.Password(sshproperties.Password)},
		}
	}
	sshConfig.HostKeyCallback = ssh.InsecureIgnoreHostKey() //nolint:gosec
	address := fmt.Sprintf("%s:%s", sshproperties.Host, sshproperties.Port)
	client, err := ssh.Dial("tcp", address, sshConfig)
	if err != nil {
		ow.SetError(fmt.Errorf("unable to dial the remote host. Make sure the host is up, and credentials are correct: %w", err))
		return
	}
	defer func(client *ssh.Client) {
		err := client.Close()
		if err != nil {
			log.Println(fmt.Errorf("unable to close ssh client: %w", err))
		}
	}(client)
	session, err := client.NewSession()
	if err != nil {
		ow.SetError(fmt.Errorf("unable to establish the session: %w", err))
		return
	}
	out, err := session.CombinedOutput(sshproperties.Command)
	if err != nil {
		ow.SetError(fmt.Errorf("unable to execute the command: %w", err))
		return
	}
	if sshproperties.ExpectedOutput != "" && !strings.Contains(string(out), sshproperties.ExpectedOutput) {
		ow.SetError(fmt.Errorf("%w. Output Received: %s", ErrDidNotMatchExpectedOutput, string(out)))
		return
	}

	ow.SetStatus(outcome.OUTCOME_STATUS_PASSED)
}
