package round

import "context"

type Serv interface {
	GetLastNonElapsingRound(ctx context.Context) (*Round, error)
	GetAll(ctx context.Context) ([]*Round, error)
	GetByID(ctx context.Context, id uint) (*Round, error)
	GetLastRound(ctx context.Context) (*Round, error)
}

type roundServ struct {
	repo Repo
}

func NewRoundServ(repo Repo) Serv {
	return &roundServ{
		repo: repo,
	}
}

func (svc *roundServ) GetLastNonElapsingRound(ctx context.Context) (*Round, error) {
	return svc.repo.GetLastNonElapsingRound(ctx)
}

func (svc *roundServ) GetLastElapsingRound(ctx context.Context) (*Round, error) {
	return svc.repo.GetLastElapsingRound(ctx)
}

func (svc *roundServ) GetLastRound(ctx context.Context) (*Round, error) {
	return svc.repo.GetLastRound(ctx)
}

func (svc *roundServ) Delete(ctx context.Context, id uint) error { return svc.repo.Delete(ctx, id) }

func (svc *roundServ) GetAll(ctx context.Context) ([]*Round, error) { return svc.repo.GetAll(ctx) }

func (svc *roundServ) GetByID(ctx context.Context, id uint) (*Round, error) {
	return svc.repo.GetByID(ctx, id)
}

func (svc *roundServ) Store(ctx context.Context, u *Round) error { return svc.repo.Store(ctx, u) }

func (svc *roundServ) Update(ctx context.Context, u *Round) error { return svc.repo.Update(ctx, u) }
