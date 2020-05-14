package config

type Serv interface {
	Get() (*Config, error)
	Update(*Config) error
}

type configServ struct {
	repo Repo
}

func NewConfigServ(repo Repo) Serv {
	return &configServ{
		repo: repo,
	}
}

func (svc *configServ) Get() (*Config, error) { return svc.repo.Get() }

func (svc *configServ) Update(cfg *Config) error { return svc.repo.Update(cfg) }
