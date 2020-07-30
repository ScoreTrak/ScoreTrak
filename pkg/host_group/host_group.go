package host_group

import (
	"github.com/L1ghtman2k/ScoreTrak/pkg/host"
	"github.com/gofrs/uuid"
)

// Host Group model represents a set of hosts that have a common purpose, but are in different teams. For instance team 1 web, and team 2 web would bellong to a host group Web
type HostGroup struct {
	ID uuid.UUID `json:"id,omitempty" gorm:"type:uuid;primary_key;"`

	Name string `json:"name" gorm:"not null; unique"`

	// Enables or disables scoring for a given host group. In case you want to stop scoring a set of simalar hosts, you can set this property to false
	Enabled *bool `json:"enabled,omitempty" gorm:"not null;default: false"`

	Hosts []*host.Host `json:"omitempty" gorm:"foreignkey:HostGroupID; constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT"`
}
