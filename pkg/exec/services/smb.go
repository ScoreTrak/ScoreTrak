package services

import (
	"ScoreTrak/pkg/exec"
	"errors"
	"fmt"
	"github.com/hirochachacha/go-smb2"
	"io/ioutil"
	"net"
	"os"
	"strings"
)

const (
	open          = "Open"
	create        = "Create"
	createAndOpen = create + "And" + open
)

type SMB struct {
	Username          string `json:"username"`
	Password          string `json:"password"`
	Domain            string `json:"domain"`
	Port              string `json:"port"`
	TransportProtocol string `json:"transport_protocol"`
	Share             string `json:"share"`
	FileName          string `json:"file_name"`
	Text              string `json:"text"`
	Operation         string `json:"text"`
	ExpectedOutput    string `json:"expected_output"`
}

func NewSMB() *SMB {
	f := SMB{Port: "445", TransportProtocol: "tcp", Text: "Hello World!", Operation: createAndOpen, FileName: "TestFile.txt"}
	return &f
}

func (s *SMB) Validate() error {
	if s.Operation != create && s.Operation != createAndOpen && s.Operation != open {
		return errors.New(fmt.Sprintf("parameter should Operation be either: %s, %s, or %s", create, open, createAndOpen))
	}
	if s.Share == "" {
		return errors.New("parameter Share should not be empty")
	}
	return nil
}

func (s *SMB) Execute(e exec.Exec) (passed bool, log string, err error) {
	dial := net.Dialer{}
	conn, err := dial.DialContext(e.Context, s.TransportProtocol, e.Host+":"+s.Port)
	if err != nil {
		return false, "Unable to dial the host", err
	}
	defer conn.Close()
	d := &smb2.Dialer{
		Initiator: &smb2.NTLMInitiator{
			User:     s.Username,
			Password: s.Password,
			Domain:   s.Domain,
		},
	}
	c, err := d.Dial(conn)
	if err != nil {
		return false, "Unable to dial the host", err
	}
	defer c.Logoff()

	fs, err := c.Mount(`\\` + e.Host + `\` + s.Share)
	if err != nil {
		return false, "Unable to mount the share", err
	}
	defer fs.Umount()
	if s.FileName != "" {
		var f *smb2.RemoteFile
		defer f.Close()
		if strings.Contains(s.Operation, create) {
			f, err = fs.Create(s.FileName)
			if err != nil {
				return false, "Unable to create filename on a remote computer", err
			}
			_, err = f.Write([]byte(s.Text))
			if err != nil {
				return false, "Unable to write into file", err
			}
		} else {
			f, err = fs.Open(s.FileName)
			if err != nil {
				return false, "Unable to open the filename on a remote computer", err
			}
			if s.Operation == create {
				return true, "Success!", nil
			}
		}
		_, err = f.Seek(0, os.SEEK_SET)
		if err != nil {
			return false, "Unable to read the file", err
		}
		bs, err := ioutil.ReadAll(f)
		if err != nil {
			return false, "Unable to read the file", nil
		}
		if s.ExpectedOutput != "" && string(bs) != s.ExpectedOutput {
			return false, "Contents of the file did not match expected output", nil //TODO: Make a more meaningful output
		}
	}
	return true, "Success!", nil
}
