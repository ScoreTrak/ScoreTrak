package service

type Serv interface {
	Delete(id uint32) error
	GetAll() ([]*Service, error)
	GetByID(id uint32) (*Service, error)
	Store(u *Service) error
	Update(u *Service) error
}

type serviceServ struct {
	repo Repo
}

func NewServiceServ(repo Repo) Serv {
	return &serviceServ{
		repo: repo,
	}
}

func (svc *serviceServ) Delete(id uint32) error { return svc.repo.Delete(id) }

func (svc *serviceServ) GetAll() ([]*Service, error) { return svc.repo.GetAll() }

func (svc *serviceServ) GetByID(id uint32) (*Service, error) { return svc.repo.GetByID(id) }

func (svc *serviceServ) Store(u *Service) error { return svc.repo.Store(u) }

func (svc *serviceServ) Update(u *Service) error { return svc.repo.Update(u) }
