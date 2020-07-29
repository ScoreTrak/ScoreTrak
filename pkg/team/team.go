package team

import (
	"github.com/L1ghtman2k/ScoreTrak/pkg/host"
)

// Team model represents internal team model of the scoring engine.
type Team struct {

	// this id refers to ID of a team in web.
	ID uint64 `json:"id" gorm:"primary_key"`

	Name string `json:"name" gorm:"unique;not null;default:null"`

	Enabled *bool `json:"enabled,omitempty" gorm:"not null;default: true"`

	Hosts []*host.Host `gorm:"foreignkey:TeamID;association_foreignkey:ID; constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT" json:"-"`
}
