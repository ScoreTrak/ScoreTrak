package services

import (
	"ScoreTrak/pkg/exec"
	"bytes"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type HTTP struct {
	ExpectedOutput string `json:"text"`
	Scheme         string `json:"scheme"`
	Port           string `json:"port"`
	Path           string `json:"path"`
	Subdomain      string `json:"subdomain"`
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
	httpClient := http.DefaultClient
	httpClient.Timeout = time.Until(e.Timeout)
	resp, err := httpClient.Do(req)
	if err != nil {
		return false, "Error while making the request", err
	}
	defer resp.Body.Close()

	if !(resp.StatusCode >= 200 && resp.StatusCode < 400) {
		return false, fmt.Sprintf(fmt.Sprintf("Invalid response code received: %d", resp.StatusCode)), nil
	}
	if h.ExpectedOutput != "" {
		buf := new(bytes.Buffer)
		_, err = buf.ReadFrom(resp.Body)
		if err != nil {
			return false, "Unable to read response body", err
		}
		newStr := buf.String()
		if !strings.Contains(newStr, h.ExpectedOutput) {
			return false, "the page output doesn't contain expected output", nil
		}
	}
	return true, "Success!", nil
}
