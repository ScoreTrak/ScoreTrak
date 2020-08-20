package config

import (
	"github.com/ScoreTrak/ScoreTrak/pkg/logger"
	"github.com/ScoreTrak/ScoreTrak/pkg/platform/platforming"
	"github.com/ScoreTrak/ScoreTrak/pkg/queue/queueing"
	"github.com/ScoreTrak/ScoreTrak/pkg/storage"
	"github.com/jinzhu/configor"
	"github.com/jinzhu/copier"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
)

// StaticConfig is a struct of settings that were set at the start of the application
type StaticConfig struct {
	// token specified on init of the staticConfig
	Token string `default:""`

	DB storage.Config

	Logger logger.Config

	Queue queueing.Config

	Platform platforming.Config

	Port string `default:"33333"`
}

var staticConfig StaticConfig

func GetToken() string {
	return staticConfig.Token
}

func GetLoggerConfig() logger.Config {
	return staticConfig.Logger
}

func GetPlatformConfig() platforming.Config {
	return staticConfig.Platform
}

func GetQueueConfig() queueing.Config {
	return staticConfig.Queue
}

func GetDBConfig() storage.Config {
	return staticConfig.DB
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
	tmpPath := filepath.Join(".", "public")
	os.MkdirAll(tmpPath, os.ModePerm)
	err = ioutil.WriteFile(f, b, 0600)
	if err != nil {
		return err
	}
	return nil
}
