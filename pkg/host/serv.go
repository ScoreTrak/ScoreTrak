package host

import "github.com/gofrs/uuid"

type Serv interface {
	Delete(id uuid.UUID) error
	GetAll() ([]*Host, error)
	GetByID(id uuid.UUID) (*Host, error)
	Store(u []*Host) error
	Update(u *Host) error
}

type hostServ struct {
	repo Repo
}

func NewHostServ(repo Repo) Serv {
	return &hostServ{
		repo: repo,
	}
}

func (svc *hostServ) Delete(id uuid.UUID) error { return svc.repo.Delete(id) }

func (svc *hostServ) GetAll() ([]*Host, error) { return svc.repo.GetAll() }

func (svc *hostServ) GetByID(id uuid.UUID) (*Host, error) { return svc.repo.GetByID(id) }

func (svc *hostServ) Store(u []*Host) error { return svc.repo.Store(u) }

func (svc *hostServ) Update(u *Host) error { return svc.repo.Update(u) }
