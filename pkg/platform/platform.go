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

func NewPlatform(cfg config.StaticConfig) (Platform, error) {
	switch cfg.Platform.Use {
	case Docker, Swarm:
		return docker.NewDocker(cfg)
	case Kubernetes:
		return kubernetes.NewKubernetes(cfg)
	case None:
		return nil, nil
	default:
		return nil, ErrInvalidPlatform
	}
}
