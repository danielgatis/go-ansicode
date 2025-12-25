package ansicode

// Performer is a interface for handling ANSI escape sequences.
type Performer struct {
	handler Handler

	hasPrecedingRune bool
	precedingRune    rune
}

// newPerformer creates a new Performer.
func NewPerformer(handler Handler) *Performer {
	return &Performer{
		handler: handler,
	}
}

// Unhook is used to handle unhook operations.
func (p *Performer) Unhook() {
	log.Debugf("Unhandled UNHOOK")
}

// Put is used to handle put operations.
func (p *Performer) Put(b byte) {
	log.Debugf("Unhandled PUT byte=%v", b)
}

// Hook is used to handle hook operations.
func (p *Performer) Hook(params [][]uint16, intermediates []byte, ignore bool, r rune) {
	log.Debugf("Unhandled HOOK params=%v intermediates=%v ignore=%v rune=%v", params, intermediates, ignore, r)
}

// SosPmApcDispatch is called when a SOS/PM/APC sequence is complete.
func (p *Performer) SosPmApcDispatch(kind byte, data []byte, bellTerminated bool) {
	switch kind {
	case 0: // SOS
		p.handler.StartOfStringReceived(data)
	case 1: // PM
		p.handler.PrivacyMessageReceived(data)
	case 2: // APC
		p.handler.ApplicationCommandReceived(data)
	}
}
