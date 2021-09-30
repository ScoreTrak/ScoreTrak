package team_repo

import (
	"context"

	"github.com/ScoreTrak/ScoreTrak/pkg/team"
	"github.com/gofrs/uuid"
)

type Repo interface {
	Delete(ctx context.Context, id uuid.UUID) error
	GetAll(ctx context.Context) ([]*team.Team, error)
	GetByID(ctx context.Context, id uuid.UUID) (*team.Team, error)
	GetByName(ctx context.Context, name string) (*team.Team, error)
	DeleteByName(ctx context.Context, name string) error
	Store(ctx context.Context, u []*team.Team) error
	Upsert(ctx context.Context, u []*team.Team) error
	UpdateByName(ctx context.Context, u *team.Team) error
	Update(ctx context.Context, u *team.Team) error
}
