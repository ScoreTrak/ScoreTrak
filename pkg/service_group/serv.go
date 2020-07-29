package service_group

type Serv interface {
	Delete(id uint64) error
	GetAll() ([]*ServiceGroup, error)
	GetByID(id uint64) (*ServiceGroup, error)
	Store(u *ServiceGroup) error
	Update(u *ServiceGroup) error
}

type serviceGroupServ struct {
	repo Repo
}

func NewServiceGroupServ(repo Repo) Serv {
	return &serviceGroupServ{
		repo: repo,
	}
}

func (svc *serviceGroupServ) Delete(id uint64) error { return svc.repo.Delete(id) }

func (svc *serviceGroupServ) GetAll() ([]*ServiceGroup, error) { return svc.repo.GetAll() }

func (svc *serviceGroupServ) GetByID(id uint64) (*ServiceGroup, error) { return svc.repo.GetByID(id) }

func (svc *serviceGroupServ) Store(u *ServiceGroup) error { return svc.repo.Store(u) }

func (svc *serviceGroupServ) Update(u *ServiceGroup) error { return svc.repo.Update(u) }
