package service_group

import (
	"errors"
	"fmt"
	"github.com/ScoreTrak/ScoreTrak/pkg/service"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
	"regexp"
)

// Serv Group model describes a grouping of services.
type ServiceGroup struct {
	ID uuid.UUID `json:"id,omitempty" gorm:"type:uuid;primary_key;"`

	Name string `json:"name" gorm:"not null;unique;default:null"`

	DisplayName string `json:"display_name,omitempty" gorm:"unique;default:'default'"`

	// Enables or Disables the check_service
	Enabled *bool `json:"enabled,omitempty" gorm:"not null;default:false"`

	SkipHelper bool `json:"skip_helper,omitempty" gorm:"-"`

	Label string `json:"label,omitempty"`

	Services []*service.Service `json:"services,omitempty" gorm:"foreignkey:ServiceGroupID; constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT"`
}

func (ServiceGroup) TableName() string {
	return "service_groups"
}

// BeforeSave ensures that name is set to correct value
func (s *ServiceGroup) BeforeSave(tx *gorm.DB) (err error) {
	if s.Name != "" { //https://github.com/nsqio/go-nsq/blob/04552936c57a26026c39e10a8993805e0f5a73d0/protocol.go
		if len(s.Name) > 1 && len(s.Name) <= 64 && regexp.MustCompile(`^[\.a-zA-Z0-9_-]+(#ephemeral)?$`).MatchString(s.Name) {
			return
		}
		return fmt.Errorf("name %s doesn't resolve to scorable check_service", s.Name)
	}
	return nil
}

// BeforeCreate will set a UUID rather than numeric ID.
func (s *ServiceGroup) BeforeCreate(tx *gorm.DB) (err error) {
	if s.Name == "" {
		return errors.New("field Name is a mandatory parameter")
	}
	if s.ID == uuid.Nil {
		u, err := uuid.NewV4()
		if err != nil {
			return err
		}
		s.ID = u
	}
	return nil
}
