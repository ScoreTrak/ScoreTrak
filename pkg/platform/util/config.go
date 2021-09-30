package util

import (
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/ScoreTrak/ScoreTrak/pkg/config"
	"github.com/ScoreTrak/ScoreTrak/pkg/platform/worker"
)

var QueueNotSupportedError = errors.New("selected queue is not yet supported by platform")

func GenerateConfigFile(info worker.Info) (path string, err error) {
	cnf, err := config.GetConfigCopy()
	if err != nil {
		return "", err
	}
	if cnf.Queue.Use == "nsq" {
		cnf.Queue.NSQ.Topic = info.Topic
	} else {
		return "", QueueNotSupportedError
	}
	tmpPath := filepath.Join(".", "tmp")
	err = os.MkdirAll(tmpPath, os.ModePerm)
	if err != nil {
		return "", err
	}
	path = fmt.Sprintf("tmp/config_worker_%s", info.Topic)
	err = config.SaveConfigToYamlFile(path, cnf)
	if err != nil {
		return "", err
	}
	return path, nil
}

func EncodeConfigFile(configPath string) (string, error) {
	content, err := ioutil.ReadFile(configPath)
	if err != nil {
		return "", err
	}
	cEnc := base64.StdEncoding.EncodeToString(content)
	return cEnc, nil
}
