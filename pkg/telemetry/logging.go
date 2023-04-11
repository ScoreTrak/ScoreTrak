package telemetry

import (
	"github.com/ScoreTrak/ScoreTrak/pkg/config"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.uber.org/zap"
)

func newLoggerConfig() {
}

func NewLogger(config config.Config) (*zap.SugaredLogger, error) {
	var zapLogger *zap.Logger
	// var otelZapSuggaredLogger *otelzap.SugaredLogger
	var zapSugaredLogger *zap.SugaredLogger
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

	otelZapLogger := otelzap.New(zapLogger)
	zapSugaredLogger = otelZapLogger.Sugar().SugaredLogger

	return zapSugaredLogger, nil
}
