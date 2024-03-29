package services

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"strings"

	"github.com/ScoreTrak/ScoreTrak/pkg/exec"
	"github.com/hirochachacha/go-smb2"
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

var ErrShareShouldNotBeEmpty = errors.New("parameter Share should not be empty")
var ErrParameterShouldBeEither = errors.New("parameter Operation should be one of the following")

func (s *SMB) Validate() error {
	if s.Operation != create && s.Operation != createAndOpen && s.Operation != open {
		return fmt.Errorf("%w: %s, %s, %s", ErrParameterShouldBeEither, create, open, createAndOpen)
	}
	if s.Share == "" {
		return ErrShareShouldNotBeEmpty
	}
	return nil
}

func (s *SMB) Execute(e exec.Exec) (passed bool, logOutput string, err error) {
	dial := net.Dialer{}
	conn, err := dial.DialContext(e.Context, s.TransportProtocol, e.Host+":"+s.Port)
	if err != nil {
		return false, "", fmt.Errorf("unable to dial the host: %w", err)
	}
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			log.Println(fmt.Errorf("unable to close the connection: %w", err))
		}
	}(conn)
	d := &smb2.Dialer{
		Initiator: &smb2.NTLMInitiator{
			User:     s.Username,
			Password: s.Password,
			Domain:   s.Domain,
		},
	}
	c, err := d.Dial(conn)
	if err != nil {
		return false, "", fmt.Errorf("unable to dial the host: %w", err)
	}
	defer func(c *smb2.Client) {
		err := c.Logoff()
		if err != nil {
			log.Println(fmt.Errorf("unable to logoff: %w", err))
		}
	}(c)

	fs, err := c.Mount(`\\` + e.Host + `\` + s.Share)
	if err != nil {
		return false, "", fmt.Errorf("unable to mount the share: %w", err)
	}
	defer func(fs *smb2.RemoteFileSystem) {
		err := fs.Umount()
		if err != nil {
			log.Println(fmt.Errorf("unable to unmount file system: %w", err))
		}
	}(fs)
	if s.FileName != "" {
		var f *smb2.RemoteFile
		if strings.Contains(s.Operation, create) {
			f, err = fs.Create(s.FileName)
			if err != nil {
				return false, "", fmt.Errorf("unable to create filename on a remote computer: %w", err)
			}
			_, err = f.Write([]byte(s.Text))
			if err != nil {
				return false, "", fmt.Errorf("unable to write into file: %w", err)
			}
		} else {
			f, err = fs.Open(s.FileName)
			if err != nil {
				return false, "", fmt.Errorf("unable to open the filename on a remote computer: %w", err)
			}
			if s.Operation == create {
				return true, Success, nil
			}
		}
		defer func(f *smb2.RemoteFile) {
			err := f.Close()
			if err != nil {
				log.Println(fmt.Errorf("unable to close remote file: %w", err))
			}
		}(f)
		_, err = f.Seek(0, io.SeekStart)
		if err != nil {
			return false, "", fmt.Errorf("unable to read the file: %w", err)
		}
		bs, err := ioutil.ReadAll(f)
		if err != nil {
			return false, "", fmt.Errorf("unable to read the file: %w", err)
		}
		if s.ExpectedOutput != "" && string(bs) != s.ExpectedOutput {
			return false, "", fmt.Errorf("%w. Output Received: %s", ErrDidNotMatchExpectedOutput, string(bs))
		}
	}
	return true, Success, nil
}
