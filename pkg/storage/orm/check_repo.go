package orm

import (
	"ScoreTrak/pkg/check"
	"ScoreTrak/pkg/logger"
	"github.com/jinzhu/gorm"
)

type checkRepo struct {
	db  *gorm.DB
	log logger.LogInfoFormat
}

func NewCheckRepo(db *gorm.DB, log logger.LogInfoFormat) check.Repo {
	return &checkRepo{db, log}
}

func (c *checkRepo) GetAllByTeamRoundID(t_id string, r_id uint64) ([]*check.Check, error) {
	c.log.Debug("get all the checks")

	checks := make([]*check.Check, 0)
	err := c.db.Find(&checks).Error
	if err != nil {
		c.log.Debug("not a single check found")
		return nil, err
	}
	return checks, nil
}

func (c *checkRepo) GetByTeamRoundServiceID(t_id string, r_id uint64, s_id uint64) (*check.Check, error) {
	cfg := &check.Check{}
	return cfg, nil
}
