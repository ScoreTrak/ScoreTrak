package config

import (
	"io/ioutil"

	"github.com/ScoreTrak/ScoreTrak/pkg/auth"
	"github.com/ScoreTrak/ScoreTrak/pkg/platform/platforming"
	"github.com/ScoreTrak/ScoreTrak/pkg/queue/queueing"
	"github.com/ScoreTrak/ScoreTrak/pkg/storage"
	"github.com/jinzhu/configor"
	"github.com/jinzhu/copier"
	"gopkg.in/yaml.v2"
)

// StaticConfig is a struct of settings that are set at the start of the application. It contains Configs from other packages defined under pkg/ directory.
// StaticConfig is read only at the moment, hence there is no lock / prevention to race conditions.
type StaticConfig struct {
	DB storage.Config
	// This value ideally shouldn't be larger than few seconds
	DatabaseMaxTimeDriftSeconds uint `default:"2"`
	// How frequently to pull dynamic configs
	DynamicConfigPullSeconds uint `default:"5"`

	Queue queueing.Config

	Platform platforming.Config

	PubSubConfig queueing.MasterConfig

	AdminUsername string `default:"admin"`

	AdminPassword string `default:"changeme"`

	Port string `default:"33333"`

	Prod bool `default:"false"`

	CertFile string `default:""`

	KeyFile string `default:""`

	JWT auth.Config
}

var staticConfig StaticConfig

func GetPlatformConfig() platforming.Config {
	return staticConfig.Platform
}

func GetQueueConfig() queueing.Config {
	return staticConfig.Queue
}

func GetDBConfig() storage.Config {
	return staticConfig.DB
}

func GetJWTConfig() auth.Config {
	return staticConfig.JWT
}

func GetPubSubConfig() queueing.MasterConfig {
	return staticConfig.PubSubConfig
}

func NewStaticConfig(f string) error {
	err := configor.Load(&staticConfig, f)
	if err != nil {
		return err
	}
	return nil
}

func SetStaticConfig(config StaticConfig) {
	staticConfig = config
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
