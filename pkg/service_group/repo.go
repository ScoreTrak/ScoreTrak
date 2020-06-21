package service_group

type Repo interface {
	Delete(id uint64) error
	GetAll() ([]*ServiceGroup, error)
	GetByID(id uint64) (*ServiceGroup, error)
	Store(u *ServiceGroup) error
	Update(u *ServiceGroup) error
}
