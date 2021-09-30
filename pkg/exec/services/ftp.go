package services

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/ScoreTrak/ScoreTrak/pkg/exec"
	"github.com/jlaffaye/ftp"
)

type FTP struct {
	Username string
	Password string
	Port     string
	// Text to upload to remote computer
	Text string
	// Name of the file to which write the parameter
	WriteFilename string
	// Filename to find on a remote computer
	ReadFilename string
	// Check expected output of the file
	ExpectedOutput string
}

func NewFTP() *FTP {
	f := FTP{Port: "21"}
	return &f
}

func (f *FTP) Validate() error {
	if f.Password != "" && f.Username != "" {
		if f.Text != "" || f.ReadFilename != "" {
			return nil
		}
		return errors.New("FTP check_service needs either text, or read_file parameter")
	}
	return errors.New("FTP check_service needs username and password")
}

func (f *FTP) Execute(e exec.Exec) (passed bool, l string, err error) {
	dialOptions := ftp.DialWithContext(e.Context)
	c, err := ftp.Dial(fmt.Sprintf("%s:%s", e.Host, f.Port), dialOptions) // For passive FTP allow Data Channel Port Range. In addition, Allow FTP as an APP in windows firewall, and allow port 20, 21, 1024-65535
	if err != nil {
		return false, "", fmt.Errorf("unable to dial FTP Server: %w", err)
	}
	defer func(c *ftp.ServerConn) {
		err := c.Quit()
		if err != nil {
			log.Println(fmt.Errorf("unable to close ftp connection: %w", err))
		}
	}(c)
	err = c.Login(f.Username, f.Password)
	if err != nil {
		return false, "", fmt.Errorf("unable to Login: %w", err)
	}
	defer func(c *ftp.ServerConn) {
		err := c.Logout()
		if err != nil {
			log.Println(fmt.Errorf("unable to logout from FTP: %w", err))
		}
	}(c)
	if f.Text != "" {
		data := bytes.NewBufferString(f.Text)
		if f.WriteFilename == "" {
			f.WriteFilename = "test_file.txt"
		}

		err = c.Stor(f.WriteFilename, data)
		if err != nil {
			return false, "", fmt.Errorf("unable to Store file to FTP server: %w", err)
		}
	}
	if f.ReadFilename != "" {
		r, err := c.Retr(f.ReadFilename)
		if err != nil {
			return false, "", fmt.Errorf("failed to retrieve the file from FTP: %w", err)
		}
		defer func(r *ftp.Response) {
			err := r.Close()
			if err != nil {
				log.Println(fmt.Errorf("unable to close file: %w", err))
			}
		}(r)
		if err := c.Quit(); err != nil {
			return false, "", fmt.Errorf("unable to gracefully exit FTP server: %w", err)
		}

		buf, err := ioutil.ReadAll(r)
		if err != nil {
			return false, "", fmt.Errorf("failed to read file contents, it might be corrupted: %w", err)
		}
		if f.ExpectedOutput != "" && !strings.Contains(string(buf), f.ExpectedOutput) {
			return false, "", fmt.Errorf("fetched file's contents do not match Expected Output. Output received: %s", string(buf))
		}
	}
	return true, Success, nil
}
