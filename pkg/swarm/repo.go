package swarm

type Repo interface {
	Delete(id uint64) error
	GetAll() ([]*Swarm, error)
	GetByID(id uint64) (*Swarm, error)
	Store(u *Swarm) error
	Update(u *Swarm) error
}
