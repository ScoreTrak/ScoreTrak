package services

import (
	"ScoreTrak/pkg/exec"
	"bytes"
	"errors"
	"fmt"
	"github.com/jlaffaye/ftp"
	"io/ioutil"
)

type FTP struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Port     string `json:"port"`
	//Text to upload to remote computer
	Text string `json:"text"`
	//Name of the file to which write the parameter
	WriteFilename string `json:"write_file"`
	//Filename to find on a remote computer
	ReadFilename string `json:"read_file"`
	//Check expected output of the file
	ExpectedOutput string `json:"expected_output"`
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
		return errors.New("FTP service needs either text, or read_file parameter")
	}
	return errors.New("FTP service needs username and password")
}

func (f *FTP) Execute(e exec.Exec) (passed bool, log string, err error) {
	dialOptions := ftp.DialWithContext(e.Context)
	c, err := ftp.Dial(fmt.Sprintf("%s:%s", e.Host, f.Port), dialOptions) // For passive FTP allow Data Channel Port Range. In addition, Allow FTP as an APP in windows firewall, and allow port 20, 21, 1024-65535
	if err != nil {
		return false, "Unable to dial FTP Server", err
	}
	err = c.Login(f.Username, f.Password)
	if err != nil {
		return false, "Unable to Login", err
	}
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
		defer r.Close()

		if err != nil {
			return false, "Failed to Retrieve the file from FTP", err
		}

		if err := c.Quit(); err != nil {
			return false, "Unable to gracefully exit FTP server", err
		}

		buf, err := ioutil.ReadAll(r)

		if f.ExpectedOutput != "" && string(buf) != f.ExpectedOutput {
			return false, "Fetched file's contents do not match Expected Output", nil //TODO: Make a more meaningful output
		}
	}
	return true, "Success!", nil
}
