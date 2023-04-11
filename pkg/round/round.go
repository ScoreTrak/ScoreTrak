package round

import (
	"time"

	"github.com/ScoreTrak/ScoreTrak/pkg/check"
	"gorm.io/gorm"
)

// Round Model holds the information about all the rounds that have passed
type Round struct {
	// ID can also represent the round number
	ID uint64 `json:"id,omitempty" gorm:"primary_key"`

	// Start represents the start of a given round
	StartedAt time.Time `json:"started_at,omitempty" gorm:"not null;default:CURRENT_TIMESTAMP"`

	// Note is any output from the check that is human readable.
	Note string `json:"note,omitempty"`

	// Err represents the error from the output
	Err string `json:"err,omitempty"`

	// Finish is finish time of a given round
	FinishedAt *time.Time `json:"finished_at,omitempty"`

	// Checks is a collection of child checks
	Checks    []check.Check `json:"checks,omitempty" gorm:"foreignkey:RoundID; constraint:OnUpdate:RESTRICT,OnDelete:CASCADE"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}

func SortByRoundNumber(db *gorm.DB) *gorm.DB {
	return db.Order("round_id desc")
}

func ByRoundId(roundNumber uint64) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("round_id = ?", roundNumber)
	}
}
