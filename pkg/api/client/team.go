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

func (t teamClient) Delete(id string) error {
	return genericDelete(fmt.Sprintf("/team/%s", id), t.s)
}

func (t teamClient) GetAll() ([]*team.Team, error) {
	var tm []*team.Team
	err := genericGet(&tm, "/team", t.s)
	if err != nil {
		return nil, err
	}
	return tm, nil
}

func (t teamClient) GetByID(id string) (*team.Team, error) {
	tm := &team.Team{}
	err := genericGet(tm, fmt.Sprintf("/team/%s", id), t.s)
	if err != nil {
		return nil, err
	}
	return tm, nil
}

func (t teamClient) Store(ut *team.Team) error {
	return genericStore(ut, fmt.Sprintf("/team/%s", ut.ID), t.s)
}

func (t teamClient) Update(ut *team.Team) error {
	return genericUpdate(ut, fmt.Sprintf("/team/%s", ut.ID), t.s)
}
