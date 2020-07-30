package check

import "github.com/gofrs/uuid"

type Serv interface {
	GetAllByRoundID(rID uint) ([]*Check, error)
	GetByRoundServiceID(rID uint, sID uuid.UUID) (*Check, error)
}

type checkServ struct {
	repo Repo
}

func NewCheckServ(repo Repo) Serv {
	return &checkServ{
		repo: repo,
	}
}

func (svc *checkServ) GetAllByRoundID(rID uint) ([]*Check, error) {
	return svc.repo.GetAllByRoundID(rID)
}
func (svc *checkServ) GetByRoundServiceID(rID uint, sID uuid.UUID) (*Check, error) {
	return svc.repo.GetByRoundServiceID(rID, sID)
}

func (svc *checkServ) Delete(rID uint, sID uuid.UUID) error { return svc.repo.Delete(rID, sID) }

func (svc *checkServ) GetAll() ([]*Check, error) { return svc.repo.GetAll() }

func (svc *checkServ) Store(c []*Check) error { return svc.repo.Store(c) }
