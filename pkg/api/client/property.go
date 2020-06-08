package client

import (
	"ScoreTrak/pkg/property"
	"fmt"
)

type propertyClient struct {
	s ScoretrakClient
}

func NewPropertyClient(c ScoretrakClient) property.Serv {
	return &propertyClient{c}
}

func (s propertyClient) Delete(id uint64) error {
	return genericDelete(fmt.Sprintf("/property/%d", id), s.s)
}

func (s propertyClient) GetAll() ([]*property.Property, error) {
	var sg []*property.Property
	err := genericGet(&sg, "/property", s.s)
	if err != nil {
		return nil, err
	}
	return sg, nil
}

func (s propertyClient) GetByID(id uint64) (*property.Property, error) {
	sg := &property.Property{}
	err := genericGet(sg, fmt.Sprintf("/property/%d", id), s.s)
	if err != nil {
		return nil, err
	}
	return sg, nil
}

func (s propertyClient) Store(u *property.Property) error {
	return genericStore(u, fmt.Sprintf("/property"), s.s)
}

func (s propertyClient) Update(u *property.Property) error {
	return genericUpdate(u, fmt.Sprintf("/property/%d", u.ID), s.s)
}
