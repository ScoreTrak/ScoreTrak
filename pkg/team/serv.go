package team

type Serv interface {
	Delete(id string) error
	GetAll() ([]*Team, error)
	GetByID(id string) (*Team, error)
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

func (svc *teamServ) Delete(id string) error { return svc.repo.Delete(id) }

func (svc *teamServ) GetAll() ([]*Team, error) { return svc.repo.GetAll() }

func (svc *teamServ) GetByID(id string) (*Team, error) { return svc.repo.GetByID(id) }

func (svc *teamServ) Store(u *Team) error { return svc.repo.Store(u) }

func (svc *teamServ) Update(u *Team) error { return svc.repo.Update(u) }
