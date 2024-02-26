package executors

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/scoretrak/scoretrak/pkg/scorer/outcome"
	"github.com/scoretrak/scoretrak/pkg/scorer/outcome_writer"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type HTTPProperties struct {
	Scheme               string `json:"schema" validate:"required,oneof=http https"`
	Host                 string `json:"host" validate:"required,hostname"`
	Port                 string `json:"port" validate:"required,number"`
	Path                 string `json:"path" validate:"required"`
	Subdomain            string `json:"subdomain" validate:"required"`
	ExpectedOutput       string `json:"expected_output" validate:""`
	ExpectedResponseCode string `json:"expected_response_code" validate:""`
}

var ErrInvalidResponseCodeReceived = errors.New("invalid response code received")

func ScoreHttp(ctx context.Context, ow *outcome_writer.OutcomeWriter, properties []byte) {
	httpproperties := &HTTPProperties{}
	err := json.Unmarshal(properties, &httpproperties)
	if err != nil {
		ow.SetError(fmt.Errorf("unable to unmarshall properties: %v", err))
		return
	}

	err = validate.Struct(httpproperties)
	if err != nil {
		// validationErrors := err.(validator.ValidationErrors)
		ow.SetError(fmt.Errorf("validation error: %w", err))
		return
	}

	baseURL := ConstructURI(httpproperties.Port, httpproperties.Subdomain, httpproperties.Host, httpproperties.Path, httpproperties.Scheme)
	req, err := http.NewRequest("GET", baseURL.String(), nil)
	if err != nil {
		ow.SetError(fmt.Errorf("unable to craft the request: %w", err))
		return
	}
	req = req.WithContext(ctx)
	httpClient := http.DefaultClient
	resp, err := httpClient.Do(req) //nolint:bodyclose
	if err != nil {
		ow.SetError(fmt.Errorf("unable to craft the request: %w", err))
		return
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println(fmt.Errorf("unable to close body: %w", err))
		}
	}(resp.Body)

	if !(resp.StatusCode >= 200 && resp.StatusCode < 400) {
		ow.SetError(fmt.Errorf("%w: %d", ErrInvalidResponseCodeReceived, resp.StatusCode))
		return
	}

	if httpproperties.ExpectedResponseCode != "" {
		expectedRespCode, err := strconv.Atoi(httpproperties.ExpectedResponseCode)
		if err != nil {
			ow.SetError(fmt.Errorf("unable to parse expected response code: %v", err))
			return
		}
		if resp.StatusCode != expectedRespCode {
			ow.SetError(fmt.Errorf("%w: %d", ErrInvalidResponseCodeReceived, resp.StatusCode))
			return
		}

	}

	if httpproperties.ExpectedOutput != "" {
		buf := new(bytes.Buffer)
		_, err = buf.ReadFrom(resp.Body)
		if err != nil {
			ow.SetError(fmt.Errorf("unable to read response body: %w", err))
			return
		}
		newStr := buf.String()
		if !strings.Contains(newStr, httpproperties.ExpectedOutput) {
			ow.SetError(fmt.Errorf("%w. Output Received: %s", ErrDidNotMatchExpectedOutput, newStr))
			return
		}
	}

	ow.SetStatus(outcome.OUTCOME_STATUS_PASSED)
}
