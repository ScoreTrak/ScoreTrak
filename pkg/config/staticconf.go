package config

import "github.com/jinzhu/configor"

// StaticConfig is a struct of settings that were set at the start of the application
type StaticConfig struct {
	// token specified on init of the staticConfig
	Token string `json:"-" default:""`

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
	} `json:"-"`

	Logger struct {
		Use         string `default:"zapLogger"`
		Environment string `default:"prod"`
		LogLevel    string `default:"info"`
		FileName    string `default:"scoretrak.log"`
	} `json:"-"`

	Queue struct {
		Use   string `default:""`
		Kafka struct {
		}
		NSQ struct {
			NSQD struct {
				Port string `default:"4150"`
				Host string `default:"nsqd"`
			}
			Topic              string `default:"default"`
			MaxInFlight        int    `default:"1"`
			ConcurrentHandlers int    `default:"1"`
			NSQLookupd         struct {
				Host string `default:"nsqlookupd"`
				Port string `default:"4161"`
			}
		}
	} `json:"-"`

	Port string `default:"8080" json:"-"`

	Platform string `default:"swarm" json:"-"`
}

var staticConfig StaticConfig

func GetConfig() *StaticConfig {
	return &staticConfig
}

func GetToken() string {
	return staticConfig.Token
}

func NewStaticConfig(f string) error {
	err := configor.Load(&staticConfig, f)
	if err != nil {
		return err
	}
	return nil
}

func GetStaticConfig() *StaticConfig {
	return &staticConfig
}
