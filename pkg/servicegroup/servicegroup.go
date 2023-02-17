package servicegroup

import (
	"errors"
	"regexp"

	"github.com/ScoreTrak/ScoreTrak/pkg/service"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

// ServiceGroup model describes a grouping of services.
type ServiceGroup struct {
	ID uuid.UUID `json:"id,omitempty" gorm:"type:uuid;primary_key;"`
	// Name is a name of the service group. This value is typically used when integrating with pkg/platform
	Name string `json:"name" gorm:"not null;unique;default:null"`

	// DisplayName is not used for anything other than displaying service group on web-ui
	DisplayName string `json:"display_name,omitempty" gorm:"unique"`

	// Enabled enables or disables the check_service
	Enabled *bool `json:"enabled,omitempty" gorm:"not null;default:false"`

	// SkipHelper  Skips pkg/platform automation
	SkipHelper bool `json:"skip_helper,omitempty" gorm:"-"`

	// Label is used to help with deployment of workers when using pkg/platform. For instance, specifying Label:default, will deploy all workers in k8s's daemonset where nodes have a label scoretrak_worker: default
	Label string `json:"label,omitempty"`

	// Services represents all services that are part of a given service group
	Services []*service.Service `json:"services,omitempty" gorm:"foreignkey:ServiceGroupID; constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT"`
}

var ErrNameDoesNotPassValidation = errors.New(`name must pass following regex: "^[\.a-zA-Z0-9_-]+(#ephemeral)?$", and be less than 64 characters`)

// BeforeSave ensures that name is set to correct value. In particular, the name should adhere to ^[\.a-zA-Z0-9_-]+(#ephemeral)?$ regex as per https://github.com/nsqio/go-nsq/blob/04552936c57a26026c39e10a8993805e0f5a73d0/protocol.go
func (s *ServiceGroup) BeforeSave(tx *gorm.DB) (err error) {
	if s.Name != "" {
		if len(s.Name) > 1 && len(s.Name) <= 64 && regexp.MustCompile(`^[\.a-zA-Z0-9_-]+(#ephemeral)?$`).MatchString(s.Name) {
			return
		}
		return ErrNameDoesNotPassValidation
	}
	return nil
}

var ErrNameIsRequired = errors.New("name is required")

// BeforeCreate ensures UUID is set.
func (s *ServiceGroup) BeforeCreate(tx *gorm.DB) (err error) {
	if s.Name == "" {
		return ErrNameIsRequired
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
