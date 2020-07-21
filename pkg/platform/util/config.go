package util

import (
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/L1ghtman2k/ScoreTrak/pkg/config"
	"github.com/L1ghtman2k/ScoreTrak/pkg/platform/worker"
	"io/ioutil"
)

func GenerateConfigFile(info worker.Info) (path string, err error) {
	cnf, err := config.GetConfigCopy()
	if err != nil {
		return "", err
	}
	if cnf.Queue.Use == "nsq" {
		cnf.Queue.NSQ.Topic = info.Topic
	} else {
		return "", errors.New("selected queue is not yet supported with platform Docker")
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