package report_repo

import (
	"context"
	"github.com/ScoreTrak/ScoreTrak/pkg/report"
	"github.com/gofrs/uuid"
)

type Repo interface {
	Get(ctx context.Context) (*report.Report, error)
	Update(context.Context, *report.Report) error
	CountPassedPerService(context.Context) (map[uuid.UUID]uint64, error)
}
