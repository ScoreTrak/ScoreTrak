package round

type Serv interface {
	GetLastRound() (*Round, error)
}

type roundServ struct {
	repo Repo
}

func NewRoundServ(repo Repo) Serv {
	return &roundServ{
		repo: repo,
	}
}

func (svc *roundServ) GetLastRound() (*Round, error) { return svc.repo.GetLastRound() }

func (svc *roundServ) GetElapsingRound() (*Round, error) { return svc.repo.GetElapsingRound() }

func (svc *roundServ) Delete(id uint64) error { return svc.repo.Delete(id) }

func (svc *roundServ) GetAll() ([]*Round, error) { return svc.repo.GetAll() }

func (svc *roundServ) GetByID(id uint64) (*Round, error) { return svc.repo.GetByID(id) }

func (svc *roundServ) Store(u *Round) error { return svc.repo.Store(u) }

func (svc *roundServ) Update(u *Round) error { return svc.repo.Update(u) }
