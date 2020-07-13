package client

import (
	"fmt"
	"github.com/L1ghtman2k/ScoreTrak/pkg/host_group"
)

type hostGroupClient struct {
	s ScoretrakClient
}

func NewHostGroupClient(c ScoretrakClient) host_group.Serv {
	return &hostGroupClient{c}
}

func (s hostGroupClient) Delete(id uint64) error {
	return s.s.genericDelete(fmt.Sprintf("/host_group/%d", id))
}

func (s hostGroupClient) GetAll() ([]*host_group.HostGroup, error) {
	var sg []*host_group.HostGroup
	err := s.s.genericGet(&sg, "/host_group")
	if err != nil {
		return nil, err
	}
	return sg, nil
}

func (s hostGroupClient) GetByID(id uint64) (*host_group.HostGroup, error) {
	sg := &host_group.HostGroup{}
	err := s.s.genericGet(sg, fmt.Sprintf("/host_group/%d", id))
	if err != nil {
		return nil, err
	}
	return sg, nil
}

func (s hostGroupClient) Store(u *host_group.HostGroup) error {
	return s.s.genericStore(u, fmt.Sprintf("/host_group"))
}

func (s hostGroupClient) Update(u *host_group.HostGroup) error {
	return s.s.genericUpdate(u, fmt.Sprintf("/host_group/%d", u.ID))
}
