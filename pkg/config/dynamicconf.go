package config

import (
	"github.com/jinzhu/configor"
	"sync"
)

var mu sync.RWMutex
var d DynamicConfig

// Dynamic Config model is a set of columns describing the dynamicConfig of the scoring engine
type DynamicConfig struct {
	// Describes how long each round unit takes to execute in seconds. This value shuold have a minimum value enforced (something like 20 seconds)
	RoundDuration uint64 `json:"round_durration,omitempty" default:"60"`
	// Enables or disables competition globally
	Enabled *bool `json:"enabled,omitempty" default:"false" gorm:"not null;default: false"`
}

//PullConfig retrieves the dynamicConfig from the database, and updates the shared dynamicConfig variable
func PullConfig() {
	mu.Lock()
	defer mu.Unlock()
	//query dynamicConfig from DB, update it
}

//UpdateConfig updates dynamic config variable from provided dc variable
func UpdateConfig(dc DynamicConfig) {
	mu.Lock()
	defer mu.Unlock()
	d.RoundDuration = dc.RoundDuration

	if d.Enabled != nil {
		*d.Enabled = *dc.Enabled
	}

	//Updates dynamicConfig in DB
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

func NewDynamicConfig(f string) error {
	err := configor.Load(&d, f)
	if err != nil {
		return err
	}
	return nil
}
