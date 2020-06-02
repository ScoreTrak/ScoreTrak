package client

import "ScoreTrak/pkg/service_group"

type serviceGroupClient struct {
	s ScoretrakClient
}

func NewServiceGroupClient(c ScoretrakClient) service_group.Serv {
	return &serviceGroupClient{c}
}

func (s serviceGroupClient) Delete(id uint64) error {
	panic("implement me")
}

func (s serviceGroupClient) GetAll() ([]*service_group.ServiceGroup, error) {
	panic("implement me")
}

func (s serviceGroupClient) GetByID(id uint64) (*service_group.ServiceGroup, error) {
	panic("implement me")
}

func (s serviceGroupClient) Store(u *service_group.ServiceGroup) error {
	panic("implement me")
}

func (s serviceGroupClient) Update(u *service_group.ServiceGroup) error {
	panic("implement me")
}
