package check

type Serv interface {
	GetAllByRoundID(rID uint32) ([]*Check, error)
	GetByRoundServiceID(rID uint32, sID uint32) (*Check, error)
}

type checkServ struct {
	repo Repo
}

func NewCheckServ(repo Repo) Serv {
	return &checkServ{
		repo: repo,
	}
}

func (svc *checkServ) GetAllByRoundID(rID uint32) ([]*Check, error) {
	return svc.repo.GetAllByRoundID(rID)
}
func (svc *checkServ) GetByRoundServiceID(rID uint32, sID uint32) (*Check, error) {
	return svc.repo.GetByRoundServiceID(rID, sID)
}

func (svc *checkServ) Delete(rID uint32, sID uint32) error { return svc.repo.Delete(rID, sID) }

func (svc *checkServ) GetAll() ([]*Check, error) { return svc.repo.GetAll() }

func (svc *checkServ) Store(c *Check) error { return svc.repo.Store(c) }

func (svc *checkServ) StoreMany(c []*Check) error { return svc.repo.StoreMany(c) }
