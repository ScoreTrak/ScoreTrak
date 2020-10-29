package property

import "github.com/gofrs/uuid"

type Repo interface {
	Delete(id uuid.UUID, key string) error
	GetAll() ([]*Property, error)
	Store(u []*Property) error
	Upsert(u []*Property) error
	Update(u *Property) error
	GetAllByServiceID(id uuid.UUID) ([]*Property, error)
	GetByServiceIDKey(id uuid.UUID, key string) (*Property, error)
	TruncateTable() error
}
