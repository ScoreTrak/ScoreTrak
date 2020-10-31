package team

import (
	"context"
	"github.com/gofrs/uuid"
)

type Repo interface {
	Delete(ctx context.Context, id uuid.UUID) error
	GetAll(ctx context.Context) ([]*Team, error)
	GetByID(ctx context.Context, id uuid.UUID) (*Team, error)
	GetByName(ctx context.Context, name string) (*Team, error)
	DeleteByName(ctx context.Context, name string) error
	Store(ctx context.Context, u []*Team) error
	Upsert(ctx context.Context, u []*Team) error
	UpdateByName(ctx context.Context, u *Team) error
	Update(ctx context.Context, u *Team) error
}
