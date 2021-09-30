package hostgroupservice

import (
	"context"

	"github.com/ScoreTrak/ScoreTrak/pkg/hostgroup"
	repo2 "github.com/ScoreTrak/ScoreTrak/pkg/hostgroup/hostgrouprepo"
	"github.com/gofrs/uuid"
)

type Serv interface {
	Delete(ctx context.Context, id uuid.UUID) error
	GetAll(ctx context.Context) ([]*hostgroup.HostGroup, error)
	GetByID(ctx context.Context, id uuid.UUID) (*hostgroup.HostGroup, error)
	Store(ctx context.Context, u []*hostgroup.HostGroup) error
	Update(ctx context.Context, u *hostgroup.HostGroup) error
}

type hostGroupServ struct {
	repo repo2.Repo
}

func NewHostGroupServ(repo repo2.Repo) Serv {
	return &hostGroupServ{
		repo: repo,
	}
}

func (svc *hostGroupServ) Delete(ctx context.Context, id uuid.UUID) error {
	return svc.repo.Delete(ctx, id)
}

func (svc *hostGroupServ) GetAll(ctx context.Context) ([]*hostgroup.HostGroup, error) {
	return svc.repo.GetAll(ctx)
}

func (svc *hostGroupServ) GetByID(ctx context.Context, id uuid.UUID) (*hostgroup.HostGroup, error) {
	return svc.repo.GetByID(ctx, id)
}

func (svc *hostGroupServ) Store(ctx context.Context, u []*hostgroup.HostGroup) error {
	return svc.repo.Store(ctx, u)
}

func (svc *hostGroupServ) Update(ctx context.Context, u *hostgroup.HostGroup) error {
	return svc.repo.Update(ctx, u)
}
