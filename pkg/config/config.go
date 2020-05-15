package config

import (
	"github.com/jinzhu/configor"
	"sync"
)

// Config model is a set of columns describing the config of the scoring engine
type Config struct {

	// Describes how long each round unit takes to execute in seconds. This value shuold have a minimum value enforced (something like 20 seconds)
	RoundDuration uint64 `json:"round_durration,omitempty" default:"60"`

	// Enables or disables competition globally
	Enabled *bool `json:"enabled,omitempty" default:"false" gorm:"not null default: false"`

	// token specified on init of the config
	Token string `json:"-" default:"" gorm:"-"`

	DB struct {
		Use       string `default:"cockroach"`
		Cockroach struct {
			Enabled  bool   `default:"true"`
			Host     string `default:"cockroach"`
			Port     string `default:"26257"`
			UserName string `default:"root"`
			Password string `default:""`
			Database string `default:"scoretrak"`
		}
	} `json:"-" gorm:"-"`

	Logger struct {
		Use         string `default:"zapLogger"`
		Environment string `default:"prod"`
		LogLevel    string `default:"info"`
		FileName    string `default:"scoretrak.log"`
	} `json:"-" gorm:"-"`

	Queue struct {
		Use   string `default:"nats"`
		Kafka struct {
		}
		Nats struct {
		}
	} `json:"-" gorm:"-"`

	Port string `default:"8080" json:"-" gorm:"-"`
}

var config Config

func NewConfig() (*Config, error) {
	err := configor.Load(&config, "configs/config.yml")
	if err != nil {
		return nil, err
	}
	return &config, nil
}

var mu sync.RWMutex

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

func Token() string {
	return config.Token
}
