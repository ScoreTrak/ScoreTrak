package util

import (
	"encoding/base64"
	"errors"
	"flag"
	"github.com/ScoreTrak/ScoreTrak/pkg/config"
	"github.com/jinzhu/copier"
	"os"
	"path/filepath"
)

func SetupConfig(f string) config.StaticConfig {
	err := config.NewStaticConfig(f)
	if err != nil {
		panic(err)
	}
	return config.GetStaticConfig()
}

func NewConfigClone(c config.StaticConfig) config.StaticConfig {
	cnf := config.StaticConfig{}
	err := copier.Copy(&cnf, &c)
	if err != nil {
		panic(err)
	}
	return cnf
}

func ConfigFlagParser() (string, error) {
	path := flag.Lookup("config").Value.(flag.Getter).Get().(string)
	encodedConfig := flag.Lookup("encoded-config").Value.(flag.Getter).Get().(string)
	if encodedConfig != "" {
		dec, err := base64.StdEncoding.DecodeString(encodedConfig)
		if err != nil {
			return "", err
		}
		tmpPath := filepath.Join(".", "configs")
		os.MkdirAll(tmpPath, os.ModePerm)
		path = "configs/config-encoded.yml"
		f, err := os.Create(path)
		if err != nil {
			return "", err
		}
		defer f.Close()
		_, err = f.Write(dec)
		if err != nil {
			return "", err
		}
		err = f.Sync()
		if err != nil {
			return "", err
		}
	} else if !ConfigExists(path) {
		err := CreateFile(path)
		if err != nil {
			return "", errors.New("unable to create the config file")
		}
	}
	return path, nil
}

func ConfigExists(f string) bool {
	file, err := os.Stat(f)
	if os.IsNotExist(err) {
		return false
	}
	return !file.IsDir()
}

func CreateFile(name string) error {
	file, err := os.OpenFile(name, os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	return file.Close()
}
