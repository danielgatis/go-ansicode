package ansicode

import (
	"fmt"
	"log/slog"
)

var log = &Logger{}

// Logger wraps the slog.
type Logger struct{}

// Debugf logs a message at Debug level.
func (l *Logger) Debugf(msg string, args ...interface{}) {
	slog.Debug(fmt.Sprintf(msg, args...))
}
