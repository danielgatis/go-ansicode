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

// StandardCharset is used to select a standard charset.
type StandardCharset int

// StandardCharset values.
const (
	StandardCharsetASCII StandardCharset = iota
	StandardCharsetSpecialCharacterAndLineDrawing
)

// Char returns the rune for the given rune in the given charset.
func (sc StandardCharset) Char(c rune) rune {
	switch sc {
	case StandardCharsetASCII:
		return c
	case StandardCharsetSpecialCharacterAndLineDrawing:
		switch c {
		case '_':
			return ' '
		case '`':
			return '◆'
		case 'a':
			return '▒'
		case 'b':
			return '␉'
		case 'c':
			return '␌'
		case 'd':
			return '␍'
		case 'e':
			return '␊'
		case 'f':
			return '°'
		case 'g':
			return '±'
		case 'h':
			return '␤'
		case 'i':
			return '␋'
		case 'j':
			return '┘'
		case 'k':
			return '┐'
		case 'l':
			return '┌'
		case 'm':
			return '└'
		case 'n':
			return '┼'
		case 'o':
			return '⎺'
		case 'p':
			return '⎻'
		case 'q':
			return '─'
		case 'r':
			return '⎼'
		case 's':
			return '⎽'
		case 't':
			return '├'
		case 'u':
			return '┤'
		case 'v':
			return '┴'
		case 'w':
			return '┬'
		case 'x':
			return '│'
		case 'y':
			return '≤'
		case 'z':
			return '≥'
		case '{':
			return 'π'
		case '|':
			return '≠'
		case '}':
			return '£'
		case '~':
			return '·'
		default:
			return c
		}
	}
	return c
}

// EscDispatch is used to handle esc operations.
func (p *Performer) EscDispatch(intermediates []byte, ignore bool, b byte) {
	switch true {
	case b == 'B' && len(intermediates) > 0:
		switch intermediates[0] {
		case '(':
			p.handler.ConfigureCharset(CharsetIndexG0, StandardCharsetASCII)

		case ')':
			p.handler.ConfigureCharset(CharsetIndexG1, StandardCharsetASCII)

		case '*':
			p.handler.ConfigureCharset(CharsetIndexG2, StandardCharsetASCII)

		case '+':
			p.handler.ConfigureCharset(CharsetIndexG3, StandardCharsetASCII)

		default:
			p.logger.Tracef("Unhandled ESC B intermediates=%v ignore=%v byte=%v", intermediates, ignore, b)
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
			p.handler.ConfigureCharset(CharsetIndexG0, StandardCharsetSpecialCharacterAndLineDrawing)

		case ')':
			p.handler.ConfigureCharset(CharsetIndexG1, StandardCharsetSpecialCharacterAndLineDrawing)

		case '*':
			p.handler.ConfigureCharset(CharsetIndexG2, StandardCharsetSpecialCharacterAndLineDrawing)

		case '+':
			p.handler.ConfigureCharset(CharsetIndexG3, StandardCharsetSpecialCharacterAndLineDrawing)

		default:
			p.logger.Tracef("Unhandled ESC B intermediates=%v ignore=%v byte=%v", intermediates, ignore, b)
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
		p.logger.Tracef("Unhandled ESC pintermediates=%v ignore=%v byte=%v", intermediates, ignore, b)
	}
}
