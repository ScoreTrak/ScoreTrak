package config

import (
	"errors"
	"fmt"
	"github.com/jinzhu/configor"
	"github.com/jinzhu/gorm"
	"reflect"
)

// Dynamic Config model is a set of columns describing the dynamicConfig of the scoring engine
type DynamicConfig struct {
	ID uint64 `json:"id,omitempty"`
	// Describes how long each round unit takes to execute in seconds. This value shuold have a minimum value enforced (something like 20 seconds)
	RoundDuration uint64 `json:"round_duration,omitempty" default:"60"`
	// Enables or disables competition globally
	Enabled *bool `json:"enabled,omitempty" default:"false" gorm:"not null;default: false"`
}

func (d DynamicConfig) Validate(db *gorm.DB) {
	if d.RoundDuration != 0 && d.RoundDuration < uint64(MinRoundDuration.Seconds()) {
		db.AddError(errors.New(fmt.Sprintf("Round Duration should not be larger than MinRoundDuration, which is %d", uint64(MinRoundDuration.Seconds()))))
	}
}

func (d DynamicConfig) TableName() string {
	return "config"
}

//NewDynamicConfig initializes global config d, but it doesn't need any locking because it is assumed that NewDynamicConfig is ran once at the start of the application
func NewDynamicConfig(f string) (*DynamicConfig, error) {
	d := &DynamicConfig{}
	err := configor.Load(d, f)
	if err != nil {
		return nil, err
	}
	d.ID = 1
	return d, nil
}

func (d *DynamicConfig) IsEqual(dc *DynamicConfig) bool {
	return reflect.DeepEqual(dc, d)
}
