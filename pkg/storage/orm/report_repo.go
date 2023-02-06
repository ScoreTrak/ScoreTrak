package orm

import (
	"context"
	"gorm.io/gorm/clause"

	"github.com/ScoreTrak/ScoreTrak/pkg/check"
	"github.com/ScoreTrak/ScoreTrak/pkg/report/reportrepo"
	"github.com/gofrs/uuid"

	"github.com/ScoreTrak/ScoreTrak/pkg/report"
	"gorm.io/gorm"
)

type reportRepo struct {
	db *gorm.DB
}

func NewReportRepo(db *gorm.DB) reportrepo.Repo {
	return &reportRepo{db}
}

type totalSuccessfulPerService struct {
	ServiceID uuid.UUID
	Total     uint64
}

func (c *reportRepo) Create(ctx context.Context, r *report.Report) error {
	err := c.db.WithContext(ctx).Create(r).Error
	if err != nil {
		return err
	}
	return nil
}

func (c *reportRepo) CountPassedPerService(ctx context.Context) (map[uuid.UUID]uint64, error) {
	var serviceToSuccess []*totalSuccessfulPerService
	ret := make(map[uuid.UUID]uint64)
	err := c.db.WithContext(ctx).Model(&check.Check{}).Distinct("service_id, COUNT(*) as total").Where("passed = ?", true).Group("service_id").Scan(&serviceToSuccess).Error
	if err != nil {
		return nil, err
	}
	for i := range serviceToSuccess {
		ret[serviceToSuccess[i].ServiceID] = serviceToSuccess[i].Total
	}
	return ret, nil
}

func (c *reportRepo) Get(ctx context.Context) (*report.Report, error) {
	cfg := &report.Report{}
	cfg.ID = 1
	c.db.WithContext(ctx).Take(cfg)
	return cfg, nil
}

func (c *reportRepo) Upsert(ctx context.Context, r *report.Report) error {
	err := c.db.WithContext(ctx).Clauses(clause.OnConflict{DoNothing: true}).Create(r).Error
	if err != nil {
		return err
	}
	return nil
}

func (c *reportRepo) Update(ctx context.Context, cfg *report.Report) error {
	cfg.ID = 1
	err := c.db.WithContext(ctx).Model(cfg).Updates(report.Report{Cache: cfg.Cache}).Error
	if err != nil {
		return err
	}
	return nil
}
