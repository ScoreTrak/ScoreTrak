package telemetry

import (
	"github.com/ScoreTrak/ScoreTrak/pkg/config"
	"go.uber.org/zap"
)

func NewLogger(config config.StaticConfig) (*zap.Logger, error) {
	var zapLogger *zap.Logger
	var err error
	if config.Prod {
		zapLogger, err = zap.NewProduction()
		if err != nil {
			return nil, err
		}
	} else {
		zapLogger, err = zap.NewDevelopment()
		if err != nil {
			return nil, err
		}
	}

	return zapLogger, err
}
