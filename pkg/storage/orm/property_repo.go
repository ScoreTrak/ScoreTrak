package orm

import (
	"errors"
	"fmt"
	"github.com/ScoreTrak/ScoreTrak/pkg/logger"
	"github.com/ScoreTrak/ScoreTrak/pkg/property"
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

func (p *propertyRepo) Delete(id uuid.UUID, key string) error {
	p.log.Debugf("deleting the property with service_id : %d", id)
	result := p.db.Delete(&property.Property{}, "service_id = ? AND key = ?", id, key)

	if result.Error != nil {
		errMsg := fmt.Sprintf("error while deleting the property with service_id : %d and key: %s", id, key)
		p.log.Errorf(errMsg)
		return errors.New(errMsg)
	}

	if result.RowsAffected == 0 {
		return &NoRowsAffected{"no model found"}
	}

	return nil

}

func (p *propertyRepo) GetAll() ([]*property.Property, error) {
	p.log.Debug("get all the properties")
	properties := make([]*property.Property, 0)
	err := p.db.Find(&properties).Error
	if err != nil {
		p.log.Debug("not a single property found")
		return nil, err
	}
	return properties, nil
}

func (p *propertyRepo) GetAllByServiceID(id uuid.UUID) ([]*property.Property, error) {
	p.log.Debugf("get property details by service_id : %s", id)
	properties := make([]*property.Property, 0)
	err := p.db.Where("service_id = ?", id).Find(&properties).Error
	if err != nil {
		p.log.Errorf("property not found with service_id : %d, reason : %v", id, err)
		return nil, err
	}
	return properties, nil
}

func (p *propertyRepo) GetByServiceIDKey(serviceID uuid.UUID, key string) (*property.Property, error) {
	prop := &property.Property{}
	err := p.db.Where("service_id = ? AND key = ?", serviceID, key).First(prop).Error
	if err != nil {
		p.log.Errorf("property not found with service_id : %d, key : %s, reason : %v", serviceID, key, err)
		return nil, err
	}
	return prop, nil
}

func (p *propertyRepo) Store(prop []*property.Property) error {
	err := p.db.Create(prop).Error
	if err != nil {
		p.log.Errorf("error while creating the property, reason : %v", err)
		return err
	}
	return nil
}

func (p *propertyRepo) Upsert(prop []*property.Property) error {
	err := p.db.Clauses(clause.OnConflict{DoNothing: true}).Create(prop).Error
	if err != nil {
		p.log.Errorf("error while creating the user, reason : %v", err)
		return err
	}
	return nil
}

func (p *propertyRepo) Update(prop *property.Property) error {
	err := p.db.Model(prop).Updates(property.Property{Value: prop.Value,
		Status: prop.Status, Description: prop.Description,
	}).Error
	if err != nil {
		p.log.Errorf("error while updating the property, reason : %v", err)
		return err
	}
	return nil
}
