package client

import (
	"fmt"
	"net/http"
	"net/url"
)

type ScoretrakClient struct {
	BaseURL    *url.URL
	token      string
	httpClient *http.Client
}

func NewScoretrakClient(url *url.URL, token string, client *http.Client) ScoretrakClient {
	return ScoretrakClient{BaseURL: url, token: token, httpClient: client}
}

//Every Client could be fed directly to resource struct in gobuffalo

func responseValidator(r *http.Response) error {
	if r.StatusCode >= 200 && r.StatusCode < 400 {
		return nil
	}
	return &InvalidResponse{fmt.Sprintf("Invalid response code received: %d", r.StatusCode)}

}

type InvalidResponse struct {
	msg string // description of error
}

func (e *InvalidResponse) Error() string { return e.msg }
