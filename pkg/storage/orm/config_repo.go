package orm

import (
	"context"

	"github.com/ScoreTrak/ScoreTrak/pkg/config"
	"github.com/ScoreTrak/ScoreTrak/pkg/config/configrepo"
	"gorm.io/gorm"
)

type configRepo struct {
	db *gorm.DB
}

func NewConfigRepo(db *gorm.DB) configrepo.Repo {
	return &configRepo{db}
}

func (c *configRepo) Get(ctx context.Context) (*config.DynamicConfig, error) {
	cfg := &config.DynamicConfig{}
	cfg.ID = 1
	c.db.WithContext(ctx).Take(cfg)
	return cfg, nil
}

func (c *configRepo) Update(ctx context.Context, cfg *config.DynamicConfig) error {
	cfg.ID = 1
	err := c.db.WithContext(ctx).Model(cfg).Updates(config.DynamicConfig{RoundDuration: cfg.RoundDuration, Enabled: cfg.Enabled}).Error
	if err != nil {
		return err
	}
	return nil
}
