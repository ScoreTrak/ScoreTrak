package client

import "ScoreTrak/pkg/team"

type teamClient struct {
	s ScoretrakClient
}

func NewTeamClient(c ScoretrakClient) team.Serv {
	return &teamClient{c}
}

func (t teamClient) Delete(id string) error {
	panic("implement me")
}

func (t teamClient) GetAll() ([]*team.Team, error) {
	panic("implement me")
}

func (t teamClient) GetByID(id string) (*team.Team, error) {
	panic("implement me")
}

func (t teamClient) Store(u *team.Team) error {
	panic("implement me")
}

func (t teamClient) Update(u *team.Team) error {
	panic("implement me")
}
