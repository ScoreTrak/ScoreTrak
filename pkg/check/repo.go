package check

type Repo interface {
	GetAllByRoundID(rID uint32) ([]*Check, error)
	GetByRoundServiceID(rID uint32, sID uint32) (*Check, error)
	Delete(rID uint32, sID uint32) error
	GetAll() ([]*Check, error)
	Store(u *Check) error
	StoreMany(u []*Check) error
}
