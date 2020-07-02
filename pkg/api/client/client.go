package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type ScoretrakClient struct {
	baseURL    *url.URL
	token      string
	httpClient *http.Client
}

func NewScoretrakClient(url *url.URL, token string, client *http.Client) ScoretrakClient {
	return ScoretrakClient{baseURL: url, token: token, httpClient: client}
}

//Every Client could be fed directly to resource struct in gobuffalo

func responseValidator(r *http.Response) error {
	if r.StatusCode >= 200 && r.StatusCode < 400 {
		return nil
	}
	bodyBytes, err := ioutil.ReadAll(r.Body)
	var body string
	if err == nil {
		body = string(bodyBytes)
	}
	msg := fmt.Sprintf("Invalid response code received: %d", r.StatusCode)
	if body != "" {
		msg += fmt.Sprintf("\nResponse body: %s", body)
	}
	return &InvalidResponse{msg, r.StatusCode, body}

}

type InvalidResponse struct {
	msg          string // description of error
	ResponseCode int
	ResponseBody string
}

func (e *InvalidResponse) Error() string { return e.msg }

func (s ScoretrakClient) setAuthToken(req *http.Request) {
	req.Header.Set("x-access-token", s.token)
}

func (s ScoretrakClient) genericGet(obj interface{}, p string) error {
	req, err := s.prepareRequest(obj, p, "GET")
	if err != nil {
		return err
	}
	resp, err := s.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	err = responseValidator(resp)
	if err != nil {
		return err
	}
	if obj != nil {
		err = json.NewDecoder(resp.Body).Decode(obj)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s ScoretrakClient) genericUpdate(obj interface{}, p string) error {
	return s.genericPut(obj, p, "PATCH")
}

func (s ScoretrakClient) genericStore(obj interface{}, p string) error {
	return s.genericPut(obj, p, "POST")
}

func (s ScoretrakClient) genericPut(obj interface{}, p string, m string) error {
	req, err := s.prepareRequest(obj, p, m)
	if err != nil {
		return err
	}
	resp, err := s.httpClient.Do(req)
	if err != nil {
		return err
	}
	resp.Body.Close()
	return responseValidator(resp)
}

func (s ScoretrakClient) genericDelete(p string) error {
	req, err := s.prepareRequest(nil, p, "DELETE")
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

func (s ScoretrakClient) prepareRequest(obj interface{}, p string, m string) (*http.Request, error) {
	rel := &url.URL{Path: p}
	u := s.baseURL.ResolveReference(rel)
	e, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	b := bytes.NewBuffer(e)
	req, err := http.NewRequest(m, u.String(), b)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}
