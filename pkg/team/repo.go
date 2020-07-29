package team

type Repo interface {
	Delete(id uint32) error
	GetAll() ([]*Team, error)
	GetByID(id uint32) (*Team, error)
	GetByName(name string) (*Team, error)
	DeleteByName(name string) error
	Store(u *Team) error
	UpdateByName(u *Team) error
	Update(u *Team) error
}
