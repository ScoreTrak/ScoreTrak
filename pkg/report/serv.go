package report

type Serv interface {
	Get() (*Report, error)
}

type reportServ struct {
	repo Repo
}

func NewReportServ(repo Repo) Serv {
	return &reportServ{
		repo: repo,
	}
}

func (svc *reportServ) Get() (*Report, error) { return svc.repo.Get() }
