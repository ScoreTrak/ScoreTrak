package userservice

import (
	"context"

	"github.com/ScoreTrak/ScoreTrak/pkg/user"
	repo2 "github.com/ScoreTrak/ScoreTrak/pkg/user/userrepo"
	"github.com/gofrs/uuid"
)

type Serv interface {
	Delete(ctx context.Context, id uuid.UUID) error
	GetAll(ctx context.Context) ([]*user.User, error)
	GetByID(ctx context.Context, id uuid.UUID) (*user.User, error)
	GetByUsername(ctx context.Context, username string) (*user.User, error)
	Store(ctx context.Context, u []*user.User) error
	Update(ctx context.Context, u *user.User) error
}

type userServ struct {
	repo repo2.Repo
}

func NewUserServ(repo repo2.Repo) Serv {
	return &userServ{
		repo: repo,
	}
}

func (svc *userServ) Delete(ctx context.Context, id uuid.UUID) error { return svc.repo.Delete(ctx, id) }

func (svc *userServ) GetAll(ctx context.Context) ([]*user.User, error) { return svc.repo.GetAll(ctx) }

func (svc *userServ) GetByID(ctx context.Context, id uuid.UUID) (*user.User, error) {
	return svc.repo.GetByID(ctx, id)
}

func (svc *userServ) GetByUsername(ctx context.Context, username string) (*user.User, error) {
	return svc.repo.GetByUsername(ctx, username)
}

func (svc *userServ) Store(ctx context.Context, u []*user.User) error { return svc.repo.Store(ctx, u) }

func (svc *userServ) Update(ctx context.Context, u *user.User) error { return svc.repo.Update(ctx, u) }
