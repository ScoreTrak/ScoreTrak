package util

import (
	"errors"
	"github.com/ScoreTrak/ScoreTrak/pkg/config"
	"github.com/ScoreTrak/ScoreTrak/pkg/platform/worker"
	"github.com/jinzhu/copier"
)

var ErrQueueNotSupported = errors.New("selected queue is not yet supported by platform")

func GenerateWorkerConfig(originalCfg config.StaticConfig, info worker.Info) (workerCfg config.StaticConfig, err error) {
	workerCfg = config.StaticConfig{}
	err = copier.Copy(originalCfg, workerCfg)
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
