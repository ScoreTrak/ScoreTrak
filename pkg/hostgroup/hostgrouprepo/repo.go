package hostgrouprepo

import (
	"context"

	"github.com/ScoreTrak/ScoreTrak/pkg/hostgroup"
	"github.com/gofrs/uuid"
)

type Repo interface {
	Delete(ctx context.Context, id uuid.UUID) error
	GetAll(ctx context.Context) ([]*hostgroup.HostGroup, error)
	GetByID(ctx context.Context, id uuid.UUID) (*hostgroup.HostGroup, error)
	Store(ctx context.Context, u []*hostgroup.HostGroup) error
	Upsert(ctx context.Context, u []*hostgroup.HostGroup) error
	Update(ctx context.Context, u *hostgroup.HostGroup) error
	TruncateTable(ctx context.Context) error
}
