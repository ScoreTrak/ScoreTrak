package service

import (
	"context"
	"github.com/gofrs/uuid"
)

type Serv interface {
	Delete(ctx context.Context, id uuid.UUID) error
	GetAll(ctx context.Context) ([]*Service, error)
	GetByID(ctx context.Context, id uuid.UUID) (*Service, error)
	Store(ctx context.Context, u []*Service) error
	Update(ctx context.Context, u *Service) error
}

type serviceServ struct {
	repo Repo
}

func NewServiceServ(repo Repo) Serv {
	return &serviceServ{
		repo: repo,
	}
}

func (svc *serviceServ) Delete(ctx context.Context, id uuid.UUID) error {
	return svc.repo.Delete(ctx, id)
}

func (svc *serviceServ) GetAll(ctx context.Context) ([]*Service, error) { return svc.repo.GetAll(ctx) }

func (svc *serviceServ) GetByID(ctx context.Context, id uuid.UUID) (*Service, error) {
	return svc.repo.GetByID(ctx, id)
}

func (svc *serviceServ) Store(ctx context.Context, u []*Service) error { return svc.repo.Store(ctx, u) }

func (svc *serviceServ) Update(ctx context.Context, u *Service) error { return svc.repo.Update(ctx, u) }
