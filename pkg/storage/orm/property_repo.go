package orm

import (
	"context"
	"errors"
	"fmt"
	"github.com/ScoreTrak/ScoreTrak/pkg/logger"
	"github.com/ScoreTrak/ScoreTrak/pkg/property"
	"github.com/ScoreTrak/ScoreTrak/pkg/storage/orm/util"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type propertyRepo struct {
	db  *gorm.DB
	log logger.LogInfoFormat
}

func NewPropertyRepo(db *gorm.DB, log logger.LogInfoFormat) property.Repo {
	return &propertyRepo{db, log}
}

func (p *propertyRepo) Delete(ctx context.Context, serviceID uuid.UUID, key string) error {
	p.log.Debugf("deleting the property with service_id : %d", serviceID)
	result := p.db.WithContext(ctx).Delete(&property.Property{}, "service_id = ? AND key = ?", serviceID, key)

	if result.Error != nil {
		errMsg := fmt.Sprintf("error while deleting the property with service_id : %d and key: %s", serviceID, key)
		p.log.Errorf(errMsg)
		return errors.New(errMsg)
	}

	if result.RowsAffected == 0 {
		return &NoRowsAffected{"no model found"}
	}

	return nil

}

func (p *propertyRepo) GetAll(ctx context.Context) ([]*property.Property, error) {
	p.log.Debug("get all the properties")
	properties := make([]*property.Property, 0)
	err := p.db.WithContext(ctx).Find(&properties).Error
	if err != nil {
		p.log.Debug("not a single property found")
		return nil, err
	}
	return properties, nil
}

func (p *propertyRepo) GetAllByServiceID(ctx context.Context, serviceID uuid.UUID) ([]*property.Property, error) {
	p.log.Debugf("get property details by service_id : %s", serviceID)
	properties := make([]*property.Property, 0)
	err := p.db.WithContext(ctx).Where("service_id = ?", serviceID).Find(&properties).Error
	if err != nil {
		p.log.Errorf("property not found with service_id : %d, reason : %v", serviceID, err)
		return nil, err
	}
	return properties, nil
}

func (p *propertyRepo) GetByServiceIDKey(ctx context.Context, serviceID uuid.UUID, key string) (*property.Property, error) {
	prop := &property.Property{}
	err := p.db.WithContext(ctx).Where("service_id = ? AND key = ?", serviceID, key).First(prop).Error
	if err != nil {
		p.log.Errorf("property not found with service_id : %d, key : %s, reason : %v", serviceID, key, err)
		return nil, err
	}
	return prop, nil
}

func (p *propertyRepo) Store(ctx context.Context, prop []*property.Property) error {
	err := p.db.WithContext(ctx).Create(prop).Error
	if err != nil {
		p.log.Errorf("error while creating the property, reason : %v", err)
		return err
	}
	return nil
}

func (p *propertyRepo) Upsert(ctx context.Context, prop []*property.Property) error {
	err := p.db.WithContext(ctx).Clauses(clause.OnConflict{DoNothing: true}).Create(prop).Error
	if err != nil {
		p.log.Errorf("error while creating the user, reason : %v", err)
		return err
	}
	return nil
}

func (p *propertyRepo) Update(ctx context.Context, prop *property.Property) error {
	err := p.db.WithContext(ctx).Model(prop).Updates(property.Property{Value: prop.Value,
		Status: prop.Status,
	}).Error
	if err != nil {
		p.log.Errorf("error while updating the property, reason : %v", err)
		return err
	}
	return nil
}

func (p *propertyRepo) TruncateTable(ctx context.Context) (err error) {
	err = util.TruncateTable(ctx, &property.Property{}, p.db)
	if err != nil {
		return err
	}
	return nil
}
