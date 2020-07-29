package service

type Repo interface {
	Delete(id uint64) error
	GetAll() ([]*Service, error)
	GetByID(id uint64) (*Service, error)
	Store(u *Service) error
	Update(u *Service) error
}
