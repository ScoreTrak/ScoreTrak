package client

import (
	"ScoreTrak/pkg/team"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type teamClient struct {
	s ScoretrakClient
}

func NewTeamClient(c ScoretrakClient) team.Serv {
	return &teamClient{c}
}

func (t teamClient) Delete(id string) error {
	rel := &url.URL{Path: fmt.Sprintf("/team/%s", id)}
	u := t.s.BaseURL.ResolveReference(rel)
	req, err := http.NewRequest("DELETE", u.String(), nil)
	if err != nil {
		return err
	}
	req.Header.Set("Accept", "application/json")
	resp, err := t.s.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return err
}

func (t teamClient) GetAll() ([]*team.Team, error) {
	rel := &url.URL{Path: "/team"}
	u := t.s.BaseURL.ResolveReference(rel)
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	resp, err := t.s.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	err = responseValidator(resp)
	if err != nil {
		return nil, err
	}
	var teams []*team.Team
	err = json.NewDecoder(resp.Body).Decode(&teams)
	return teams, err
}

func (t teamClient) GetByID(id string) (*team.Team, error) {
	rel := &url.URL{Path: fmt.Sprintf("/team/%s", id)}
	u := t.s.BaseURL.ResolveReference(rel)
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	resp, err := t.s.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var tm *team.Team
	err = responseValidator(resp)
	if err != nil {
		return nil, err
	}
	err = json.NewDecoder(resp.Body).Decode(&tm)
	if err != nil {
		return nil, err
	}
	return tm, nil
}

func (t teamClient) Store(ut *team.Team) error {
	rel := &url.URL{Path: fmt.Sprintf("/team/%s", ut.ID)}
	u := t.s.BaseURL.ResolveReference(rel)
	e, err := json.Marshal(ut)
	if err != nil {
		return nil
	}
	b := bytes.NewBuffer(e)
	req, err := http.NewRequest("POST", u.String(), b)
	if err != nil {
		return err
	}
	req.Header.Set("Accept", "application/json")
	resp, err := t.s.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return responseValidator(resp)
}

func (t teamClient) Update(ut *team.Team) error {
	rel := &url.URL{Path: fmt.Sprintf("/team/%s", ut.ID)}
	u := t.s.BaseURL.ResolveReference(rel)
	e, err := json.Marshal(ut)
	if err != nil {
		return nil
	}
	b := bytes.NewBuffer(e)
	req, err := http.NewRequest("PATCH", u.String(), b)
	if err != nil {
		return err
	}
	req.Header.Set("Accept", "application/json")
	resp, err := t.s.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return responseValidator(resp)
}
