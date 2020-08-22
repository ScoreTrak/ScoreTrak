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

// Property model describes a single key value pair for a service(parameters). An example could be a port for HTTP checking
type Property struct {
	ServiceID uuid.UUID `json:"service_id" gorm:"type:uuid;not null;primary_key"`

	Key string `json:"key" gorm:"not null;primary_key"`

	Value string `json:"value" gorm:"not null; default: null"`

	Description string `json:"description"` //Todo: Description of a property could be moved to a separate table

	Status string `json:"status,omitempty" gorm:"not null;default:'View'"`
}

func (Property) TableName() string {
	return "properties"
}

func (p Property) Validate(db *gorm.DB) {
	if p.Status != "" {
		for _, item := range []string{View, Edit, Hide} {
			if item == p.Status {
				return
			}
		}
		db.AddError(errors.New("property Status should either be View, Edit, or Hide"))
		return
	}
}
