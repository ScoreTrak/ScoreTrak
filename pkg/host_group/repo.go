package host_group

import "github.com/gofrs/uuid"

type Repo interface {
	Delete(id uuid.UUID) error
	GetAll() ([]*HostGroup, error)
	GetByID(id uuid.UUID) (*HostGroup, error)
	Store(u []*HostGroup) error
	Update(u *HostGroup) error
}
