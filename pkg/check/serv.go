package check

import (
	"context"
	"github.com/gofrs/uuid"
)

type Serv interface {
	GetAllByRoundID(ctx context.Context, roundID uint) ([]*Check, error)
	GetByRoundServiceID(ctx context.Context, roundID uint, serviceID uuid.UUID) (*Check, error)
	GetAllByServiceID(ctx context.Context, serviceID uuid.UUID) ([]*Check, error)
}

type checkServ struct {
	repo Repo
}

func NewCheckServ(repo Repo) Serv {
	return &checkServ{
		repo: repo,
	}
}

func (svc *checkServ) GetAllByRoundID(ctx context.Context, roundID uint) ([]*Check, error) {
	return svc.repo.GetAllByRoundID(ctx, roundID)
}
func (svc *checkServ) GetByRoundServiceID(ctx context.Context, roundID uint, serviceID uuid.UUID) (*Check, error) {
	return svc.repo.GetByRoundServiceID(ctx, roundID, serviceID)
}
func (svc *checkServ) GetAllByServiceID(ctx context.Context, serviceID uuid.UUID) ([]*Check, error) {
	return svc.repo.GetAllByServiceID(ctx, serviceID)
}

func (svc *checkServ) Delete(ctx context.Context, roundID uint, serviceID uuid.UUID) error {
	return svc.repo.Delete(ctx, roundID, serviceID)
}

func (svc *checkServ) GetAll(ctx context.Context) ([]*Check, error) { return svc.repo.GetAll(ctx) }

func (svc *checkServ) Store(ctx context.Context, c []*Check) error { return svc.repo.Store(ctx, c) }
