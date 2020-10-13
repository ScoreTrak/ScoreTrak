package client

import (
	"fmt"
	"github.com/ScoreTrak/ScoreTrak/pkg/host_group"
	"github.com/gofrs/uuid"
)

type HostGroupClient struct {
	s ScoretrakClient
}

func NewHostGroupClient(c ScoretrakClient) *HostGroupClient {
	return &HostGroupClient{c}
}

func (s HostGroupClient) Delete(id uuid.UUID) error {
	return s.s.GenericDelete(fmt.Sprintf("/host_group/%s", id.String()))
}

func (s HostGroupClient) GetAll() ([]*host_group.HostGroup, error) {
	var sg []*host_group.HostGroup
	err := s.s.GenericGet(&sg, "/host_group")
	if err != nil {
		return nil, err
	}
	return sg, nil
}

func (s HostGroupClient) GetByID(id uuid.UUID) (*host_group.HostGroup, error) {
	sg := &host_group.HostGroup{}
	err := s.s.GenericGet(sg, fmt.Sprintf("/host_group/%s", id.String()))
	if err != nil {
		return nil, err
	}
	return sg, nil
}

func (s HostGroupClient) Store(u []*host_group.HostGroup) error {
	return s.s.GenericStore(u, "/host_group")
}

func (s HostGroupClient) Update(u *host_group.HostGroup) error {
	return s.s.GenericUpdate(u, fmt.Sprintf("/host_group/%s", u.ID.String()))
}
