package orm

import (
	"context"
	"fmt"

	"github.com/ScoreTrak/ScoreTrak/pkg/servicegroup/servicegrouprepo"

	"github.com/ScoreTrak/ScoreTrak/pkg/servicegroup"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type serviceGroupRepo struct {
	db *gorm.DB
}

func NewServiceGroupRepo(db *gorm.DB) servicegrouprepo.Repo {
	return &serviceGroupRepo{db}
}

func (s *serviceGroupRepo) Delete(ctx context.Context, id uuid.UUID) error {
	result := s.db.WithContext(ctx).Delete(&servicegroup.ServiceGroup{}, "id = ?", id)
	if result.Error != nil {
		return fmt.Errorf("error while deleting the Service Group with id: %d, err: %w", id, result.Error)
	}
	if result.RowsAffected == 0 {
		return &NoRowsAffectedError{"no model found"}
	}
	return nil
}

func (s *serviceGroupRepo) GetAll(ctx context.Context) ([]*servicegroup.ServiceGroup, error) {
	serviceGroups := make([]*servicegroup.ServiceGroup, 0)
	err := s.db.WithContext(ctx).Find(&serviceGroups).Error
	if err != nil {
		return nil, err
	}
	return serviceGroups, nil
}

func (s *serviceGroupRepo) GetByID(ctx context.Context, id uuid.UUID) (*servicegroup.ServiceGroup, error) {
	sgr := &servicegroup.ServiceGroup{}
	err := s.db.WithContext(ctx).Where("id = ?", id).First(sgr).Error
	if err != nil {
		return nil, err
	}
	return sgr, nil
}

func (s *serviceGroupRepo) Store(ctx context.Context, sgr *servicegroup.ServiceGroup) error {
	err := s.db.WithContext(ctx).Create(sgr).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *serviceGroupRepo) Upsert(ctx context.Context, sgr *servicegroup.ServiceGroup) error {
	err := s.db.WithContext(ctx).Clauses(clause.OnConflict{DoNothing: true}).Create(sgr).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *serviceGroupRepo) Update(ctx context.Context, sgr *servicegroup.ServiceGroup) error {
	err := s.db.WithContext(ctx).Model(sgr).Updates(servicegroup.ServiceGroup{Enabled: sgr.Enabled, DisplayName: sgr.DisplayName}).Error // Updating check_service group names is not supported because check_service group name tightly coupled with platform operations
	if err != nil {
		return err
	}
	return nil
}
