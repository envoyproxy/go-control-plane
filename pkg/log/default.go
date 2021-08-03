package log

import (
	"github.com/rs/zerolog/log"
)

// DefaultLogger is enabled when no consuming clients provide
// a logger to the server/cache subsystem
type DefaultLogger struct{}

func NewDefaultLogger() *DefaultLogger {
	return &DefaultLogger{}
}

// Debugf logs a message at level debug on the standard logger.
func (l *DefaultLogger) Debugf(format string, args ...interface{}) {
	log.Debug().Msgf(format, args...)
}

// Infof logs a message at level info on the standard logger.
func (l *DefaultLogger) Infof(format string, args ...interface{}) {
	log.Info().Msgf(format, args...)
}

// Warnf logs a message at level warn on the standard logger.
func (l *DefaultLogger) Warnf(format string, args ...interface{}) {
	log.Warn().Msgf(format, args...)
}

// Errorf logs a message at level error on the standard logger.
func (l *DefaultLogger) Errorf(format string, args ...interface{}) {
	log.Error().Msgf(format, args...)
}
