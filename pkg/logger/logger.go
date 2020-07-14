package logger

import (
	"errors"
	"go.uber.org/zap"
	"log"
)

type (
	LogInfo interface {
		Debug(args ...interface{})
		Info(args ...interface{})
		Warn(args ...interface{})
		Error(args ...interface{})
		Panic(args ...interface{})
		Fatal(args ...interface{})
	}

	LogFormat interface {
		Debugf(template string, args ...interface{})
		Infof(template string, args ...interface{})
		Warnf(template string, args ...interface{})
		Errorf(template string, args ...interface{})
		Panicf(template string, args ...interface{})
		Fatalf(template string, args ...interface{})
	}

	LogInfoFormat interface {
		LogInfo
		LogFormat
	}
)

type Config struct {
	Use         string `default:"zapLogger"`
	Environment string `default:"prod"`
	LogLevel    string `default:"info"`
	FileName    string `default:"scoretrak.log"`
}

type Logger struct {
	zapSugarLogger *zap.SugaredLogger
}

func NewLogger(c Config) (LogInfoFormat, error) {
	if c.Use == "zapLogger" {
		z, er := NewZapLogger(c)
		if er != nil {
			log.Fatalf("can't initialize zap logger: %v", er)
			return nil, er
		}
		return &Logger{zapSugarLogger: z}, nil

	}
	return nil, errors.New("logger not supported : " + c.Use)
}

func (l *Logger) Debug(args ...interface{}) {
	l.zapSugarLogger.Debug(args)
}

func (l *Logger) Info(args ...interface{}) {
	l.zapSugarLogger.Info(args)
}

func (l *Logger) Warn(args ...interface{}) {
	l.zapSugarLogger.Warn(args)
}

func (l *Logger) Error(args ...interface{}) {
	l.zapSugarLogger.Error(args)
}

func (l *Logger) Panic(args ...interface{}) {
	l.zapSugarLogger.Panic(args)
}

func (l *Logger) Fatal(args ...interface{}) {
	l.zapSugarLogger.Fatal(args)
}

func (l *Logger) Debugf(template string, args ...interface{}) {
	l.zapSugarLogger.Debugf(template, args)
}

func (l *Logger) Infof(template string, args ...interface{}) {
	l.zapSugarLogger.Infof(template, args)
}

func (l *Logger) Warnf(template string, args ...interface{}) {
	l.zapSugarLogger.Warnf(template, args)
}

func (l *Logger) Errorf(template string, args ...interface{}) {
	l.zapSugarLogger.Errorf(template, args)
}

func (l *Logger) Panicf(template string, args ...interface{}) {
	l.zapSugarLogger.Panicf(template, args)
}

func (l *Logger) Fatalf(template string, args ...interface{}) {
	l.zapSugarLogger.Fatalf(template, args)
}
