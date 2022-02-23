package util

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"github.com/ScoreTrak/ScoreTrak/pkg/config"
	"github.com/ScoreTrak/ScoreTrak/pkg/platform/worker"
	"github.com/jinzhu/copier"
)

var ErrQueueNotSupported = errors.New("selected queue is not yet supported by platform")

func GenerateWorkerCfg(originalCfg config.StaticConfig, info worker.Info) (workerCfg config.StaticConfig, err error) {
	workerCfg = config.StaticConfig{}
	err = copier.Copy(&workerCfg, &originalCfg)
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

func EncodeCfg(config config.StaticConfig) (string, error) {
	out, err := json.Marshal(config)
	if err != nil {
		return "", err
	}
	encodedCfg := base64.StdEncoding.EncodeToString(out)
	return encodedCfg, nil
}

func GenerateEncodedWorkerCfg(originalCfg config.StaticConfig, info worker.Info) (string, error) {
	workerCfg, err := GenerateWorkerCfg(originalCfg, info)
	if err != nil {
		return "", err
	}

	encodedWorkerCfg, err := EncodeCfg(workerCfg)
	if err != nil {
		return "", err
	}
	return encodedWorkerCfg, nil
}
