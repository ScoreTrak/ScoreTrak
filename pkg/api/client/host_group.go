package client

import (
	"ScoreTrak/pkg/host_group"
	"fmt"
)

type hostGroupClient struct {
	s ScoretrakClient
}

func NewHostGroupClient(c ScoretrakClient) host_group.Serv {
	return &hostGroupClient{c}
}

func (s hostGroupClient) Delete(id uint64) error {
	return genericDelete(fmt.Sprintf("/host_group/%d", id), s.s)
}

func (s hostGroupClient) GetAll() ([]*host_group.HostGroup, error) {
	var sg []*host_group.HostGroup
	err := genericGet(&sg, "/host_group", s.s)
	if err != nil {
		return nil, err
	}
	return sg, nil
}

func (s hostGroupClient) GetByID(id uint64) (*host_group.HostGroup, error) {
	sg := &host_group.HostGroup{}
	err := genericGet(sg, fmt.Sprintf("/host_group/%d", id), s.s)
	if err != nil {
		return nil, err
	}
	return sg, nil
}

func (s hostGroupClient) Store(u *host_group.HostGroup) error {
	return genericStore(u, fmt.Sprintf("/host_group"), s.s)
}

func (s hostGroupClient) Update(u *host_group.HostGroup) error {
	return genericUpdate(u, fmt.Sprintf("/host_group/%d", u.ID), s.s)
}
