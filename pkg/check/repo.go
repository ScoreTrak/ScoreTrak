package check

type Repo interface {
	GetAllByTeamRoundID(t_id string, r_id uint64) ([]*Check, error)
	GetByTeamRoundServiceID(t_id string, r_id uint64, s_id uint64) (*Check, error)
	Delete(id uint64) error
	GetAll() ([]*Check, error)
	GetByID(id uint64) (*Check, error)
	Store(u *Check) error
	Update(u *Check) error
}
