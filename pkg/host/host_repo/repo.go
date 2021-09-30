package host_repo

import (
	"context"

	"github.com/ScoreTrak/ScoreTrak/pkg/host"
	"github.com/gofrs/uuid"
)

type Repo interface {
	Delete(ctx context.Context, id uuid.UUID) error
	GetAll(ctx context.Context) ([]*host.Host, error)
	GetByID(ctx context.Context, id uuid.UUID) (*host.Host, error)
	Store(ctx context.Context, u []*host.Host) error
	Upsert(ctx context.Context, u []*host.Host) error
	Update(ctx context.Context, u *host.Host) error
	TruncateTable(ctx context.Context) error
}
