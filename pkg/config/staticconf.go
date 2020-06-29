package config

import (
	"github.com/jinzhu/configor"
	"github.com/jinzhu/copier"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

// StaticConfig is a struct of settings that were set at the start of the application
type StaticConfig struct {
	// token specified on init of the staticConfig
	Token string `json:"-" default:""`

	DB struct {
		Use       string `default:"cockroach"`
		Cockroach struct {
			Enabled           bool   `default:"true"`
			Host              string `default:"cockroach"`
			Port              string `default:"26257"`
			UserName          string `default:"root"`
			Password          string `default:""`
			Database          string `default:"scoretrak"`
			ConfigureZones    bool   `default:"true"`
			DefaultZoneConfig struct {
				GcTtlseconds                    uint64 `default:"600"`
				BackpressureRangeSizeMultiplier uint64 `default:"0"`
			}
		}
	} `json:"-"`

	Logger struct {
		Use         string `default:"zapLogger"`
		Environment string `default:"prod"`
		LogLevel    string `default:"info"`
		FileName    string `default:"scoretrak.log"`
	} `json:"-"`

	Queue struct {
		Use   string `default:"none"`
		Kafka struct {
		}
		NSQ struct {
			NSQD struct {
				Port string `default:"4150"`
				Host string `default:"nsqd"`
			}
			Topic              string `default:"default"`
			MaxInFlight        int    `default:"100"`
			ConcurrentHandlers int    `default:"100"`
			NSQLookupd         struct {
				Hosts []string `default:"[\"nsqlookupd\"]"`
				Port  string   `default:"4161"`
			}
		}
	} `json:"-"`

	Port     string `default:"33333" json:"-"`
	Platform struct {
		Use    string `default:"none"`
		Docker struct {
			Name    string `default:"scoretrak"`
			Host    string `default:"unix:///var/run/docker.sock"`
			Network string `default:"default"`
		}
	} `json:"-"`
}

var staticConfig StaticConfig

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

func GetStaticConfig() StaticConfig {
	return staticConfig
}

func GetConfigCopy() (StaticConfig, error) {
	cp := StaticConfig{}
	err := copier.Copy(&cp, &staticConfig)
	if err != nil {
		return cp, err
	}
	return cp, nil
}

func SaveConfigToYamlFile(f string, config StaticConfig) error {
	b, err := yaml.Marshal(&config)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(f, b, 0600)
	if err != nil {
		return err
	}
	return nil
}
