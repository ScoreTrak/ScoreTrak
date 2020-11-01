package repo

import (
	"context"
	"github.com/ScoreTrak/ScoreTrak/pkg/host_group"
	"github.com/gofrs/uuid"
)

type Repo interface {
	Delete(ctx context.Context, id uuid.UUID) error
	GetAll(ctx context.Context) ([]*host_group.HostGroup, error)
	GetByID(ctx context.Context, id uuid.UUID) (*host_group.HostGroup, error)
	Store(ctx context.Context, u []*host_group.HostGroup) error
	Upsert(ctx context.Context, u []*host_group.HostGroup) error
	Update(ctx context.Context, u *host_group.HostGroup) error
	TruncateTable(ctx context.Context) error
}
