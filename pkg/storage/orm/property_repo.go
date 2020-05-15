package orm

import (
	"ScoreTrak/pkg/logger"
	"ScoreTrak/pkg/property"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
)

type propertyRepo struct {
	db  *gorm.DB
	log logger.LogInfoFormat
}

func NewPropertyRepo(db *gorm.DB, log logger.LogInfoFormat) property.Repo {
	return &propertyRepo{db, log}
}

func (p *propertyRepo) Delete(id uint64) error {
	p.log.Debugf("deleting the property with id : %s", id)

	if p.db.Delete(&property.Property{}, "property_id = ?", id).Error != nil {
		errMsg := fmt.Sprintf("error while deleting the property with id : %s", id)
		p.log.Errorf(errMsg)
		return errors.New(errMsg)
	}
	return nil
}

func (p *propertyRepo) GetAll() ([]*property.Property, error) {
	p.log.Debug("get all the propertys")

	propertys := make([]*property.Property, 0)
	err := p.db.Find(&propertys).Error
	if err != nil {
		p.log.Debug("not a single property found")
		return nil, err
	}
	return propertys, nil
}

func (p *propertyRepo) GetByID(id uint64) (*property.Property, error) {
	p.log.Debugf("get property details by id : %s", id)

	prop := &property.Property{}
	err := p.db.Where("property_id = ?", id).First(&prop).Error
	if err != nil {
		p.log.Errorf("property not found with id : %s, reason : %v", id, err)
		return nil, err
	}
	return prop, nil
}

func (p *propertyRepo) Store(prop *property.Property) error {
	p.log.Debugf("creating the property with id : %v", prop.ID)

	err := p.db.Create(&prop).Error
	if err != nil {
		p.log.Errorf("error while creating the property, reason : %v", err)
		return err
	}
	return nil
}

func (p *propertyRepo) Update(prop *property.Property) error {
	p.log.Debugf("updating the property, property_id : %v", prop.ID)
	err := p.db.Model(&prop).Updates(property.Property{
		ServiceID: prop.ServiceID, Key: prop.Key, Value: prop.Value,
		Status: prop.Status,
	}).Error
	if err != nil {
		p.log.Errorf("error while updating the property, reason : %v", err)
		return err
	}
	return nil
}
