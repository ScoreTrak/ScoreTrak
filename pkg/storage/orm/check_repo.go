package orm

import (
	"context"
	"errors"
	"fmt"
	"github.com/ScoreTrak/ScoreTrak/pkg/check"
	"github.com/ScoreTrak/ScoreTrak/pkg/logger"
	"github.com/ScoreTrak/ScoreTrak/pkg/storage/orm/util"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type checkRepo struct {
	db  *gorm.DB
	log logger.LogInfoFormat
}

func NewCheckRepo(db *gorm.DB, log logger.LogInfoFormat) check.Repo {
	return &checkRepo{db, log}
}

func (c *checkRepo) GetAllByRoundID(ctx context.Context, roundID uint) ([]*check.Check, error) {
	c.log.Debug("get all the checks")
	var checks []*check.Check
	err := c.db.WithContext(ctx).Where("round_id = ?", roundID).Find(&checks).Error
	return checks, err
}

func (c *checkRepo) GetAllByServiceID(ctx context.Context, serviceID uuid.UUID) ([]*check.Check, error) {
	c.log.Debug("get all the checks")
	var checks []*check.Check
	err := c.db.WithContext(ctx).Where("service_id = ?", serviceID).Find(&checks).Error
	return checks, err
}

func (c *checkRepo) GetByRoundServiceID(ctx context.Context, roundID uint, serviceID uuid.UUID) (*check.Check, error) {
	c.log.Debug("get all the checks")
	chk := &check.Check{}
	err := c.db.WithContext(ctx).Where("round_id = ? AND service_id = ?", roundID, serviceID).First(&chk).Error
	if err != nil {
		c.log.Errorf("the check with rid, sid : %d, %d not found, reason : %v", roundID, serviceID, err)
		return nil, err
	}
	return chk, err
}

func (c *checkRepo) Delete(ctx context.Context, roundID uint, serviceID uuid.UUID) error {
	c.log.Debugf("deleting the check with rid, sid : %d, %d", roundID, serviceID)
	result := c.db.WithContext(ctx).Delete(&check.Check{}, "round_id = ? AND service_id = ?", roundID, serviceID)
	if result.Error != nil {
		errMsg := fmt.Sprintf("error while deleting the check with rid, sid : %d, %d", roundID, serviceID)
		c.log.Errorf(errMsg)
		return errors.New(errMsg)
	}

	if result.RowsAffected == 0 {
		return &NoRowsAffected{"no model found for round ID, and service id provided"}
	}

	return nil
}

func (c *checkRepo) GetAll(ctx context.Context) ([]*check.Check, error) {
	c.log.Debug("get all the checks")

	checks := make([]*check.Check, 0)
	err := c.db.WithContext(ctx).Find(&checks).Error
	if err != nil {
		c.log.Debug("not a single check found")
		return nil, err
	}
	return checks, nil
}

func (c *checkRepo) GetByID(ctx context.Context, roundID uint, serviceID uuid.UUID) (*check.Check, error) {
	c.log.Debugf("get the check with rid, sid : %d, %d", roundID, serviceID)

	chck := &check.Check{}
	err := c.db.WithContext(ctx).Where("round_id = ? AND service_id = ?", roundID, serviceID).First(&chck).Error
	if err != nil {
		c.log.Errorf("the check with rid, sid : %d, %d not found, reason : %v", roundID, serviceID, err)
		return nil, err
	}
	return chck, nil
}

func (c *checkRepo) Store(ctx context.Context, chck []*check.Check) error {
	err := c.db.WithContext(ctx).Create(chck).Error
	if err != nil {
		c.log.Errorf("error while creating the check, reason : %v", err)
		return err
	}
	return nil
}

func (c *checkRepo) Upsert(ctx context.Context, chck []*check.Check) error {
	err := c.db.WithContext(ctx).Clauses(clause.OnConflict{DoNothing: true}).Create(chck).Error
	if err != nil {
		c.log.Errorf("error while creating the check, reason : %v", err)
		return err
	}
	return nil
}

func (c *checkRepo) TruncateTable(ctx context.Context) (err error) {
	err = util.TruncateTable(ctx, &check.Check{}, c.db)
	if err != nil {
		return err
	}
	return nil
}
