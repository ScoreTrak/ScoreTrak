package team

import (
	"ScoreTrak/pkg/host"
)

// Team model represents internal team model of the scoring engine.
type Team struct {

	// this id refers to ID of a team in web.
	ID string `json:"id" gorm:"primary_key"`

	Enabled bool `json:"enabled,omitempty, not null"`

	Hosts []host.Host `gorm:"foreignkey:TeamID" json:"-"`
}

func (Team) TableName() string {
	return "teams"
}
