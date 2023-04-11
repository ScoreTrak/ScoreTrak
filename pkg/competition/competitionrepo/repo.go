package competitionrepo

import (
	"context"

	"github.com/ScoreTrak/ScoreTrak/pkg/competition"
)

type Repo interface {
	Create(context.Context, *competition.Competition) error
	Get() error
}
