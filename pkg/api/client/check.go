package client

import (
	"ScoreTrak/pkg/check"
)

type checkClient struct {
	s ScoretrakClient
}

func NewCheckClient(c ScoretrakClient) check.Repo {
	return &checkClient{c}
}

func (s checkClient) GetAllByRoundID(rID uint64) ([]*check.Check, error) {
	panic("implement me")
}

func (s checkClient) GetByRoundServiceID(rID uint64, sID uint64) ([]*check.Check, error) {
	panic("implement me")
}

func (s checkClient) Delete(id uint64) error {
	panic("implement me")
}

func (s checkClient) GetAll() ([]*check.Check, error) {
	panic("implement me")
}

func (s checkClient) GetByID(id uint64) (*check.Check, error) {
	panic("implement me")
}

func (s checkClient) Store(u *check.Check) error {
	panic("implement me")
}

func (s checkClient) StoreMany(u []*check.Check) error {
	panic("implement me")
}
