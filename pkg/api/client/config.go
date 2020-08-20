package client

import (
	"fmt"
	"github.com/ScoreTrak/ScoreTrak/pkg/config"
)

type ConfigClient struct {
	s ScoretrakClient
}

func NewConfigClient(c ScoretrakClient) config.Serv {
	return &ConfigClient{c}
}

func (c ConfigClient) Get() (*config.DynamicConfig, error) {
	conf := &config.DynamicConfig{}
	err := c.s.GenericGet(conf, fmt.Sprintf("/config"))
	if err != nil {
		return nil, err
	}
	return conf, nil
}

func (c ConfigClient) Update(d *config.DynamicConfig) error {
	return c.s.GenericUpdate(d, fmt.Sprintf("/config"))
}

func (c ConfigClient) DeleteCompetition() error {
	return c.s.GenericDelete(fmt.Sprintf("/config/delete_competition"))
}

func (c ConfigClient) ResetScores() error {
	return c.s.GenericDelete(fmt.Sprintf("/config/reset_competition"))
}

func NewStaticConfigClient(c ScoretrakClient) config.StaticServ {
	return &StaticConfigClient{c}
}

type StaticConfigClient struct {
	s ScoretrakClient
}

func (c StaticConfigClient) Get() (*config.StaticConfig, error) {
	conf := &config.StaticConfig{}
	err := c.s.GenericGet(conf, fmt.Sprintf("/static_config"))
	if err != nil {
		return nil, err
	}
	return conf, nil
}
