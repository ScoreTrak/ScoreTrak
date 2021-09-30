package services

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/ScoreTrak/ScoreTrak/pkg/exec"
	"github.com/emersion/go-webdav"
)

type CalDav struct {
	Scheme         string
	Port           string
	Path           string
	Subdomain      string
	ExpectedOutput string
	Username       string
	Password       string
}

func NewCalDav() *CalDav {
	return &CalDav{Scheme: "http"}
}

func (h *CalDav) Validate() error {
	return nil
}

func (h *CalDav) Execute(e exec.Exec) (passed bool, logOutput string, err error) {
	authClient := webdav.HTTPClientWithBasicAuth(&http.Client{Timeout: time.Until(e.Deadline())}, h.Username, h.Password)
	baseURL := ConstructURI(h.Port, h.Subdomain, e.Host, h.Path, h.Scheme)
	client, err := webdav.NewClient(authClient, baseURL.String())
	if err != nil {
		return false, "", fmt.Errorf("unable to create client :%w", err)
	}
	usr, err := client.FindCurrentUserPrincipal()
	if err != nil {
		return false, "", fmt.Errorf("unable to retrieve current user principal: %w", err)
	}
	if h.ExpectedOutput != "" && !strings.Contains(usr, h.ExpectedOutput) {
		return false, "", fmt.Errorf("%w. Output Received: %s", ErrDidNotMatchExpectedOutput, usr)
	}
	return true, "", nil
}
