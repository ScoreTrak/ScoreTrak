package client

import (
	"ScoreTrak/pkg/config"
	"fmt"
)

type configClient struct {
	s ScoretrakClient
}

func NewConfigClient(c ScoretrakClient) config.Serv {
	return &configClient{c}
}

func (c configClient) Get() (*config.DynamicConfig, error) {
	conf := &config.DynamicConfig{}
	err := genericGet(conf, fmt.Sprintf("/config"), c.s)
	if err != nil {
		return nil, err
	}
	return conf, nil
}

func (c configClient) Update(d *config.DynamicConfig) error {
	return genericUpdate(d, fmt.Sprintf("/config"), c.s)
}
