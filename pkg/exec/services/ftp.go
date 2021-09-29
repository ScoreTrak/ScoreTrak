package services

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/ScoreTrak/ScoreTrak/pkg/exec"
	"github.com/jlaffaye/ftp"
	"io/ioutil"
	"strings"
)

type FTP struct {
	Username string
	Password string
	Port     string
	//Text to upload to remote computer
	Text string
	//Name of the file to which write the parameter
	WriteFilename string
	//Filename to find on a remote computer
	ReadFilename string
	//Check expected output of the file
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

func (f *FTP) Execute(e exec.Exec) (passed bool, log string, err error) {
	dialOptions := ftp.DialWithContext(e.Context)
	c, err := ftp.Dial(fmt.Sprintf("%s:%s", e.Host, f.Port), dialOptions) // For passive FTP allow Data Channel Port Range. In addition, Allow FTP as an APP in windows firewall, and allow port 20, 21, 1024-65535
	if err != nil {
		return false, "Unable to dial FTP Server", err
	}
	defer c.Quit()
	err = c.Login(f.Username, f.Password)
	if err != nil {
		return false, "Unable to Login", err
	}
	defer c.Logout()
	if f.Text != "" {
		data := bytes.NewBufferString(f.Text)
		if f.WriteFilename == "" {
			f.WriteFilename = "test_file.txt"
		}

		err = c.Stor(f.WriteFilename, data)
		if err != nil {
			return false, "Unable to Store file to FTP server", err
		}
	}
	if f.ReadFilename != "" {
		r, err := c.Retr(f.ReadFilename)
		if err != nil {
			return false, "Failed to Retrieve the file from FTP", err
		}
		defer func(r *ftp.Response) {
			err := r.Close()
			if err != nil {
				fmt.Errorf("unable to close file: %w", err)
			}
		}(r)
		if err := c.Quit(); err != nil {
			return false, "Unable to gracefully exit FTP server", err
		}

		buf, err := ioutil.ReadAll(r)
		if err != nil {
			return false, "Failed to read file contents, it might be corrupted", err
		}
		if f.ExpectedOutput != "" && !strings.Contains(string(buf), f.ExpectedOutput) {
			return false, fmt.Sprintf("Fetched file's contents do not match Expected Output. Output received: %s", string(buf)), nil
		}
	}
	return true, "Success!", nil
}
