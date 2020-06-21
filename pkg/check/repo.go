package check

type Repo interface {
	GetAllByRoundID(rID uint64) ([]*Check, error)
	GetByRoundServiceID(rID uint64, sID uint64) (*Check, error)
	Delete(rID uint64, sID uint64) error
	GetAll() ([]*Check, error)
	Store(u *Check) error
	StoreMany(u []*Check) error
}
