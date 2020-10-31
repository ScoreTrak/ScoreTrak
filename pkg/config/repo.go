package config

import "context"

type Repo interface {
	Get(ctx context.Context) (*DynamicConfig, error)
	Update(context.Context, *DynamicConfig) error
}
