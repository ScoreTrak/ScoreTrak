package property

type Repo interface {
	Delete(id uint32) error
	GetAll() ([]*Property, error)
	GetByID(id uint32) (*Property, error)
	Store(u *Property) error
	Update(u *Property) error
	GetAllByServiceID(id uint32) ([]*Property, error)
}
