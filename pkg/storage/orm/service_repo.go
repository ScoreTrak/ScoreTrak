package orm

import (
	"errors"
	"fmt"
	"github.com/L1ghtman2k/ScoreTrak/pkg/logger"
	"github.com/L1ghtman2k/ScoreTrak/pkg/service"
	"gorm.io/gorm"
)

type serviceRepo struct {
	db  *gorm.DB
	log logger.LogInfoFormat
}

func NewServiceRepo(db *gorm.DB, log logger.LogInfoFormat) service.Repo {
	return &serviceRepo{db, log}
}

func (s *serviceRepo) Delete(id uint32) error {
	s.log.Debugf("deleting the service with id : %d", id)

	result := s.db.Delete(&service.Service{}, "id = ?", id)

	if result.Error != nil {
		errMsg := fmt.Sprintf("error while deleting the service with id : %d", id)
		s.log.Errorf(errMsg)
		return errors.New(errMsg)
	}

	if result.RowsAffected == 0 {
		return &NoRowsAffected{"no model found"}
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

func (s *serviceRepo) GetByID(id uint32) (*service.Service, error) {
	s.log.Debugf("get service details by id : %s", id)

	ser := &service.Service{}
	err := s.db.Where("id = ?", id).First(&ser).Error
	if err != nil {
		s.log.Errorf("service not found with id : %d, reason : %v", id, err)
		return nil, err
	}
	return ser, nil
}

func (s *serviceRepo) Store(swm *service.Service) error {
	s.log.Debugf("creating the service with id : %v", swm.ID)
	err := s.db.Create(swm).Error
	if err != nil {
		s.log.Errorf("error while creating the service, reason : %v", err)
		return err
	}
	return nil
}

func (s *serviceRepo) Update(swm *service.Service) error {
	s.log.Debugf("updating the service, id : %v", swm.ID)
	err := s.db.Model(swm).Updates(service.Service{Enabled: swm.Enabled,
		Name: swm.Name, Points: swm.Points, PointsBoost: swm.PointsBoost, RoundDelay: swm.RoundDelay,
		RoundUnits: swm.RoundUnits, ServiceGroupID: swm.ServiceGroupID,
		HostID: swm.HostID, DisplayName: swm.DisplayName}).Error
	if err != nil {
		s.log.Errorf("error while updating the service, reason : %v", err)
		return err
	}
	return nil
}
