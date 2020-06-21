package team

type Serv interface {
	GetAll() ([]*Team, error)
	GetByName(name string) (*Team, error)
	DeleteByName(name string) error
	Store(u *Team) error
	UpdateByName(u *Team) error
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

func (svc *teamServ) GetAll() ([]*Team, error) { return svc.repo.GetAll() }

func (svc *teamServ) GetByName(name string) (*Team, error) { return svc.repo.GetByName(name) }

func (svc *teamServ) Store(u *Team) error { return svc.repo.Store(u) }

func (svc *teamServ) UpdateByName(u *Team) error { return svc.repo.UpdateByName(u) }

func (svc *teamServ) Update(u *Team) error { return svc.repo.Update(u) }
