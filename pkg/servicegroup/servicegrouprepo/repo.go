package servicegrouprepo

import (
	"context"

	"github.com/ScoreTrak/ScoreTrak/pkg/servicegroup"
	"github.com/gofrs/uuid"
)

type Repo interface {
	Delete(ctx context.Context, id uuid.UUID) error
	GetAll(ctx context.Context) ([]*servicegroup.ServiceGroup, error)
	GetByID(ctx context.Context, id uuid.UUID) (*servicegroup.ServiceGroup, error)
	Store(ctx context.Context, u *servicegroup.ServiceGroup) error
	Upsert(ctx context.Context, u *servicegroup.ServiceGroup) error
	Update(ctx context.Context, u *servicegroup.ServiceGroup) error
}
