package repo

import (
	"context"
	"github.com/ScoreTrak/ScoreTrak/pkg/round"
)

type Repo interface {
	Delete(ctx context.Context, id uint64) error
	GetAll(ctx context.Context) ([]*round.Round, error)
	GetByID(ctx context.Context, id uint64) (*round.Round, error)
	Store(ctx context.Context, u *round.Round) error
	Upsert(ctx context.Context, u []*round.Round) error
	StoreMany(ctx context.Context, u []*round.Round) error
	Update(ctx context.Context, u *round.Round) error
	GetLastRound(ctx context.Context) (*round.Round, error)
	GetLastNonElapsingRound(ctx context.Context) (*round.Round, error)
	GetLastElapsingRound(ctx context.Context) (*round.Round, error)
	TruncateTable(ctx context.Context) error
}
