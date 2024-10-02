package ansicode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPerformer_CsiDispatch(t *testing.T) {
	type args struct {
		params        [][]uint16
		intermediates []byte
		ignore        bool
		action        rune
	}

	type want struct {
		mock *handlerMock
	}

	tests := []struct {
		name string
		args args
		want want
	}{
		{name: "ignore", args: args{ignore: true}, want: want{mock: &handlerMock{}}},
		{name: "intermediates gt 2", args: args{intermediates: []byte{0, 0}}, want: want{mock: &handlerMock{}}},

		// CSI Ps `
		{name: "CSI 3 `", args: args{params: [][]uint16{{3}}, action: '`'}, want: want{mock: m("GotoCol", 2)}},

		// CSI Ps @
		{name: "CSI 3 @", args: args{params: [][]uint16{{3}}, action: '@'}, want: want{mock: m("InsertBlank", 3)}},

		// CSI Ps A
		{name: "CSI 3 A", args: args{params: [][]uint16{{3}}, action: 'A'}, want: want{mock: m("MoveUp", 3)}},

		// CSI Ps B
		{name: "CSI 3 B", args: args{params: [][]uint16{{3}}, action: 'B'}, want: want{mock: m("MoveDown", 3)}},

		// CSI Ps b
		{name: "CSI 3 b", args: args{params: [][]uint16{{3}}, action: 'b'}, want: want{mock: ms([]string{"Input", "Input", "Input"}, 'X', 'X', 'X')}},

		// CSI Ps C
		{name: "CSI 3 C", args: args{params: [][]uint16{{3}}, action: 'C'}, want: want{mock: m("MoveForward", 3)}},

		// CSI Ps c
		{name: "CSI c", args: args{params: [][]uint16{{}}, intermediates: []byte{0}, action: 'c'}, want: want{mock: m("IdentifyTerminal", byte(0))}},

		// CSI Ps a
		{name: "CSI 3 a", args: args{params: [][]uint16{{3}}, action: 'a'}, want: want{mock: m("MoveForward", 3)}},

		// CSI Ps D
		{name: "CSI 3 D", args: args{params: [][]uint16{{3}}, action: 'D'}, want: want{mock: m("MoveBackward", 3)}},

		// CSI Ps d
		{name: "CSI 3 d", args: args{params: [][]uint16{{3}}, action: 'd'}, want: want{mock: m("GotoLine", 2)}},

		// CSI Ps E
		{name: "CSI 3 E", args: args{params: [][]uint16{{3}}, action: 'E'}, want: want{mock: m("MoveDownCr", 2)}},

		// CSI Ps e
		{name: "CSI 3 e", args: args{params: [][]uint16{{3}}, action: 'e'}, want: want{mock: m("MoveDown", 3)}},

		// CSI Ps F
		{name: "CSI 3 F", args: args{params: [][]uint16{{3}}, action: 'F'}, want: want{mock: m("MoveUpCr", 2)}},

		// CSI Ps G
		{name: "CSI 3 G", args: args{params: [][]uint16{{3}}, action: 'G'}, want: want{mock: m("GotoCol", 2)}},

		// CSI Ps g
		{name: "CSI 0 g", args: args{params: [][]uint16{{0}}, action: 'g'}, want: want{mock: m("ClearTabs", TabulationClearModeCurrent)}},
		{name: "CSI 3 g", args: args{params: [][]uint16{{3}}, action: 'g'}, want: want{mock: m("ClearTabs", TabulationClearModeAll)}},

		// CSI Ps ; Ps H
		{name: "CSI 5 ; 3 H", args: args{params: [][]uint16{{5}, {3}}, action: 'H'}, want: want{mock: m("Goto", 4, 2)}},

		// CSI Ps ; Ps f
		{name: "CSI 5 ; 3 f", args: args{params: [][]uint16{{5}, {3}}, action: 'f'}, want: want{mock: m("Goto", 4, 2)}},

		// CSI Ps I
		{name: "CSI 3 I", args: args{params: [][]uint16{{3}}, action: 'I'}, want: want{mock: m("MoveForwardTabs", 3)}},

		// CSI Ps J
		{name: "CSI 0 J", args: args{params: [][]uint16{{0}}, action: 'J'}, want: want{mock: m("ClearScreen", ClearModeBelow)}},
		{name: "CSI 1 J", args: args{params: [][]uint16{{1}}, action: 'J'}, want: want{mock: m("ClearScreen", ClearModeAbove)}},
		{name: "CSI 2 J", args: args{params: [][]uint16{{2}}, action: 'J'}, want: want{mock: m("ClearScreen", ClearModeAll)}},
		{name: "CSI 3 J", args: args{params: [][]uint16{{3}}, action: 'J'}, want: want{mock: m("ClearScreen", ClearModeSaved)}},

		// CSI Ps K
		{name: "CSI 0 K", args: args{params: [][]uint16{{0}}, action: 'K'}, want: want{mock: m("ClearLine", LineClearModeRight)}},
		{name: "CSI 1 K", args: args{params: [][]uint16{{1}}, action: 'K'}, want: want{mock: m("ClearLine", LineClearModeLeft)}},
		{name: "CSI 2 K", args: args{params: [][]uint16{{2}}, action: 'K'}, want: want{mock: m("ClearLine", LineClearModeAll)}},

		// CSI Ps L
		{name: "CSI 3 L", args: args{params: [][]uint16{{3}}, action: 'L'}, want: want{mock: m("InsertBlankLines", 3)}},

		// CSI Ps h
		{name: "CSI 4 h", args: args{params: [][]uint16{{4}}, action: 'h'}, want: want{mock: m("SetMode", TerminalModeInsert)}},
		{name: "CSI 20 h", args: args{params: [][]uint16{{20}}, action: 'h'}, want: want{mock: m("SetMode", TerminalModeLineFeedNewLine)}},

		// CSI ? Ps h
		{name: "CSI ? 1 h", args: args{params: [][]uint16{{1}}, intermediates: []byte{'?'}, action: 'h'}, want: want{mock: m("SetMode", TerminalModeCursorKeys)}},
		{name: "CSI ? 3 h", args: args{params: [][]uint16{{3}}, intermediates: []byte{'?'}, action: 'h'}, want: want{mock: m("SetMode", TerminalModeColumnMode)}},
		{name: "CSI ? 6 h", args: args{params: [][]uint16{{6}}, intermediates: []byte{'?'}, action: 'h'}, want: want{mock: m("SetMode", TerminalModeOrigin)}},
		{name: "CSI ? 7 h", args: args{params: [][]uint16{{7}}, intermediates: []byte{'?'}, action: 'h'}, want: want{mock: m("SetMode", TerminalModeLineWrap)}},
		{name: "CSI ? 12 h", args: args{params: [][]uint16{{12}}, intermediates: []byte{'?'}, action: 'h'}, want: want{mock: m("SetMode", TerminalModeBlinkingCursor)}},
		{name: "CSI ? 25 h", args: args{params: [][]uint16{{25}}, intermediates: []byte{'?'}, action: 'h'}, want: want{mock: m("SetMode", TerminalModeShowCursor)}},
		{name: "CSI ? 1000 h", args: args{params: [][]uint16{{1000}}, intermediates: []byte{'?'}, action: 'h'}, want: want{mock: m("SetMode", TerminalModeReportMouseClicks)}},
		{name: "CSI ? 1002 h", args: args{params: [][]uint16{{1002}}, intermediates: []byte{'?'}, action: 'h'}, want: want{mock: m("SetMode", TerminalModeReportCellMouseMotion)}},
		{name: "CSI ? 1003 h", args: args{params: [][]uint16{{1003}}, intermediates: []byte{'?'}, action: 'h'}, want: want{mock: m("SetMode", TerminalModeReportAllMouseMotion)}},
		{name: "CSI ? 1004 h", args: args{params: [][]uint16{{1004}}, intermediates: []byte{'?'}, action: 'h'}, want: want{mock: m("SetMode", TerminalModeReportFocusInOut)}},
		{name: "CSI ? 1005 h", args: args{params: [][]uint16{{1005}}, intermediates: []byte{'?'}, action: 'h'}, want: want{mock: m("SetMode", TerminalModeUTF8Mouse)}},
		{name: "CSI ? 1006 h", args: args{params: [][]uint16{{1006}}, intermediates: []byte{'?'}, action: 'h'}, want: want{mock: m("SetMode", TerminalModeSGRMouse)}},
		{name: "CSI ? 1007 h", args: args{params: [][]uint16{{1007}}, intermediates: []byte{'?'}, action: 'h'}, want: want{mock: m("SetMode", TerminalModeAlternateScroll)}},
		{name: "CSI ? 1042 h", args: args{params: [][]uint16{{1042}}, intermediates: []byte{'?'}, action: 'h'}, want: want{mock: m("SetMode", TerminalModeUrgencyHints)}},
		{name: "CSI ? 1049 h", args: args{params: [][]uint16{{1049}}, intermediates: []byte{'?'}, action: 'h'}, want: want{mock: m("SetMode", TerminalModeSwapScreenAndSetRestoreCursor)}},
		{name: "CSI ? 2004 h", args: args{params: [][]uint16{{2004}}, intermediates: []byte{'?'}, action: 'h'}, want: want{mock: m("SetMode", TerminalModeBracketedPaste)}},

		// CSI Ps l
		{name: "CSI 4 l", args: args{params: [][]uint16{{4}}, action: 'l'}, want: want{mock: m("UnsetMode", TerminalModeInsert)}},
		{name: "CSI 20 l", args: args{params: [][]uint16{{20}}, action: 'l'}, want: want{mock: m("UnsetMode", TerminalModeLineFeedNewLine)}},

		// CSI ? Ps l
		{name: "CSI ? 1 l", args: args{params: [][]uint16{{1}}, intermediates: []byte{'?'}, action: 'l'}, want: want{mock: m("UnsetMode", TerminalModeCursorKeys)}},
		{name: "CSI ? 3 l", args: args{params: [][]uint16{{3}}, intermediates: []byte{'?'}, action: 'l'}, want: want{mock: m("UnsetMode", TerminalModeColumnMode)}},
		{name: "CSI ? 6 l", args: args{params: [][]uint16{{6}}, intermediates: []byte{'?'}, action: 'l'}, want: want{mock: m("UnsetMode", TerminalModeOrigin)}},
		{name: "CSI ? 7 l", args: args{params: [][]uint16{{7}}, intermediates: []byte{'?'}, action: 'l'}, want: want{mock: m("UnsetMode", TerminalModeLineWrap)}},
		{name: "CSI ? 12 l", args: args{params: [][]uint16{{12}}, intermediates: []byte{'?'}, action: 'l'}, want: want{mock: m("UnsetMode", TerminalModeBlinkingCursor)}},
		{name: "CSI ? 25 l", args: args{params: [][]uint16{{25}}, intermediates: []byte{'?'}, action: 'l'}, want: want{mock: m("UnsetMode", TerminalModeShowCursor)}},
		{name: "CSI ? 1000 l", args: args{params: [][]uint16{{1000}}, intermediates: []byte{'?'}, action: 'l'}, want: want{mock: m("UnsetMode", TerminalModeReportMouseClicks)}},
		{name: "CSI ? 1002 l", args: args{params: [][]uint16{{1002}}, intermediates: []byte{'?'}, action: 'l'}, want: want{mock: m("UnsetMode", TerminalModeReportCellMouseMotion)}},
		{name: "CSI ? 1003 l", args: args{params: [][]uint16{{1003}}, intermediates: []byte{'?'}, action: 'l'}, want: want{mock: m("UnsetMode", TerminalModeReportAllMouseMotion)}},
		{name: "CSI ? 1004 l", args: args{params: [][]uint16{{1004}}, intermediates: []byte{'?'}, action: 'l'}, want: want{mock: m("UnsetMode", TerminalModeReportFocusInOut)}},
		{name: "CSI ? 1005 l", args: args{params: [][]uint16{{1005}}, intermediates: []byte{'?'}, action: 'l'}, want: want{mock: m("UnsetMode", TerminalModeUTF8Mouse)}},
		{name: "CSI ? 1006 l", args: args{params: [][]uint16{{1006}}, intermediates: []byte{'?'}, action: 'l'}, want: want{mock: m("UnsetMode", TerminalModeSGRMouse)}},
		{name: "CSI ? 1007 l", args: args{params: [][]uint16{{1007}}, intermediates: []byte{'?'}, action: 'l'}, want: want{mock: m("UnsetMode", TerminalModeAlternateScroll)}},
		{name: "CSI ? 1042 l", args: args{params: [][]uint16{{1042}}, intermediates: []byte{'?'}, action: 'l'}, want: want{mock: m("UnsetMode", TerminalModeUrgencyHints)}},
		{name: "CSI ? 1049 l", args: args{params: [][]uint16{{1049}}, intermediates: []byte{'?'}, action: 'l'}, want: want{mock: m("UnsetMode", TerminalModeSwapScreenAndSetRestoreCursor)}},
		{name: "CSI ? 2004 l", args: args{params: [][]uint16{{2004}}, intermediates: []byte{'?'}, action: 'l'}, want: want{mock: m("UnsetMode", TerminalModeBracketedPaste)}},

		// CSI Ps M
		{name: "CSI 3 M", args: args{params: [][]uint16{{3}}, action: 'M'}, want: want{mock: m("DeleteLines", 3)}},

		// CSI Pm m
		{name: "CSI 0 m", args: args{params: [][]uint16{{0}}, action: 'm'}, want: want{mock: m("SetTerminalCharAttribute", attr(CharAttributeReset))}},
		{name: "CSI 1 m", args: args{params: [][]uint16{{1}}, action: 'm'}, want: want{mock: m("SetTerminalCharAttribute", attr(CharAttributeBold))}},
		{name: "CSI 2 m", args: args{params: [][]uint16{{2}}, action: 'm'}, want: want{mock: m("SetTerminalCharAttribute", attr(CharAttributeDim))}},
		{name: "CSI 3 m", args: args{params: [][]uint16{{3}}, action: 'm'}, want: want{mock: m("SetTerminalCharAttribute", attr(CharAttributeItalic))}},
		{name: "CSI 4 ; 0 m", args: args{params: [][]uint16{{4}, {0}}, action: 'm'}, want: want{mock: m("SetTerminalCharAttribute", attr(CharAttributeCancelUnderline))}},
		{name: "CSI 4 ; 2 m", args: args{params: [][]uint16{{4}, {2}}, action: 'm'}, want: want{mock: m("SetTerminalCharAttribute", attr(CharAttributeDoubleUnderline))}},
		{name: "CSI 4 ; 3 m", args: args{params: [][]uint16{{4}, {3}}, action: 'm'}, want: want{mock: m("SetTerminalCharAttribute", attr(CharAttributeCurlyUnderline))}},
		{name: "CSI 4 ; 4 m", args: args{params: [][]uint16{{4}, {4}}, action: 'm'}, want: want{mock: m("SetTerminalCharAttribute", attr(CharAttributeDottedUnderline))}},
		{name: "CSI 4 ; 5 m", args: args{params: [][]uint16{{4}, {5}}, action: 'm'}, want: want{mock: m("SetTerminalCharAttribute", attr(CharAttributeDashedUnderline))}},
		{name: "CSI 4 ; 6 m", args: args{params: [][]uint16{{4}, {6}}, action: 'm'}, want: want{mock: m("SetTerminalCharAttribute", attr(CharAttributeUnderline))}},
		{name: "CSI 5 m", args: args{params: [][]uint16{{5}}, action: 'm'}, want: want{mock: m("SetTerminalCharAttribute", attr(CharAttributeBlinkSlow))}},
		{name: "CSI 6 m", args: args{params: [][]uint16{{6}}, action: 'm'}, want: want{mock: m("SetTerminalCharAttribute", attr(CharAttributeBlinkFast))}},
		{name: "CSI 7 m", args: args{params: [][]uint16{{7}}, action: 'm'}, want: want{mock: m("SetTerminalCharAttribute", attr(CharAttributeReverse))}},
		{name: "CSI 8 m", args: args{params: [][]uint16{{8}}, action: 'm'}, want: want{mock: m("SetTerminalCharAttribute", attr(CharAttributeHidden))}},
		{name: "CSI 9 m", args: args{params: [][]uint16{{9}}, action: 'm'}, want: want{mock: m("SetTerminalCharAttribute", attr(CharAttributeStrike))}},
		{name: "CSI 21 m", args: args{params: [][]uint16{{21}}, action: 'm'}, want: want{mock: m("SetTerminalCharAttribute", attr(CharAttributeCancelBold))}},
		{name: "CSI 22 m", args: args{params: [][]uint16{{22}}, action: 'm'}, want: want{mock: m("SetTerminalCharAttribute", attr(CharAttributeCancelBoldDim))}},
		{name: "CSI 23 m", args: args{params: [][]uint16{{23}}, action: 'm'}, want: want{mock: m("SetTerminalCharAttribute", attr(CharAttributeCancelItalic))}},
		{name: "CSI 24 m", args: args{params: [][]uint16{{24}}, action: 'm'}, want: want{mock: m("SetTerminalCharAttribute", attr(CharAttributeCancelUnderline))}},
		{name: "CSI 25 m", args: args{params: [][]uint16{{25}}, action: 'm'}, want: want{mock: m("SetTerminalCharAttribute", attr(CharAttributeCancelBlink))}},
		{name: "CSI 27 m", args: args{params: [][]uint16{{27}}, action: 'm'}, want: want{mock: m("SetTerminalCharAttribute", attr(CharAttributeCancelReverse))}},
		{name: "CSI 28 m", args: args{params: [][]uint16{{28}}, action: 'm'}, want: want{mock: m("SetTerminalCharAttribute", attr(CharAttributeCancelHidden))}},
		{name: "CSI 29 m", args: args{params: [][]uint16{{29}}, action: 'm'}, want: want{mock: m("SetTerminalCharAttribute", attr(CharAttributeCancelStrike))}},
		{name: "CSI 30 m", args: args{params: [][]uint16{{30}}, action: 'm'}, want: want{mock: m("SetTerminalCharAttribute", attrWithNamedColor(CharAttributeForeground, NamedColorBlack))}},
		{name: "CSI 31 m", args: args{params: [][]uint16{{31}}, action: 'm'}, want: want{mock: m("SetTerminalCharAttribute", attrWithNamedColor(CharAttributeForeground, NamedColorRed))}},
		{name: "CSI 32 m", args: args{params: [][]uint16{{32}}, action: 'm'}, want: want{mock: m("SetTerminalCharAttribute", attrWithNamedColor(CharAttributeForeground, NamedColorGreen))}},
		{name: "CSI 33 m", args: args{params: [][]uint16{{33}}, action: 'm'}, want: want{mock: m("SetTerminalCharAttribute", attrWithNamedColor(CharAttributeForeground, NamedColorYellow))}},
		{name: "CSI 34 m", args: args{params: [][]uint16{{34}}, action: 'm'}, want: want{mock: m("SetTerminalCharAttribute", attrWithNamedColor(CharAttributeForeground, NamedColorBlue))}},
		{name: "CSI 35 m", args: args{params: [][]uint16{{35}}, action: 'm'}, want: want{mock: m("SetTerminalCharAttribute", attrWithNamedColor(CharAttributeForeground, NamedColorMagenta))}},
		{name: "CSI 36 m", args: args{params: [][]uint16{{36}}, action: 'm'}, want: want{mock: m("SetTerminalCharAttribute", attrWithNamedColor(CharAttributeForeground, NamedColorCyan))}},
		{name: "CSI 37 m", args: args{params: [][]uint16{{37}}, action: 'm'}, want: want{mock: m("SetTerminalCharAttribute", attrWithNamedColor(CharAttributeForeground, NamedColorWhite))}},
		{name: "CSI 38 ; 5 ; 31 m", args: args{params: [][]uint16{{38}, {5}, {31}}, action: 'm'}, want: want{mock: m("SetTerminalCharAttribute", attrWithIndexedColor(CharAttributeForeground, 31))}},
		{name: "CSI 38 ; 2 ; 3 ; 3 ; 3 m", args: args{params: [][]uint16{{38}, {2}, {3}, {3}, {3}}, action: 'm'}, want: want{mock: m("SetTerminalCharAttribute", attrWithRGBColor(CharAttributeForeground, 3, 3, 3))}},
		{name: "CSI 39 m", args: args{params: [][]uint16{{39}}, action: 'm'}, want: want{mock: m("SetTerminalCharAttribute", attrWithNamedColor(CharAttributeForeground, NamedColorForeground))}},
		{name: "CSI 40 m", args: args{params: [][]uint16{{40}}, action: 'm'}, want: want{mock: m("SetTerminalCharAttribute", attrWithNamedColor(CharAttributeBackground, NamedColorBlack))}},
		{name: "CSI 41 m", args: args{params: [][]uint16{{41}}, action: 'm'}, want: want{mock: m("SetTerminalCharAttribute", attrWithNamedColor(CharAttributeBackground, NamedColorRed))}},
		{name: "CSI 42 m", args: args{params: [][]uint16{{42}}, action: 'm'}, want: want{mock: m("SetTerminalCharAttribute", attrWithNamedColor(CharAttributeBackground, NamedColorGreen))}},
		{name: "CSI 43 m", args: args{params: [][]uint16{{43}}, action: 'm'}, want: want{mock: m("SetTerminalCharAttribute", attrWithNamedColor(CharAttributeBackground, NamedColorYellow))}},
		{name: "CSI 44 m", args: args{params: [][]uint16{{44}}, action: 'm'}, want: want{mock: m("SetTerminalCharAttribute", attrWithNamedColor(CharAttributeBackground, NamedColorBlue))}},
		{name: "CSI 45 m", args: args{params: [][]uint16{{45}}, action: 'm'}, want: want{mock: m("SetTerminalCharAttribute", attrWithNamedColor(CharAttributeBackground, NamedColorMagenta))}},
		{name: "CSI 46 m", args: args{params: [][]uint16{{46}}, action: 'm'}, want: want{mock: m("SetTerminalCharAttribute", attrWithNamedColor(CharAttributeBackground, NamedColorCyan))}},
		{name: "CSI 47 m", args: args{params: [][]uint16{{47}}, action: 'm'}, want: want{mock: m("SetTerminalCharAttribute", attrWithNamedColor(CharAttributeBackground, NamedColorWhite))}},
		{name: "CSI 48 ; 5 ; 31 m", args: args{params: [][]uint16{{48}, {5}, {31}}, action: 'm'}, want: want{mock: m("SetTerminalCharAttribute", attrWithIndexedColor(CharAttributeBackground, 31))}},
		{name: "CSI 48 ; 2 ; 3 ; 3 ; 3 m", args: args{params: [][]uint16{{48}, {2}, {3}, {3}, {3}}, action: 'm'}, want: want{mock: m("SetTerminalCharAttribute", attrWithRGBColor(CharAttributeBackground, 3, 3, 3))}},
		{name: "CSI 49 m", args: args{params: [][]uint16{{49}}, action: 'm'}, want: want{mock: m("SetTerminalCharAttribute", attrWithNamedColor(CharAttributeBackground, NamedColorBackground))}},
		{name: "CSI 58 ; 5 ; 31 m", args: args{params: [][]uint16{{58}, {5}, {31}}, action: 'm'}, want: want{mock: m("SetTerminalCharAttribute", attrWithIndexedColor(CharAttributeUnderlineColor, 31))}},
		{name: "CSI 58 ; 2 ; 3 ; 3 ; 3 m", args: args{params: [][]uint16{{58}, {2}, {3}, {3}, {3}}, action: 'm'}, want: want{mock: m("SetTerminalCharAttribute", attrWithRGBColor(CharAttributeUnderlineColor, 3, 3, 3))}},
		{name: "CSI 59 m", args: args{params: [][]uint16{{59}}, action: 'm'}, want: want{mock: m("SetTerminalCharAttribute", attr(CharAttributeUnderlineColor))}},
		{name: "CSI 90 m", args: args{params: [][]uint16{{90}}, action: 'm'}, want: want{mock: m("SetTerminalCharAttribute", attrWithNamedColor(CharAttributeForeground, NamedColorBrightBlack))}},
		{name: "CSI 91 m", args: args{params: [][]uint16{{91}}, action: 'm'}, want: want{mock: m("SetTerminalCharAttribute", attrWithNamedColor(CharAttributeForeground, NamedColorBrightRed))}},
		{name: "CSI 92 m", args: args{params: [][]uint16{{92}}, action: 'm'}, want: want{mock: m("SetTerminalCharAttribute", attrWithNamedColor(CharAttributeForeground, NamedColorBrightGreen))}},
		{name: "CSI 93 m", args: args{params: [][]uint16{{93}}, action: 'm'}, want: want{mock: m("SetTerminalCharAttribute", attrWithNamedColor(CharAttributeForeground, NamedColorBrightYellow))}},
		{name: "CSI 94 m", args: args{params: [][]uint16{{94}}, action: 'm'}, want: want{mock: m("SetTerminalCharAttribute", attrWithNamedColor(CharAttributeForeground, NamedColorBrightBlue))}},
		{name: "CSI 95 m", args: args{params: [][]uint16{{95}}, action: 'm'}, want: want{mock: m("SetTerminalCharAttribute", attrWithNamedColor(CharAttributeForeground, NamedColorBrightMagenta))}},
		{name: "CSI 96 m", args: args{params: [][]uint16{{96}}, action: 'm'}, want: want{mock: m("SetTerminalCharAttribute", attrWithNamedColor(CharAttributeForeground, NamedColorBrightCyan))}},
		{name: "CSI 97 m", args: args{params: [][]uint16{{97}}, action: 'm'}, want: want{mock: m("SetTerminalCharAttribute", attrWithNamedColor(CharAttributeForeground, NamedColorBrightWhite))}},
		{name: "CSI 100 m", args: args{params: [][]uint16{{100}}, action: 'm'}, want: want{mock: m("SetTerminalCharAttribute", attrWithNamedColor(CharAttributeBackground, NamedColorBrightBlack))}},
		{name: "CSI 101 m", args: args{params: [][]uint16{{101}}, action: 'm'}, want: want{mock: m("SetTerminalCharAttribute", attrWithNamedColor(CharAttributeBackground, NamedColorBrightRed))}},
		{name: "CSI 102 m", args: args{params: [][]uint16{{102}}, action: 'm'}, want: want{mock: m("SetTerminalCharAttribute", attrWithNamedColor(CharAttributeBackground, NamedColorBrightGreen))}},
		{name: "CSI 103 m", args: args{params: [][]uint16{{103}}, action: 'm'}, want: want{mock: m("SetTerminalCharAttribute", attrWithNamedColor(CharAttributeBackground, NamedColorBrightYellow))}},
		{name: "CSI 104 m", args: args{params: [][]uint16{{104}}, action: 'm'}, want: want{mock: m("SetTerminalCharAttribute", attrWithNamedColor(CharAttributeBackground, NamedColorBrightBlue))}},
		{name: "CSI 105 m", args: args{params: [][]uint16{{105}}, action: 'm'}, want: want{mock: m("SetTerminalCharAttribute", attrWithNamedColor(CharAttributeBackground, NamedColorBrightMagenta))}},
		{name: "CSI 106 m", args: args{params: [][]uint16{{106}}, action: 'm'}, want: want{mock: m("SetTerminalCharAttribute", attrWithNamedColor(CharAttributeBackground, NamedColorBrightCyan))}},
		{name: "CSI 107 m", args: args{params: [][]uint16{{107}}, action: 'm'}, want: want{mock: m("SetTerminalCharAttribute", attrWithNamedColor(CharAttributeBackground, NamedColorBrightWhite))}},

		// CSI > Pp ; Pv m
		{name: "CSI > 4 2 m", args: args{params: [][]uint16{{4}, {2}}, action: 'm', intermediates: []byte{'>'}}, want: want{mock: m("SetModifyOtherKeys", ModifyOtherKeysResetEnableAll)}},

		// CSI ? Pp m
		{name: "CSI ? 4 m", args: args{params: [][]uint16{{4}}, action: 'm', intermediates: []byte{'?'}}, want: want{mock: m("ReportModifyOtherKeys")}},

		// CSI Ps n
		{name: "CSI 3 n", args: args{params: [][]uint16{{3}}, action: 'n'}, want: want{mock: m("DeviceStatus", 3)}},

		// CSI Ps P
		{name: "CSI 3 P", args: args{params: [][]uint16{{3}}, action: 'P'}, want: want{mock: m("DeleteChars", 3)}},

		// CSI Ps SP q
		{name: "CSI 3 SP q", args: args{params: [][]uint16{{3}}, action: 'q', intermediates: []byte{' '}}, want: want{mock: m("SetCursorStyle", CursorStyleBlinkingUnderline)}},

		// CSI Ps ; Ps r
		{name: "CSI 3 ; 3 r", args: args{params: [][]uint16{{3}, {3}}, action: 'r'}, want: want{mock: m("SetScrollingRegion", 3, 3)}},

		// CSI Ps S
		{name: "CSI 3 S", args: args{params: [][]uint16{{3}}, action: 'S'}, want: want{mock: m("ScrollUp", 3)}},

		// CSI s
		{name: "CSI s", args: args{params: [][]uint16{{}}, action: 's'}, want: want{mock: m("SaveCursorPosition")}},

		// CSI Ps T
		{name: "CSI 3 T", args: args{params: [][]uint16{{3}}, action: 'T'}, want: want{mock: m("ScrollDown", 3)}},

		// CSI Ps t
		{name: "CSI 14 t", args: args{params: [][]uint16{{14}}, action: 't'}, want: want{mock: m("TextAreaSizePixels")}},
		{name: "CSI 18 t", args: args{params: [][]uint16{{18}}, action: 't'}, want: want{mock: m("TextAreaSizeChars")}},
		{name: "CSI 22 t", args: args{params: [][]uint16{{22}}, action: 't'}, want: want{mock: m("PushTitle")}},
		{name: "CSI 23 t", args: args{params: [][]uint16{{23}}, action: 't'}, want: want{mock: m("PopTitle")}},

		// CSI u
		{name: "CSI u", args: args{params: [][]uint16{{}}, action: 'u'}, want: want{mock: m("RestoreCursorPosition")}},

		// CSI ? u -- kitty keyboard protocol
		{name: "CSI ? u", args: args{params: [][]uint16{{}}, intermediates: []byte{'?'}, action: 'u'}, want: want{mock: m("ReportKeyboardMode")}},

		// CSI = Ps ; Ps u -- kitty keyboard protocol
		{name: "CSI = 2 ; 3 u", args: args{params: [][]uint16{{2}, {3}}, intermediates: []byte{'='}, action: 'u'}, want: want{mock: m("SetKeyboardMode", KeyboardModeReportEventTypes, KeyboardModeBehaviorDifference)}},

		// CSI > Ps u -- kitty keyboard protocol
		{name: "CSI > 2 u", args: args{params: [][]uint16{{2}}, intermediates: []byte{'>'}, action: 'u'}, want: want{mock: m("PushKeyboardMode", KeyboardModeReportEventTypes)}},

		// CSI < Ps u -- kitty keyboard protocol
		{name: "CSI < 2 u", args: args{params: [][]uint16{{2}}, intermediates: []byte{'<'}, action: 'u'}, want: want{mock: m("PopKeyboardMode", 2)}},

		// CSI Ps X
		{name: "CSI 3 X", args: args{params: [][]uint16{{3}}, action: 'X'}, want: want{mock: m("EraseChars", 3)}},

		// CSI Ps Z
		{name: "CSI 3 Z", args: args{params: [][]uint16{{3}}, action: 'Z'}, want: want{mock: m("MoveBackwardTabs", 3)}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := &handlerMock{}

			performer := NewPerformer(handler)
			performer.precedingRune = 'X'
			performer.hasPrecedingRune = true
			performer.CsiDispatch(tt.args.params, tt.args.intermediates, tt.args.ignore, tt.args.action)

			assert.Equal(t, tt.want.mock, handler)
		})
	}
}

