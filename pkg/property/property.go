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
	ID uuid.UUID `json:"id,omitempty" gorm:"type:uuid;primary_key;"`

	ServiceID uuid.UUID `json:"service_id" gorm:"type:uuid;not null"`

	Key string `json:"key" gorm:"not null; default: null"`

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

func (p *Property) BeforeCreate(tx *gorm.DB) (err error) {
	if p.ID == uuid.Nil {
		u, err := uuid.NewV4()
		if err != nil {
			return err
		}
		p.ID = u
	}
	return nil
}
