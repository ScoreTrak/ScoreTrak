package check

import (
	"context"
	"github.com/gofrs/uuid"
)

type Repo interface {
	GetAllByRoundID(ctx context.Context, roundID uint) ([]*Check, error)
	GetAllByServiceID(ctx context.Context, serviceID uuid.UUID) ([]*Check, error)
	GetByRoundServiceID(ctx context.Context, roundID uint, serviceID uuid.UUID) (*Check, error)
	Delete(ctx context.Context, roundID uint, serviceID uuid.UUID) error
	GetAll(ctx context.Context) ([]*Check, error)
	Store(ctx context.Context, u []*Check) error
	Upsert(ctx context.Context, u []*Check) error
	TruncateTable(ctx context.Context) error
}
