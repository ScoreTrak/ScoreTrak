package check

import "github.com/gofrs/uuid"

type Serv interface {
	GetAllByRoundID(roundID uint) ([]*Check, error)
	GetByRoundServiceID(roundID uint, serviceID uuid.UUID) (*Check, error)
	GetAllByServiceID(serviceID uuid.UUID) ([]*Check, error)
}

type checkServ struct {
	repo Repo
}

func NewCheckServ(repo Repo) Serv {
	return &checkServ{
		repo: repo,
	}
}

func (svc *checkServ) GetAllByRoundID(roundID uint) ([]*Check, error) {
	return svc.repo.GetAllByRoundID(roundID)
}
func (svc *checkServ) GetByRoundServiceID(roundID uint, serviceID uuid.UUID) (*Check, error) {
	return svc.repo.GetByRoundServiceID(roundID, serviceID)
}
func (svc *checkServ) GetAllByServiceID(serviceID uuid.UUID) ([]*Check, error) {
	return svc.repo.GetAllByServiceID(serviceID)
}

func (svc *checkServ) Delete(roundID uint, serviceID uuid.UUID) error {
	return svc.repo.Delete(roundID, serviceID)
}

func (svc *checkServ) GetAll() ([]*Check, error) { return svc.repo.GetAll() }

func (svc *checkServ) Store(c []*Check) error { return svc.repo.Store(c) }
