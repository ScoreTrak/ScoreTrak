package check

type Repo interface {
	GetAllByRoundID(r_id uint64) ([]*Check, error)
	GetByRoundServiceID(r_id uint64, s_id uint64) ([]*Check, error)
	Delete(id uint64) error
	GetAll() ([]*Check, error)
	GetByID(id uint64) (*Check, error)
	Store(u *Check) error
	StoreMany(u []*Check) error
	Update(u *Check) error
}
