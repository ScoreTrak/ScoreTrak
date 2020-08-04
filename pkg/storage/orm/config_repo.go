package orm

import (
	"github.com/ScoreTrak/ScoreTrak/pkg/check"
	"github.com/ScoreTrak/ScoreTrak/pkg/config"
	"github.com/ScoreTrak/ScoreTrak/pkg/host"
	"github.com/ScoreTrak/ScoreTrak/pkg/host_group"
	"github.com/ScoreTrak/ScoreTrak/pkg/logger"
	"github.com/ScoreTrak/ScoreTrak/pkg/property"
	"github.com/ScoreTrak/ScoreTrak/pkg/round"
	"github.com/ScoreTrak/ScoreTrak/pkg/service"
	"github.com/ScoreTrak/ScoreTrak/pkg/service_group"
	"github.com/ScoreTrak/ScoreTrak/pkg/team"
	"gorm.io/gorm"
)

type configRepo struct {
	db  *gorm.DB
	log logger.LogInfoFormat
}

func NewConfigRepo(db *gorm.DB, log logger.LogInfoFormat) config.Repo {
	return &configRepo{db, log}
}

func (c *configRepo) Get() (*config.DynamicConfig, error) {
	cfg := &config.DynamicConfig{}
	cfg.ID = 1
	c.db.Take(cfg)
	return cfg, nil
}

func (c *configRepo) ResetCompetition() error {
	err := c.db.Where("1 = 1").Delete(&check.Check{}).Error
	if err != nil {
		return err
	}
	err = c.db.Where("1 = 1").Delete(&round.Round{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (c *configRepo) DeleteCompetition() error {
	err := c.db.Where("1 = 1").Delete(&check.Check{}).Error
	if err != nil {
		return err
	}
	err = c.db.Where("1 = 1").Delete(&round.Round{}).Error
	if err != nil {
		return err
	}
	err = c.db.Where("1 = 1").Delete(&property.Property{}).Error
	if err != nil {
		return err
	}
	err = c.db.Where("1 = 1").Delete(&service.Service{}).Error
	if err != nil {
		return err
	}
	err = c.db.Where("1 = 1").Delete(&service_group.ServiceGroup{}).Error
	if err != nil {
		return err
	}
	err = c.db.Where("1 = 1").Delete(&host.Host{}).Error
	if err != nil {
		return err
	}
	err = c.db.Where("1 = 1").Delete(&host_group.HostGroup{}).Error
	if err != nil {
		return err
	}
	err = c.db.Where("1 = 1").Delete(&team.Team{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (c *configRepo) Update(cfg *config.DynamicConfig) error {
	c.log.Debugf("updating the config")
	cfg.ID = 1
	err := c.db.Model(cfg).Updates(config.DynamicConfig{RoundDuration: cfg.RoundDuration, Enabled: cfg.Enabled}).Error
	if err != nil {
		c.log.Errorf("error while updating the config, reason : %v", err)
		return err
	}
	return nil
}
