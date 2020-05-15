package orm

import (
	"ScoreTrak/pkg/check"
	"ScoreTrak/pkg/logger"
	"errors"
	"fmt"
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

func (s *checkRepo) Delete(id uint64) error {
	s.log.Debugf("deleting the check with id : %s", id)

	if s.db.Delete(&check.Check{}, "id = ?", id).Error != nil {
		errMsg := fmt.Sprintf("error while deleting the check with id : %s", id)
		s.log.Errorf(errMsg)
		return errors.New(errMsg)
	}
	return nil
}

func (s *checkRepo) GetAll() ([]*check.Check, error) {
	s.log.Debug("get all the checks")

	checks := make([]*check.Check, 0)
	err := s.db.Find(&checks).Error
	if err != nil {
		s.log.Debug("not a single check found")
		return nil, err
	}
	return checks, nil
}

func (s *checkRepo) GetByID(id uint64) (*check.Check, error) {
	s.log.Debugf("get check details by id : %s", id)

	chck := &check.Check{}
	err := s.db.Where("id = ?", id).First(&chck).Error
	if err != nil {
		s.log.Errorf("check not found with id : %s, reason : %v", id, err)
		return nil, err
	}
	return chck, nil
}

func (s *checkRepo) Store(chck *check.Check) error {
	s.log.Debugf("creating the check with id : %v", chck.ID)

	err := s.db.Create(&chck).Error
	if err != nil {
		s.log.Errorf("error while creating the check, reason : %v", err)
		return err
	}
	return nil
}

func (s *checkRepo) Update(chck *check.Check) error {
	s.log.Debugf("updating the check, id : %v", chck.ID)
	err := s.db.Model(&chck).Updates(check.Check{ServiceID: chck.ServiceID,
		RoundID: chck.RoundID, Log: chck.Log, Passed: chck.Passed,
	}).Error
	if err != nil {
		s.log.Errorf("error while updating the check, reason : %v", err)
		return err
	}
	return nil
}
