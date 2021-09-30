package config_repo

import (
	"context"

	"github.com/ScoreTrak/ScoreTrak/pkg/config"
)

type Repo interface {
	Get(ctx context.Context) (*config.DynamicConfig, error)
	Update(context.Context, *config.DynamicConfig) error
}
