package services

import (
	"crypto/tls"
	"errors"
	"fmt"
	"log"
	"net"

	"github.com/ScoreTrak/ScoreTrak/pkg/exec"
	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
)

type IMAP struct {
	Username string
	Password string
	Port     string
	Scheme   string
}

func NewIMAP() *IMAP {
	f := IMAP{}
	return &f
}

func (i *IMAP) Validate() error {
	if i.Password != "" && i.Username != "" {
		return nil
	}
	return errors.New("IMAP check_service needs password, and username to operate")
}

func (i *IMAP) Execute(e exec.Exec) (passed bool, logOutput string, err error) {
	var c *client.Client
	if isHTTPS := IsSecure(i.Scheme); isHTTPS {
		if i.Port == "" {
			i.Port = "993"
		}
		c, err = client.DialWithDialerTLS(&net.Dialer{Deadline: e.Deadline()}, e.Host+":"+i.Port, &tls.Config{InsecureSkipVerify: true})
	} else {
		if i.Port == "" {
			i.Port = "143"
		}
		c, err = client.DialWithDialer(&net.Dialer{Deadline: e.Deadline()}, e.Host+":"+i.Port)
	}
	if err != nil {
		return false, "", fmt.Errorf("unable to dial the remote server: %w", err)
	}
	defer func(c *client.Client) {
		err := c.Close()
		if err != nil {
			log.Println(fmt.Errorf("unable to close client: %w", err))
		}
	}(c)
	if err := c.Login(i.Username, i.Password); err != nil {
		return false, "", fmt.Errorf("unable to login with the credentials provided: %w", err)
	}
	defer func(c *client.Client) {
		err := c.Logout()
		if err != nil {
			log.Println(fmt.Errorf("unable to logout: %w", err))
		}
	}(c)
	// List mailboxes
	mailboxes := make(chan *imap.MailboxInfo, 10)
	done := make(chan error, 1)
	go func() {
		done <- c.List("", "*", mailboxes)
	}()
	logOutput = Success + "\n"
	for m := range mailboxes {
		logOutput += "* " + m.Name
	}
	return true, logOutput, nil
}
