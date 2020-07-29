package host_group

type Repo interface {
	Delete(id uint64) error
	GetAll() ([]*HostGroup, error)
	GetByID(id uint64) (*HostGroup, error)
	Store(u *HostGroup) error
	Update(u *HostGroup) error
}
