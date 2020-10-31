package property

import (
	"context"
	"github.com/gofrs/uuid"
)

type Serv interface {
	Delete(ctx context.Context, serviceID uuid.UUID, key string) error
	GetAll(ctx context.Context) ([]*Property, error)
	Store(ctx context.Context, u []*Property) error
	Update(ctx context.Context, u *Property) error
	GetByServiceIDKey(ctx context.Context, serviceID uuid.UUID, key string) (*Property, error)
	GetAllByServiceID(ctx context.Context, serviceID uuid.UUID) ([]*Property, error)
}

type propertyServ struct {
	repo Repo
}

func NewPropertyServ(repo Repo) Serv {
	return &propertyServ{
		repo: repo,
	}
}

func (svc *propertyServ) Delete(ctx context.Context, serviceID uuid.UUID, key string) error {
	return svc.repo.Delete(ctx, serviceID, key)
}

func (svc *propertyServ) GetAll(ctx context.Context) ([]*Property, error) {
	return svc.repo.GetAll(ctx)
}

func (svc *propertyServ) Store(ctx context.Context, u []*Property) error {
	return svc.repo.Store(ctx, u)
}

func (svc *propertyServ) Update(ctx context.Context, u *Property) error {
	return svc.repo.Update(ctx, u)
}

func (svc *propertyServ) GetAllByServiceID(ctx context.Context, serviceID uuid.UUID) ([]*Property, error) {
	return svc.repo.GetAllByServiceID(ctx, serviceID)
}

func (svc *propertyServ) GetByServiceIDKey(ctx context.Context, id uuid.UUID, key string) (*Property, error) {
	return svc.repo.GetByServiceIDKey(ctx, id, key)
}
