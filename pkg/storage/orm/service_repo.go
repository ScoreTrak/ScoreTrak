package orm

import (
	"context"
	"errors"
	"fmt"
	"github.com/ScoreTrak/ScoreTrak/pkg/service/service_repo"

	"github.com/ScoreTrak/ScoreTrak/pkg/service"
	"github.com/ScoreTrak/ScoreTrak/pkg/storage/orm/testutil"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type serviceRepo struct {
	db *gorm.DB
}

func NewServiceRepo(db *gorm.DB) service_repo.Repo {
	return &serviceRepo{db}
}

func (s *serviceRepo) Delete(ctx context.Context, id uuid.UUID) error {
	result := s.db.WithContext(ctx).Delete(&service.Service{}, "id = ?", id)

	if result.Error != nil {
		errMsg := fmt.Sprintf("error while deleting the check_service with id : %d", id)
		return errors.New(errMsg)
	}

	if result.RowsAffected == 0 {
		return &NoRowsAffected{"no model found"}
	}
	return nil
}

func (s *serviceRepo) GetAll(ctx context.Context) ([]*service.Service, error) {
	services := make([]*service.Service, 0)
	err := s.db.WithContext(ctx).Find(&services).Error
	if err != nil {
		return nil, err
	}
	return services, nil
}

func (s *serviceRepo) GetByID(ctx context.Context, id uuid.UUID) (*service.Service, error) {
	ser := &service.Service{}
	err := s.db.WithContext(ctx).Where("id = ?", id).First(&ser).Error
	if err != nil {
		return nil, err
	}
	return ser, nil
}

func (s *serviceRepo) Store(ctx context.Context, swm []*service.Service) error {
	err := s.db.WithContext(ctx).Create(swm).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *serviceRepo) Upsert(ctx context.Context, swm []*service.Service) error {
	err := s.db.WithContext(ctx).Clauses(clause.OnConflict{DoNothing: true}).Create(swm).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *serviceRepo) Update(ctx context.Context, swm *service.Service) error {
	err := s.db.WithContext(ctx).Model(swm).Updates(service.Service{Pause: swm.Pause, Hide: swm.Hide,
		Name: swm.Name, Weight: swm.Weight, PointsBoost: swm.PointsBoost, RoundDelay: swm.RoundDelay,
		RoundUnits: swm.RoundUnits, ServiceGroupID: swm.ServiceGroupID,
		HostID: swm.HostID, DisplayName: swm.DisplayName}).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *serviceRepo) TruncateTable(ctx context.Context) (err error) {
	err = testutil.TruncateTable(ctx, &service.Service{}, s.db)
	if err != nil {
		return err
	}
	return nil
}
