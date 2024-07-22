package ansicode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
		{name: "ESC ( B", args: args{intermediates: []byte{'('}, b: 'B'}, want: want{mock: m("ConfigureCharset", CharsetIndexG0, CharsetASCII)}},

		// ESC ) C
		{name: "ESC ) B", args: args{intermediates: []byte{')'}, b: 'B'}, want: want{mock: m("ConfigureCharset", CharsetIndexG1, CharsetASCII)}},

		// ESC * C
		{name: "ESC * B", args: args{intermediates: []byte{'*'}, b: 'B'}, want: want{mock: m("ConfigureCharset", CharsetIndexG2, CharsetASCII)}},

		// ESC + C
		{name: "ESC + B", args: args{intermediates: []byte{'+'}, b: 'B'}, want: want{mock: m("ConfigureCharset", CharsetIndexG3, CharsetASCII)}},

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
		{name: "ESC ( 0", args: args{intermediates: []byte{'('}, b: '0'}, want: want{mock: m("ConfigureCharset", CharsetIndexG0, CharsetLineDrawing)}},

		// ESC ) 0
		{name: "ESC ) 0", args: args{intermediates: []byte{')'}, b: '0'}, want: want{mock: m("ConfigureCharset", CharsetIndexG1, CharsetLineDrawing)}},

		// ESC * 0
		{name: "ESC * 0", args: args{intermediates: []byte{'*'}, b: '0'}, want: want{mock: m("ConfigureCharset", CharsetIndexG2, CharsetLineDrawing)}},

		// ESC + 0
		{name: "ESC + 0", args: args{intermediates: []byte{'+'}, b: '0'}, want: want{mock: m("ConfigureCharset", CharsetIndexG3, CharsetLineDrawing)}},

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
