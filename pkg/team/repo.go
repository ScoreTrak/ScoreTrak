package team

type Repo interface {
	Delete(id string) error
	GetAll() ([]*Team, error)
	GetByID(id string) (*Team, error)
	Store(u *Team) error
	Update(u *Team) error
}
