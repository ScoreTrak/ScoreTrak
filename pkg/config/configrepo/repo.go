package configrepo

import (
	"context"

	"github.com/ScoreTrak/ScoreTrak/pkg/config"
)

type Repo interface {
	Create(context.Context, *config.DynamicConfig) error
	Get(context.Context) (*config.DynamicConfig, error)
	Upsert(context.Context, *config.DynamicConfig) error
	Update(context.Context, *config.DynamicConfig) error
}
