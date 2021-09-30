package host_group

import (
	"github.com/ScoreTrak/ScoreTrak/pkg/host"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

// HostGroup model represents a set of hosts that have a common purpose, but are in different teams. For instance team 1 web, and team 2 web would bellong to a host group Web
type HostGroup struct {
	ID uuid.UUID `json:"id,omitempty" gorm:"type:uuid;primary_key;"`

	Name string `json:"name" gorm:"not null; unique"`
	// Hide is responsible for hiding the Host Group on scoring table
	Hide *bool `json:"pause,omitempty" gorm:"not null;default:false"`
	// Pause is responsible for pausing the Host Group on scoring table
	Pause *bool `json:"hide,omitempty" gorm:"not null;default:false"`

	Hosts []*host.Host `json:"hosts,omitempty" gorm:"foreignkey:HostGroupID; constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT"`
}

func (h *HostGroup) BeforeCreate(tx *gorm.DB) (err error) {
	if h.ID == uuid.Nil {
		u, err := uuid.NewV4()
		if err != nil {
			return err
		}
		h.ID = u
	}
	return nil
}
