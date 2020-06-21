package client

import (
	"ScoreTrak/pkg/host"
	"fmt"
)

type hostClient struct {
	s ScoretrakClient
}

func NewHostClient(c ScoretrakClient) host.Serv {
	return &hostClient{c}
}

func (s hostClient) Delete(id uint64) error {
	return s.s.genericDelete(fmt.Sprintf("/host/%d", id))
}

func (s hostClient) GetAll() ([]*host.Host, error) {
	var sg []*host.Host
	err := s.s.genericGet(&sg, "/host")
	if err != nil {
		return nil, err
	}
	return sg, nil
}

func (s hostClient) GetByID(id uint64) (*host.Host, error) {
	sg := &host.Host{}
	err := s.s.genericGet(sg, fmt.Sprintf("/host/%d", id))
	if err != nil {
		return nil, err
	}
	return sg, nil
}

func (s hostClient) Store(u *host.Host) error {
	return s.s.genericStore(u, fmt.Sprintf("/host"))
}

func (s hostClient) Update(u *host.Host) error {
	return s.s.genericUpdate(u, fmt.Sprintf("/host/%d", u.ID))
}
