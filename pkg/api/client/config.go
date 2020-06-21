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
	err := c.s.genericGet(conf, fmt.Sprintf("/config"))
	if err != nil {
		return nil, err
	}
	return conf, nil
}

func (c configClient) Update(d *config.DynamicConfig) error {
	return c.s.genericUpdate(d, fmt.Sprintf("/config"))
}
