package platform

import (
	"errors"
	"github.com/L1ghtman2k/ScoreTrak/pkg/logger"
	"github.com/L1ghtman2k/ScoreTrak/pkg/platform/docker"
	"github.com/L1ghtman2k/ScoreTrak/pkg/platform/kubernetes"
	"github.com/L1ghtman2k/ScoreTrak/pkg/platform/platforming"
	"github.com/L1ghtman2k/ScoreTrak/pkg/platform/worker"
)

type Platform interface {
	DeployWorkers(info worker.Info) error
	RemoveWorkers(info worker.Info) error
}

func NewPlatform(config platforming.Config, l logger.LogInfoFormat) (Platform, error) {
	if config.Use == "docker" || config.Use == "swarm" {
		return docker.NewDocker(config, l)
	} else if config.Use == "kubernetes" {
		return kubernetes.NewKubernetes(config, l)
	} else if config.Use == "none" {
		return nil, nil
	}
	return nil, errors.New("invalid platform specified")
}
