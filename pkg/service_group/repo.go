package service_group

import (
	"context"
	"github.com/gofrs/uuid"
)

type Repo interface {
	Delete(ctx context.Context, id uuid.UUID) error
	GetAll(ctx context.Context) ([]*ServiceGroup, error)
	GetByID(ctx context.Context, id uuid.UUID) (*ServiceGroup, error)
	Store(ctx context.Context, u *ServiceGroup) error
	Upsert(ctx context.Context, u *ServiceGroup) error
	Update(ctx context.Context, u *ServiceGroup) error
}
