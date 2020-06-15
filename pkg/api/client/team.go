package client

import (
	"ScoreTrak/pkg/team"
	"fmt"
)

type teamClient struct {
	s ScoretrakClient
}

func NewTeamClient(c ScoretrakClient) team.Serv {
	return &teamClient{c}
}

func (t teamClient) DeleteByName(name string) error {
	return genericDelete(fmt.Sprintf("/team/%s", name), t.s)
}

func (t teamClient) GetAll() ([]*team.Team, error) {
	var tm []*team.Team
	err := genericGet(&tm, "/team", t.s)
	if err != nil {
		return nil, err
	}
	return tm, nil
}

func (t teamClient) GetByName(name string) (*team.Team, error) {
	tm := &team.Team{}
	err := genericGet(tm, fmt.Sprintf("/team/%s", name), t.s)
	if err != nil {
		return nil, err
	}
	return tm, nil
}

func (t teamClient) Store(u *team.Team) error {
	return genericStore(u, fmt.Sprintf("/team"), t.s)
}

func (t teamClient) Update(u *team.Team) error {
	return genericUpdate(u, fmt.Sprintf("/team/%s", u.ID), t.s)
}
