package userrepo

import (
	"context"

	"github.com/ScoreTrak/ScoreTrak/pkg/user"
	"github.com/gofrs/uuid"
)

type Repo interface {
	GetAll(ctx context.Context) ([]*user.User, error)
	GetByID(ctx context.Context, id uuid.UUID) (*user.User, error)
	GetByUsername(ctx context.Context, username string) (*user.User, error)
	Store(ctx context.Context, u []*user.User) error
	Update(ctx context.Context, u *user.User) error
	Upsert(ctx context.Context, u []*user.User) error
	Delete(ctx context.Context, id uuid.UUID) error
}
