package client

import (
	"fmt"
	"github.com/L1ghtman2k/ScoreTrak/pkg/service_group"
)

type serviceGroupClient struct {
	s ScoretrakClient
}

func NewServiceGroupClient(c ScoretrakClient) service_group.Serv {
	return &serviceGroupClient{c}
}

func (s serviceGroupClient) Delete(id uint64) error {
	return s.s.genericDelete(fmt.Sprintf("/service_group/%d", id))
}

func (s serviceGroupClient) GetAll() ([]*service_group.ServiceGroup, error) {
	var sg []*service_group.ServiceGroup
	err := s.s.genericGet(&sg, "/service_group")
	if err != nil {
		return nil, err
	}
	return sg, nil
}

func (s serviceGroupClient) GetByID(id uint64) (*service_group.ServiceGroup, error) {
	sg := &service_group.ServiceGroup{}
	err := s.s.genericGet(sg, fmt.Sprintf("/service_group/%d", id))
	if err != nil {
		return nil, err
	}
	return sg, nil
}

func (s serviceGroupClient) Store(u *service_group.ServiceGroup) error {
	return s.s.genericStore(u, fmt.Sprintf("/service_group"))
}

func (s serviceGroupClient) Update(u *service_group.ServiceGroup) error {
	return s.s.genericUpdate(u, fmt.Sprintf("/service_group/%d", u.ID))
}
