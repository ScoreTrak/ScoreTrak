package round

import (
	"github.com/ScoreTrak/ScoreTrak/pkg/check"
	"time"
)

// Round Model holds the information about all the rounds that have passed
type Round struct {

	// ID can also represent the round number
	ID uint64 `json:"id,omitempty" gorm:"primary_key"`

	// Start represents the start of a given round
	Start time.Time `json:"start,omitempty" gorm:"not null;default:CURRENT_TIMESTAMP"`

	// Note is any output from the check that is human readable.
	Note string `json:"note,omitempty"`

	// Err represents the error from the output
	Err string `json:"err,omitempty"`

	// Finish is finish time of a given round
	Finish *time.Time `json:"finish,omitempty"`

	// Checks is a collection of child checks
	Checks []check.Check `json:"checks,omitempty" gorm:"foreignkey:RoundID; constraint:OnUpdate:RESTRICT,OnDelete:CASCADE"`
}
