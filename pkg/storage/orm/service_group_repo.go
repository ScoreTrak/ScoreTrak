package orm

import (
	"ScoreTrak/pkg/logger"
	"ScoreTrak/pkg/service_group"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
)

type serviceGroupRepo struct {
	db  *gorm.DB
	log logger.LogInfoFormat
}

func NewServiceGroupRepo(db *gorm.DB, log logger.LogInfoFormat) service_group.Repo {
	return &serviceGroupRepo{db, log}
}

func (s *serviceGroupRepo) Delete(id uint64) error {
	s.log.Debugf("deleting the Service Group with id : %s", id)

	if s.db.Delete(&service_group.ServiceGroup{}, "id = ?", id).Error != nil {
		errMsg := fmt.Sprintf("error while deleting the Service Group with id : %s", id)
		s.log.Errorf(errMsg)
		return errors.New(errMsg)
	}
	return nil
}

func (s *serviceGroupRepo) GetAll() ([]*service_group.ServiceGroup, error) {
	s.log.Debug("get all the serviceGroups")

	serviceGroups := make([]*service_group.ServiceGroup, 0)
	err := s.db.Find(&serviceGroups).Error
	if err != nil {
		s.log.Debug("not a single Service Group found")
		return nil, err
	}
	return serviceGroups, nil
}

func (s *serviceGroupRepo) GetByID(id uint64) (*service_group.ServiceGroup, error) {
	s.log.Debugf("get Service Group details by id : %s", id)

	sgr := &service_group.ServiceGroup{}
	err := s.db.Where("id = ?", id).First(&sgr).Error
	if err != nil {
		s.log.Errorf("serviceGroup not found with id : %s, reason : %v", id, err)
		return nil, err
	}
	return sgr, nil
}

func (s *serviceGroupRepo) Store(sgr *service_group.ServiceGroup) error {
	s.log.Debugf("creating the Service Group with id : %v", sgr.ID)

	err := s.db.Create(&sgr).Error
	if err != nil {
		s.log.Errorf("error while creating the Service Group, reason : %v", err)
		return err
	}
	return nil
}

func (s *serviceGroupRepo) Update(sgr *service_group.ServiceGroup) error {
	s.log.Debugf("updating the Service Group, id : %v", sgr.ID)
	err := s.db.Model(&sgr).Updates(service_group.ServiceGroup{}).Error
	if err != nil {
		s.log.Errorf("error while updating the Service Group, reason : %v", err)
		return err
	}
	return nil
}