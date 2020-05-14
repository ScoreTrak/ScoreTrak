package swarm

type Serv interface {
	Delete(id uint64) error
	GetAll() ([]*Swarm, error)
	GetByID(id uint64) (*Swarm, error)
	Store(u *Swarm) error
	Update(u *Swarm) error
}

type swarmServ struct {
	repo Repo
}

func NewSwarmServ(repo Repo) Serv {
	return &swarmServ{
		repo: repo,
	}
}

func (svc *swarmServ) Delete(id uint64) error { return svc.repo.Delete(id) }

func (svc *swarmServ) GetAll() ([]*Swarm, error) { return svc.repo.GetAll() }

func (svc *swarmServ) GetByID(id uint64) (*Swarm, error) { return svc.repo.GetByID(id) }

func (svc *swarmServ) Store(u *Swarm) error { return svc.repo.Store(u) }

func (svc *swarmServ) Update(u *Swarm) error { return svc.repo.Update(u) }
