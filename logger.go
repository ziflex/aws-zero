package awszero

import (
	"github.com/aws/smithy-go/logging"
	"github.com/rs/zerolog"
)

type Logger struct {
	log zerolog.Logger
}

// New creates a Zerolog-based implementation of Smithy Logger.
func New(log zerolog.Logger) logging.Logger {
	return &Logger{log}
}

// Logf is expected to support the standard fmt package "verbs".
func (l *Logger) Logf(classification logging.Classification, format string, v ...interface{}) {
	var evt *zerolog.Event

	switch classification {
	case logging.Warn:
		evt = l.log.Warn()
	case logging.Debug:
		evt = l.log.Debug()
	default:
		evt = l.log.Trace()
	}

	evt.Msgf(format, v...)
}
