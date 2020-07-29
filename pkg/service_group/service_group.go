package service_group

import (
	"errors"
	"fmt"
	"github.com/L1ghtman2k/ScoreTrak/pkg/service"
	"gorm.io/gorm"
	"regexp"
)

// Serv Group model describes a grouping of services.
type ServiceGroup struct {
	ID uint32 `json:"id,omitempty"`

	Name string `json:"name" gorm:"not null;unique;default:null"`

	DisplayName string `json:"display_name,omitempty" gorm:"default:null"`

	// Enables or Disables the service
	Enabled *bool `json:"enabled,omitempty" gorm:"not null;default: false"`

	SkipHelper bool `json:"skip_helper,omitempty" gorm:"-"`

	Label string `json:"label,omitempty" gorm:"-"`

	Services []*service.Service `json:"services,omitempty" gorm:"foreignkey:ServiceGroupID; constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT"`
}

func (ServiceGroup) TableName() string {
	return "service_groups"
}

func (s ServiceGroup) Validate(db *gorm.DB) {
	if s.Name != "" { //https://github.com/nsqio/go-nsq/blob/04552936c57a26026c39e10a8993805e0f5a73d0/protocol.go
		if len(s.Name) > 1 && len(s.Name) <= 64 && regexp.MustCompile(`^[\.a-zA-Z0-9_-]+(#ephemeral)?$`).MatchString(s.Name) {
			return
		}
		db.AddError(errors.New(fmt.Sprintf("name %s doesn't resolve to scorable service", s.Name)))
	}
}
