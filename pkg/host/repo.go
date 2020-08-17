package host

import "github.com/gofrs/uuid"

type Repo interface {
	Delete(id uuid.UUID) error
	GetAll() ([]*Host, error)
	GetByID(id uuid.UUID) (*Host, error)
	Store(u []*Host) error
	Upsert(u []*Host) error
	Update(u *Host) error
}
