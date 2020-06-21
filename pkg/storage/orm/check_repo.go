package orm

import (
	"errors"
	"fmt"
	"github.com/L1ghtman2k/ScoreTrak/pkg/check"
	"github.com/L1ghtman2k/ScoreTrak/pkg/logger"
	"github.com/jinzhu/gorm"
)

type checkRepo struct {
	db  *gorm.DB
	log logger.LogInfoFormat
}

func NewCheckRepo(db *gorm.DB, log logger.LogInfoFormat) check.Repo {
	return &checkRepo{db, log}
}

func (c *checkRepo) GetAllByRoundID(rID uint64) ([]*check.Check, error) {
	c.log.Debug("get all the checks")
	var checks []*check.Check
	err := c.db.Where("round_id = ?", rID).Find(&checks).Error
	return checks, err
}

func (c *checkRepo) GetByRoundServiceID(rID uint64, sID uint64) (*check.Check, error) {
	c.log.Debug("get all the checks")
	chk := &check.Check{}
	err := c.db.Where("round_id = ? AND service_id = ?", rID, sID).First(&chk).Error
	if err != nil {
		c.log.Errorf("the check with rid, sid : %d, %d not found, reason : %v", rID, sID, err)
		return nil, err
	}
	return chk, err
}

func (c *checkRepo) Delete(rID uint64, sID uint64) error {
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

func (c *checkRepo) GetByID(rID uint64, sID uint64) (*check.Check, error) {
	c.log.Debugf("get the check with rid, sid : %d, %d", rID, sID)

	chck := &check.Check{}
	err := c.db.Where("round_id = ? AND service_id = ?", rID, sID).First(&chck).Error
	if err != nil {
		c.log.Errorf("the check with rid, sid : %d, %d not found, reason : %v", rID, sID, err)
		return nil, err
	}
	return chck, nil
}

func (c *checkRepo) Store(chck *check.Check) error {
	err := c.db.Create(chck).Error
	if err != nil {
		c.log.Errorf("error while creating the check, reason : %v", err)
		return err
	}
	return nil
}

//This method could allow for optimization as per gorm's https://github.com/jinzhu/gorm/issues/255#issuecomment-590287329
func (c *checkRepo) StoreMany(chcks []*check.Check) error {
	for _, chck := range chcks {
		err := c.Store(chck)
		if err != nil {
			return err
		}
	}
	return nil
}
