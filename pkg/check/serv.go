package check

type Serv interface {
	GetAllByRoundID(r_id uint64) ([]*Check, error)
	GetByRoundServiceID(r_id uint64, s_id uint64) ([]*Check, error)
	Delete(id uint64) error
	GetAll() ([]*Check, error)
	GetByID(id uint64) (*Check, error)
	Store(u *Check) error
	Update(u *Check) error
}

type checkServ struct {
	repo Repo
}

func NewCheckServ(repo Repo) Serv {
	return &checkServ{
		repo: repo,
	}
}

func (svc *checkServ) GetAllByRoundID(r_id uint64) ([]*Check, error) {
	return svc.repo.GetAllByRoundID(r_id)
}
func (svc *checkServ) GetByRoundServiceID(r_id uint64, s_id uint64) ([]*Check, error) {
	return svc.repo.GetByRoundServiceID(r_id, s_id)
}

func (svc *checkServ) Delete(id uint64) error { return svc.repo.Delete(id) }

func (svc *checkServ) GetAll() ([]*Check, error) { return svc.repo.GetAll() }

func (svc *checkServ) GetByID(id uint64) (*Check, error) { return svc.repo.GetByID(id) }

func (svc *checkServ) Store(u *Check) error { return svc.repo.Store(u) }

func (svc *checkServ) Update(u *Check) error { return svc.repo.Update(u) }
