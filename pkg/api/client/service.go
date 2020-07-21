package client

import (
	"fmt"
	"github.com/L1ghtman2k/ScoreTrak/pkg/queue/queueing"
	"github.com/L1ghtman2k/ScoreTrak/pkg/service"
)

type ServiceClient struct {
	s ScoretrakClient
}

func NewServiceClient(c ScoretrakClient) service.Serv {
	return &ServiceClient{c}
}

func (s ServiceClient) Delete(id uint64) error {
	return s.s.genericDelete(fmt.Sprintf("/service/%d", id))
}

func (s ServiceClient) GetAll() ([]*service.Service, error) {
	var sg []*service.Service
	err := s.s.genericGet(&sg, "/service")
	if err != nil {
		return nil, err
	}
	return sg, nil
}

func (s ServiceClient) GetByID(id uint64) (*service.Service, error) {
	sg := &service.Service{}
	err := s.s.genericGet(sg, fmt.Sprintf("/service/%d", id))
	if err != nil {
		return nil, err
	}
	return sg, nil
}

func (s ServiceClient) Store(u *service.Service) error {
	return s.s.genericStore(u, fmt.Sprintf("/service"))
}

func (s ServiceClient) Update(u *service.Service) error {
	return s.s.genericUpdate(u, fmt.Sprintf("/service/%d", u.ID))
}

func (s ServiceClient) TestService(id uint64) (*queueing.ScoringData, error) {
	sd := &queueing.ScoringData{}
	err := s.s.genericGet(sd, fmt.Sprintf("/service/test/%d", id))
	if err != nil {
		return nil, err
	}
	return sd, nil
}
