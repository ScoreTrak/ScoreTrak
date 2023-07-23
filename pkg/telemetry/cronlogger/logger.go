package cronlogger

import (
	"context"
	"github.com/robfig/cron/v3"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.uber.org/zap"
)

// Logger implements cron.Logger.
type Logger struct {
	otelZaplogger *otelzap.Logger
}

var _ cron.Logger = (*Logger)(nil)

func NewLogger(z *otelzap.Logger) Logger {
	return Logger{otelZaplogger: z}
}

func (l Logger) Error(err error, msg string, keysAndValues ...interface{}) {
	l.otelZaplogger.ErrorContext(context.Background(), msg, mapKeysAndValues(keysAndValues)...)
}

func (l Logger) Info(msg string, keysAndValues ...interface{}) {
	l.otelZaplogger.InfoContext(context.Background(), msg, mapKeysAndValues(keysAndValues)...)
}

func mapKeysAndValues(fields ...interface{}) []zap.Field {
	ret := make([]zap.Field, 0, len(fields))
	for _, v := range fields {
		if kv, ok := v.([]interface{}); ok {
			for i := 0; i < len(kv); i += 2 {
				key := kv[i].(string)
				ret = append(ret, zap.Any(key, kv[i+1]))
			}
		}
	}
	return ret
}
