package logger

import (
	"errors"
	"strings"

	"github.com/L1ghtman2k/ScoreTrak/pkg/config"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewZapLogger(config config.Logger) (*zap.SugaredLogger, error) {
	var cfg zap.Config

	switch strings.ToLower(config.Environment) {
	case "dev", "development":
		cfg = zap.NewDevelopmentConfig()
	case "prod", "production":
		cfg = zap.NewProductionConfig()
	default:
		return nil, errors.New("logger environment not supported.")
	}

	cfg.Level = zap.NewAtomicLevelAt(getLevel(config.LogLevel))
	cfg.OutputPaths = []string{config.FileName}
	log, err := cfg.Build()
	if err != nil {
		return nil, errors.New("zap logger build constructs failed.")
	}
	return log.Sugar(), nil
}

func getLevel(level string) zapcore.Level {
	var zapLevel zapcore.Level
	switch strings.ToLower(level) {
	case "debug":
		zapLevel = zap.DebugLevel

	case "info":
		zapLevel = zap.InfoLevel

	case "warn", "warning":
		zapLevel = zap.WarnLevel

	case "error":
		zapLevel = zap.ErrorLevel
	}
	return zapLevel
}
