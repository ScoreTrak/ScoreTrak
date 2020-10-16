package orm

import (
	"fmt"
	"github.com/ScoreTrak/ScoreTrak/pkg/check"
	"github.com/ScoreTrak/ScoreTrak/pkg/config"
	"github.com/ScoreTrak/ScoreTrak/pkg/host"
	"github.com/ScoreTrak/ScoreTrak/pkg/host_group"
	"github.com/ScoreTrak/ScoreTrak/pkg/logger"
	"github.com/ScoreTrak/ScoreTrak/pkg/property"
	"github.com/ScoreTrak/ScoreTrak/pkg/report"
	"github.com/ScoreTrak/ScoreTrak/pkg/round"
	"github.com/ScoreTrak/ScoreTrak/pkg/service"
	"github.com/ScoreTrak/ScoreTrak/pkg/service_group"
	"github.com/ScoreTrak/ScoreTrak/pkg/storage/orm/util"
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

func (c *configRepo) TruncateTable(v interface{}) error {
	stmt := &gorm.Statement{DB: c.db}
	err := stmt.Parse(v)
	if err != nil {
		return err
	}
	return c.db.Exec(fmt.Sprintf("SET FOREIGN_KEY_CHECKS = 0; TRUNCATE TABLE %s ; SET FOREIGN_KEY_CHECKS = 1;", stmt.Schema.Table)).Error
}

func (c *configRepo) ResetScores() error {
	err := c.TruncateTable(&check.Check{})
	if err != nil {
		return err
	}
	err = c.TruncateTable(&round.Round{})
	if err != nil {
		return err
	}
	err = c.TruncateTable(&report.Report{})
	if err != nil {
		return err
	}
	err = util.LoadReport(c.db)
	if err != nil {
		return err
	}
	return nil
}

func (c *configRepo) DeleteCompetition() error {
	err := c.ResetScores()
	if err != nil {
		return err
	}
	err = c.TruncateTable(&property.Property{})
	if err != nil {
		return err
	}
	err = c.TruncateTable(&service.Service{})
	if err != nil {
		return err
	}
	err = c.TruncateTable(&service_group.ServiceGroup{})
	if err != nil {
		return err
	}
	err = c.TruncateTable(&host.Host{})
	if err != nil {
		return err
	}
	err = c.TruncateTable(&host_group.HostGroup{})
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
