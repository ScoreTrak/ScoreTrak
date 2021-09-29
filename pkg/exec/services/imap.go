package services

import (
	"crypto/tls"
	"errors"
	"github.com/ScoreTrak/ScoreTrak/pkg/exec"
	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
	"net"
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

func (i *IMAP) Execute(e exec.Exec) (passed bool, log string, err error) {
	isHttps := exec.IsSecure(i.Scheme)
	if i.Port == "" {
		if isHttps {
			i.Port = "993"
		} else {
			i.Port = "143"
		}
	}
	var c *client.Client
	if isHttps {
		c, err = client.DialWithDialerTLS(&net.Dialer{Deadline: e.Deadline()}, e.Host+":"+i.Port, &tls.Config{InsecureSkipVerify: true})
	} else {
		c, err = client.DialWithDialer(&net.Dialer{Deadline: e.Deadline()}, e.Host+":"+i.Port)
	}
	if err != nil {
		return false, "Unable to pass Dial the remote server", err
	}
	defer func(){ _ = c.Close() }()
	if err := c.Login(i.Username, i.Password); err != nil {
		return false, "Unable to login with the credentials provided", err
	}
	defer func(){ _ = c.Logout() }()
	// List mailboxes
	mailboxes := make(chan *imap.MailboxInfo, 10)
	done := make(chan error, 1)
	go func() {
		done <- c.List("", "*", mailboxes)
	}()
	log = "Success!\n"
	for m := range mailboxes {
		log += "* " + m.Name
	}
	return true, log, nil
}
