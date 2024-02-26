package telemetry

import (
	"github.com/ThreeDotsLabs/watermill"
	"github.com/scoretrak/scoretrak/pkg/telemetry/cronlogger"
	"github.com/scoretrak/scoretrak/pkg/telemetry/watermilllogger"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

func NewLogger() (*otelzap.SugaredLogger, error) {
	var zapLogger *zap.Logger
	var err error
	zapLogger, err = zap.NewProduction()
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

// Logging Keys
const (
	COMPETITION_ID = "competition_id"
	TEAM_ID        = "team_id"
	ROUND_ID       = "round_id"
	ROUND_NUMBER   = "round_number"
)
