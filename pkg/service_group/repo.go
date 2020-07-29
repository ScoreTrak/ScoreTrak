package service_group

type Repo interface {
	Delete(id uint32) error
	GetAll() ([]*ServiceGroup, error)
	GetByID(id uint32) (*ServiceGroup, error)
	Store(u *ServiceGroup) error
	Update(u *ServiceGroup) error
}
