package service

import (
	"errors"
	"fmt"
	"github.com/ScoreTrak/ScoreTrak/pkg/check"
	"github.com/ScoreTrak/ScoreTrak/pkg/exec/resolver"
	"github.com/ScoreTrak/ScoreTrak/pkg/property"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

// Service Model represents a check_service that is being scored for a given host
type Service struct {
	ID uuid.UUID `json:"id,omitempty" gorm:"type:uuid;primary_key;"`

	//Name of the check_service that will be checked against known services
	Name string `json:"name" gorm:"not null;default:null"`

	DisplayName string `json:"display_name,omitempty"`

	// Weight granted for a successful check
	Weight *uint64 `json:"weight" gorm:"not null;default:0"`

	PointsBoost *uint64 `json:"points_boost" gorm:"not null;default:0"`

	// The frequency of a check_service check. If round_units is 5 and round_delay is 0, then check_service checks will happen on every 5th round. (5,10, etc)
	RoundUnits uint64 `json:"round_units,omitempty" gorm:"not null;default:1"`

	// The frequency of a check_service check. If round_units is 7 and round_delay is 3, then check_service checks will happen on every 7th round with an offset of 3. (10,17, etc)
	RoundDelay *uint64 `json:"round_delay,omitempty" gorm:"not null;default:0"`

	// ID of a check_service group the check_service belongs to
	ServiceGroupID uuid.UUID `json:"service_group_id" gorm:"type:uuid;not null"`

	// ID of a host the check_service belongs to
	HostID uuid.UUID `json:"host_id" gorm:"type:uuid;not null"`

	// Enables or Disables the check_service
	Enabled *bool `json:"enabled,omitempty" gorm:"not null;default:true"`

	Properties []*property.Property `json:"properties,omitempty" gorm:"foreignkey:ServiceID; constraint:OnUpdate:RESTRICT,OnDelete:CASCADE"`

	Checks []*check.Check `json:"checks,omitempty" gorm:"foreignkey:ServiceID; constraint:OnUpdate:RESTRICT,OnDelete:CASCADE"`
}

// BeforeCreate will set a UUID rather than numeric ID.
func (s *Service) BeforeCreate(tx *gorm.DB) (err error) {
	if s.ID == uuid.Nil {
		u, err := uuid.NewV4()
		if err != nil {
			return err
		}
		s.ID = u
	}
	return nil
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
		if se.RoundDelay != nil && *(se.RoundDelay) >= s.RoundUnits {
			db.AddError(errors.New("round delay should not be larger than round unit"))
		}
	}

	if s.Name != "" && resolver.ExecutableByName(s.Name) == nil {
		db.AddError(fmt.Errorf("name %s doesn't resolve to scorable check_service", s.Name))
	}
}
