package repo

import (
	"context"
	"github.com/ScoreTrak/ScoreTrak/pkg/check"
	"github.com/gofrs/uuid"
)

type Repo interface {
	GetAllByRoundID(ctx context.Context, roundID uint64) ([]*check.Check, error)
	GetAllByServiceID(ctx context.Context, serviceID uuid.UUID) ([]*check.Check, error)
	GetByRoundServiceID(ctx context.Context, roundID uint64, serviceID uuid.UUID) (*check.Check, error)
	Delete(ctx context.Context, roundID uint64, serviceID uuid.UUID) error
	GetAll(ctx context.Context) ([]*check.Check, error)
	Store(ctx context.Context, u []*check.Check) error
	Upsert(ctx context.Context, u []*check.Check) error
	TruncateTable(ctx context.Context) error
}
