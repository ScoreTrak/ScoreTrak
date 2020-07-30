package host_group

import "github.com/gofrs/uuid"

type Serv interface {
	Delete(id uuid.UUID) error
	GetAll() ([]*HostGroup, error)
	GetByID(id uuid.UUID) (*HostGroup, error)
	Store(u []*HostGroup) error
	Update(u *HostGroup) error
}

type hostGroupServ struct {
	repo Repo
}

func NewHostGroupServ(repo Repo) Serv {
	return &hostGroupServ{
		repo: repo,
	}
}

func (svc *hostGroupServ) Delete(id uuid.UUID) error { return svc.repo.Delete(id) }

func (svc *hostGroupServ) GetAll() ([]*HostGroup, error) { return svc.repo.GetAll() }

func (svc *hostGroupServ) GetByID(id uuid.UUID) (*HostGroup, error) { return svc.repo.GetByID(id) }

func (svc *hostGroupServ) Store(u []*HostGroup) error { return svc.repo.Store(u) }

func (svc *hostGroupServ) Update(u *HostGroup) error { return svc.repo.Update(u) }
