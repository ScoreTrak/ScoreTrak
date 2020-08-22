package client

import (
	"fmt"
	"github.com/ScoreTrak/ScoreTrak/pkg/property"
	"github.com/gofrs/uuid"
)

type PropertyClient struct {
	s ScoretrakClient
}

func NewPropertyClient(c ScoretrakClient) *PropertyClient {
	return &PropertyClient{c}
}

func (s PropertyClient) Delete(id uuid.UUID, key string) error {
	return s.s.GenericDelete(fmt.Sprintf("/property/%s/%s", id.String(), key))
}

func (s PropertyClient) GetAll() ([]*property.Property, error) {
	var sg []*property.Property
	err := s.s.GenericGet(&sg, "/property")
	if err != nil {
		return nil, err
	}
	return sg, nil
}

func (s PropertyClient) GetByServiceIDKey(id uuid.UUID, key string) (*property.Property, error) {
	sg := &property.Property{}
	err := s.s.GenericGet(sg, fmt.Sprintf("/property/%s/%s", id.String(), key))
	if err != nil {
		return nil, err
	}
	return sg, nil
}

func (s PropertyClient) GetAllByServiceID(id uuid.UUID) ([]*property.Property, error) {
	var sg []*property.Property
	err := s.s.GenericGet(&sg, fmt.Sprintf("/property/%s", id.String()))
	if err != nil {
		return nil, err
	}
	return sg, nil
}

func (s PropertyClient) Store(u []*property.Property) error {
	return s.s.GenericStore(u, fmt.Sprintf("/property"))
}

func (s PropertyClient) Update(u *property.Property) error {
	return s.s.GenericUpdate(u, fmt.Sprintf("/property/%s/%s", u.ServiceID.String(), u.Key))
}
