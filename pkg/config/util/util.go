package util

import (
	"fmt"
	"github.com/ScoreTrak/ScoreTrak/pkg/config"
	"github.com/spf13/viper"
	"os"
)

func LoadViperConfig(path string) (config config.StaticConfig, err error) {
	viper.SetConfigFile(path)

	if err := viper.ReadInConfig(); err == nil {
		_, err := fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
		if err != nil {
			return config, err
		}
	}

	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}

	return config, nil
}
