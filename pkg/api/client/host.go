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
	return genericDelete(fmt.Sprintf("/host/%d", id), s.s)
}

func (s hostClient) GetAll() ([]*host.Host, error) {
	var sg []*host.Host
	err := genericGet(&sg, "/host", s.s)
	if err != nil {
		return nil, err
	}
	return sg, nil
}

func (s hostClient) GetByID(id uint64) (*host.Host, error) {
	sg := &host.Host{}
	err := genericGet(sg, fmt.Sprintf("/host/%d", id), s.s)
	if err != nil {
		return nil, err
	}
	return sg, nil
}

func (s hostClient) Store(u *host.Host) error {
	return genericStore(u, fmt.Sprintf("/host/%d", u.ID), s.s)
}

func (s hostClient) Update(u *host.Host) error {
	return genericUpdate(u, fmt.Sprintf("/host/%d", u.ID), s.s)
}
