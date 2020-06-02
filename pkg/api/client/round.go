package client

import "ScoreTrak/pkg/round"

type roundClient struct {
	s ScoretrakClient
}

func NewRoundClient(c ScoretrakClient) round.Serv {
	return &roundClient{c}
}

func (r roundClient) Delete(id uint64) error {
	panic("implement me")
}

func (r roundClient) GetAll() ([]*round.Round, error) {
	panic("implement me")
}

func (r roundClient) GetByID(id uint64) (*round.Round, error) {
	panic("implement me")
}

func (r roundClient) Store(u *round.Round) error {
	panic("implement me")
}

func (r roundClient) Update(u *round.Round) error {
	panic("implement me")
}

func (r roundClient) GetLastRound() (*round.Round, error) {
	panic("implement me")
}
