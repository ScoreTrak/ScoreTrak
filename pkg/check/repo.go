package check

import "github.com/gofrs/uuid"

type Repo interface {
	GetAllByRoundID(roundID uint) ([]*Check, error)
	GetAllByServiceID(serviceID uuid.UUID) ([]*Check, error)
	GetByRoundServiceID(roundID uint, serviceID uuid.UUID) (*Check, error)
	Delete(roundID uint, serviceID uuid.UUID) error
	GetAll() ([]*Check, error)
	Store(u []*Check) error
	Upsert(u []*Check) error
	TruncateTable() error
}
