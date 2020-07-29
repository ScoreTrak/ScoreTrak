package service

type Repo interface {
	Delete(id uint32) error
	GetAll() ([]*Service, error)
	GetByID(id uint32) (*Service, error)
	Store(u *Service) error
	Update(u *Service) error
}
