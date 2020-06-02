package client

import "ScoreTrak/pkg/host"

type hostClient struct {
	s ScoretrakClient
}

func NewHostClient(c ScoretrakClient) host.Serv {
	return &hostClient{c}
}

func (h hostClient) Delete(id uint64) error {
	panic("implement me")
}

func (h hostClient) GetAll() ([]*host.Host, error) {
	panic("implement me")
}

func (h hostClient) GetByID(id uint64) (*host.Host, error) {
	panic("implement me")
}

func (h hostClient) Store(u *host.Host) error {
	panic("implement me")
}

func (h hostClient) Update(u *host.Host) error {
	panic("implement me")
}
