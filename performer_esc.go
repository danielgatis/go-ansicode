package ansicode

// CharsetIndex is used to index into the charset table.
type CharsetIndex int

// CharsetIndex values.
const (
	CharsetIndexG0 CharsetIndex = iota
	CharsetIndexG1
	CharsetIndexG2
	CharsetIndexG3
)

// Charset is used to select a standard charset.
type Charset int

// Charset values.
const (
	CharsetASCII Charset = iota
	CharsetLineDrawing
)

// EscDispatch is used to handle esc operations.
func (p *Performer) EscDispatch(intermediates []byte, ignore bool, b byte) {
	switch true {
	case b == 'B' && len(intermediates) > 0:
		switch intermediates[0] {
		case '(':
			p.handler.ConfigureCharset(CharsetIndexG0, CharsetASCII)

		case ')':
			p.handler.ConfigureCharset(CharsetIndexG1, CharsetASCII)

		case '*':
			p.handler.ConfigureCharset(CharsetIndexG2, CharsetASCII)

		case '+':
			p.handler.ConfigureCharset(CharsetIndexG3, CharsetASCII)

		default:
			log.Debugf("Unhandled ESC B intermediates=%v ignore=%v byte=%v", intermediates, ignore, b)
		}

	case b == 'D' && len(intermediates) == 0:
		p.handler.LineFeed()

	case b == 'E' && len(intermediates) == 0:
		p.handler.LineFeed()
		p.handler.CarriageReturn()

	case b == 'H' && len(intermediates) == 0:
		p.handler.HorizontalTabSet()

	case b == 'M' && len(intermediates) == 0:
		p.handler.ReverseIndex()

	case b == 'Z' && len(intermediates) == 0:
		p.handler.IdentifyTerminal(0)

	case b == 'c' && len(intermediates) == 0:
		p.handler.ResetState()

	case b == '0' && len(intermediates) > 0:
		switch intermediates[0] {
		case '(':
			p.handler.ConfigureCharset(CharsetIndexG0, CharsetLineDrawing)

		case ')':
			p.handler.ConfigureCharset(CharsetIndexG1, CharsetLineDrawing)

		case '*':
			p.handler.ConfigureCharset(CharsetIndexG2, CharsetLineDrawing)

		case '+':
			p.handler.ConfigureCharset(CharsetIndexG3, CharsetLineDrawing)

		default:
			log.Debugf("Unhandled ESC B intermediates=%v ignore=%v byte=%v", intermediates, ignore, b)
		}

	case b == '7' && len(intermediates) == 0:
		p.handler.SaveCursorPosition()

	case b == '8' && len(intermediates) > 0 && intermediates[0] == '#':
		p.handler.Decaln()

	case b == '8' && len(intermediates) == 0:
		p.handler.RestoreCursorPosition()

	case b == '=' && len(intermediates) == 0:
		p.handler.SetKeypadApplicationMode()

	case b == '>' && len(intermediates) == 0:
		p.handler.UnsetKeypadApplicationMode()

	case b == '\\' && len(intermediates) == 0:
		// Do nothing.

	default:
		log.Debugf("Unhandled ESC pintermediates=%v ignore=%v byte=%v", intermediates, ignore, b)
	}
}
