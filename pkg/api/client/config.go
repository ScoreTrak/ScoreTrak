package client

import (
	"ScoreTrak/pkg/config"
)

type configClient struct {
	s ScoretrakClient
}

func NewConfigClient(c ScoretrakClient) config.Serv {
	return &configClient{c}
}

func (c configClient) Get() (*config.DynamicConfig, error) {
	panic("implement me")
}

func (c configClient) Update(dynamicConfig *config.DynamicConfig) error {
	panic("implement me")
}
