package hostservice

import (
	"context"

	"github.com/ScoreTrak/ScoreTrak/pkg/host"
	repo2 "github.com/ScoreTrak/ScoreTrak/pkg/host/hostrepo"
	"github.com/gofrs/uuid"
)

type Serv interface {
	Delete(ctx context.Context, id uuid.UUID) error
	GetAll(ctx context.Context) ([]*host.Host, error)
	GetByID(ctx context.Context, id uuid.UUID) (*host.Host, error)
	Store(ctx context.Context, u []*host.Host) error
	Update(ctx context.Context, u *host.Host) error
}

type hostServ struct {
	repo repo2.Repo
}

func NewHostServ(repo repo2.Repo) Serv {
	return &hostServ{
		repo: repo,
	}
}

func (svc *hostServ) Delete(ctx context.Context, id uuid.UUID) error { return svc.repo.Delete(ctx, id) }

func (svc *hostServ) GetAll(ctx context.Context) ([]*host.Host, error) { return svc.repo.GetAll(ctx) }

func (svc *hostServ) GetByID(ctx context.Context, id uuid.UUID) (*host.Host, error) {
	return svc.repo.GetByID(ctx, id)
}

func (svc *hostServ) Store(ctx context.Context, u []*host.Host) error { return svc.repo.Store(ctx, u) }

func (svc *hostServ) Update(ctx context.Context, u *host.Host) error { return svc.repo.Update(ctx, u) }
