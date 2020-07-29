package client

import (
	"fmt"
	"github.com/L1ghtman2k/ScoreTrak/pkg/host_group"
)

type HostGroupClient struct {
	s ScoretrakClient
}

func NewHostGroupClient(c ScoretrakClient) host_group.Serv {
	return &HostGroupClient{c}
}

func (s HostGroupClient) Delete(id uint32) error {
	return s.s.GenericDelete(fmt.Sprintf("/host_group/%d", id))
}

func (s HostGroupClient) GetAll() ([]*host_group.HostGroup, error) {
	var sg []*host_group.HostGroup
	err := s.s.GenericGet(&sg, "/host_group")
	if err != nil {
		return nil, err
	}
	return sg, nil
}

func (s HostGroupClient) GetByID(id uint32) (*host_group.HostGroup, error) {
	sg := &host_group.HostGroup{}
	err := s.s.GenericGet(sg, fmt.Sprintf("/host_group/%d", id))
	if err != nil {
		return nil, err
	}
	return sg, nil
}

func (s HostGroupClient) Store(u *host_group.HostGroup) error {
	return s.s.GenericStore(u, fmt.Sprintf("/host_group"))
}

func (s HostGroupClient) Update(u *host_group.HostGroup) error {
	return s.s.GenericUpdate(u, fmt.Sprintf("/host_group/%d", u.ID))
}
