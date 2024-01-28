package ansicode

// Performer is a interface for handling ANSI escape sequences.
type Performer struct {
	logger  Logger
	handler Handler

	hasPrecedingRune bool
	precedingRune    rune
}

// newPerformer creates a new Performer.
func NewPerformer(handler Handler, logger Logger) *Performer {
	return &Performer{
		handler: handler,
		logger:  logger,
	}
}

// Unhook is used to handle unhook operations.
func (p *Performer) Unhook() {
	p.logger.Tracef("Unhandled UNHOOK")
}

// Put is used to handle put operations.
func (p *Performer) Put(b byte) {
	p.logger.Tracef("Unhandled PUT byte=%v", b)
}

// Hook is used to handle hook operations.
func (p *Performer) Hook(params [][]uint16, intermediates []byte, ignore bool, r rune) {
	p.logger.Tracef("Unhandled HOOK params=%v intermediates=%v ignore=%v rune=%v", params, intermediates, ignore, r)
}
