package client

import (
	"fmt"
	"github.com/L1ghtman2k/ScoreTrak/pkg/team"
)

type TeamClient struct {
	s ScoretrakClient
}

func NewTeamClient(c ScoretrakClient) team.Serv {
	return &TeamClient{c}
}

func (t TeamClient) Delete(id uint32) error {
	return t.s.GenericDelete(fmt.Sprintf("/team/%d", id))
}

func (t TeamClient) GetAll() ([]*team.Team, error) {
	var tm []*team.Team
	err := t.s.GenericGet(&tm, "/team")
	if err != nil {
		return nil, err
	}
	return tm, nil
}

func (t TeamClient) GetByID(id uint32) (*team.Team, error) {
	tm := &team.Team{}
	err := t.s.GenericGet(tm, fmt.Sprintf("/team/%d", id))
	if err != nil {
		return nil, err
	}
	return tm, nil
}

func (t TeamClient) Store(u *team.Team) error {
	return t.s.GenericStore(u, fmt.Sprintf("/team"))
}

func (t TeamClient) Update(u *team.Team) error {
	return t.s.GenericUpdate(u, fmt.Sprintf("/team/%d", u.ID))
}
