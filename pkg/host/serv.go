package host

import (
	"context"
	"github.com/gofrs/uuid"
)

type Serv interface {
	Delete(ctx context.Context, id uuid.UUID) error
	GetAll(ctx context.Context) ([]*Host, error)
	GetByID(ctx context.Context, id uuid.UUID) (*Host, error)
	Store(ctx context.Context, u []*Host) error
	Update(ctx context.Context, u *Host) error
}

type hostServ struct {
	repo Repo
}

func NewHostServ(repo Repo) Serv {
	return &hostServ{
		repo: repo,
	}
}

func (svc *hostServ) Delete(ctx context.Context, id uuid.UUID) error { return svc.repo.Delete(ctx, id) }

func (svc *hostServ) GetAll(ctx context.Context) ([]*Host, error) { return svc.repo.GetAll(ctx) }

func (svc *hostServ) GetByID(ctx context.Context, id uuid.UUID) (*Host, error) {
	return svc.repo.GetByID(ctx, id)
}

func (svc *hostServ) Store(ctx context.Context, u []*Host) error { return svc.repo.Store(ctx, u) }

func (svc *hostServ) Update(ctx context.Context, u *Host) error { return svc.repo.Update(ctx, u) }
