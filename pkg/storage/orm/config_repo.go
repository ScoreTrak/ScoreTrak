package orm

import (
	"context"
	"github.com/ScoreTrak/ScoreTrak/pkg/config"
	"github.com/ScoreTrak/ScoreTrak/pkg/logger"
	"gorm.io/gorm"
)

type configRepo struct {
	db  *gorm.DB
	log logger.LogInfoFormat
}

func NewConfigRepo(db *gorm.DB, log logger.LogInfoFormat) config.Repo {
	return &configRepo{db, log}
}

func (c *configRepo) Get(ctx context.Context) (*config.DynamicConfig, error) {
	cfg := &config.DynamicConfig{}
	cfg.ID = 1
	c.db.WithContext(ctx).Take(cfg)
	return cfg, nil
}

func (c *configRepo) Update(ctx context.Context, cfg *config.DynamicConfig) error {
	c.log.Debugf("updating the config")
	cfg.ID = 1
	err := c.db.WithContext(ctx).Model(cfg).Updates(config.DynamicConfig{RoundDuration: cfg.RoundDuration, Enabled: cfg.Enabled}).Error
	if err != nil {
		c.log.Errorf("error while updating the config, reason : %v", err)
		return err
	}
	return nil
}
