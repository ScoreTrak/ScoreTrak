package config

import (
	"errors"
	"fmt"
	"reflect"

	"gorm.io/gorm"
)

// DynamicConfig model is a set of columns describing the dynamicConfig of the scoring engine
type DynamicConfig struct {
	ID uint64 `json:"id,omitempty"`
	// Describes how long each round unit takes to execute in seconds. This value shuold have a minimum value enforced (something like 20 seconds)
	RoundDuration uint64 `json:"round_duration,omitempty" default:"60"`
	// Enables or disables competition globally
	Enabled *bool `json:"enabled,omitempty" default:"false" gorm:"not null;default:false"`
}

var ErrRoundDurationLargerThanMinRoundDuration = errors.New("round Duration should not be larger than MinRoundDuration")

func (d *DynamicConfig) BeforeSave(tx *gorm.DB) (err error) {
	d.ID = 1
	if d.RoundDuration != 0 && d.RoundDuration < uint64(MinRoundDuration.Seconds()) {
		return fmt.Errorf("%w, MinRoundDuration: %d", ErrRoundDurationLargerThanMinRoundDuration, uint64(MinRoundDuration.Seconds()))
	}
	return nil
}

func (d *DynamicConfig) IsEqual(dc *DynamicConfig) bool {
	return reflect.DeepEqual(dc, d)
}
