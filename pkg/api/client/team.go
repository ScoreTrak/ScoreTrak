package client

import (
	"fmt"
	"github.com/ScoreTrak/ScoreTrak/pkg/team"
	"github.com/gofrs/uuid"
)

type TeamClient struct {
	s ScoretrakClient
}

func NewTeamClient(c ScoretrakClient) team.Serv {
	return &TeamClient{c}
}

func (t TeamClient) Delete(id uuid.UUID) error {
	return t.s.GenericDelete(fmt.Sprintf("/team/%s", id.String()))
}

func (t TeamClient) GetAll() ([]*team.Team, error) {
	var tm []*team.Team
	err := t.s.GenericGet(&tm, "/team")
	if err != nil {
		return nil, err
	}
	return tm, nil
}

func (t TeamClient) GetByID(id uuid.UUID) (*team.Team, error) {
	tm := &team.Team{}
	err := t.s.GenericGet(tm, fmt.Sprintf("/team/%s", id.String()))
	if err != nil {
		return nil, err
	}
	return tm, nil
}

func (t TeamClient) Store(u []*team.Team) error {
	return t.s.GenericStore(u, fmt.Sprintf("/team"))
}

func (t TeamClient) Update(u *team.Team) error {
	return t.s.GenericUpdate(u, fmt.Sprintf("/team/%s", u.ID.String()))
}
