package platform

import (
	"context"
	"errors"
	"github.com/ScoreTrak/ScoreTrak/pkg/config"

	"github.com/ScoreTrak/ScoreTrak/pkg/platform/docker"
	"github.com/ScoreTrak/ScoreTrak/pkg/platform/kubernetes"
	"github.com/ScoreTrak/ScoreTrak/pkg/platform/worker"
)

// Platform is an interface that allows ScoreTrak to deploy/remove the worker containers on a given environment like docker, docker swarm, or kubernetes.
type Platform interface {
	DeployWorkers(ctx context.Context, info worker.Info) error
	RemoveWorkers(ctx context.Context, info worker.Info) error
}

var ErrInvalidPlatform = errors.New("invalid platform specified")

const (
	Docker     = "docker"
	Swarm      = "swarm"
	Kubernetes = "kubernetes"
	None       = "none"
)

func NewPlatform(config platforming.Config) (Platform, error) {
	switch config.Use {
	case Docker, Swarm:
		return docker.NewDocker(config)
	case Kubernetes:
		return kubernetes.NewKubernetes(config)
	case None:
		return nil, nil
	default:
		return nil, ErrInvalidPlatform
	}
}
