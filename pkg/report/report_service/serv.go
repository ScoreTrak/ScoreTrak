package report_service

import (
	"context"
	"github.com/ScoreTrak/ScoreTrak/pkg/report"
	repo2 "github.com/ScoreTrak/ScoreTrak/pkg/report/report_repo"
)

type Serv interface {
	Get(ctx context.Context) (*report.Report, error)
}

type reportServ struct {
	repo repo2.Repo
}

func NewReportServ(repo repo2.Repo) Serv {
	return &reportServ{
		repo: repo,
	}
}

func (svc *reportServ) Get(ctx context.Context) (*report.Report, error) { return svc.repo.Get(ctx) }
