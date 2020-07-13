package client

import (
	"fmt"
	"github.com/L1ghtman2k/ScoreTrak/pkg/config"
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

func NewStaticConfigClient(c ScoretrakClient) config.StaticServ {
	return &staticConfigClient{c}
}

type staticConfigClient struct {
	s ScoretrakClient
}

func (c staticConfigClient) Get() (*config.StaticConfig, error) {
	conf := &config.StaticConfig{}
	err := c.s.genericGet(conf, fmt.Sprintf("/static_config"))
	if err != nil {
		return nil, err
	}
	return conf, nil
}
