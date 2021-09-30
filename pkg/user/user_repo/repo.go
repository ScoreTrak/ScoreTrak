package user_repo

import (
	"context"

	"github.com/ScoreTrak/ScoreTrak/pkg/user"
	"github.com/gofrs/uuid"
)

type Repo interface {
	Delete(ctx context.Context, id uuid.UUID) error
	GetAll(ctx context.Context) ([]*user.User, error)
	GetByID(ctx context.Context, id uuid.UUID) (*user.User, error)
	GetByUsername(ctx context.Context, username string) (*user.User, error)
	Store(ctx context.Context, u []*user.User) error
	Upsert(ctx context.Context, u []*user.User) error
	Update(ctx context.Context, u *user.User) error
}
