package config

import (
	"errors"
	"fmt"
	"github.com/jinzhu/configor"
	"github.com/jinzhu/copier"
	"github.com/jinzhu/gorm"
	"reflect"
	"sync"
)

var mu sync.RWMutex
var d *DynamicConfig

// Dynamic Config model is a set of columns describing the dynamicConfig of the scoring engine
type DynamicConfig struct {
	ID uint64 `json:"id,omitempty"`
	// Describes how long each round unit takes to execute in seconds. This value shuold have a minimum value enforced (something like 20 seconds)
	RoundDuration uint64 `json:"round_durration,omitempty" default:"60"`
	// Enables or disables competition globally
	Enabled *bool `json:"enabled,omitempty" default:"false" gorm:"not null;default: false"`
}

func (dc DynamicConfig) Validate(db *gorm.DB) {
	if dc.RoundDuration != 0 && dc.RoundDuration < MinRoundDuration {
		db.AddError(errors.New(fmt.Sprintf("Round Duration should not be larger than MinRoundDuration, which is %d", MinRoundDuration)))
	}
}

func (DynamicConfig) TableName() string {
	return "config"
}

//PullConfig retrieves the dynamicConfig from the database, and updates the shared dynamicConfig variable
func PullConfig() {
	mu.Lock()
	defer mu.Unlock()
	//query dynamicConfig from DB, update it
}

//UpdateConfig updates dynamic config variable from provided dc variable
func UpdateConfig(dc *DynamicConfig) {
	mu.Lock()
	defer mu.Unlock()
	copier.Copy(&d, &dc)
}

func GetRoundDuration() uint64 {
	mu.RLock()
	defer mu.RUnlock()
	return d.RoundDuration
}

func GetEnabled() bool {
	mu.RLock()
	defer mu.RUnlock()
	return *d.Enabled
}

//NewDynamicConfig initializes global config d, but it doesn't need any locking because it is assumed that NewDynamicConfig is ran once at the start of the application
func NewDynamicConfig(f string) error {
	if d != nil {
		return errors.New("you shouldn't be initializing the config twice")
	}
	err := configor.Load(&d, f)
	if err != nil {
		return err
	}
	d.ID = 1
	return nil
}

func GetConfigCopy() *DynamicConfig {
	mu.RLock()
	defer mu.RUnlock()
	dc := DynamicConfig{}
	copier.Copy(&dc, &d)
	return &dc
}

func IsEqual(dc *DynamicConfig) bool {
	mu.RLock()
	defer mu.RUnlock()
	return reflect.DeepEqual(&dc, &d)
}
