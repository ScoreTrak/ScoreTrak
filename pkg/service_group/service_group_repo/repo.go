package service_group_repo

import (
	"context"

	"github.com/ScoreTrak/ScoreTrak/pkg/service_group"
	"github.com/gofrs/uuid"
)

type Repo interface {
	Delete(ctx context.Context, id uuid.UUID) error
	GetAll(ctx context.Context) ([]*service_group.ServiceGroup, error)
	GetByID(ctx context.Context, id uuid.UUID) (*service_group.ServiceGroup, error)
	Store(ctx context.Context, u *service_group.ServiceGroup) error
	Upsert(ctx context.Context, u *service_group.ServiceGroup) error
	Update(ctx context.Context, u *service_group.ServiceGroup) error
}
