package property_repo

import (
	"context"
	"github.com/ScoreTrak/ScoreTrak/pkg/property"
	"github.com/gofrs/uuid"
)

type Repo interface {
	Delete(ctx context.Context, serviceID uuid.UUID, key string) error
	GetAll(ctx context.Context) ([]*property.Property, error)
	Store(ctx context.Context, u []*property.Property) error
	Upsert(ctx context.Context, u []*property.Property) error
	Update(ctx context.Context, u *property.Property) error
	GetAllByServiceID(ctx context.Context, id uuid.UUID) ([]*property.Property, error)
	GetByServiceIDKey(ctx context.Context, id uuid.UUID, key string) (*property.Property, error)
	TruncateTable(ctx context.Context) error
}
