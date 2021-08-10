package log

import (
	"go.uber.org/zap"
)

// DefaultLogger is enabled when no consuming clients provide
// a logger to the server/cache subsystem
type DefaultLogger struct {
	log *zap.Logger
}

// NewDefaultLogger creates a DefaultLogger. It also provides the option to override the generic
// zap.NewProduction() logger with a custom instance.
func NewDefaultLogger(log *zap.Logger) *DefaultLogger {
	if log == nil {
		log, _ = zap.NewProduction()
	}

	return &DefaultLogger{log: log}
}

// Debugf logs a message at level debug on the standard logger.
func (l *DefaultLogger) Debugf(format string, args ...interface{}) {
	l.log.Sugar().Debugf(format, args...)
}

// Infof logs a message at level info on the standard logger.
func (l *DefaultLogger) Infof(format string, args ...interface{}) {
	l.log.Sugar().Infof(format, args...)
}

// Warnf logs a message at level warn on the standard logger.
func (l *DefaultLogger) Warnf(format string, args ...interface{}) {
	l.log.Sugar().Warnf(format, args...)
}

// Errorf logs a message at level error on the standard logger.
func (l *DefaultLogger) Errorf(format string, args ...interface{}) {
	l.log.Sugar().Errorf(format, args...)
}
