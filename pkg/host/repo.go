package host

type Repo interface {
	Delete(id uint32) error
	GetAll() ([]*Host, error)
	GetByID(id uint32) (*Host, error)
	Store(u *Host) error
	Update(u *Host) error
}
