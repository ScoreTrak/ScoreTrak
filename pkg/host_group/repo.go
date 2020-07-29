package host_group

type Repo interface {
	Delete(id uint32) error
	GetAll() ([]*HostGroup, error)
	GetByID(id uint32) (*HostGroup, error)
	Store(u *HostGroup) error
	Update(u *HostGroup) error
}
