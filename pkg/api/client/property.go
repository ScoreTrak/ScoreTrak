package client

import (
	"fmt"
	"github.com/L1ghtman2k/ScoreTrak/pkg/property"
	"github.com/gofrs/uuid"
)

type PropertyClient struct {
	s ScoretrakClient
}

func NewPropertyClient(c ScoretrakClient) property.Serv {
	return &PropertyClient{c}
}

func (s PropertyClient) Delete(id uuid.UUID) error {
	return s.s.GenericDelete(fmt.Sprintf("/property/%s", id.String()))
}

func (s PropertyClient) GetAll() ([]*property.Property, error) {
	var sg []*property.Property
	err := s.s.GenericGet(&sg, "/property")
	if err != nil {
		return nil, err
	}
	return sg, nil
}

func (s PropertyClient) GetByID(id uuid.UUID) (*property.Property, error) {
	sg := &property.Property{}
	err := s.s.GenericGet(sg, fmt.Sprintf("/property/%s", id.String()))
	if err != nil {
		return nil, err
	}
	return sg, nil
}

func (s PropertyClient) Store(u []*property.Property) error {
	return s.s.GenericStore(u, fmt.Sprintf("/property"))
}

func (s PropertyClient) Update(u *property.Property) error {
	return s.s.GenericUpdate(u, fmt.Sprintf("/property/%s", u.ID.String()))
}
