package host

type Repo interface {
	Delete(id uint64) error
	GetAll() ([]*Host, error)
	GetByID(id uint64) (*Host, error)
	Store(u *Host) error
	Update(u *Host) error
}
