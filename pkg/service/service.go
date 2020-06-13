package service

import (
	"ScoreTrak/pkg/check"
	"ScoreTrak/pkg/exec/resolver"
	"ScoreTrak/pkg/property"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
)

// Service Model represents a service that is being scored for a given host
type Service struct {
	ID uint64 `json:"id,omitempty" gorm:"primary_key"`

	//Name of the service that will be checked against known services
	Name string `json:"name" gorm:"not null;default:null"`

	DisplayName string `json:"display_name,omitempty"`

	// Points granted for a successful check
	Points uint64 `json:"points" gorm:"not null;default: 0"`

	// The frequency of a service check. If round_units is 5 and round_delay is 0, then service checks will happen on every 5th round. (5,10, etc)
	RoundUnits uint64 `json:"round_units,omitempty" gorm:"not null;default:1"`

	// The frequency of a service check. If round_units is 7 and round_delay is 3, then service checks will happen on every 7th round with an offset of 3. (10,17, etc)
	RoundDelay *uint64 `json:"round_delay,omitempty" gorm:"not null;default: 0"`

	// ID of a service group the service belongs to
	ServiceGroupID uint64 `json:"service_group_id" gorm:"not null"`

	// ID of a host the service belongs to
	HostID uint64 `json:"host_id" gorm:"not null"`

	// Enables or Disables the service
	Enabled *bool `json:"enabled,omitempty" gorm:"not null;default: false"`

	Properties []property.Property `json:"-" gorm:"foreignkey:ServiceID"`

	Checks []check.Check `json:"-" gorm:"foreignkey:ServiceID"`
}

func (s Service) Validate(db *gorm.DB) {
	if s.RoundDelay != nil && *(s.RoundDelay) != 0 {
		if s.RoundUnits != 0 {
			if *(s.RoundDelay) >= s.RoundUnits {
				db.AddError(errors.New("round delay should not be larger than round unit"))
			}
		} else {
			se := Service{}
			db.Where("id = ?", s.ID).First(&se)
			if *(s.RoundDelay) >= se.RoundUnits {
				db.AddError(errors.New("round delay should not be larger than round unit"))
			}
		}
		return
	}

	if s.RoundUnits != 0 {
		se := Service{}
		db.Where("id = ?", s.ID).First(&se)
		if *(se.RoundDelay) >= s.RoundUnits {
			db.AddError(errors.New("round delay should not be larger than round unit"))
		}
	}

	if s.Name != "" && resolver.ExecutableByName(s.Name) == nil {
		db.AddError(errors.New(fmt.Sprintf("name %s doesn't resolve to scorable service", s.Name)))
	}
}
