package team

type Serv interface {
	GetAll() ([]*Team, error)
	GetByID(id uint64) (*Team, error)
	Delete(id uint64) error
	Store(u *Team) error
	Update(u *Team) error
}

type teamServ struct {
	repo Repo
}

func NewTeamServ(repo Repo) Serv {
	return &teamServ{
		repo: repo,
	}
}

func (svc *teamServ) DeleteByName(name string) error { return svc.repo.DeleteByName(name) }

func (svc *teamServ) Delete(id uint64) error { return svc.repo.Delete(id) }

func (svc *teamServ) GetAll() ([]*Team, error) { return svc.repo.GetAll() }

func (svc *teamServ) GetByName(name string) (*Team, error) { return svc.repo.GetByName(name) }

func (svc *teamServ) GetByID(id uint64) (*Team, error) { return svc.repo.GetByID(id) }

func (svc *teamServ) Store(u *Team) error { return svc.repo.Store(u) }

func (svc *teamServ) UpdateByName(u *Team) error { return svc.repo.UpdateByName(u) }

func (svc *teamServ) Update(u *Team) error { return svc.repo.Update(u) }
