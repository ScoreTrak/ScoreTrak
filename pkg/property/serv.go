package property

import "github.com/gofrs/uuid"

type Serv interface {
	Delete(serviceID uuid.UUID, key string) error
	GetAll() ([]*Property, error)
	Store(u []*Property) error
	Update(u *Property) error
	GetByServiceIDKey(serviceID uuid.UUID, key string) (*Property, error)
	GetAllByServiceID(serviceID uuid.UUID) ([]*Property, error)
}

type propertyServ struct {
	repo Repo
}

func NewPropertyServ(repo Repo) Serv {
	return &propertyServ{
		repo: repo,
	}
}

func (svc *propertyServ) Delete(serviceID uuid.UUID, key string) error {
	return svc.repo.Delete(serviceID, key)
}

func (svc *propertyServ) GetAll() ([]*Property, error) { return svc.repo.GetAll() }

func (svc *propertyServ) Store(u []*Property) error { return svc.repo.Store(u) }

func (svc *propertyServ) Update(u *Property) error { return svc.repo.Update(u) }

func (svc *propertyServ) GetAllByServiceID(serviceID uuid.UUID) ([]*Property, error) {
	return svc.repo.GetAllByServiceID(serviceID)
}

func (svc *propertyServ) GetByServiceIDKey(id uuid.UUID, key string) (*Property, error) {
	return svc.repo.GetByServiceIDKey(id, key)
}
