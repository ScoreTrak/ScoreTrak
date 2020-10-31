package property

import (
	"context"
	"github.com/gofrs/uuid"
)

type Repo interface {
	Delete(ctx context.Context, serviceID uuid.UUID, key string) error
	GetAll(ctx context.Context) ([]*Property, error)
	Store(ctx context.Context, u []*Property) error
	Upsert(ctx context.Context, u []*Property) error
	Update(ctx context.Context, u *Property) error
	GetAllByServiceID(ctx context.Context, id uuid.UUID) ([]*Property, error)
	GetByServiceIDKey(ctx context.Context, id uuid.UUID, key string) (*Property, error)
	TruncateTable(ctx context.Context) error
}
