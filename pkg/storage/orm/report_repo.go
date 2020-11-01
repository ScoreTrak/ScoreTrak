package orm

import (
	"context"
	"github.com/ScoreTrak/ScoreTrak/pkg/report/repo"

	"github.com/ScoreTrak/ScoreTrak/pkg/report"
	"gorm.io/gorm"
)

type reportRepo struct {
	db *gorm.DB
}

func NewReportRepo(db *gorm.DB) repo.Repo {
	return &reportRepo{db}
}

func (c *reportRepo) Get(ctx context.Context) (*report.Report, error) {
	cfg := &report.Report{}
	cfg.ID = 1
	c.db.WithContext(ctx).Take(cfg)
	return cfg, nil
}

func (c *reportRepo) Update(ctx context.Context, cfg *report.Report) error {
	cfg.ID = 1
	err := c.db.WithContext(ctx).Model(cfg).Updates(report.Report{Cache: cfg.Cache}).Error
	if err != nil {
		return err
	}
	return nil
}
