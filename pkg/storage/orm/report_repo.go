package orm

import (
	"github.com/ScoreTrak/ScoreTrak/pkg/check"
	"github.com/ScoreTrak/ScoreTrak/pkg/logger"
	"github.com/ScoreTrak/ScoreTrak/pkg/report"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type reportRepo struct {
	db  *gorm.DB
	log logger.LogInfoFormat
}

func NewReportRepo(db *gorm.DB, log logger.LogInfoFormat) report.Repo {
	return &reportRepo{db, log}
}

type totalSuccessfulPerService struct {
	serviceId uuid.UUID
	total     uint64
}

func (c *reportRepo) CountPassedPerService() (map[uuid.UUID]uint64, error) {
	var serviceToSuccess []*totalSuccessfulPerService
	ret := make(map[uuid.UUID]uint64)
	err := c.db.Model(&check.Check{}).Select("service_id, COUNT(*) as total").Group("service_id").Having("passed = ?", true).Scan(&serviceToSuccess).Error
	if err != nil {
		return nil, err
	}
	for i := range serviceToSuccess {
		ret[serviceToSuccess[i].serviceId] = serviceToSuccess[i].total
	}
	return ret, nil
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
