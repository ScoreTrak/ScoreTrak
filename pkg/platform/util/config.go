package util

import (
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/ScoreTrak/ScoreTrak/pkg/config"
	"github.com/ScoreTrak/ScoreTrak/pkg/platform/worker"
	"github.com/jinzhu/copier"
)

var ErrQueueNotSupported = errors.New("selected queue is not yet supported by platform")

func GenerateWorkerCfg(originalCfg config.StaticConfig, info worker.Info) (workerCfg config.StaticConfig, err error) {
	workerCfg = config.StaticConfig{}
	err = copier.Copy(&originalCfg, &workerCfg)
	if err != nil {
		return workerCfg, err
	}

	if workerCfg.Queue.Use == "nsq" {
		workerCfg.Queue.NSQ.Topic = info.Topic
	} else {
		return workerCfg, ErrQueueNotSupported
	}

	return workerCfg, nil
}

func EncodeCfg(config config.StaticConfig) string {
	cfgString := fmt.Sprintf("%v", config)
	encodedCfg := base64.StdEncoding.EncodeToString([]byte(cfgString))
	return encodedCfg
}

func GenerateEncodedWorkerCfg(originalCfg config.StaticConfig, info worker.Info) (string, error) {
	workerCfg, err := GenerateWorkerCfg(originalCfg, info)
	if err != nil {
		return "", err
	}

	encodedWorkerCfg := EncodeCfg(workerCfg)
	return encodedWorkerCfg, nil
}
