package round

import "context"

type Repo interface {
	Delete(ctx context.Context, id uint) error
	GetAll(ctx context.Context) ([]*Round, error)
	GetByID(ctx context.Context, id uint) (*Round, error)
	Store(ctx context.Context, u *Round) error
	Upsert(ctx context.Context, u []*Round) error
	StoreMany(ctx context.Context, u []*Round) error
	Update(ctx context.Context, u *Round) error
	GetLastRound(ctx context.Context) (*Round, error)
	GetLastNonElapsingRound(ctx context.Context) (*Round, error)
	GetLastElapsingRound(ctx context.Context) (*Round, error)
	TruncateTable(ctx context.Context) error
}
