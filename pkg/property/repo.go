package property

type Repo interface {
	Delete(id uint64) error
	GetAll() ([]*Property, error)
	GetByID(id uint64) (*Property, error)
	Store(u *Property) error
	Update(u *Property) error
	GetAllByServiceID(id uint64) ([]*Property, error)
}
