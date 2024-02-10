package ansicode

import (
	"github.com/golang/glog"
)

var log = &Logger{}

// Logger wraps the glog.Logger.
type Logger struct{}

// Info logs a message at Info level.
func (l *Logger) Tracef(args ...interface{}) {
	glog.V(2).Info(args...)
}
