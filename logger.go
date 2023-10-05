package ansicode

// Logger is a interface for logging.
type Logger interface {

	// Tracef logs a message.
	Tracef(format string, args ...interface{})
}
