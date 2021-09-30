package platform

import (
	"context"
	"errors"

	"github.com/ScoreTrak/ScoreTrak/pkg/platform/docker"
	"github.com/ScoreTrak/ScoreTrak/pkg/platform/kubernetes"
	"github.com/ScoreTrak/ScoreTrak/pkg/platform/platforming"
	"github.com/ScoreTrak/ScoreTrak/pkg/platform/worker"
)

//Platform is an interface that allows ScoreTrak to deploy/remove the worker containers on a given environment like docker, docker swarm, or kubernetes.
type Platform interface {
	DeployWorkers(ctx context.Context, info worker.Info) error
	RemoveWorkers(ctx context.Context, info worker.Info) error
}

var ErrInvalidPlatform = errors.New("invalid platform specified")

func NewPlatform(config platforming.Config) (Platform, error) {
	if config.Use == "docker" || config.Use == "swarm" {
		return docker.NewDocker(config)
	} else if config.Use == "kubernetes" {
		return kubernetes.NewKubernetes(config)
	} else if config.Use == "none" {
		return nil, nil
	}
	return nil, ErrInvalidPlatform
}
