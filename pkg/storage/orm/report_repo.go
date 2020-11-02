package orm

import (
	"context"
	"github.com/ScoreTrak/ScoreTrak/pkg/check"
	"github.com/ScoreTrak/ScoreTrak/pkg/report/repo"
	"github.com/gofrs/uuid"

	"github.com/ScoreTrak/ScoreTrak/pkg/report"
	"gorm.io/gorm"
)

type reportRepo struct {
	db *gorm.DB
}

func NewReportRepo(db *gorm.DB) repo.Repo {
	return &reportRepo{db}
}

type totalSuccessfulPerService struct {
	id          uuid.UUID
	totalPassed uint64
}

func (c *reportRepo) CountPassedPerService(ctx context.Context) (map[uuid.UUID]uint64, error) {
	var serviceToSuccess []*totalSuccessfulPerService
	ret := make(map[uuid.UUID]uint64)
	err := c.db.WithContext(ctx).Model(&check.Check{}).Select("service_id, sum(Passed) as total").Group("service_id").Having("passed = ?", true).Scan(&serviceToSuccess).Error
	if err != nil {
		return nil, err
	}
	for i := range serviceToSuccess {
		ret[serviceToSuccess[i].id] = serviceToSuccess[i].totalPassed
	}
	return ret, nil
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
