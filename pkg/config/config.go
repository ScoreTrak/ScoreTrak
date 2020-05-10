package config

import "sync"

// Engine model is a set of columns describing the config of the scoring engine
type Config struct {

	// Describes how long each round unit takes to execute in seconds. This value shuold have a minimum value enforced (something like 20 seconds)
	RoundDurration int64 `json:"round_durration,omitempty"`

	// Enables or disables competition globally
	Enabled bool `json:"enabled,omitempty"`

	// Auth Token specified on init of the config
	Token string `json:"-"`
}

var config Config
var mu sync.RWMutex

//Initializes the auth token
func InitAuthToken(t string) {
	mu.RLock()
	defer mu.RUnlock()
	config.Token = t
}

func GetToken() string {
	mu.RLock()
	defer mu.RUnlock()
	return config.Token
}

//PullConfig retrieves the config from the database, and updates the shared config variable
func PullConfig() {
	mu.Lock()
	defer mu.Unlock()

	//query config from DB, update it
}

//PushConfig pushes the provided config to database, AND updates config variable
func PushConfig() {
	mu.Lock()
	defer mu.Unlock()

	//Updates config in DB
}
