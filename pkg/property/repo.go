package property

import "github.com/gofrs/uuid"

type Repo interface {
	Delete(id uuid.UUID) error
	GetAll() ([]*Property, error)
	GetByID(id uuid.UUID) (*Property, error)
	Store(u []*Property) error
	Update(u *Property) error
	GetAllByServiceID(id uuid.UUID) ([]*Property, error)
}
