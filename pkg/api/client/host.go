package client

import (
	"fmt"
	"github.com/ScoreTrak/ScoreTrak/pkg/host"
	"github.com/gofrs/uuid"
)

type HostClient struct {
	s ScoretrakClient
}

func NewHostClient(c ScoretrakClient) host.Serv {
	return &HostClient{c}
}

func (s HostClient) Delete(id uuid.UUID) error {
	return s.s.GenericDelete(fmt.Sprintf("/host/%s", id.String()))
}

func (s HostClient) GetAll() ([]*host.Host, error) {
	var sg []*host.Host
	err := s.s.GenericGet(&sg, "/host")
	if err != nil {
		return nil, err
	}
	return sg, nil
}

func (s HostClient) GetByID(id uuid.UUID) (*host.Host, error) {
	sg := &host.Host{}
	err := s.s.GenericGet(sg, fmt.Sprintf("/host/%s", id.String()))
	if err != nil {
		return nil, err
	}
	return sg, nil
}

func (s HostClient) Store(u []*host.Host) error {
	return s.s.GenericStore(u, fmt.Sprintf("/host"))
}

func (s HostClient) Update(u *host.Host) error {
	return s.s.GenericUpdate(u, fmt.Sprintf("/host/%s", u.ID.String()))
}
