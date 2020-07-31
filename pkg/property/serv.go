package property

import "github.com/gofrs/uuid"

type Serv interface {
	Delete(id uuid.UUID) error
	GetAll() ([]*Property, error)
	GetByID(id uuid.UUID) (*Property, error)
	Store(u []*Property) error
	Update(u *Property) error
	GetAllByTeamID(TeamID uuid.UUID) ([]*Property, error)
	GetAllByHostID(HostID uuid.UUID) ([]*Property, error)
}

type propertyServ struct {
	repo Repo
}

func NewPropertyServ(repo Repo) Serv {
	return &propertyServ{
		repo: repo,
	}
}

func (svc *propertyServ) Delete(id uuid.UUID) error { return svc.repo.Delete(id) }

func (svc *propertyServ) GetAll() ([]*Property, error) { return svc.repo.GetAll() }

func (svc *propertyServ) GetByID(id uuid.UUID) (*Property, error) { return svc.repo.GetByID(id) }

func (svc *propertyServ) Store(u []*Property) error { return svc.repo.Store(u) }

func (svc *propertyServ) Update(u *Property) error { return svc.repo.Update(u) }

func (svc *propertyServ) GetAllByTeamID(TeamID uuid.UUID) ([]*Property, error) {
	return svc.repo.GetAllByTeamID(TeamID)
}
func (svc *propertyServ) GetAllByHostID(HostID uuid.UUID) ([]*Property, error) {
	return svc.repo.GetAllByTeamHostID(HostID)
}
