package report

type Serv interface {
	Get() (*Report, error)
}

type checkServ struct {
	repo Repo
}

func NewReportServ(repo Repo) Serv {
	return &checkServ{
		repo: repo,
	}
}
func (svc *checkServ) Get() (*Report, error) { return svc.repo.Get() }
