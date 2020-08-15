package orm

import (
	"github.com/ScoreTrak/ScoreTrak/pkg/logger"
	"github.com/ScoreTrak/ScoreTrak/pkg/report"
	"gorm.io/gorm"
)

type reportRepo struct {
	db  *gorm.DB
	log logger.LogInfoFormat
}

func NewReportRepo(db *gorm.DB, log logger.LogInfoFormat) report.Repo {
	return &reportRepo{db, log}
}

func (c *reportRepo) Get() (*report.Report, error) {
	cfg := &report.Report{}
	cfg.ID = 1
	c.db.Take(cfg)
	return cfg, nil
}

func (c *reportRepo) Update(cfg *report.Report) error {
	c.log.Debugf("updating the report")
	cfg.ID = 1
	err := c.db.Model(cfg).Updates(report.Report{Cache: cfg.Cache}).Error
	if err != nil {
		c.log.Errorf("error while updating the config, reason : %v", err)
		return err
	}
	return nil
}
