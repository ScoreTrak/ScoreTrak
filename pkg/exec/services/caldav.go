package services

import (
	"fmt"
	"github.com/ScoreTrak/ScoreTrak/pkg/exec"
	"github.com/emersion/go-webdav"
	"net/http"
	"strings"
	"time"
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

func (h *CalDav) Execute(e exec.Exec) (passed bool, log string, err error) {
	authClient := webdav.HTTPClientWithBasicAuth(&http.Client{Timeout: time.Until(e.Deadline())}, h.Username, h.Password)
	baseURL := exec.ConstructURI(h.Port, h.Subdomain, e.Host, h.Path, h.Scheme)
	client, err := webdav.NewClient(authClient, baseURL.String())
	if err != nil {
		return false, "Unable to create client", err
	}
	usr, err := client.FindCurrentUserPrincipal()
	if err != nil {
		return false, "Unable to retrieve current user principal", err
	}
	if h.ExpectedOutput != "" && !strings.Contains(usr, h.ExpectedOutput) {
		return false, fmt.Sprintf("User Principal did not match expected output. Output received: %s", usr), nil
	}
	return true, "", nil
}
