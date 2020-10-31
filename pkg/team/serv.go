package team

import (
	"context"
	"github.com/gofrs/uuid"
)

type Serv interface {
	GetAll(ctx context.Context) ([]*Team, error)
	GetByID(ctx context.Context, id uuid.UUID) (*Team, error)
	Delete(ctx context.Context, id uuid.UUID) error
	Store(ctx context.Context, u []*Team) error
	Update(ctx context.Context, u *Team) error
}

type teamServ struct {
	repo Repo
}

func NewTeamServ(repo Repo) Serv {
	return &teamServ{
		repo: repo,
	}
}

func (svc *teamServ) DeleteByName(ctx context.Context, name string) error {
	return svc.repo.DeleteByName(ctx, name)
}
func (svc *teamServ) Delete(ctx context.Context, id uuid.UUID) error { return svc.repo.Delete(ctx, id) }
func (svc *teamServ) GetAll(ctx context.Context) ([]*Team, error)    { return svc.repo.GetAll(ctx) }
func (svc *teamServ) GetByName(ctx context.Context, name string) (*Team, error) {
	return svc.repo.GetByName(ctx, name)
}
func (svc *teamServ) GetByID(ctx context.Context, id uuid.UUID) (*Team, error) {
	return svc.repo.GetByID(ctx, id)
}
func (svc *teamServ) Store(ctx context.Context, u []*Team) error { return svc.repo.Store(ctx, u) }
func (svc *teamServ) UpdateByName(ctx context.Context, u *Team) error {
	return svc.repo.UpdateByName(ctx, u)
}
func (svc *teamServ) Update(ctx context.Context, u *Team) error { return svc.repo.Update(ctx, u) }
