package host

import (
	"context"
	"github.com/gofrs/uuid"
)

type Repo interface {
	Delete(ctx context.Context, id uuid.UUID) error
	GetAll(ctx context.Context) ([]*Host, error)
	GetByID(ctx context.Context, id uuid.UUID) (*Host, error)
	Store(ctx context.Context, u []*Host) error
	Upsert(ctx context.Context, u []*Host) error
	Update(ctx context.Context, u *Host) error
	TruncateTable(ctx context.Context) error
}
