package orm

import (
	"context"
	"errors"
	"fmt"
	"github.com/ScoreTrak/ScoreTrak/pkg/property/property_repo"

	"github.com/ScoreTrak/ScoreTrak/pkg/property"
	"github.com/ScoreTrak/ScoreTrak/pkg/storage/orm/testutil"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type propertyRepo struct {
	db *gorm.DB
}

func NewPropertyRepo(db *gorm.DB) property_repo.Repo {
	return &propertyRepo{db}
}

func (p *propertyRepo) Delete(ctx context.Context, serviceID uuid.UUID, key string) error {
	result := p.db.WithContext(ctx).Delete(&property.Property{}, "service_id = ? AND key = ?", serviceID, key)

	if result.Error != nil {
		errMsg := fmt.Sprintf("error while deleting the property with service_id : %d and key: %s", serviceID, key)
		return errors.New(errMsg)
	}

	if result.RowsAffected == 0 {
		return &NoRowsAffected{"no model found"}
	}

	return nil

}

func (p *propertyRepo) GetAll(ctx context.Context) ([]*property.Property, error) {
	properties := make([]*property.Property, 0)
	err := p.db.WithContext(ctx).Find(&properties).Error
	if err != nil {
		return nil, err
	}
	return properties, nil
}

func (p *propertyRepo) GetAllByServiceID(ctx context.Context, serviceID uuid.UUID) ([]*property.Property, error) {
	properties := make([]*property.Property, 0)
	err := p.db.WithContext(ctx).Where("service_id = ?", serviceID).Find(&properties).Error
	if err != nil {
		return nil, err
	}
	return properties, nil
}

func (p *propertyRepo) GetByServiceIDKey(ctx context.Context, serviceID uuid.UUID, key string) (*property.Property, error) {
	prop := &property.Property{}
	err := p.db.WithContext(ctx).Where("service_id = ? AND key = ?", serviceID, key).First(prop).Error
	if err != nil {
		return nil, err
	}
	return prop, nil
}

func (p *propertyRepo) Store(ctx context.Context, prop []*property.Property) error {
	err := p.db.WithContext(ctx).Create(prop).Error
	if err != nil {
		return err
	}
	return nil
}

func (p *propertyRepo) Upsert(ctx context.Context, prop []*property.Property) error {
	err := p.db.WithContext(ctx).Clauses(clause.OnConflict{DoNothing: true}).Create(prop).Error
	if err != nil {
		return err
	}
	return nil
}

func (p *propertyRepo) Update(ctx context.Context, prop *property.Property) error {
	err := p.db.WithContext(ctx).Model(prop).Updates(property.Property{Value: prop.Value,
		Status: prop.Status,
	}).Error
	if err != nil {
		return err
	}
	return nil
}

func (p *propertyRepo) TruncateTable(ctx context.Context) (err error) {
	err = testutil.TruncateTable(ctx, &property.Property{}, p.db)
	if err != nil {
		return err
	}
	return nil
}
