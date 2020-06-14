package host

import (
	"ScoreTrak/pkg/service"
	"errors"
	"github.com/jinzhu/gorm"
)

// Host model represents a single machine. This could be an IP address or a resolvable hostname
type Host struct {
	ID uint64 `json:"id,omitempty"`

	Address *string `json:"address" gorm:"not null;default:null" valid:"host,optional"`

	// The ID of a host group that the host belongs to.
	HostGroupID uint64 `json:"host_group_id,omitempty" gorm:"default: null"`

	// The ID of a team that this host belongs too. This parameter is optional, however is needed to appear on the scoring engine.
	TeamID string `json:"team_id,omitempty" gorm:"default: null"`

	// Enables or disables scoring for a single host
	Enabled *bool `json:"enabled,omitempty" gorm:"not null;default: false"`

	// Enables to Edit the hostname. If a single host needs to be eddited for one service, and kept only visible for other service, you can make 2 services that point to same address, and have different edit_host properties.
	EditHost *bool `json:"edit_host,omitempty" gorm:"not null;default: false"`

	Services []service.Service `gorm:"foreignkey:HostID"`
}

func (h Host) Validate(db *gorm.DB) {
	if h.Address != nil {
		var editHost bool
		if h.EditHost != nil {
			editHost = *h.EditHost
		} else {
			he := Host{}
			db.Where("id = ?", h.ID).First(&he)
			if he.EditHost == nil {
				editHost = false
			} else {
				editHost = *he.EditHost
			}
		}
		if !editHost {
			db.AddError(errors.New("you can not edit address until EditHost is True"))
		}
		return
	}
}
