package services

import (
	"bytes"
	"fmt"
	"github.com/ScoreTrak/ScoreTrak/pkg/exec"
	"net/http"
	"net/url"
	"strings"
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

func (h *HTTP) Execute(e exec.Exec) (passed bool, log string, err error) {
	if h.Port == "" {
		if strings.ToLower(h.Scheme) == "https" {
			h.Port = "443"
		} else {
			h.Port = "80"
		}
	}
	if h.Subdomain != "" && h.Subdomain[len(h.Subdomain)-1:] != "." {
		h.Subdomain += "."
	}
	host := h.Subdomain + e.Host + ":" + h.Port
	baseURL := url.URL{Path: h.Path, Scheme: h.Scheme, Host: host}
	req, err := http.NewRequest("GET", baseURL.String(), nil)
	if err != nil {
		return false, "Error while crafting the request", err
	}
	req = req.WithContext(e.Context)
	httpClient := http.DefaultClient
	resp, err := httpClient.Do(req)
	if err != nil {
		return false, "Error while making the request", err
	}
	defer resp.Body.Close()

	if !(resp.StatusCode >= 200 && resp.StatusCode < 400) {
		return false, fmt.Sprintf("Invalid response code received: %d", resp.StatusCode), nil
	}
	if h.ExpectedOutput != "" {
		buf := new(bytes.Buffer)
		_, err = buf.ReadFrom(resp.Body)
		if err != nil {
			return false, "Unable to read response body", err
		}
		newStr := buf.String()
		if !strings.Contains(newStr, h.ExpectedOutput) {
			return false, "the page output doesn't contain expected output", nil //TODO: Make a more meaningful output
		}
	}
	return true, "Success!", nil
}
