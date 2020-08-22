package property

import "github.com/gofrs/uuid"

type Serv interface {
	Delete(id uuid.UUID, key string) error
	GetAll() ([]*Property, error)
	Store(u []*Property) error
	Update(u *Property) error
	GetByServiceIDKey(id uuid.UUID, key string) (*Property, error)
	GetAllByServiceID(id uuid.UUID) ([]*Property, error)
}

type propertyServ struct {
	repo Repo
}

func NewPropertyServ(repo Repo) Serv {
	return &propertyServ{
		repo: repo,
	}
}

func (svc *propertyServ) Delete(id uuid.UUID, key string) error { return svc.repo.Delete(id, key) }

func (svc *propertyServ) GetAll() ([]*Property, error) { return svc.repo.GetAll() }

func (svc *propertyServ) Store(u []*Property) error { return svc.repo.Store(u) }

func (svc *propertyServ) Update(u *Property) error { return svc.repo.Update(u) }

func (svc *propertyServ) GetAllByServiceID(id uuid.UUID) ([]*Property, error) {
	return svc.repo.GetAllByServiceID(id)
}

func (svc *propertyServ) GetByServiceIDKey(id uuid.UUID, key string) (*Property, error) {
	return svc.repo.GetByServiceIDKey(id, key)
}
