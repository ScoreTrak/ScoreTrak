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
	Protocol       string `json:"protocol"`
	Port           string `json:"port"`
	Path           string `json:"path"`
	Subdomain      string `json:"subdomain"`
}

func NewHTTP() *HTTP {
	f := HTTP{Port: "80", Protocol: "http"}
	return &f
}

func (h *HTTP) Validate() error {
	return nil
}

func (h *HTTP) Execute(e exec.Exec) (passed bool, log string, err error) {
	if h.Subdomain != "" && h.Subdomain[len(h.Subdomain)-1:] != "." {
		h.Subdomain += "."
	}
	host := h.Subdomain + e.Host + ":" + h.Port
	baseURL := url.URL{Path: h.Path, Scheme: h.Protocol, Host: host}
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
		buf.ReadFrom(resp.Body)
		newStr := buf.String()
		if !strings.Contains(newStr, h.ExpectedOutput) {
			return false, "the page output doesn't contain expected output", nil
		}
	}
	return true, "Success!", nil
}
