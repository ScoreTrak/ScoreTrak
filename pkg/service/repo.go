package service

import "github.com/gofrs/uuid"

type Repo interface {
	Delete(id uuid.UUID) error
	GetAll() ([]*Service, error)
	GetByID(id uuid.UUID) (*Service, error)
	Store(u []*Service) error
	Update(u *Service) error
}
