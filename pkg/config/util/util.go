package util

import (
	"github.com/ScoreTrak/ScoreTrak/pkg/config"
	"github.com/jinzhu/copier"
)

func SetupConfig(f string) config.StaticConfig {
	err := config.NewStaticConfig(f)
	if err != nil {
		panic(err)
	}
	return config.GetStaticConfig()
}

func NewTestConfigClone(testConfPath string) config.StaticConfig {
	return NewConfigClone(SetupConfig(testConfPath))
}

func NewConfigClone(c config.StaticConfig) config.StaticConfig {
	cnf := config.StaticConfig{}
	if err := copier.Copy(&cnf, &c); err != nil {
		panic(err)
	}
	return cnf
}
