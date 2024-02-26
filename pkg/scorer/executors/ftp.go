package executors

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jlaffaye/ftp"
	"github.com/scoretrak/scoretrak/pkg/scorer/outcome"
	"github.com/scoretrak/scoretrak/pkg/scorer/outcome_writer"
	"io/ioutil"
	"log"
	"strings"
)

type FTPProperties struct {
	Username       string
	Password       string
	Host           string
	Port           string `default:"21" validate:"number"`
	Text           string
	WriteFilename  string
	ReadFilename   string
	ExpectedOutput string `json:"expected_output" validate:"required"`
}

const TestFileName = "test_file.scoretrak.txt"

var ErrFTPNeedsUsernameAndPassword = errors.New("FTP check_service needs username and password")
var ErrFTPNeedsEitherTextOrReadFile = errors.New("FTP check_service needs either text, or read_file parameter")

func ScoreFTP(ctx context.Context, ow *outcome_writer.OutcomeWriter, properties []byte) {
	ftpproperties := &FTPProperties{}
	err := json.Unmarshal(properties, &ftpproperties)
	if err != nil {
		ow.SetError(fmt.Errorf("unable to unmarshall properties: %v", err))
		return
	}

	err = validate.Struct(ftpproperties)
	if err != nil {
		// validationErrors := err.(validator.ValidationErrors)
		ow.SetError(fmt.Errorf("validation error: %w", err))
		return
	}

	address := fmt.Sprintf("%s:%s", ftpproperties.Host, ftpproperties.Port)
	ftpConn, err := ftp.Dial(address)
	if err != nil {
		ow.SetError(fmt.Errorf("unable to dial FTP Server: %w", err))
		return
	}

	defer func(c *ftp.ServerConn) {
		err := c.Quit()
		if err != nil {
			log.Println(fmt.Errorf("unable to close ftp connection: %w", err))
		}
	}(ftpConn)

	// Login if credentials are provided
	if ftpproperties.Username != "" && ftpproperties.Password != "" {
		err = ftpConn.Login(ftpproperties.Username, ftpproperties.Password)
		if err != nil {
			ow.SetError(fmt.Errorf("unable to Login: %w", err))
			return
		}

		defer func(c *ftp.ServerConn) {
			err := c.Logout()
			if err != nil {
				log.Println(fmt.Errorf("unable to logout from FTP: %w", err))
			}
		}(ftpConn)
	}

	// Test writing to a file
	if ftpproperties.Text != "" {
		data := bytes.NewBufferString(ftpproperties.Text)
		if ftpproperties.WriteFilename == "" {
			ftpproperties.WriteFilename = TestFileName
		}

		err = ftpConn.Stor(ftpproperties.WriteFilename, data)
		if err != nil {
			ow.SetError(fmt.Errorf("unable to store file to FTP server: %w", err))
			return
		}
	}

	// Test reading to a file
	if ftpproperties.ReadFilename != "" {
		r, err := ftpConn.Retr(ftpproperties.ReadFilename)
		if err != nil {
			ow.SetError(fmt.Errorf("failed to retrieve the file from FTP: %w", err))
			return
		}

		defer func(r *ftp.Response) {
			err := r.Close()
			if err != nil {
				log.Println(fmt.Errorf("unable to close file: %w", err))
			}
		}(r)

		buf, err := ioutil.ReadAll(r)
		if err != nil {
			ow.SetError(fmt.Errorf("failed to read file contents, it might be corrupted: %w", err))
			return
		}
		if ftpproperties.ExpectedOutput != "" && !strings.Contains(string(buf), ftpproperties.ExpectedOutput) {
			ow.SetError(fmt.Errorf("%w. Output Received: %s", ErrDidNotMatchExpectedOutput, string(buf)))
			return
		}
	}

	ow.SetStatus(outcome.OUTCOME_STATUS_PASSED)
}
