package client

import (
	"fmt"
	"github.com/L1ghtman2k/ScoreTrak/pkg/host"
)

type HostClient struct {
	s ScoretrakClient
}

func NewHostClient(c ScoretrakClient) host.Serv {
	return &HostClient{c}
}

func (s HostClient) Delete(id uint32) error {
	return s.s.GenericDelete(fmt.Sprintf("/host/%d", id))
}

func (s HostClient) GetAll() ([]*host.Host, error) {
	var sg []*host.Host
	err := s.s.GenericGet(&sg, "/host")
	if err != nil {
		return nil, err
	}
	return sg, nil
}

func (s HostClient) GetByID(id uint32) (*host.Host, error) {
	sg := &host.Host{}
	err := s.s.GenericGet(sg, fmt.Sprintf("/host/%d", id))
	if err != nil {
		return nil, err
	}
	return sg, nil
}

func (s HostClient) Store(u *host.Host) error {
	return s.s.GenericStore(u, fmt.Sprintf("/host"))
}

func (s HostClient) Update(u *host.Host) error {
	return s.s.GenericUpdate(u, fmt.Sprintf("/host/%d", u.ID))
}
