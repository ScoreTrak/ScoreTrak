package repo

import (
	"context"
	"github.com/ScoreTrak/ScoreTrak/pkg/report"
)

type Repo interface {
	Get(ctx context.Context) (*report.Report, error)
	Update(context.Context, *report.Report) error
}
