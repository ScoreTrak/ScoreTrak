package config

type Serv interface {
	Get() (*DynamicConfig, error)
	Update(*DynamicConfig) error
}

type configServ struct {
	repo Repo
}

func NewConfigServ(repo Repo) Serv {
	return &configServ{
		repo: repo,
	}
}

func (svc *configServ) Get() (*DynamicConfig, error) { return svc.repo.Get() }

func (svc *configServ) Update(cfg *DynamicConfig) error { return svc.repo.Update(cfg) }
