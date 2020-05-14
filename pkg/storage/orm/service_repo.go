package orm

import (
	"ScoreTrak/pkg/logger"
	"ScoreTrak/pkg/service"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
)

type serviceRepo struct {
	db  *gorm.DB
	log logger.LogInfoFormat
}

func NewServiceRepo(db *gorm.DB, log logger.LogInfoFormat) service.Repo {
	return &serviceRepo{db, log}
}

func (s *serviceRepo) Delete(id uint64) error {
	s.log.Debugf("deleting the service with id : %s", id)

	if s.db.Delete(&service.Service{}, "id = ?", id).Error != nil {
		errMsg := fmt.Sprintf("error while deleting the service with id : %s", id)
		s.log.Errorf(errMsg)
		return errors.New(errMsg)
	}
	return nil
}

func (s *serviceRepo) GetAll() ([]*service.Service, error) {
	s.log.Debug("get all the services")

	services := make([]*service.Service, 0)
	err := s.db.Find(&services).Error
	if err != nil {
		s.log.Debug("not a single service found")
		return nil, err
	}
	return services, nil
}

func (s *serviceRepo) GetByID(id uint64) (*service.Service, error) {
	s.log.Debugf("get service details by id : %s", id)

	ser := &service.Service{}
	err := s.db.Where("id = ?", id).First(&ser).Error
	if err != nil {
		s.log.Errorf("service not found with id : %s, reason : %v", id, err)
		return nil, err
	}
	return ser, nil
}

func (s *serviceRepo) Store(swm *service.Service) error {
	s.log.Debugf("creating the service with id : %v", swm.ID)
	err := s.db.Create(&swm).Error
	if err != nil {
		s.log.Errorf("error while creating the service, reason : %v", err)
		return err
	}
	return nil
}

func (s *serviceRepo) Update(swm *service.Service) error {
	s.log.Debugf("updating the service, id : %v", swm.ID)
	err := s.db.Save(&swm).Error
	if err != nil {
		s.log.Errorf("error while updating the service, reason : %v", err)
		return err
	}
	return nil
}
