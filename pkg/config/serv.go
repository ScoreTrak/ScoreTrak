package config

import "context"

type Serv interface {
	Get(ctx context.Context) (*DynamicConfig, error)
	Update(context.Context, *DynamicConfig) error
}

type configServ struct {
	repo Repo
}

func NewConfigServ(repo Repo) Serv {
	return &configServ{
		repo: repo,
	}
}

func (svc *configServ) Get(ctx context.Context) (*DynamicConfig, error) { return svc.repo.Get(ctx) }

func (svc *configServ) Update(ctx context.Context, cfg *DynamicConfig) error {
	return svc.repo.Update(ctx, cfg)
}

type StaticServ interface {
	Get() (*StaticConfig, error)
}

type configStaticServ struct{}

func NewStaticConfigServ() StaticServ {
	return &configStaticServ{}
}

func (svc *configStaticServ) Get() (*StaticConfig, error) { sc := GetStaticConfig(); return &sc, nil }
