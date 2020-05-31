package orm

import (
	"ScoreTrak/pkg/logger"
	"ScoreTrak/pkg/swarm"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
)

type swarmRepo struct {
	db  *gorm.DB
	log logger.LogInfoFormat
}

func NewSwarmRepo(db *gorm.DB, log logger.LogInfoFormat) swarm.Repo {
	return &swarmRepo{db, log}
}

func (s *swarmRepo) Delete(id uint64) error {
	s.log.Debugf("deleting the swarm with id : %d", id)

	result := s.db.Delete(&swarm.Swarm{}, "id = ?", id)

	if result.Error != nil {
		errMsg := fmt.Sprintf("error while deleting the swarm with id : %d", id)
		s.log.Errorf(errMsg)
		return errors.New(errMsg)
	}

	if result.RowsAffected == 0 {
		return &NoRowsAffected{"no model found for id"}
	}

	return nil
}

func (s *swarmRepo) GetAll() ([]*swarm.Swarm, error) {
	s.log.Debug("get all the swarms")

	swarms := make([]*swarm.Swarm, 0)
	err := s.db.Find(&swarms).Error
	if err != nil {
		s.log.Debug("not a single swarm found")
		return nil, err
	}
	return swarms, nil
}

func (s *swarmRepo) GetByID(id uint64) (*swarm.Swarm, error) {
	s.log.Debugf("get swarm details by id : %s", id)

	swm := &swarm.Swarm{}
	err := s.db.Where("id = ?", id).First(&swm).Error
	if err != nil {
		s.log.Errorf("swarm not found with id : %d, reason : %v", id, err)
		return nil, err
	}
	return swm, nil
}

func (s *swarmRepo) Store(swm *swarm.Swarm) error {
	s.log.Debugf("creating the swarm with id : %v", swm.ID)

	err := s.db.Create(&swm).Error
	if err != nil {
		s.log.Errorf("error while creating the swarm, reason : %v", err)
		return err
	}
	return nil
}

func (s *swarmRepo) Update(swm *swarm.Swarm) error {
	s.log.Debugf("updating the swarm, id : %v", swm.ID)
	err := s.db.Model(&swm).Updates(swarm.Swarm{Label: swm.Label}).Error
	if err != nil {
		s.log.Errorf("error while updating the swarm, reason : %v", err)
		return err
	}
	return nil
}
