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
	DatabaseMaxTimeDriftSeconds uint
	// How frequently to pull dynamic configs
	DynamicConfigPullSeconds uint

	Queue queueing.Config

	Platform platforming.Config

	PubSubConfig queueing.MasterConfig

	AdminUsername string

	AdminPassword string

	Port string

	Prod bool

	CertFile string

	KeyFile string

	JWT auth.Config
}

var staticConfig StaticConfig

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
