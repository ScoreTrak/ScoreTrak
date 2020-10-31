package service

import (
	"context"
	"github.com/gofrs/uuid"
)

type Repo interface {
	Delete(ctx context.Context, id uuid.UUID) error
	GetAll(ctx context.Context) ([]*Service, error)
	GetByID(ctx context.Context, id uuid.UUID) (*Service, error)
	Store(ctx context.Context, u []*Service) error
	Upsert(ctx context.Context, u []*Service) error
	Update(ctx context.Context, u *Service) error
	TruncateTable(ctx context.Context) error
}
