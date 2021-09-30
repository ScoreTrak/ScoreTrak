package config_service

import (
	"context"

	"github.com/ScoreTrak/ScoreTrak/pkg/config"
	repo2 "github.com/ScoreTrak/ScoreTrak/pkg/config/config_repo"
)

type Serv interface {
	Get(ctx context.Context) (*config.DynamicConfig, error)
	Update(context.Context, *config.DynamicConfig) error
}

type configServ struct {
	repo repo2.Repo
}

func NewConfigServ(repo repo2.Repo) Serv {
	return &configServ{
		repo: repo,
	}
}

func (svc *configServ) Get(ctx context.Context) (*config.DynamicConfig, error) {
	return svc.repo.Get(ctx)
}

func (svc *configServ) Update(ctx context.Context, cfg *config.DynamicConfig) error {
	return svc.repo.Update(ctx, cfg)
}

type StaticServ interface {
	Get() (*config.StaticConfig, error)
}

type configStaticServ struct{}

func NewStaticConfigServ() StaticServ {
	return &configStaticServ{}
}

func (svc *configStaticServ) Get() (*config.StaticConfig, error) {
	sc := config.GetStaticConfig()
	return &sc, nil
}
