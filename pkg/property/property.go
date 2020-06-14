package property

import (
	"errors"
	"github.com/jinzhu/gorm"
)

// Property model describes a single key value pair for a service(parameters). An example could be a port for HTTP checking
type Property struct {
	ID uint64 `json:"id,omitempty"`

	ServiceID uint64 `json:"service_id" gorm:"not null"`

	Key string `json:"key" gorm:"not null; default: null"`

	Value string `json:"value" gorm:"not null; default: null"`

	Description string `json:"description"`

	Status string `json:"status,omitempty" gorm:"not null;default:'View'"`
}

func (Property) TableName() string {
	return "properties"
}

func (p Property) Validate(db *gorm.DB) {
	if p.Status != "" {
		for _, item := range []string{"View", "Edit", "Hide"} {
			if item == p.Status {
				goto NextCheck
			}
		}
		db.AddError(errors.New("property Status should either be View, Edit, or Hide"))
		return
	}

NextCheck:
	if p.Value != "" {
		var status string
		if p.Status == "" {
			pe := Property{}
			db.Where("id = ?", p.ID).First(&pe)
			if pe.Status == "" {
				status = "View"
			} else {
				status = pe.Status
			}
		} else {
			status = p.Status
		}
		if status != "Edit" {
			db.AddError(errors.New("you cannot edit value if Status is not Edit"))
		}
	}
	return
}