func TestKeyboardMode(t *testing.T) {
	tests := []struct {
		name  string
		param uint16
		want  KeyboardMode
	}{
		{
			name:  "KeyboardMode 0",
			param: 0,
			want:  KeyboardMode(0),
		},
		{
			name:  "KeyboardMode 1",
			param: 1,
			want:  KeyboardMode(1),
		},
		{
			name:  "KeyboardMode 31",
			param: 31,
			want:  KeyboardMode(31),
		},
		{
			name:  "KeyboardMode 32",
			param: 32,
			want:  KeyboardMode(0),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := keyboardMode(tt.param)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestKeyboardBehavior(t *testing.T) {
	tests := []struct {
		name  string
		param uint16
		want  KeyboardModeBehavior
	}{
		{
			name:  "KeyboardModeBehaviorReplace",
			param: 0,
			want:  KeyboardModeBehaviorReplace,
		},
		{
			name:  "KeyboardModeBehaviorUnion",
			param: 2,
			want:  KeyboardModeBehaviorUnion,
		},
		{
			name:  "KeyboardModeBehaviorDifference",
			param: 3,
			want:  KeyboardModeBehaviorDifference,
		},
		{
			name:  "KeyboardModeBehaviorReplace (default)",
			param: 4,
			want:  KeyboardModeBehaviorReplace,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := keyboardBehavior(tt.param)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestTerminalMode(t *testing.T) {
	tests := []struct {
		name    string
		private bool
		num     uint16
		want    TerminalMode
		wantOk  bool
	}{
		{
			name:    "Private mode 1",
			private: true,
			num:     1,
			want:    TerminalModeCursorKeys,
			wantOk:  true,
		},
		{
			name:    "Private mode 3",
			private: true,
			num:     3,
			want:    TerminalModeColumnMode,
			wantOk:  true,
		},
		{
			name:    "Private mode 6",
			private: true,
			num:     6,
			want:    TerminalModeOrigin,
			wantOk:  true,
		},
		{
			name:    "Private mode 7",
			private: true,
			num:     7,
			want:    TerminalModeLineWrap,
			wantOk:  true,
		},
		{
			name:    "Private mode 12",
			private: true,
			num:     12,
			want:    TerminalModeBlinkingCursor,
			wantOk:  true,
		},
		{
			name:    "Private mode 25",
			private: true,
			num:     25,
			want:    TerminalModeShowCursor,
			wantOk:  true,
		},
		{
			name:    "Private mode 1000",
			private: true,
			num:     1000,
			want:    TerminalModeReportMouseClicks,
			wantOk:  true,
		},
		{
			name:    "Private mode 1002",
			private: true,
			num:     1002,
			want:    TerminalModeReportCellMouseMotion,
			wantOk:  true,
		},
		{
			name:    "Private mode 1003",
			private: true,
			num:     1003,
			want:    TerminalModeReportAllMouseMotion,
			wantOk:  true,
		},
		{
			name:    "Private mode 1004",
			private: true,
			num:     1004,
			want:    TerminalModeReportFocusInOut,
			wantOk:  true,
		},
		{
			name:    "Private mode 1005",
			private: true,
			num:     1005,
			want:    TerminalModeUTF8Mouse,
			wantOk:  true,
		},
		{
			name:    "Private mode 1006",
			private: true,
			num:     1006,
			want:    TerminalModeSGRMouse,
			wantOk:  true,
		},
		{
			name:    "Private mode 1007",
			private: true,
			num:     1007,
			want:    TerminalModeAlternateScroll,
			wantOk:  true,
		},
		{
			name:    "Private mode 1042",
			private: true,
			num:     1042,
			want:    TerminalModeUrgencyHints,
			wantOk:  true,
		},
		{
			name:    "Private mode 1049",
			private: true,
			num:     1049,
			want:    TerminalModeSwapScreenAndSetRestoreCursor,
			wantOk:  true,
		},
		{
			name:    "Private mode 2004",
			private: true,
			num:     2004,
			want:    TerminalModeBracketedPaste,
			wantOk:  true,
		},
		{
			name:    "Non-private mode 4",
			private: false,
			num:     4,
			want:    TerminalModeInsert,
			wantOk:  true,
		},
		{
			name:    "Non-private mode 20",
			private: false,
			num:     20,
			want:    TerminalModeLineFeedNewLine,
			wantOk:  true,
		},
		{
			name:    "Non-existent mode",
			private: true,
			num:     999,
			want:    TerminalMode(-1),
			wantOk:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotOk := terminalMode(tt.num, tt.private)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantOk, gotOk)
		})
	}
}
