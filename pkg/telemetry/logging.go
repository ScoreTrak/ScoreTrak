package telemetry

import (
	"github.com/ScoreTrak/ScoreTrak/pkg/config"
	"github.com/ScoreTrak/ScoreTrak/pkg/telemetry/cronlogger"
	"github.com/ScoreTrak/ScoreTrak/pkg/telemetry/watermilllogger"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

func NewLogger(cfg *config.Config) (*otelzap.SugaredLogger, error) {
	var zapLogger *zap.Logger
	var err error
	if cfg.Dev {
		zapLogger, err = zap.NewDevelopment()
	} else {
		zapLogger, err = zap.NewProduction()
	}
	if err != nil {
		return nil, err
	}

	otelZapLogger := otelzap.New(zapLogger).Sugar()

	return otelZapLogger, nil
}

func NewFxEventLogger(logger *otelzap.SugaredLogger) fxevent.Logger {
	fxeventLogger := fxevent.ZapLogger{
		Logger: logger.SugaredLogger.Desugar(),
	}
	return &fxeventLogger
}

func NewWatermillLogger(logger *otelzap.SugaredLogger) watermill.LoggerAdapter {
	return watermilllogger.NewLogger(logger.Desugar())
}

func NewCronLogger(logger *otelzap.SugaredLogger) cronlogger.Logger {
	return cronlogger.NewLogger(logger.Desugar())
}
