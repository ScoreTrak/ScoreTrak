package host_group

import (
	"context"
	"github.com/gofrs/uuid"
)

type Serv interface {
	Delete(ctx context.Context, id uuid.UUID) error
	GetAll(ctx context.Context) ([]*HostGroup, error)
	GetByID(ctx context.Context, id uuid.UUID) (*HostGroup, error)
	Store(ctx context.Context, u []*HostGroup) error
	Update(ctx context.Context, u *HostGroup) error
}

type hostGroupServ struct {
	repo Repo
}

func NewHostGroupServ(repo Repo) Serv {
	return &hostGroupServ{
		repo: repo,
	}
}

func (svc *hostGroupServ) Delete(ctx context.Context, id uuid.UUID) error {
	return svc.repo.Delete(ctx, id)
}

func (svc *hostGroupServ) GetAll(ctx context.Context) ([]*HostGroup, error) {
	return svc.repo.GetAll(ctx)
}

func (svc *hostGroupServ) GetByID(ctx context.Context, id uuid.UUID) (*HostGroup, error) {
	return svc.repo.GetByID(ctx, id)
}

func (svc *hostGroupServ) Store(ctx context.Context, u []*HostGroup) error {
	return svc.repo.Store(ctx, u)
}

func (svc *hostGroupServ) Update(ctx context.Context, u *HostGroup) error {
	return svc.repo.Update(ctx, u)
}
