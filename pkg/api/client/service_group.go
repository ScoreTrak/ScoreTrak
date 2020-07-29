package client

import (
	"fmt"
	"github.com/L1ghtman2k/ScoreTrak/pkg/service_group"
)

type ServiceGroupClient struct {
	s ScoretrakClient
}

func NewServiceGroupClient(c ScoretrakClient) service_group.Serv {
	return &ServiceGroupClient{c}
}

func (s ServiceGroupClient) Delete(id uint32) error {
	return s.s.GenericDelete(fmt.Sprintf("/service_group/%d", id))
}

func (s ServiceGroupClient) GetAll() ([]*service_group.ServiceGroup, error) {
	var sg []*service_group.ServiceGroup
	err := s.s.GenericGet(&sg, "/service_group")
	if err != nil {
		return nil, err
	}
	return sg, nil
}

func (s ServiceGroupClient) GetByID(id uint32) (*service_group.ServiceGroup, error) {
	sg := &service_group.ServiceGroup{}
	err := s.s.GenericGet(sg, fmt.Sprintf("/service_group/%d", id))
	if err != nil {
		return nil, err
	}
	return sg, nil
}

func (s ServiceGroupClient) Store(u *service_group.ServiceGroup) error {
	return s.s.GenericStore(u, fmt.Sprintf("/service_group"))
}

func (s ServiceGroupClient) Update(u *service_group.ServiceGroup) error {
	return s.s.GenericUpdate(u, fmt.Sprintf("/service_group/%d", u.ID))
}

func (s ServiceGroupClient) Redeploy(id uint32) error {
	err := s.s.GenericGet(nil, fmt.Sprintf("/service_group/%d/redeploy", id))
	if err != nil {
		return err
	}
	return nil
}
