package service_group

import (
	"context"
	"github.com/gofrs/uuid"
)

type Serv interface {
	Delete(ctx context.Context, id uuid.UUID) error
	GetAll(ctx context.Context) ([]*ServiceGroup, error)
	GetByID(ctx context.Context, id uuid.UUID) (*ServiceGroup, error)
	Store(ctx context.Context, u *ServiceGroup) error
	Update(ctx context.Context, u *ServiceGroup) error
}

type serviceGroupServ struct {
	repo Repo
}

func NewServiceGroupServ(repo Repo) Serv {
	return &serviceGroupServ{
		repo: repo,
	}
}

func (svc *serviceGroupServ) Delete(ctx context.Context, id uuid.UUID) error {
	return svc.repo.Delete(ctx, id)
}

func (svc *serviceGroupServ) GetAll(ctx context.Context) ([]*ServiceGroup, error) {
	return svc.repo.GetAll(ctx)
}

func (svc *serviceGroupServ) GetByID(ctx context.Context, id uuid.UUID) (*ServiceGroup, error) {
	return svc.repo.GetByID(ctx, id)
}

func (svc *serviceGroupServ) Store(ctx context.Context, u *ServiceGroup) error {
	return svc.repo.Store(ctx, u)
}

func (svc *serviceGroupServ) Update(ctx context.Context, u *ServiceGroup) error {
	return svc.repo.Update(ctx, u)
}
