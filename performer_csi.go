package ansicode

import (
	"github.com/danielgatis/go-iterator"
)

// ModifyOtherKeys is the XTerm modify other keys mode.
type ModifyOtherKeys int

// ModifyOtherKeys values.
const (
	ModifyOtherKeysReset ModifyOtherKeys = iota
	ModifyOtherKeysEnableExceptWellDefined
	ModifyOtherKeysResetEnableAll
)

// CursorStyle is the cursor style.
type CursorStyle int

// CursorStyle values.
const (
	CursorStyleBlinkingBlock CursorStyle = iota
	CursorStyleSteadyBlock
	CursorStyleBlinkingUnderline
	CursorStyleSteadyUnderline
	CursorStyleBlinkingBar
	CursorStyleSteadyBar
)

// KeyboardMode is the keyboard mode.
type KeyboardMode byte

// KeyboardMode values.
const (
	KeyboardModeNoMode               KeyboardMode = 0b0000_0000
	KeyboardModeDisambiguateEscCodes KeyboardMode = 0b0000_0001
	KeyboardModeReportEventTypes     KeyboardMode = 0b0000_0010
	KeyboardModeReportAlternateKeys  KeyboardMode = 0b0000_0100
	KeyboardModeReportAllKeysAsEsc   KeyboardMode = 0b0000_1000
	KeyboardModeReportAssociatedText KeyboardMode = 0b0001_0000
)

// KeyboardModeBehavior is the keyboard mode behavior.
type KeyboardModeBehavior int

// KeyboardModeBehavior values.
const (
	KeyboardModeBehaviorReplace KeyboardModeBehavior = iota
	KeyboardModeBehaviorUnion
	KeyboardModeBehaviorDifference
)

// LineClearMode is the line clear mode.
type LineClearMode int

// LineClearMode values.
const (
	LineClearModeRight LineClearMode = iota
	LineClearModeLeft
	LineClearModeAll
)

// ClearMode is the clear mode.
type ClearMode int

// ClearMode values.
const (
	ClearModeBelow ClearMode = iota
	ClearModeAbove
	ClearModeAll
	ClearModeSaved
)

// TabulationClearMode is the tabulation clear mode.
type TabulationClearMode int

// TabulationClearMode values.
const (
	TabulationClearModeCurrent TabulationClearMode = iota
	TabulationClearModeAll
)

// TerminalMode is the terminal mode.
type TerminalMode int

// TerminalMode values.
const (
	TerminalModeCursorKeys                    TerminalMode = 1
	TerminalModeColumnMode                    TerminalMode = 3
	TerminalModeInsert                        TerminalMode = 4
	TerminalModeOrigin                        TerminalMode = 6
	TerminalModeLineWrap                      TerminalMode = 7
	TerminalModeBlinkingCursor                TerminalMode = 12
	TerminalModeLineFeedNewLine               TerminalMode = 20
	TerminalModeShowCursor                    TerminalMode = 25
	TerminalModeReportMouseClicks             TerminalMode = 1000
	TerminalModeReportCellMouseMotion         TerminalMode = 1002
	TerminalModeReportAllMouseMotion          TerminalMode = 1003
	TerminalModeReportFocusInOut              TerminalMode = 1004
	TerminalModeUTF8Mouse                     TerminalMode = 1005
	TerminalModeSGRMouse                      TerminalMode = 1006
	TerminalModeAlternateScroll               TerminalMode = 1007
	TerminalModeUrgencyHints                  TerminalMode = 1042
	TerminalModeSwapScreenAndSetRestoreCursor TerminalMode = 1049
	TerminalModeBracketedPaste                TerminalMode = 2004
)

// CharAttribute is the character attribute.
type CharAttribute int

