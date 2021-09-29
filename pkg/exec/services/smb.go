package services

import (
	"errors"
	"fmt"
	"github.com/ScoreTrak/ScoreTrak/pkg/exec"
	"github.com/hirochachacha/go-smb2"
	"io"
	"io/ioutil"
	"net"
	"strings"
)

const (
	open          = "Open"
	create        = "Create"
	createAndOpen = create + "And" + open
)

type SMB struct {
	Username          string
	Password          string
	Domain            string
	Port              string
	TransportProtocol string
	Share             string
	FileName          string
	Text              string
	Operation         string
	ExpectedOutput    string
}

func NewSMB() *SMB {
	f := SMB{Port: "445", TransportProtocol: "tcp", Text: "Hello World!", Operation: createAndOpen, FileName: "TestFile.txt"}
	return &f
}

func (s *SMB) Validate() error {
	if s.Operation != create && s.Operation != createAndOpen && s.Operation != open {
		return fmt.Errorf("parameter should Operation be either: %s, %s, or %s", create, open, createAndOpen)
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
	defer func(){ _ = conn.Close() }()
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
	defer func(){ _ = c.Logoff()}()

	fs, err := c.Mount(`\\` + e.Host + `\` + s.Share)
	if err != nil {
		return false, "Unable to mount the share", err
	}
	defer func(){ _ = fs.Umount() }()
	if s.FileName != "" {
		var f *smb2.RemoteFile
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
		defer func(){ _ = f.Close() }()
		_, err = f.Seek(0, io.SeekStart)
		if err != nil {
			return false, "Unable to read the file", err
		}
		bs, err := ioutil.ReadAll(f)
		if err != nil {
			return false, "Unable to read the file", nil
		}
		if s.ExpectedOutput != "" && string(bs) != s.ExpectedOutput {
			return false, fmt.Sprintf("Contents of the file did not match expected output. Output Received: %s", string(bs)), nil
		}
	}
	return true, "Success!", nil
}
