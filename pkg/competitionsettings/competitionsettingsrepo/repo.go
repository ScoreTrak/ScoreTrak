package competitionsettingsrepo

import (
	"context"

	"github.com/ScoreTrak/ScoreTrak/pkg/competitionsettings"
)

type Repo interface {
	Create(context.Context, *competitionsettings.CompetitionSettings) error
	Get(context.Context) (*competitionsettings.CompetitionSettings, error)
	Upsert(context.Context, *competitionsettings.CompetitionSettings) error
	Update(context.Context, *competitionsettings.CompetitionSettings) error
}
