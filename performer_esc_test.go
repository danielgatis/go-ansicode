package ansicode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStandardCharset_Char(t *testing.T) {
	tests := []struct {
		name string
		sc   StandardCharset
		c    rune
		want rune
	}{
		{
			name: "Ascii",
			sc:   StandardCharsetASCII,
			c:    'A',
			want: 'A',
		},
		{
			name: "SpecialCharacterAndLineDrawing",
			sc:   StandardCharsetSpecialCharacterAndLineDrawing,
			c:    '_',
			want: ' ',
		},
		{
			name: "SpecialCharacterAndLineDrawing",
			sc:   StandardCharsetSpecialCharacterAndLineDrawing,
			c:    '`',
			want: '◆',
		},
		{
			name: "SpecialCharacterAndLineDrawing",
			sc:   StandardCharsetSpecialCharacterAndLineDrawing,
			c:    'a',
			want: '▒',
		},
		{
			name: "SpecialCharacterAndLineDrawing",
			sc:   StandardCharsetSpecialCharacterAndLineDrawing,
			c:    'b',
			want: '␉',
		},
		{
			name: "SpecialCharacterAndLineDrawing",
			sc:   StandardCharsetSpecialCharacterAndLineDrawing,
			c:    'c',
			want: '␌',
		},
		{
			name: "SpecialCharacterAndLineDrawing",
			sc:   StandardCharsetSpecialCharacterAndLineDrawing,
			c:    'd',
			want: '␍',
		},
		{
			name: "SpecialCharacterAndLineDrawing",
			sc:   StandardCharsetSpecialCharacterAndLineDrawing,
			c:    'e',
			want: '␊',
		},
		{
			name: "SpecialCharacterAndLineDrawing",
			sc:   StandardCharsetSpecialCharacterAndLineDrawing,
			c:    'f',
			want: '°',
		},
		{
			name: "SpecialCharacterAndLineDrawing",
			sc:   StandardCharsetSpecialCharacterAndLineDrawing,
			c:    'g',
			want: '±',
		},
		{
			name: "SpecialCharacterAndLineDrawing",
			sc:   StandardCharsetSpecialCharacterAndLineDrawing,
			c:    'h',
			want: '␤',
		},
		{
			name: "SpecialCharacterAndLineDrawing",
			sc:   StandardCharsetSpecialCharacterAndLineDrawing,
			c:    'i',
			want: '␋',
		},
		{
			name: "SpecialCharacterAndLineDrawing",
			sc:   StandardCharsetSpecialCharacterAndLineDrawing,
			c:    'j',
			want: '┘',
		},
		{
			name: "SpecialCharacterAndLineDrawing",
			sc:   StandardCharsetSpecialCharacterAndLineDrawing,
			c:    'k',
			want: '┐',
		},
		{
			name: "SpecialCharacterAndLineDrawing",
			sc:   StandardCharsetSpecialCharacterAndLineDrawing,
			c:    'l',
			want: '┌',
		},
		{
			name: "SpecialCharacterAndLineDrawing",
			sc:   StandardCharsetSpecialCharacterAndLineDrawing,
			c:    'm',
			want: '└',
		},
		{
			name: "SpecialCharacterAndLineDrawing",
			sc:   StandardCharsetSpecialCharacterAndLineDrawing,
			c:    'n',
			want: '┼',
		},
		{
			name: "SpecialCharacterAndLineDrawing",
			sc:   StandardCharsetSpecialCharacterAndLineDrawing,
			c:    'o',
			want: '⎺',
		},
		{
			name: "SpecialCharacterAndLineDrawing",
			sc:   StandardCharsetSpecialCharacterAndLineDrawing,
			c:    'p',
			want: '⎻',
		},
		{
			name: "SpecialCharacterAndLineDrawing",
			sc:   StandardCharsetSpecialCharacterAndLineDrawing,
			c:    'q',
			want: '─',
		},
		{
			name: "SpecialCharacterAndLineDrawing",
			sc:   StandardCharsetSpecialCharacterAndLineDrawing,
			c:    'r',
			want: '⎼',
		},
		{
			name: "SpecialCharacterAndLineDrawing",
			sc:   StandardCharsetSpecialCharacterAndLineDrawing,
			c:    's',
			want: '⎽',
		},
		{
			name: "SpecialCharacterAndLineDrawing",
			sc:   StandardCharsetSpecialCharacterAndLineDrawing,
			c:    't',
			want: '├',
		},
		{
			name: "SpecialCharacterAndLineDrawing",
			sc:   StandardCharsetSpecialCharacterAndLineDrawing,
			c:    'u',
			want: '┤',
		},
		{
			name: "SpecialCharacterAndLineDrawing",
			sc:   StandardCharsetSpecialCharacterAndLineDrawing,
			c:    'v',
			want: '┴',
		},
		{
			name: "SpecialCharacterAndLineDrawing",
			sc:   StandardCharsetSpecialCharacterAndLineDrawing,
			c:    'w',
			want: '┬',
		},
		{
			name: "SpecialCharacterAndLineDrawing",
			sc:   StandardCharsetSpecialCharacterAndLineDrawing,
			c:    'x',
			want: '│',
		},
		{
			name: "SpecialCharacterAndLineDrawing",
			sc:   StandardCharsetSpecialCharacterAndLineDrawing,
			c:    'y',
			want: '≤',
		},
		{
			name: "SpecialCharacterAndLineDrawing",
			sc:   StandardCharsetSpecialCharacterAndLineDrawing,
			c:    'z',
			want: '≥',
		},
		{
			name: "SpecialCharacterAndLineDrawing",
			sc:   StandardCharsetSpecialCharacterAndLineDrawing,
			c:    '{',
			want: 'π',
		},
		{
			name: "SpecialCharacterAndLineDrawing",
			sc:   StandardCharsetSpecialCharacterAndLineDrawing,
			c:    '|',
			want: '≠',
		},
		{
			name: "SpecialCharacterAndLineDrawing",
			sc:   StandardCharsetSpecialCharacterAndLineDrawing,
			c:    '}',
			want: '£',
		},
		{
			name: "SpecialCharacterAndLineDrawing",
			sc:   StandardCharsetSpecialCharacterAndLineDrawing,
			c:    '~',
			want: '·',
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := StandardCharset(tt.sc)
			assert.Equal(t, tt.want, s.Char(tt.c))
		})
	}
}

