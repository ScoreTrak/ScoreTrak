package host

import (
	"github.com/L1ghtman2k/ScoreTrak/pkg/service"
)

// Host model represents a single machine. This could be an IP address or a resolvable hostname
type Host struct {
	ID uint64 `json:"id,omitempty"`

	Address *string `json:"address" gorm:"not null;default:null" valid:"host,optional"`

	// The ID of a host group that the host belongs to.
	HostGroupID uint64 `json:"host_group_id,omitempty" gorm:"default: null"`

	// The ID of a team that this host belongs too. This parameter is optional, however is needed to appear on the scoring engine.
	TeamName string `json:"team_name,omitempty" gorm:"default: null"`

	// Enables or disables scoring for a single host
	Enabled *bool `json:"enabled,omitempty" gorm:"not null;default: true"`

	// Enables to Edit the hostname. If a single host needs to be eddited for one service, and kept only visible for other service, you can make 2 services that point to same address, and have different edit_host properties.
	EditHost *bool `json:"edit_host,omitempty" gorm:"not null;default: false"`

	Services []*service.Service `json:"omitempty" gorm:"foreignkey:HostID"`
}
