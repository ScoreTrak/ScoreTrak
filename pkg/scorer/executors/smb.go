package executors

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/hirochachacha/go-smb2"
	"github.com/scoretrak/scoretrak/pkg/scorer/outcome"
	"github.com/scoretrak/scoretrak/pkg/scorer/outcome_writer"
)

const (
	SMB_OPERATION_OPEN            = "Open"
	SMB_OPERATION_CREATE          = "Create"
	SMB_OPERATION_CREATE_AND_OPEN = "OpenAndCreate"
)

type SMBProperties struct {
	Username          string
	Password          string
	Domain            string
	Host              string
	Port              string
	TransportProtocol string
	Share             string `validate:"required,smb_operation"`
	FileName          string
	Text              string
	Operation         string
	ExpectedOutput    string
}

var ErrShareShouldNotBeEmpty = errors.New("parameter Share should not be empty")
var ErrParameterShouldBeEither = errors.New("parameter Operation should be one of the following")

func validateSmbOperation(fl validator.FieldLevel) bool {
	k, ok := fl.Field().Interface().(string)
	if !ok {
		return false
	}
	if k != SMB_OPERATION_CREATE && k != SMB_OPERATION_CREATE_AND_OPEN && k != SMB_OPERATION_OPEN {
		return false
	}

	return true
}

func ScoreSmb(ctx context.Context, ow *outcome_writer.OutcomeWriter, properties []byte) {
	smbproperties := &SMBProperties{}
	err := json.Unmarshal(properties, &smbproperties)
	if err != nil {
		ow.SetError(fmt.Errorf("unable to unmarshall properties: %v", err))
		return
	}

	err = validate.Struct(smbproperties)
	if err != nil {
		// validationErrors := err.(validator.ValidationErrors)
		ow.SetError(fmt.Errorf("validation error: %w", err))
		return
	}

	address := fmt.Sprintf("%s:%s", smbproperties.Host, smbproperties.Port)
	dial := net.Dialer{}
	conn, err := dial.DialContext(ctx, smbproperties.TransportProtocol, address)
	if err != nil {
		ow.SetError(fmt.Errorf("unable to dial the host: %w", err))
		return
	}
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {

			log.Println(fmt.Errorf("unable to close the connection: %w", err))
		}
	}(conn)
	d := &smb2.Dialer{
		Initiator: &smb2.NTLMInitiator{
			User:     smbproperties.Username,
			Password: smbproperties.Password,
			Domain:   smbproperties.Domain,
		},
	}
	c, err := d.Dial(conn)
	if err != nil {
		ow.SetError(fmt.Errorf("unable to dial the host: %w", err))
		return
	}
	defer func(c *smb2.Client) {
		err := c.Logoff()
		if err != nil {
			log.Println(fmt.Errorf("unable to logoff: %w", err))
		}
	}(c)

	fs, err := c.Mount(`\\` + smbproperties.Host + `\` + smbproperties.Share)
	if err != nil {
		ow.SetError(fmt.Errorf("unable to mount the share: %w", err))
		return
	}
	defer func(fs *smb2.RemoteFileSystem) {
		err := fs.Umount()
		if err != nil {
			log.Println(fmt.Errorf("unable to unmount file system: %w", err))
		}
	}(fs)
	if smbproperties.FileName != "" {
		var f *smb2.RemoteFile
		if strings.Contains(smbproperties.Operation, SMB_OPERATION_CREATE) {
			f, err = fs.Create(smbproperties.FileName)
			if err != nil {
				ow.SetError(fmt.Errorf("unable to create filename on a remote computer: %w", err))
				return
			}
			_, err = f.Write([]byte(smbproperties.Text))
			if err != nil {
				ow.SetError(fmt.Errorf("unable to write into file: %w", err))
				return
			}
		} else {
			f, err = fs.Open(smbproperties.FileName)
			if err != nil {
				ow.SetError(fmt.Errorf("unable to open the filename on a remote computer: %w", err))
				return
			}
			if smbproperties.Operation == SMB_OPERATION_CREATE {
				ow.SetError(nil)
				return
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
			ow.SetError(fmt.Errorf("unable to read the file: %w", err))
			return
		}
		bs, err := ioutil.ReadAll(f)
		if err != nil {
			ow.SetError(fmt.Errorf("unable to read the file: %w", err))
			return
		}
		if smbproperties.ExpectedOutput != "" && string(bs) != smbproperties.ExpectedOutput {
			ow.SetError(fmt.Errorf("%w. Output Received: %s", ErrDidNotMatchExpectedOutput, string(bs)))
			return
		}
	}

	ow.SetStatus(outcome.OUTCOME_STATUS_PASSED)
}
