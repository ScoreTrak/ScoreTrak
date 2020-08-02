package round

import (
	"github.com/ScoreTrak/ScoreTrak/pkg/check"
	"time"
)

// Round Model holds the information about all the rounds that have passed
type Round struct {

	// Round ID can also represent the round number
	ID uint `json:"id,omitempty" gorm:"primary_key"`

	Start time.Time `json:"start,omitempty" gorm:"not null; default:CURRENT_TIMESTAMP"`

	Note string `json:"note,omitempty"`

	Err string `json:"err,omitempty"`

	Finish *time.Time `json:"finish,omitempty"`

	Checks []check.Check `json:"checks,omitempty" gorm:"foreignkey:RoundID; constraint:OnUpdate:RESTRICT,OnDelete:CASCADE"`
}
