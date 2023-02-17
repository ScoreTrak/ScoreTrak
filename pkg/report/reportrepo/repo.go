package reportrepo

import (
	"context"
	"github.com/ScoreTrak/ScoreTrak/pkg/report"
	"github.com/gofrs/uuid"
)

type Repo interface {
	Create(context.Context, *report.Report) error
	Get(ctx context.Context) (*report.Report, error)
	Update(context.Context, *report.Report) error
	Upsert(ctx context.Context, r *report.Report) error
	CountPassedPerService(context.Context) (map[uuid.UUID]uint64, error)
}
