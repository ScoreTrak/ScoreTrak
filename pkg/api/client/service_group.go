package client

import (
	"ScoreTrak/pkg/service_group"
	"fmt"
)

type serviceGroupClient struct {
	s ScoretrakClient
}

func NewServiceGroupClient(c ScoretrakClient) service_group.Serv {
	return &serviceGroupClient{c}
}

func (s serviceGroupClient) Delete(id uint64) error {
	return genericDelete(fmt.Sprintf("/service_group/%d", id), s.s)
}

func (s serviceGroupClient) GetAll() ([]*service_group.ServiceGroup, error) {
	var sg []*service_group.ServiceGroup
	err := genericGet(&sg, "/service_group", s.s)
	if err != nil {
		return nil, err
	}
	return sg, nil
}

func (s serviceGroupClient) GetByID(id uint64) (*service_group.ServiceGroup, error) {
	sg := &service_group.ServiceGroup{}
	err := genericGet(sg, fmt.Sprintf("/service_group/%d", id), s.s)
	if err != nil {
		return nil, err
	}
	return sg, nil
}

func (s serviceGroupClient) Store(u *service_group.ServiceGroup) error {
	return genericStore(u, fmt.Sprintf("/service_group/%d", u.ID), s.s)
}

func (s serviceGroupClient) Update(u *service_group.ServiceGroup) error {
	return genericUpdate(u, fmt.Sprintf("/service_group/%d", u.ID), s.s)
}
