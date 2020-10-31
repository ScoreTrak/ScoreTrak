package orm

import (
	"context"
	"errors"
	"fmt"
	"github.com/ScoreTrak/ScoreTrak/pkg/logger"
	"github.com/ScoreTrak/ScoreTrak/pkg/service"
	"github.com/ScoreTrak/ScoreTrak/pkg/storage/orm/util"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type serviceRepo struct {
	db  *gorm.DB
	log logger.LogInfoFormat
}

func NewServiceRepo(db *gorm.DB, log logger.LogInfoFormat) service.Repo {
	return &serviceRepo{db, log}
}

func (s *serviceRepo) Delete(ctx context.Context, id uuid.UUID) error {
	s.log.Debugf("deleting the service with id : %d", id)

	result := s.db.WithContext(ctx).Delete(&service.Service{}, "id = ?", id)

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

func (s *serviceRepo) GetAll(ctx context.Context) ([]*service.Service, error) {
	s.log.Debug("get all the services")

	services := make([]*service.Service, 0)
	err := s.db.WithContext(ctx).Find(&services).Error
	if err != nil {
		s.log.Debug("not a single service found")
		return nil, err
	}
	return services, nil
}

func (s *serviceRepo) GetByID(ctx context.Context, id uuid.UUID) (*service.Service, error) {
	s.log.Debugf("get service details by id : %s", id)

	ser := &service.Service{}
	err := s.db.WithContext(ctx).Where("id = ?", id).First(&ser).Error
	if err != nil {
		s.log.Errorf("service not found with id : %d, reason : %v", id, err)
		return nil, err
	}
	return ser, nil
}

func (s *serviceRepo) Store(ctx context.Context, swm []*service.Service) error {
	err := s.db.WithContext(ctx).Create(swm).Error
	if err != nil {
		s.log.Errorf("error while creating the service, reason : %v", err)
		return err
	}
	return nil
}

func (s *serviceRepo) Upsert(ctx context.Context, swm []*service.Service) error {
	err := s.db.WithContext(ctx).Clauses(clause.OnConflict{DoNothing: true}).Create(swm).Error
	if err != nil {
		s.log.Errorf("error while creating the user, reason : %v", err)
		return err
	}
	return nil
}

func (s *serviceRepo) Update(ctx context.Context, swm *service.Service) error {
	s.log.Debugf("updating the service, id : %v", swm.ID)
	err := s.db.WithContext(ctx).Model(swm).Updates(service.Service{Enabled: swm.Enabled,
		Name: swm.Name, Weight: swm.Weight, PointsBoost: swm.PointsBoost, RoundDelay: swm.RoundDelay,
		RoundUnits: swm.RoundUnits, ServiceGroupID: swm.ServiceGroupID,
		HostID: swm.HostID, DisplayName: swm.DisplayName}).Error
	if err != nil {
		s.log.Errorf("error while updating the service, reason : %v", err)
		return err
	}
	return nil
}

func (s *serviceRepo) TruncateTable(ctx context.Context) (err error) {
	err = util.TruncateTable(ctx, &service.Service{}, s.db)
	if err != nil {
		return err
	}
	return nil
}
