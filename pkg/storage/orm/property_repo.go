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
	p.log.Debugf("deleting the property with id : %d", id)
	result := p.db.Delete(&property.Property{}, "id = ?", id)

	if result.Error != nil {
		errMsg := fmt.Sprintf("error while deleting the property with id : %d", id)
		p.log.Errorf(errMsg)
		return errors.New(errMsg)
	}

	if result.RowsAffected == 0 {
		return &NoRowsAffected{"no model found for id"}
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

func (p *propertyRepo) GetByID(id uint64) (*property.Property, error) {
	p.log.Debugf("get property details by id : %s", id)

	prop := &property.Property{}
	err := p.db.Where("id = ?", id).First(&prop).Error
	if err != nil {
		p.log.Errorf("property not found with id : %d, reason : %v", id, err)
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
	p.log.Debugf("updating the property, id : %v", prop.ID)
	err := p.db.Model(&prop).Updates(property.Property{Value: prop.Value,
		Status: prop.Status, Description: prop.Description,
	}).Error
	if err != nil {
		p.log.Errorf("error while updating the property, reason : %v", err)
		return err
	}
	return nil
}
