package orm

import (
	"github.com/L1ghtman2k/ScoreTrak/pkg/logger"
	"github.com/L1ghtman2k/ScoreTrak/pkg/report"
	"github.com/jinzhu/gorm"
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
