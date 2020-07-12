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
	Token string `default:""`

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
	}

	Logger struct {
		Use         string `default:"zapLogger"`
		Environment string `default:"prod"`
		LogLevel    string `default:"info"`
		FileName    string `default:"scoretrak.log"`
	}

	Queue struct {
		Use   string `default:"none"`
		Kafka struct {
		}
		NSQ struct {
			NSQD struct {
				Port string `default:"4150"`
				Host string `default:"nsqd"`
			}
			IgnoreAllScoresIfWorkerFails bool   `default:"true"`
			Topic                        string `default:"default"`
			MaxInFlight                  int    `default:"200"`
			ConcurrentHandlers           int    `default:"200"`
			NSQLookupd                   struct {
				Hosts []string `default:"[\"nsqlookupd\"]"`
				Port  string   `default:"4161"`
			}
		}
	}

	Port     string `default:"33333"`
	Platform struct {
		Use    string `default:"none"`
		Docker struct {
			Name    string `default:"scoretrak"`
			Host    string `default:"unix:///var/run/docker.sock"`
			Network string `default:"default"`
		}
		Kubernetes struct {
			Namespace string `default:"default"`
		}
	}
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
