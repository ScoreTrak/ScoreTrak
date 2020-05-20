package config

import (
	"github.com/jinzhu/configor"
	"sync"
)

var mu sync.RWMutex
var dc DynamicConfig

// Dynamic Config model is a set of columns describing the staticConfig of the scoring engine
type DynamicConfig struct {
	// Describes how long each round unit takes to execute in seconds. This value shuold have a minimum value enforced (something like 20 seconds)
	RoundDuration uint64 `json:"round_durration,omitempty" default:"60"`
	// Enables or disables competition globally
	Enabled bool `json:"enabled,omitempty" default:"false" gorm:"not null default: false"`
}

//PullConfig retrieves the staticConfig from the database, and updates the shared staticConfig variable
func PullConfig() {
	mu.Lock()
	defer mu.Unlock()
	//query staticConfig from DB, update it
}

//PushConfig pushes the provided staticConfig to database, AND updates staticConfig variable
func PushConfig() {
	mu.Lock()
	defer mu.Unlock()

	//Updates staticConfig in DB
}

func GetRoundDuration() uint64 {
	mu.RLock()
	defer mu.RUnlock()
	return dc.RoundDuration
}

func GetEnabled() bool {
	mu.RLock()
	defer mu.RUnlock()
	return dc.Enabled
}

func NewDynamicConfig() error {
	err := configor.Load(&staticConfig, "configs/config.yml")
	if err != nil {
		return err
	}
	return nil
}
