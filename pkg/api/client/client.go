package client

import (
	"bytes"
	"encoding/json"
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

func genericGet(obj interface{}, p string, s ScoretrakClient) error {
	rel := &url.URL{Path: p}
	u := s.BaseURL.ResolveReference(rel)
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return err
	}
	req.Header.Set("Accept", "application/json")
	resp, err := s.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	err = responseValidator(resp)
	if err != nil {
		return err
	}
	err = json.NewDecoder(resp.Body).Decode(obj)
	if err != nil {
		return err
	}
	return nil
}

func genericUpdate(obj interface{}, p string, s ScoretrakClient) error {
	return genericPut(obj, p, s, "PATCH")
}

func genericStore(obj interface{}, p string, s ScoretrakClient) error {
	return genericPut(obj, p, s, "POST")
}

func genericPut(obj interface{}, p string, s ScoretrakClient, m string) error {
	rel := &url.URL{Path: p}
	u := s.BaseURL.ResolveReference(rel)
	e, err := json.Marshal(obj)
	if err != nil {
		return nil
	}
	b := bytes.NewBuffer(e)
	req, err := http.NewRequest(m, u.String(), b)
	if err != nil {
		return err
	}
	req.Header.Set("Accept", "application/json")
	resp, err := s.httpClient.Do(req)
	if err != nil {
		return err
	}
	resp.Body.Close()
	return responseValidator(resp)
}

func genericDelete(p string, s ScoretrakClient) error {
	rel := &url.URL{Path: p}
	u := s.BaseURL.ResolveReference(rel)
	req, err := http.NewRequest("DELETE", u.String(), nil)
	if err != nil {
		return err
	}
	req.Header.Set("Accept", "application/json")
	resp, err := s.httpClient.Do(req)
	if err != nil {
		return err
	}
	resp.Body.Close()
	return nil
}
