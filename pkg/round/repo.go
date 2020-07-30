package round

type Repo interface {
	Delete(id uint) error
	GetAll() ([]*Round, error)
	GetByID(id uint) (*Round, error)
	Store(u *Round) error
	Update(u *Round) error
	GetLastRound() (*Round, error)
	GetLastNonElapsingRound() (*Round, error)
	GetLastElapsingRound() (*Round, error)
}
