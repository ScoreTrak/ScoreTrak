package watermilllogger

import (
	"context"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.uber.org/zap"
)

// Logger implements watermill.LoggerAdapter with *otelzap.Logger.
type Logger struct {
	backend *otelzap.Logger
	fields  watermill.LogFields
}

// NewLogger returns new watermill.LoggerAdapter using passed *otelzap.Logger as backend.
func NewLogger(z *otelzap.Logger) watermill.LoggerAdapter {
	return &Logger{backend: z}
}

// Error writes error log with message, error and some fields.
func (l *Logger) Error(msg string, err error, fields watermill.LogFields) {
	fields = l.fields.Add(fields)
	fs := make([]zap.Field, 0, len(fields)+1)
	fs = append(fs, zap.Error(err))
	for k, v := range fields {
		fs = append(fs, zap.Any(k, v))
	}
	l.backend.ErrorContext(context.Background(), msg, fs...)
}

// Info writes info log with message and some fields.
func (l *Logger) Info(msg string, fields watermill.LogFields) {
	fields = l.fields.Add(fields)
	fs := make([]zap.Field, 0, len(fields)+1)
	for k, v := range fields {
		fs = append(fs, zap.Any(k, v))
	}
	l.backend.InfoContext(context.Background(), msg, fs...)
}

// Debug writes debug log with message and some fields.
func (l *Logger) Debug(msg string, fields watermill.LogFields) {
	fields = l.fields.Add(fields)
	fs := make([]zap.Field, 0, len(fields)+1)
	for k, v := range fields {
		fs = append(fs, zap.Any(k, v))
	}
	l.backend.DebugContext(context.Background(), msg, fs...)
}

// Trace writes debug log instead of trace log because zap does not support trace level logging.
func (l *Logger) Trace(msg string, fields watermill.LogFields) {
	fields = l.fields.Add(fields)
	fs := make([]zap.Field, 0, len(fields)+1)
	for k, v := range fields {
		fs = append(fs, zap.Any(k, v))
	}
	l.backend.DebugContext(context.Background(), msg, fs...)
}

// With returns new LoggerAdapter with passed fields.
func (l *Logger) With(fields watermill.LogFields) watermill.LoggerAdapter {
	return &Logger{
		backend: l.backend,
		fields:  l.fields.Add(fields),
	}
}
