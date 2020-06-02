package client

import "ScoreTrak/pkg/property"

type propertyClient struct {
	s ScoretrakClient
}

func NewPropertyClient(c ScoretrakClient) property.Serv {
	return &propertyClient{c}
}

func (p propertyClient) Delete(id uint64) error {
	panic("implement me")
}

func (p propertyClient) GetAll() ([]*property.Property, error) {
	panic("implement me")
}

func (p propertyClient) GetByID(id uint64) (*property.Property, error) {
	panic("implement me")
}

func (p propertyClient) Store(u *property.Property) error {
	panic("implement me")
}

func (p propertyClient) Update(u *property.Property) error {
	panic("implement me")
}
