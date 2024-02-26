package executors

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/emersion/go-imap/client"
	"github.com/scoretrak/scoretrak/pkg/scorer/outcome"
	"github.com/scoretrak/scoretrak/pkg/scorer/outcome_writer"
	"log"
)

type IMAPProperties struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Host     string
	Port     string `default:"993"`
	Scheme   string
}

var ErrIMAPRequiresUsernameAndPassword = errors.New("IMAP check_service needs password, and username to operate")

func ScoreImap(ctx context.Context, ow *outcome_writer.OutcomeWriter, properties []byte) {
	imapproperties := &IMAPProperties{}
	err := json.Unmarshal(properties, &imapproperties)
	if err != nil {
		ow.SetError(fmt.Errorf("unable to unmarshall properties: %v", err))
		return
	}

	err = validate.Struct(imapproperties)
	if err != nil {
		// validationErrors := err.(validator.ValidationErrors)
		ow.SetError(fmt.Errorf("validation error: %w", err))
		return
	}

	address := fmt.Sprintf("%s:%s", imapproperties.Host, imapproperties.Port)
	var c *client.Client
	if isHTTPS := IsSecure(imapproperties.Scheme); isHTTPS {
		if imapproperties.Port == "" {
			imapproperties.Port = "993"
		}
		c, err = client.DialTLS(address, &tls.Config{InsecureSkipVerify: true}) //nolint:gosec
	} else {
		if imapproperties.Port == "" {
			imapproperties.Port = "143"
		}
		c, err = client.Dial(address)
	}

	if err != nil {
		ow.SetError(fmt.Errorf("unable to dial the remote server: %w", err))
		return
	}
	defer func(c *client.Client) {
		err := c.Close()
		if err != nil {
			log.Println(fmt.Errorf("unable to close client: %w", err))
		}
	}(c)
	if err := c.Login(imapproperties.Username, imapproperties.Password); err != nil {
		ow.SetError(fmt.Errorf("unable to login with the credentials provided: %w", err))
		return
	}
	defer func(c *client.Client) {
		err := c.Logout()
		if err != nil {
			log.Println(fmt.Errorf("unable to logout: %w", err))
		}
	}(c)
	// List mailboxes
	//mailboxes := make(chan *imap.MailboxInfo, 10)
	//done := make(chan error, 1)
	//go func() {
	//	done <- c.List("", "*", mailboxes)
	//}()
	//logOutput = Success + "\n"
	//for m := range mailboxes {
	//	logOutput += "* " + m.Name
	//}
	ow.SetStatus(outcome.OUTCOME_STATUS_PASSED)
}
