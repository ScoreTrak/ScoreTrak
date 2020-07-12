package platform

import (
	"errors"
	"github.com/L1ghtman2k/ScoreTrak/pkg/config"
	"github.com/L1ghtman2k/ScoreTrak/pkg/logger"
	"github.com/L1ghtman2k/ScoreTrak/pkg/platform/docker"
	"github.com/L1ghtman2k/ScoreTrak/pkg/platform/kubernetes"
	"github.com/L1ghtman2k/ScoreTrak/pkg/platform/worker"
)

type Platform interface {
	DeployWorkers(info worker.Info) error
	RemoveWorkers(info worker.Info) error
}

func NewPlatform(config config.StaticConfig, l logger.LogInfoFormat) (Platform, error) {
	if config.Platform.Use == "docker" || config.Platform.Use == "swarm" {
		return docker.NewDocker(config, l)
	} else if config.Platform.Use == "kubernetes" {
		return kubernetes.NewKubernetes(config, l)
	} else if config.Platform.Use == "none" {
		return nil, nil
	}
	return nil, errors.New("invalid platform specified")
}
