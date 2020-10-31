package host_group

import (
	"context"
	"github.com/gofrs/uuid"
)

type Repo interface {
	Delete(ctx context.Context, id uuid.UUID) error
	GetAll(ctx context.Context) ([]*HostGroup, error)
	GetByID(ctx context.Context, id uuid.UUID) (*HostGroup, error)
	Store(ctx context.Context, u []*HostGroup) error
	Upsert(ctx context.Context, u []*HostGroup) error
	Update(ctx context.Context, u *HostGroup) error
	TruncateTable(ctx context.Context) error
}