// CharAttribute values.
const (
	CharAttributeReset CharAttribute = iota
	CharAttributeBold
	CharAttributeDim
	CharAttributeItalic
	CharAttributeUnderline
	CharAttributeDoubleUnderline
	CharAttributeCurlyUnderline
	CharAttributeDottedUnderline
	CharAttributeDashedUnderline
	CharAttributeBlinkSlow
	CharAttributeBlinkFast
	CharAttributeReverse
	CharAttributeHidden
	CharAttributeStrike
	CharAttributeCancelBold
	CharAttributeCancelBoldDim
	CharAttributeCancelItalic
	CharAttributeCancelUnderline
	CharAttributeCancelBlink
	CharAttributeCancelReverse
	CharAttributeCancelHidden
	CharAttributeCancelStrike
	CharAttributeForeground
	CharAttributeBackground
	CharAttributeUnderlineColor
)

// TerminalCharAttribute is the terminal character attribute.
type TerminalCharAttribute struct {
	Attr         CharAttribute
	NamedColor   *NamedColor
	RGBColor     *RGBColor
	IndexedColor *IndexedColor
}

// CsiDispatch is used to handle csi operations.
func (p *Performer) CsiDispatch(params [][]uint16, intermediates []byte, ignore bool, action rune) {
	if ignore || len(intermediates) > 1 {
		return
	}

	flatParams := make([]uint16, 0, len(params))
	for _, param := range params {
		flatParams = append(flatParams, param...)
	}

	paramsIter := iterator.NewIterator(flatParams)

	switch true {
	case action == '@' && len(intermediates) == 0:
		n := paramsIter.GetNextOrDefault(1)
		p.handler.InsertBlank(int(n))

	case action == 'A' && len(intermediates) == 0:
		n := paramsIter.GetNextOrDefault(1)
		p.handler.MoveUp(int(n))

	case action == 'B' && len(intermediates) == 0, action == 'e' && len(intermediates) == 0:
		n := paramsIter.GetNextOrDefault(1)
		p.handler.MoveDown(int(n))

	case action == 'b' && len(intermediates) == 0:
		if !p.hasPrecedingRune {
			return
		}

		n := paramsIter.GetNextOrDefault(1)
		for i := 0; i < int(n); i++ {
			p.handler.Input(p.precedingRune)
		}

	case action == 'C' && len(intermediates) == 0, action == 'a' && len(intermediates) == 0:
		n := paramsIter.GetNextOrDefault(1)
		p.handler.MoveForward(int(n))

	case action == 'c':
		n := paramsIter.GetNextOrDefault(0)
		if n == 0 {
			if len(intermediates) > 0 {
				p.handler.IdentifyTerminal(intermediates[0])
			} else {
				p.handler.IdentifyTerminal(0)
			}
		}

	case action == 'D' && len(intermediates) == 0:
		n := paramsIter.GetNextOrDefault(1)
		p.handler.MoveBackward(int(n))

	case action == 'd' && len(intermediates) == 0:
		n := paramsIter.GetNextOrDefault(1)
		p.handler.GotoLine(int(n - 1))

	case action == 'E' && len(intermediates) == 0:
		n := paramsIter.GetNextOrDefault(1)
		p.handler.MoveDownCr(int(n - 1))

	case action == 'F' && len(intermediates) == 0:
		n := paramsIter.GetNextOrDefault(1)
		p.handler.MoveUpCr(int(n - 1))

	case action == 'G' && len(intermediates) == 0, action == '`' && len(intermediates) == 0:
		n := paramsIter.GetNextOrDefault(1)
		p.handler.GotoCol(int(n - 1))

	case action == 'g' && len(intermediates) == 0:
		mode := paramsIter.GetNextOrDefault(0)
		if int(mode) == 0 {
			p.handler.ClearTabs(TabulationClearModeCurrent)
		}

		if int(mode) == 3 {
			p.handler.ClearTabs(TabulationClearModeAll)
		}

	case action == 'H' && len(intermediates) == 0, action == 'f' && len(intermediates) == 0:
		y := paramsIter.GetNextOrDefault(1)
		x := paramsIter.GetNextOrDefault(1)

		p.handler.Goto(int(y-1), int(x-1))

	case action == 'I' && len(intermediates) == 0:
		n := paramsIter.GetNextOrDefault(1)
		p.handler.MoveForwardTabs(int(n))

	case action == 'J' && len(intermediates) == 0:
		mode := paramsIter.GetNextOrDefault(0)
		if int(mode) == 0 {
			p.handler.ClearScreen(ClearModeBelow)
		}

		if int(mode) == 1 {
			p.handler.ClearScreen(ClearModeAbove)
		}

		if int(mode) == 2 {
			p.handler.ClearScreen(ClearModeAll)
		}

		if int(mode) == 3 {
			p.handler.ClearScreen(ClearModeSaved)
		}

	case action == 'K' && len(intermediates) == 0:
		mode := paramsIter.GetNextOrDefault(0)
		if int(mode) == 0 {
			p.handler.ClearLine(LineClearModeRight)
		}

		if int(mode) == 1 {
			p.handler.ClearLine(LineClearModeLeft)
		}

		if int(mode) == 2 {
			p.handler.ClearLine(LineClearModeAll)
		}

	case action == 'L' && len(intermediates) == 0:
		n := paramsIter.GetNextOrDefault(1)
		p.handler.InsertBlankLines(int(n))

	case action == 'l' && len(intermediates) == 0:
		for _, param := range flatParams {
			mode, ok := terminalMode(param, false)
			if ok {
				p.handler.UnsetMode(mode)
			} else {
				log.Tracef("Unhandled CSI params=%v intermediates=%v ignore=%v action=%v", params, intermediates, ignore, action)
			}
		}

	case action == 'l' && len(intermediates) > 0 && intermediates[0] == '?':
		for _, param := range flatParams {
			mode, ok := terminalMode(param, true)
			if ok {
				p.handler.UnsetMode(mode)
			} else {
				log.Tracef("Unhandled CSI params=%v intermediates=%v ignore=%v action=%v", params, intermediates, ignore, action)
			}
		}

	case action == 'M' && len(intermediates) == 0:
		n := paramsIter.GetNextOrDefault(1)
		p.handler.DeleteLines(int(n))

	case action == 'm' && len(intermediates) == 0:
		if len(flatParams) == 0 {
			p.handler.SetTerminalCharAttribute(attr(CharAttributeReset))
			return
		}

		for paramsIter.HasNext() {
			param, ok := paramsIter.GetNext()
			if !ok {
				log.Tracef("Unhandled CSI params=%v intermediates=%v ignore=%v action=%v", params, intermediates, ignore, action)
				continue
			}

			switch int(param) {
			case 0:
				p.handler.SetTerminalCharAttribute(attr(CharAttributeReset))

			case 1:
				p.handler.SetTerminalCharAttribute(attr(CharAttributeBold))

			case 2:
				p.handler.SetTerminalCharAttribute(attr(CharAttributeDim))

			case 3:
				p.handler.SetTerminalCharAttribute(attr(CharAttributeItalic))

			case 4:
				switch paramsIter.GetNextOrDefault(0) {
				case 0:
					p.handler.SetTerminalCharAttribute(attr(CharAttributeCancelUnderline))

				case 2:
					p.handler.SetTerminalCharAttribute(attr(CharAttributeDoubleUnderline))

				case 3:
					p.handler.SetTerminalCharAttribute(attr(CharAttributeCurlyUnderline))

				case 4:
					p.handler.SetTerminalCharAttribute(attr(CharAttributeDottedUnderline))

				case 5:
					p.handler.SetTerminalCharAttribute(attr(CharAttributeDashedUnderline))

				default:
					p.handler.SetTerminalCharAttribute(attr(CharAttributeUnderline))
				}

			case 5:
				p.handler.SetTerminalCharAttribute(attr(CharAttributeBlinkSlow))

			case 6:
				p.handler.SetTerminalCharAttribute(attr(CharAttributeBlinkFast))

			case 7:
				p.handler.SetTerminalCharAttribute(attr(CharAttributeReverse))

			case 8:
				p.handler.SetTerminalCharAttribute(attr(CharAttributeHidden))

			case 9:
				p.handler.SetTerminalCharAttribute(attr(CharAttributeStrike))

			case 21:
				p.handler.SetTerminalCharAttribute(attr(CharAttributeCancelBold))

			case 22:
				p.handler.SetTerminalCharAttribute(attr(CharAttributeCancelBoldDim))

			case 23:
				p.handler.SetTerminalCharAttribute(attr(CharAttributeCancelItalic))

			case 24:
				p.handler.SetTerminalCharAttribute(attr(CharAttributeCancelUnderline))

			case 25:
				p.handler.SetTerminalCharAttribute(attr(CharAttributeCancelBlink))

			case 27:
				p.handler.SetTerminalCharAttribute(attr(CharAttributeCancelReverse))

			case 28:
				p.handler.SetTerminalCharAttribute(attr(CharAttributeCancelHidden))

			case 29:
				p.handler.SetTerminalCharAttribute(attr(CharAttributeCancelStrike))

			case 30:
				p.handler.SetTerminalCharAttribute(attrWithNamedColor(CharAttributeForeground, NamedColorBlack))

			case 31:
				p.handler.SetTerminalCharAttribute(attrWithNamedColor(CharAttributeForeground, NamedColorRed))

			case 32:
				p.handler.SetTerminalCharAttribute(attrWithNamedColor(CharAttributeForeground, NamedColorGreen))

			case 33:
				p.handler.SetTerminalCharAttribute(attrWithNamedColor(CharAttributeForeground, NamedColorYellow))

			case 34:
				p.handler.SetTerminalCharAttribute(attrWithNamedColor(CharAttributeForeground, NamedColorBlue))

			case 35:
				p.handler.SetTerminalCharAttribute(attrWithNamedColor(CharAttributeForeground, NamedColorMagenta))

			case 36:
				p.handler.SetTerminalCharAttribute(attrWithNamedColor(CharAttributeForeground, NamedColorCyan))

			case 37:
				p.handler.SetTerminalCharAttribute(attrWithNamedColor(CharAttributeForeground, NamedColorWhite))

			case 38:
				for paramsIter.HasNext() {
					m, ok := paramsIter.GetNext()
					if !ok {
						continue
					}

					switch m {
					case 2:
						r := paramsIter.GetNextOrDefault(0)
						g := paramsIter.GetNextOrDefault(0)
						b := paramsIter.GetNextOrDefault(0)

						p.handler.SetTerminalCharAttribute(attrWithRGBColor(CharAttributeForeground, r, g, b))

					case 5:
						i := paramsIter.GetNextOrDefault(0)
						p.handler.SetTerminalCharAttribute(attrWithIndexedColor(CharAttributeForeground, i))
					}
				}

			case 39:
				p.handler.SetTerminalCharAttribute(attrWithNamedColor(CharAttributeForeground, NamedColorForeground))

			case 40:
				p.handler.SetTerminalCharAttribute(attrWithNamedColor(CharAttributeBackground, NamedColorBlack))

			case 41:
				p.handler.SetTerminalCharAttribute(attrWithNamedColor(CharAttributeBackground, NamedColorRed))

			case 42:
				p.handler.SetTerminalCharAttribute(attrWithNamedColor(CharAttributeBackground, NamedColorGreen))

			case 43:
				p.handler.SetTerminalCharAttribute(attrWithNamedColor(CharAttributeBackground, NamedColorYellow))

			case 44:
				p.handler.SetTerminalCharAttribute(attrWithNamedColor(CharAttributeBackground, NamedColorBlue))

			case 45:
				p.handler.SetTerminalCharAttribute(attrWithNamedColor(CharAttributeBackground, NamedColorMagenta))

			case 46:
				p.handler.SetTerminalCharAttribute(attrWithNamedColor(CharAttributeBackground, NamedColorCyan))

			case 47:
				p.handler.SetTerminalCharAttribute(attrWithNamedColor(CharAttributeBackground, NamedColorWhite))

			case 48:
				for paramsIter.HasNext() {
					m, ok := paramsIter.GetNext()
					if !ok {
						continue
					}

					switch m {
					case 2:
						r := paramsIter.GetNextOrDefault(0)
						g := paramsIter.GetNextOrDefault(0)
						b := paramsIter.GetNextOrDefault(0)

						p.handler.SetTerminalCharAttribute(attrWithRGBColor(CharAttributeBackground, r, g, b))

					case 5:
						i := paramsIter.GetNextOrDefault(0)
						p.handler.SetTerminalCharAttribute(attrWithIndexedColor(CharAttributeBackground, i))
					}
				}

			case 49:
				p.handler.SetTerminalCharAttribute(attrWithNamedColor(CharAttributeBackground, NamedColorBackground))

			case 58:
				for paramsIter.HasNext() {
					m, ok := paramsIter.GetNext()
					if !ok {
						continue
					}

					switch m {
					case 2:
						r := paramsIter.GetNextOrDefault(0)
						g := paramsIter.GetNextOrDefault(0)
						b := paramsIter.GetNextOrDefault(0)

						p.handler.SetTerminalCharAttribute(attrWithRGBColor(CharAttributeUnderlineColor, r, g, b))
					case 5:
						i := paramsIter.GetNextOrDefault(0)
						p.handler.SetTerminalCharAttribute(attrWithIndexedColor(CharAttributeUnderlineColor, i))
					}
				}

			case 59:
				p.handler.SetTerminalCharAttribute(attr(CharAttributeUnderlineColor))

			case 90:
				p.handler.SetTerminalCharAttribute(attrWithNamedColor(CharAttributeForeground, NamedColorBrightBlack))

			case 91:
				p.handler.SetTerminalCharAttribute(attrWithNamedColor(CharAttributeForeground, NamedColorBrightRed))

			case 92:
				p.handler.SetTerminalCharAttribute(attrWithNamedColor(CharAttributeForeground, NamedColorBrightGreen))

			case 93:
				p.handler.SetTerminalCharAttribute(attrWithNamedColor(CharAttributeForeground, NamedColorBrightYellow))

			case 94:
				p.handler.SetTerminalCharAttribute(attrWithNamedColor(CharAttributeForeground, NamedColorBrightBlue))

			case 95:
				p.handler.SetTerminalCharAttribute(attrWithNamedColor(CharAttributeForeground, NamedColorBrightMagenta))

			case 96:
				p.handler.SetTerminalCharAttribute(attrWithNamedColor(CharAttributeForeground, NamedColorBrightCyan))

			case 97:
				p.handler.SetTerminalCharAttribute(attrWithNamedColor(CharAttributeForeground, NamedColorBrightWhite))

			case 100:
				p.handler.SetTerminalCharAttribute(attrWithNamedColor(CharAttributeBackground, NamedColorBrightBlack))

			case 101:
				p.handler.SetTerminalCharAttribute(attrWithNamedColor(CharAttributeBackground, NamedColorBrightRed))

			case 102:
				p.handler.SetTerminalCharAttribute(attrWithNamedColor(CharAttributeBackground, NamedColorBrightGreen))

			case 103:
				p.handler.SetTerminalCharAttribute(attrWithNamedColor(CharAttributeBackground, NamedColorBrightYellow))

			case 104:
				p.handler.SetTerminalCharAttribute(attrWithNamedColor(CharAttributeBackground, NamedColorBrightBlue))

			case 105:
				p.handler.SetTerminalCharAttribute(attrWithNamedColor(CharAttributeBackground, NamedColorBrightMagenta))

			case 106:
				p.handler.SetTerminalCharAttribute(attrWithNamedColor(CharAttributeBackground, NamedColorBrightCyan))

			case 107:
				p.handler.SetTerminalCharAttribute(attrWithNamedColor(CharAttributeBackground, NamedColorBrightWhite))

			default:
				log.Tracef("Unhandled CSI params=%v intermediates=%v ignore=%v action=%v", params, intermediates, ignore, action)
			}
		}

	case action == 'm' && len(intermediates) > 0 && intermediates[0] == '>':
		n := paramsIter.GetNextOrDefault(0)
		if n != 4 {
			log.Tracef("Unhandled CSI params=%v intermediates=%v ignore=%v action=%v", params, intermediates, ignore, action)
			return
		}

		m := paramsIter.GetNextOrDefault(0)

		switch int(m) {
		case 0:
			p.handler.SetModifyOtherKeys(ModifyOtherKeysReset)

		case 1:
			p.handler.SetModifyOtherKeys(ModifyOtherKeysEnableExceptWellDefined)

		case 2:
			p.handler.SetModifyOtherKeys(ModifyOtherKeysResetEnableAll)

		default:
			log.Tracef("Unhandled CSI params=%v intermediates=%v ignore=%v action=%v", params, intermediates, ignore, action)
		}

	case action == 'm' && len(intermediates) > 0 && intermediates[0] == '?':
		n := paramsIter.GetNextOrDefault(0)
		if n != 4 {
			log.Tracef("Unhandled CSI params=%v intermediates=%v ignore=%v action=%v", params, intermediates, ignore, action)
			return
		}

		p.handler.ReportModifyOtherKeys()

	case action == 'n' && len(intermediates) == 0:
		n := paramsIter.GetNextOrDefault(0)
		p.handler.DeviceStatus(int(n))

	case action == 'P' && len(intermediates) == 0:
		n := paramsIter.GetNextOrDefault(1)
		p.handler.DeleteChars(int(n))

	case action == 'q' && len(intermediates) > 0 && intermediates[0] == ' ':
		n := paramsIter.GetNextOrDefault(0)

		switch int(n) {
		case 0:
			p.handler.SetCursorStyle(CursorStyleBlinkingBlock)

		case 1:
			p.handler.SetCursorStyle(CursorStyleBlinkingBlock)

		case 2:
			p.handler.SetCursorStyle(CursorStyleSteadyBlock)

		case 3:
			p.handler.SetCursorStyle(CursorStyleBlinkingUnderline)

		case 4:
			p.handler.SetCursorStyle(CursorStyleSteadyUnderline)

		case 5:
			p.handler.SetCursorStyle(CursorStyleBlinkingBar)

		case 6:
			p.handler.SetCursorStyle(CursorStyleSteadyBar)

		default:
			log.Tracef("Unhandled CSI params=%v intermediates=%v ignore=%v action=%v", params, intermediates, ignore, action)
		}

	case action == 'r' && len(intermediates) == 0:
		top := paramsIter.GetNextOrDefault(1)
		bottom := paramsIter.GetNextOrDefault(1)

		p.handler.SetScrollingRegion(int(top), int(bottom))

	case action == 'S' && len(intermediates) == 0:
		n := paramsIter.GetNextOrDefault(1)
		p.handler.ScrollUp(int(n))

	case action == 's' && len(intermediates) == 0:
		p.handler.SaveCursorPosition()

	case action == 'T' && len(intermediates) == 0:
		n := paramsIter.GetNextOrDefault(1)
		p.handler.ScrollDown(int(n))

	case action == 't' && len(intermediates) == 0:
		n := paramsIter.GetNextOrDefault(1)
		switch int(n) {
		case 14:
			p.handler.TextAreaSizePixels()

		case 18:
			p.handler.TextAreaSizeChars()

		case 22:
			p.handler.PushTitle()

		case 23:
			p.handler.PopTitle()
		}

	case action == 'u' && len(intermediates) > 0 && intermediates[0] == '?':
		p.handler.ReportKeyboardMode()

	case action == 'u' && len(intermediates) > 0 && intermediates[0] == '=':
		n := paramsIter.GetNextOrDefault(0)
		mode := keyboardMode(n)

		m := paramsIter.GetNextOrDefault(1)
		behavior := keyboardBehavior(m)

		p.handler.SetKeyboardMode(mode, behavior)

	case action == 'u' && len(intermediates) > 0 && intermediates[0] == '>':
		n := paramsIter.GetNextOrDefault(0)
		mode := keyboardMode(n)
		p.handler.PushKeyboardMode(mode)

	case action == 'u' && len(intermediates) > 0 && intermediates[0] == '<':
		n := paramsIter.GetNextOrDefault(1)
		p.handler.PopKeyboardMode(int(n))

	case action == 'u' && len(intermediates) == 0:
		p.handler.RestoreCursorPosition()

	case action == 'X' && len(intermediates) == 0:
		n := paramsIter.GetNextOrDefault(1)
		p.handler.EraseChars(int(n))

	case action == 'Z' && len(intermediates) == 0:
		n := paramsIter.GetNextOrDefault(1)
		p.handler.MoveBackwardTabs(int(n))

	default:
		log.Tracef("Unhandled CSI params=%v intermediates=%v ignore=%v action=%v", params, intermediates, ignore, action)
	}
}

