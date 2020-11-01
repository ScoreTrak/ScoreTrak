package orm

import (
	"context"
	"errors"
	"fmt"
	"github.com/ScoreTrak/ScoreTrak/pkg/service_group/repo"

	"github.com/ScoreTrak/ScoreTrak/pkg/service_group"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type serviceGroupRepo struct {
	db *gorm.DB
}

func NewServiceGroupRepo(db *gorm.DB) repo.Repo {
	return &serviceGroupRepo{db}
}

func (s *serviceGroupRepo) Delete(ctx context.Context, id uuid.UUID) error {
	result := s.db.WithContext(ctx).Delete(&service_group.ServiceGroup{}, "id = ?", id)
	if result.Error != nil {
		errMsg := fmt.Sprintf("error while deleting the Service Group with id : %d", id)
		return errors.New(errMsg)
	}
	if result.RowsAffected == 0 {
		return &NoRowsAffected{"no model found"}
	}
	return nil
}

func (s *serviceGroupRepo) GetAll(ctx context.Context) ([]*service_group.ServiceGroup, error) {
	serviceGroups := make([]*service_group.ServiceGroup, 0)
	err := s.db.WithContext(ctx).Find(&serviceGroups).Error
	if err != nil {
		return nil, err
	}
	return serviceGroups, nil
}

func (s *serviceGroupRepo) GetByID(ctx context.Context, id uuid.UUID) (*service_group.ServiceGroup, error) {
	sgr := &service_group.ServiceGroup{}
	err := s.db.WithContext(ctx).Where("id = ?", id).First(sgr).Error
	if err != nil {
		return nil, err
	}
	return sgr, nil
}

func (s *serviceGroupRepo) Store(ctx context.Context, sgr *service_group.ServiceGroup) error {
	err := s.db.WithContext(ctx).Create(sgr).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *serviceGroupRepo) Upsert(ctx context.Context, sgr *service_group.ServiceGroup) error {
	err := s.db.WithContext(ctx).Clauses(clause.OnConflict{DoNothing: true}).Create(sgr).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *serviceGroupRepo) Update(ctx context.Context, sgr *service_group.ServiceGroup) error {
	err := s.db.WithContext(ctx).Model(sgr).Updates(service_group.ServiceGroup{Enabled: sgr.Enabled, DisplayName: sgr.DisplayName}).Error //Updating service group names is not supported because service group name tightly coupled with platform operations
	if err != nil {
		return err
	}
	return nil
}
