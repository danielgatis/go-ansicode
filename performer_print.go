package ansicode

// Print is used to handle print operations.
func (p *Performer) Print(r rune) {
	p.handler.Input(r)

	p.hasPrecedingRune = true
	p.precedingRune = r
}
