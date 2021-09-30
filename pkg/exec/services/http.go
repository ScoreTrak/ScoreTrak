package services

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/ScoreTrak/ScoreTrak/pkg/exec"
)

type HTTP struct {
	ExpectedOutput string
	Scheme         string
	Port           string
	Path           string
	Subdomain      string
}

func NewHTTP() *HTTP {
	f := HTTP{Scheme: "http"}
	return &f
}

func (h *HTTP) Validate() error {
	return nil
}

func (h *HTTP) Execute(e exec.Exec) (passed bool, logOutput string, err error) {
	baseURL := ConstructURI(h.Port, h.Subdomain, e.Host, h.Path, h.Scheme)
	req, err := http.NewRequest("GET", baseURL.String(), nil)
	if err != nil {
		return false, "", fmt.Errorf("unable to craft the request: %w", err)
	}
	req = req.WithContext(e.Context)
	httpClient := http.DefaultClient
	resp, err := httpClient.Do(req)
	if err != nil {
		return false, "", fmt.Errorf("unable to craft the request: %w", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println(fmt.Errorf("unable to close body: %w", err))
		}
	}(resp.Body)

	if !(resp.StatusCode >= 200 && resp.StatusCode < 400) {
		return false, fmt.Sprintf("Invalid response code received: %d", resp.StatusCode), nil
	}
	if h.ExpectedOutput != "" {
		buf := new(bytes.Buffer)
		_, err = buf.ReadFrom(resp.Body)
		if err != nil {
			return false, "", fmt.Errorf("unable to read response body: %w", err)
		}
		newStr := buf.String()
		if !strings.Contains(newStr, h.ExpectedOutput) {
			return false, fmt.Sprintf("the page output doesn't contain expected output. Output received: %s", newStr), nil
		}
	}
	return true, Success, nil
}
