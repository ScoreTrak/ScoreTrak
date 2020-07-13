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

type StaticServ interface {
	Get() (*StaticConfig, error)
}

type configStaticServ struct{}

func NewStaticConfigServ() StaticServ {
	return &configStaticServ{}
}

func (svc *configStaticServ) Get() (*StaticConfig, error) { sc := GetStaticConfig(); return &sc, nil }
