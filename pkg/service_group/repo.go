package service_group

import "github.com/gofrs/uuid"

type Repo interface {
	Delete(id uuid.UUID) error
	GetAll() ([]*ServiceGroup, error)
	GetByID(id uuid.UUID) (*ServiceGroup, error)
	Store(u *ServiceGroup) error
	Upsert(u *ServiceGroup) error
	Update(u *ServiceGroup) error
}
