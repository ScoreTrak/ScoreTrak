package client

import (
	"fmt"
	"github.com/L1ghtman2k/ScoreTrak/pkg/queue/queueing"
	"github.com/L1ghtman2k/ScoreTrak/pkg/service"
	"github.com/gofrs/uuid"
)

type ServiceClient struct {
	s ScoretrakClient
}

func NewServiceClient(c ScoretrakClient) service.Serv {
	return &ServiceClient{c}
}

func (s ServiceClient) Delete(id uuid.UUID) error {
	return s.s.GenericDelete(fmt.Sprintf("/service/%s", id.String()))
}

func (s ServiceClient) GetAll() ([]*service.Service, error) {
	var sg []*service.Service
	err := s.s.GenericGet(&sg, "/service")
	if err != nil {
		return nil, err
	}
	return sg, nil
}

func (s ServiceClient) GetByID(id uuid.UUID) (*service.Service, error) {
	sg := &service.Service{}
	err := s.s.GenericGet(sg, fmt.Sprintf("/service/%s", id.String()))
	if err != nil {
		return nil, err
	}
	return sg, nil
}

func (s ServiceClient) Store(u []*service.Service) error {
	return s.s.GenericStore(u, fmt.Sprintf("/service"))
}

func (s ServiceClient) Update(u *service.Service) error {
	return s.s.GenericUpdate(u, fmt.Sprintf("/service/%s", u.ID.String()))
}

func (s ServiceClient) TestService(id uuid.UUID) (*queueing.ScoringData, error) {
	sd := &queueing.ScoringData{}
	err := s.s.GenericGet(sd, fmt.Sprintf("/service/test/%s", id.String()))
	if err != nil {
		return nil, err
	}
	return sd, nil
}
