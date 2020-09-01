package host

import (
	"github.com/ScoreTrak/ScoreTrak/pkg/service"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

// Host model represents a single machine. This could be an IP address or a resolvable hostname
type Host struct {
	ID uuid.UUID `json:"id,omitempty" gorm:"type:uuid;primary_key;"`

	Address *string `json:"address" gorm:"not null;default:null" valid:"host,optional"`

	// The ID of a host group that the host belongs to.
	HostGroupID *uuid.UUID `json:"host_group_id,omitempty" gorm:"type:uuid"`

	// The ID of a team that this host belongs too.
	TeamID uuid.UUID `json:"team_id,omitempty" gorm:"type:uuid;not null"`

	// Enables or disables scoring for a single host
	Enabled *bool `json:"enabled,omitempty" gorm:"not null;default:true"`

	// Enables to Edit the hostname. If a single host needs to be eddited for one service, and kept only visible for other service, you can make 2 services that point to same address, and have different edit_host properties.
	EditHost *bool `json:"edit_host,omitempty" gorm:"not null;default: false"`

	Services []*service.Service `json:"services,omitempty" gorm:"foreignkey:HostID; constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT"`
}

func (p *Host) BeforeCreate(tx *gorm.DB) (err error) {
	if p.ID == uuid.Nil {
		u, err := uuid.NewV4()
		if err != nil {
			return err
		}
		p.ID = u
	}
	return nil
}
