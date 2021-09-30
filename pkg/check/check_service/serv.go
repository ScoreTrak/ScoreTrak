package check_service

import (
	"context"

	"github.com/ScoreTrak/ScoreTrak/pkg/check"
	repo2 "github.com/ScoreTrak/ScoreTrak/pkg/check/check_repo"
	"github.com/gofrs/uuid"
)

type Serv interface {
	GetAllByRoundID(ctx context.Context, roundID uint64) ([]*check.Check, error)
	GetByRoundServiceID(ctx context.Context, roundID uint64, serviceID uuid.UUID) (*check.Check, error)
	GetAllByServiceID(ctx context.Context, serviceID uuid.UUID) ([]*check.Check, error)
}

type checkServ struct {
	repo repo2.Repo
}

func NewCheckServ(repo repo2.Repo) Serv {
	return &checkServ{
		repo: repo,
	}
}

func (svc *checkServ) GetAllByRoundID(ctx context.Context, roundID uint64) ([]*check.Check, error) {
	return svc.repo.GetAllByRoundID(ctx, roundID)
}
func (svc *checkServ) GetByRoundServiceID(ctx context.Context, roundID uint64, serviceID uuid.UUID) (*check.Check, error) {
	return svc.repo.GetByRoundServiceID(ctx, roundID, serviceID)
}
func (svc *checkServ) GetAllByServiceID(ctx context.Context, serviceID uuid.UUID) ([]*check.Check, error) {
	return svc.repo.GetAllByServiceID(ctx, serviceID)
}

func (svc *checkServ) Delete(ctx context.Context, roundID uint64, serviceID uuid.UUID) error {
	return svc.repo.Delete(ctx, roundID, serviceID)
}

func (svc *checkServ) GetAll(ctx context.Context) ([]*check.Check, error) {
	return svc.repo.GetAll(ctx)
}

func (svc *checkServ) Store(ctx context.Context, c []*check.Check) error {
	return svc.repo.Store(ctx, c)
}
