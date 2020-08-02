package client

import (
	"fmt"
	"github.com/ScoreTrak/ScoreTrak/pkg/service_group"
	"github.com/gofrs/uuid"
)

type ServiceGroupClient struct {
	s ScoretrakClient
}

func NewServiceGroupClient(c ScoretrakClient) service_group.Serv {
	return &ServiceGroupClient{c}
}

func (s ServiceGroupClient) Delete(id uuid.UUID) error {
	return s.s.GenericDelete(fmt.Sprintf("/service_group/%s", id.String()))
}

func (s ServiceGroupClient) GetAll() ([]*service_group.ServiceGroup, error) {
	var sg []*service_group.ServiceGroup
	err := s.s.GenericGet(&sg, "/service_group")
	if err != nil {
		return nil, err
	}
	return sg, nil
}

func (s ServiceGroupClient) GetByID(id uuid.UUID) (*service_group.ServiceGroup, error) {
	sg := &service_group.ServiceGroup{}
	err := s.s.GenericGet(sg, fmt.Sprintf("/service_group/%s", id.String()))
	if err != nil {
		return nil, err
	}
	return sg, nil
}

func (s ServiceGroupClient) Store(u *service_group.ServiceGroup) error {
	return s.s.GenericStore(u, fmt.Sprintf("/service_group"))
}

func (s ServiceGroupClient) Update(u *service_group.ServiceGroup) error {
	return s.s.GenericUpdate(u, fmt.Sprintf("/service_group/%s", u.ID.String()))
}

func (s ServiceGroupClient) Redeploy(id uuid.UUID) error {
	err := s.s.GenericGet(nil, fmt.Sprintf("/service_group/%s/redeploy", id.String()))
	if err != nil {
		return err
	}
	return nil
}
