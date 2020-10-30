package client

import (
	"github.com/ScoreTrak/ScoreTrak/pkg/config"
)

type ConfigClient struct {
	s ScoretrakClient
}

func NewConfigClient(c ScoretrakClient) *ConfigClient {
	return &ConfigClient{c}
}

func (c ConfigClient) Get() (*config.DynamicConfig, error) {
	conf := &config.DynamicConfig{}
	err := c.s.GenericGet(conf, "/config")
	if err != nil {
		return nil, err
	}
	return conf, nil
}

func (c ConfigClient) Update(d *config.DynamicConfig) error {
	return c.s.GenericUpdate(d, "/config")
}

func (c ConfigClient) DeleteCompetition() error {
	return c.s.GenericDelete("/config/delete_competition")
}

func (c ConfigClient) ResetScores() error {
	return c.s.GenericDelete("/config/reset_competition")
}

func NewStaticConfigClient(c ScoretrakClient) *StaticConfigClient {
	return &StaticConfigClient{c}
}

type StaticConfigClient struct {
	s ScoretrakClient
}

func (c StaticConfigClient) Get() (*config.StaticConfig, error) {
	conf := &config.StaticConfig{}
	err := c.s.GenericGet(conf, "/static_config")
	if err != nil {
		return nil, err
	}
	return conf, nil
}
