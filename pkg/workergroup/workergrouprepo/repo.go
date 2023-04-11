package workergrouprepo

import (
	"context"

	"github.com/ScoreTrak/ScoreTrak/pkg/workergroup"
	"github.com/gofrs/uuid"
)

type Repo interface {
	Delete(ctx context.Context, id uuid.UUID) error
	GetAll(ctx context.Context) ([]*workergroup.WorkerGroup, error)
	GetByID(ctx context.Context, id uuid.UUID) (*workergroup.WorkerGroup, error)
	Store(ctx context.Context, u *workergroup.WorkerGroup) error
	Upsert(ctx context.Context, u *workergroup.WorkerGroup) error
	Update(ctx context.Context, u *workergroup.WorkerGroup) error
}
