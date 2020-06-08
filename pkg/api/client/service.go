package client

import (
	"ScoreTrak/pkg/service"
	"fmt"
)

type serviceClient struct {
	s ScoretrakClient
}

func NewServiceClient(c ScoretrakClient) service.Serv {
	return &serviceClient{c}
}

func (s serviceClient) Delete(id uint64) error {
	return genericDelete(fmt.Sprintf("/service/%d", id), s.s)
}

func (s serviceClient) GetAll() ([]*service.Service, error) {
	var sg []*service.Service
	err := genericGet(&sg, "/service", s.s)
	if err != nil {
		return nil, err
	}
	return sg, nil
}

func (s serviceClient) GetByID(id uint64) (*service.Service, error) {
	sg := &service.Service{}
	err := genericGet(sg, fmt.Sprintf("/service/%d", id), s.s)
	if err != nil {
		return nil, err
	}
	return sg, nil
}

func (s serviceClient) Store(u *service.Service) error {
	return genericStore(u, fmt.Sprintf("/service"), s.s)
}

func (s serviceClient) Update(u *service.Service) error {
	return genericUpdate(u, fmt.Sprintf("/service/%d", u.ID), s.s)
}