func keyboardMode(param uint16) KeyboardMode {
	return KeyboardMode(param & 0b0001_1111)
}

func keyboardBehavior(param uint16) KeyboardModeBehavior {
	switch param {

	case 3:
		return KeyboardModeBehaviorDifference

	case 2:
		return KeyboardModeBehaviorUnion

	default:
		return KeyboardModeBehaviorReplace
	}
}

func terminalMode(num uint16, private bool) (TerminalMode, bool) {
	if private {
		switch num {
		case 1:
			return TerminalModeCursorKeys, true

		case 3:
			return TerminalModeColumnMode, true

		case 6:
			return TerminalModeOrigin, true

		case 7:
			return TerminalModeLineWrap, true

		case 12:
			return TerminalModeBlinkingCursor, true

		case 25:
			return TerminalModeShowCursor, true

		case 1000:
			return TerminalModeReportMouseClicks, true

		case 1002:
			return TerminalModeReportCellMouseMotion, true

		case 1003:
			return TerminalModeReportAllMouseMotion, true

		case 1004:
			return TerminalModeReportFocusInOut, true

		case 1005:
			return TerminalModeUTF8Mouse, true

		case 1006:
			return TerminalModeSGRMouse, true

		case 1007:
			return TerminalModeAlternateScroll, true

		case 1042:
			return TerminalModeUrgencyHints, true

		case 1049:
			return TerminalModeSwapScreenAndSetRestoreCursor, true

		case 2004:
			return TerminalModeBracketedPaste, true

		default:
			return TerminalMode(-1), false
		}
	} else {
		switch num {
		case 4:
			return TerminalModeInsert, true

		case 20:
			return TerminalModeLineFeedNewLine, true

		default:
			return TerminalMode(-1), false
		}
	}
}

func attr(attr CharAttribute) TerminalCharAttribute {
	return TerminalCharAttribute{
		Attr: attr,
	}
}

func attrWithNamedColor(attr CharAttribute, color NamedColor) TerminalCharAttribute {
	return TerminalCharAttribute{
		Attr:       attr,
		NamedColor: &color,
	}
}

func attrWithRGBColor(attr CharAttribute, r, g, b uint16) TerminalCharAttribute {
	return TerminalCharAttribute{
		Attr:     attr,
		RGBColor: &RGBColor{uint8(r), uint8(g), uint8(b)},
	}
}

func attrWithIndexedColor(attr CharAttribute, i uint16) TerminalCharAttribute {
	return TerminalCharAttribute{
		Attr:         attr,
		IndexedColor: &IndexedColor{byte(i)},
	}
}
