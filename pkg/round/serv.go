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
