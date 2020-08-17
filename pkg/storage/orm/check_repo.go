package orm

import (
	"errors"
	"fmt"
	"github.com/ScoreTrak/ScoreTrak/pkg/check"
	"github.com/ScoreTrak/ScoreTrak/pkg/logger"
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

func (c *checkRepo) GetAllByRoundID(rID uint) ([]*check.Check, error) {
	c.log.Debug("get all the checks")
	var checks []*check.Check
	err := c.db.Where("round_id = ?", rID).Find(&checks).Error
	return checks, err
}

func (c *checkRepo) GetByRoundServiceID(rID uint, sID uuid.UUID) (*check.Check, error) {
	c.log.Debug("get all the checks")
	chk := &check.Check{}
	err := c.db.Where("round_id = ? AND service_id = ?", rID, sID).First(&chk).Error
	if err != nil {
		c.log.Errorf("the check with rid, sid : %d, %d not found, reason : %v", rID, sID, err)
		return nil, err
	}
	return chk, err
}

func (c *checkRepo) Delete(rID uint, sID uuid.UUID) error {
	c.log.Debugf("deleting the check with rid, sid : %d, %d", rID, sID)
	result := c.db.Delete(&check.Check{}, "round_id = ? AND service_id = ?", rID, sID)
	if result.Error != nil {
		errMsg := fmt.Sprintf("error while deleting the check with rid, sid : %d, %d", rID, sID)
		c.log.Errorf(errMsg)
		return errors.New(errMsg)
	}

	if result.RowsAffected == 0 {
		return &NoRowsAffected{"no model found for round ID, and service id provided"}
	}

	return nil
}

func (c *checkRepo) GetAll() ([]*check.Check, error) {
	c.log.Debug("get all the checks")

	checks := make([]*check.Check, 0)
	err := c.db.Find(&checks).Error
	if err != nil {
		c.log.Debug("not a single check found")
		return nil, err
	}
	return checks, nil
}

func (c *checkRepo) GetByID(rID uint, sID uuid.UUID) (*check.Check, error) {
	c.log.Debugf("get the check with rid, sid : %d, %d", rID, sID)

	chck := &check.Check{}
	err := c.db.Where("round_id = ? AND service_id = ?", rID, sID).First(&chck).Error
	if err != nil {
		c.log.Errorf("the check with rid, sid : %d, %d not found, reason : %v", rID, sID, err)
		return nil, err
	}
	return chck, nil
}

func (c *checkRepo) Store(chck []*check.Check) error {
	err := c.db.Create(chck).Error
	if err != nil {
		c.log.Errorf("error while creating the check, reason : %v", err)
		return err
	}
	return nil
}

func (c *checkRepo) Upsert(chck []*check.Check) error {
	err := c.db.Clauses(clause.OnConflict{DoNothing: true}).Create(chck).Error
	if err != nil {
		c.log.Errorf("error while creating the user, reason : %v", err)
		return err
	}
	return nil
}