func TestPerformer_EscDispatch(t *testing.T) {
	type args struct {
		intermediates []byte
		ignore        bool
		b             byte
	}

	type want struct {
		mock *handlerMock
	}

	tests := []struct {
		name string
		args args
		want want
	}{
		// ESC ( C
		{name: "ESC ( B", args: args{intermediates: []byte{'('}, b: 'B'}, want: want{mock: m("ConfigureCharset", CharsetIndexG0, StandardCharsetASCII)}},

		// ESC ) C
		{name: "ESC ) B", args: args{intermediates: []byte{')'}, b: 'B'}, want: want{mock: m("ConfigureCharset", CharsetIndexG1, StandardCharsetASCII)}},

		// ESC * C
		{name: "ESC * B", args: args{intermediates: []byte{'*'}, b: 'B'}, want: want{mock: m("ConfigureCharset", CharsetIndexG2, StandardCharsetASCII)}},

		// ESC + C
		{name: "ESC + B", args: args{intermediates: []byte{'+'}, b: 'B'}, want: want{mock: m("ConfigureCharset", CharsetIndexG3, StandardCharsetASCII)}},

		// ESC D
		{name: "ESC D", args: args{b: 'D'}, want: want{mock: m("LineFeed")}},

		// ESC E
		{name: "ESC E", args: args{b: 'E'}, want: want{mock: ms([]string{"LineFeed", "CarriageReturn"})}},

		// ESC H
		{name: "ESC H", args: args{b: 'H'}, want: want{mock: m("HorizontalTabSet")}},

		// ESC M
		{name: "ESC M", args: args{b: 'M'}, want: want{mock: m("ReverseIndex")}},

		// ESC Z
		{name: "ESC Z", args: args{b: 'Z'}, want: want{mock: m("IdentifyTerminal", byte(0))}},

		// ESC c
		{name: "ESC c", args: args{b: 'c'}, want: want{mock: m("ResetState")}},

		// ESC c
		{name: "ESC c", args: args{b: 'c'}, want: want{mock: m("ResetState")}},

		// ESC ( 0
		{name: "ESC ( 0", args: args{intermediates: []byte{'('}, b: '0'}, want: want{mock: m("ConfigureCharset", CharsetIndexG0, StandardCharsetSpecialCharacterAndLineDrawing)}},

		// ESC ) 0
		{name: "ESC ) 0", args: args{intermediates: []byte{')'}, b: '0'}, want: want{mock: m("ConfigureCharset", CharsetIndexG1, StandardCharsetSpecialCharacterAndLineDrawing)}},

		// ESC * 0
		{name: "ESC * 0", args: args{intermediates: []byte{'*'}, b: '0'}, want: want{mock: m("ConfigureCharset", CharsetIndexG2, StandardCharsetSpecialCharacterAndLineDrawing)}},

		// ESC + 0
		{name: "ESC + 0", args: args{intermediates: []byte{'+'}, b: '0'}, want: want{mock: m("ConfigureCharset", CharsetIndexG3, StandardCharsetSpecialCharacterAndLineDrawing)}},

		// ESC 7
		{name: "ESC 7", args: args{b: '7'}, want: want{mock: m("SaveCursorPosition")}},

		// ESC 8
		{name: "ESC 8", args: args{b: '8'}, want: want{mock: m("RestoreCursorPosition")}},

		// ESC # 8
		{name: "ESC # 8", args: args{intermediates: []byte{'#'}, b: '8'}, want: want{mock: m("Decaln")}},

		// ESC =
		{name: "ESC =", args: args{b: '='}, want: want{mock: m("SetKeypadApplicationMode")}},

		// ESC >
		{name: "ESC >", args: args{b: '>'}, want: want{mock: m("UnsetKeypadApplicationMode")}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := &handlerMock{}

			performer := NewPerformer(handler)
			performer.EscDispatch(tt.args.intermediates, tt.args.ignore, tt.args.b)

			assert.Equal(t, tt.want.mock, handler)
		})
	}
}
