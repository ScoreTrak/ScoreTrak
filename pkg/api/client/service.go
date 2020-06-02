package client

import "ScoreTrak/pkg/service"

type serviceClient struct {
	s ScoretrakClient
}

func NewServiceClient(c ScoretrakClient) service.Serv {
	return &serviceClient{c}
}

func (s serviceClient) Delete(id uint64) error {
	panic("implement me")
}

func (s serviceClient) GetAll() ([]*service.Service, error) {
	panic("implement me")
}

func (s serviceClient) GetByID(id uint64) (*service.Service, error) {
	panic("implement me")
}

func (s serviceClient) Store(u *service.Service) error {
	panic("implement me")
}

func (s serviceClient) Update(u *service.Service) error {
	panic("implement me")
}
