package property

import (
	"errors"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

const (
	View = "View"
	Edit = "Edit"
	Hide = "Hide"
)

// Property model describes a single key value pair for a check_service(parameters). An example could be a port for HTTP checking
type Property struct {
	ServiceID uuid.UUID `json:"service_id" gorm:"type:uuid;not null;primary_key"`
	// Key represents property for a struct located under exec/services. Example: Port, or Password
	Key string `json:"key" gorm:"not null;primary_key"`
	// Value represents property value for a struct located under exec/services. Example: 80, or SOME_SECURE_PASSWORD
	Value *string `json:"value" gorm:"not null;default:''"`
	// Status is a type of a property that is either View, Edit, or Hide. View allows users to ONLY view the given property. Edit Allows to both View, and Edit the given property, and finally Hide ensures that property is hidden from the competitor's view
	Status string `json:"status,omitempty" gorm:"not null;default:'View'"`
}

var ErrInvalidStatus = errors.New("property Status should either be View, Edit, or Hide")

func (p *Property) BeforeSave(tx *gorm.DB) (err error) {
	if p.Status != "" {
		var validStatus bool
		for _, item := range []string{View, Edit, Hide} {
			if item == p.Status {
				validStatus = true
			}
		}
		if !validStatus {
			return ErrInvalidStatus
		}
	}
	return nil
}

var ErrPropertyKeyShouldNotBeEmpty = errors.New("property key should not be empty")

func (p *Property) BeforeCreate(_ *gorm.DB) (err error) {
	if p.Key == "" {
		return ErrPropertyKeyShouldNotBeEmpty
	}
	return nil
}

func PropertiesToMap(props []*Property) map[string]string {
	params := map[string]string{}
	for _, p := range props {
		params[p.Key] = *p.Value
	}
	return params
}
