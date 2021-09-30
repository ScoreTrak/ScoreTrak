package round_service

import (
	"context"

	"github.com/ScoreTrak/ScoreTrak/pkg/round"
	repo2 "github.com/ScoreTrak/ScoreTrak/pkg/round/round_repo"
)

type Serv interface {
	GetLastNonElapsingRound(ctx context.Context) (*round.Round, error)
	GetAll(ctx context.Context) ([]*round.Round, error)
	GetByID(ctx context.Context, id uint64) (*round.Round, error)
	GetLastRound(ctx context.Context) (*round.Round, error)
}

type roundServ struct {
	repo repo2.Repo
}

func NewRoundServ(repo repo2.Repo) Serv {
	return &roundServ{
		repo: repo,
	}
}

func (svc *roundServ) GetLastNonElapsingRound(ctx context.Context) (*round.Round, error) {
	return svc.repo.GetLastNonElapsingRound(ctx)
}

func (svc *roundServ) GetLastElapsingRound(ctx context.Context) (*round.Round, error) {
	return svc.repo.GetLastElapsingRound(ctx)
}

func (svc *roundServ) GetLastRound(ctx context.Context) (*round.Round, error) {
	return svc.repo.GetLastRound(ctx)
}

func (svc *roundServ) Delete(ctx context.Context, id uint64) error { return svc.repo.Delete(ctx, id) }

func (svc *roundServ) GetAll(ctx context.Context) ([]*round.Round, error) {
	return svc.repo.GetAll(ctx)
}

func (svc *roundServ) GetByID(ctx context.Context, id uint64) (*round.Round, error) {
	return svc.repo.GetByID(ctx, id)
}

func (svc *roundServ) Store(ctx context.Context, u *round.Round) error { return svc.repo.Store(ctx, u) }

func (svc *roundServ) Update(ctx context.Context, u *round.Round) error {
	return svc.repo.Update(ctx, u)
}
