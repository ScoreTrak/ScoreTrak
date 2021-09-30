package servicerepo

import (
	"context"

	"github.com/ScoreTrak/ScoreTrak/pkg/service"
	"github.com/gofrs/uuid"
)

type Repo interface {
	Delete(ctx context.Context, id uuid.UUID) error
	GetAll(ctx context.Context) ([]*service.Service, error)
	GetByID(ctx context.Context, id uuid.UUID) (*service.Service, error)
	Store(ctx context.Context, u []*service.Service) error
	Upsert(ctx context.Context, u []*service.Service) error
	Update(ctx context.Context, u *service.Service) error
	TruncateTable(ctx context.Context) error
}
