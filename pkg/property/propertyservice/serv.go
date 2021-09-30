package propertyservice

import (
	"context"

	"github.com/ScoreTrak/ScoreTrak/pkg/property"
	repo2 "github.com/ScoreTrak/ScoreTrak/pkg/property/propertyrepo"
	"github.com/gofrs/uuid"
)

type Serv interface {
	Delete(ctx context.Context, serviceID uuid.UUID, key string) error
	GetAll(ctx context.Context) ([]*property.Property, error)
	Store(ctx context.Context, u []*property.Property) error
	Update(ctx context.Context, u *property.Property) error
	GetByServiceIDKey(ctx context.Context, serviceID uuid.UUID, key string) (*property.Property, error)
	GetAllByServiceID(ctx context.Context, serviceID uuid.UUID) ([]*property.Property, error)
}

type propertyServ struct {
	repo repo2.Repo
}

func NewPropertyServ(repo repo2.Repo) Serv {
	return &propertyServ{
		repo: repo,
	}
}

func (svc *propertyServ) Delete(ctx context.Context, serviceID uuid.UUID, key string) error {
	return svc.repo.Delete(ctx, serviceID, key)
}

func (svc *propertyServ) GetAll(ctx context.Context) ([]*property.Property, error) {
	return svc.repo.GetAll(ctx)
}

func (svc *propertyServ) Store(ctx context.Context, u []*property.Property) error {
	return svc.repo.Store(ctx, u)
}

func (svc *propertyServ) Update(ctx context.Context, u *property.Property) error {
	return svc.repo.Update(ctx, u)
}

func (svc *propertyServ) GetAllByServiceID(ctx context.Context, serviceID uuid.UUID) ([]*property.Property, error) {
	return svc.repo.GetAllByServiceID(ctx, serviceID)
}

func (svc *propertyServ) GetByServiceIDKey(ctx context.Context, id uuid.UUID, key string) (*property.Property, error) {
	return svc.repo.GetByServiceIDKey(ctx, id, key)
}
