package competitionsettings

import (
	"errors"
	"fmt"
	"time"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

// DynamicConfig model is a set of columns describing the dynamicConfig of the scoring engine
type CompetitionSettings struct {
	ID uuid.UUID `json:"id,omitempty" gorm:"type:uuid;primary_key;"`
	// Describes how long each round unit takes to execute in seconds. This value shuold have a minimum value enforced (something like 20 seconds)
	RoundDuration uint64 `json:"round_duration,omitempty" default:"60"`
	// Enables or disables competition globally
	Enabled       *bool `json:"enabled,omitempty" default:"false" gorm:"not null;default:false"`
	CompetitionID uuid.UUID
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}

var ErrRoundDurationLargerThanMinRoundDuration = errors.New("round Duration should not be larger than MinRoundDuration")

const MinRoundDuration = time.Duration(20) * time.Second

func (d *CompetitionSettings) BeforeSave(tx *gorm.DB) (err error) {
	if d.RoundDuration != 0 && d.RoundDuration < uint64(MinRoundDuration.Seconds()) {
		return fmt.Errorf("%w, MinRoundDuration: %d", ErrRoundDurationLargerThanMinRoundDuration, uint64(MinRoundDuration.Seconds()))
	}
	return nil
}
