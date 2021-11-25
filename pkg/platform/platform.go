package platform

import (
	"context"
	"errors"
	"github.com/ScoreTrak/ScoreTrak/pkg/platform/docker"
	"github.com/ScoreTrak/ScoreTrak/pkg/platform/kubernetes"
	"github.com/ScoreTrak/ScoreTrak/pkg/platform/platforming"
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

func NewPlatform(mainCfg string, platformCfg platforming.Config) (Platform, error) {
	switch platformCfg.Use {
	case Docker, Swarm:
		return docker.NewDocker(mainCfg, platformCfg)
	case Kubernetes:
		return kubernetes.NewKubernetes(mainCfg, platformCfg)
	case None:
		return nil, nil
	default:
		return nil, ErrInvalidPlatform
	}
}
