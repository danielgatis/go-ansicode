package ansicode

type c0 byte

const (
	c0NUL  c0 = 0x00
	c0SOH  c0 = 0x01
	c0STX  c0 = 0x02
	c0ETX  c0 = 0x03
	c0EOT  c0 = 0x04
	c0ENQ  c0 = 0x05
	c0ACK  c0 = 0x06
	c0BEL  c0 = 0x07
	c0BS   c0 = 0x08
	c0HT   c0 = 0x09
	c0LF   c0 = 0x0A
	c0VT   c0 = 0x0B
	c0FF   c0 = 0x0C
	c0CR   c0 = 0x0D
	c0SO   c0 = 0x0E
	c0SI   c0 = 0x0F
	c0DLE  c0 = 0x10
	c0XON  c0 = 0x11
	c0DC2  c0 = 0x12
	c0XOFF c0 = 0x13
	c0DC4  c0 = 0x14
	c0NAK  c0 = 0x15
	c0SYN  c0 = 0x16
	c0ETB  c0 = 0x17
	c0CAN  c0 = 0x18
	c0EM   c0 = 0x19
	c0SUB  c0 = 0x1A
	c0ESC  c0 = 0x1B
	c0FS   c0 = 0x1C
	c0GS   c0 = 0x1D
	c0RS   c0 = 0x1E
	c0US   c0 = 0x1F
	c0DEL  c0 = 0x7f
)

// Execute is used to handle execute operations.
func (p *Performer) Execute(b byte) {
	switch c0(b) {
	case c0HT:
		p.handler.Tab(1)
	case c0BS:
		p.handler.Backspace()
	case c0CR:
		p.handler.CarriageReturn()
	case c0LF, c0VT, c0FF:
		p.handler.LineFeed()
	case c0BEL:
		p.handler.Bell()
	case c0SUB:
		p.handler.Substitute()
	case c0SI:
		p.handler.SetActiveCharset(0)
	case c0SO:
		p.handler.SetActiveCharset(1)
	default:
		log.Tracef("Unhandled EXECUTE byte=%v", b)
	}
}
