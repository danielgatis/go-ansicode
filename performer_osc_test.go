package ansicode

import (
	"image/color"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPerformer_OscDispatch(t *testing.T) {
	type args struct {
		params         [][]byte
		bellTerminated bool
	}

	type want struct {
		mock *handlerMock
	}

	tests := []struct {
		name string
		args args
		want want
	}{
		{name: "no params", args: args{}, want: want{mock: &handlerMock{}}},

		// OSC 0 ; Pt BEL
		{name: "OSC 0 ; title BEL", args: args{params: [][]byte{[]byte("0"), []byte("title")}}, want: want{mock: m("SetTitle", "title")}},

		// OSC 2 ; Pt BEL
		{name: "OSC 2 ; title BEL", args: args{params: [][]byte{[]byte("2"), []byte("title")}}, want: want{mock: m("SetTitle", "title")}},
		{name: "OSC 2 ;", args: args{params: [][]byte{[]byte("2")}}, want: want{mock: &handlerMock{}}},

		// OSC 4 ; c ; spec BEL
		{name: "OSC 4 ; 0 ; rgb:bb/ee/ff BEL", args: args{params: [][]byte{[]byte("4"), []byte("0"), []byte("rgb:bb/ee/ff")}, bellTerminated: true}, want: want{mock: m("SetColor", 0, color.RGBA{0xbb, 0xee, 0xff, 0xff})}},
		{name: "OSC 4 ; 0 ; #bbeeff BEL", args: args{params: [][]byte{[]byte("4"), []byte("0"), []byte("#bbeeff")}, bellTerminated: true}, want: want{mock: m("SetColor", 0, color.RGBA{0xbb, 0xee, 0xff, 0xff})}},
		{name: "OSC 4 ; 0 ; #bbeeff ; 1 ; #aabbcc  BEL", args: args{params: [][]byte{[]byte("4"), []byte("0"), []byte("#bbeeff"), []byte("1"), []byte("#aabbcc")}, bellTerminated: true}, want: want{mock: ms([]string{"SetColor", "SetColor"}, 0, color.RGBA{0xbb, 0xee, 0xff, 0xff}, 1, color.RGBA{0xaa, 0xbb, 0xcc, 0xff})}},
		{name: "OSC 4 ; 0 ; ? BEL", args: args{params: [][]byte{[]byte("4"), []byte("0"), []byte("?")}, bellTerminated: true}, want: want{mock: m("SetDynamicColor", "4;0", 0, "\x07")}},
		{name: "OSC 4 ; 0 ; ? ST", args: args{params: [][]byte{[]byte("4"), []byte("0"), []byte("?")}}, want: want{mock: m("SetDynamicColor", "4;0", 0, "\x1b\\")}},

		// OSC 8 ; params ; uri BEL
		{name: "OSC 8 ; id=foo:abc=def ; http://foo.com BEL", args: args{params: [][]byte{[]byte("8"), []byte("id=foo:abc=def"), []byte("http://foo.com")}, bellTerminated: true}, want: want{mock: m("SetHyperlink", &Hyperlink{ID: "foo", URI: "http://foo.com"})}},
		{name: "OSC 8 ; id=foo:abc=def BEL", args: args{params: [][]byte{[]byte("8"), []byte("id=foo:abc=def")}, bellTerminated: true}, want: want{mock: &handlerMock{}}},

		// OSC 10 ; Ps BEL
		{name: "OSC 10 ; rgb:bb/ee/ff BEL", args: args{params: [][]byte{[]byte("10"), []byte("rgb:bb/ee/ff")}, bellTerminated: true}, want: want{mock: m("SetColor", int(NamedColorForeground), color.RGBA{0xbb, 0xee, 0xff, 0xff})}},
		{name: "OSC 10 ; ? BEL", args: args{params: [][]byte{[]byte("10"), []byte("?")}, bellTerminated: true}, want: want{mock: m("SetDynamicColor", "10", int(NamedColorForeground), "\x07")}},

		// OSC 11 ; Ps BEL
		{name: "OSC 11 ; rgb:bb/ee/ff BEL", args: args{params: [][]byte{[]byte("11"), []byte("rgb:bb/ee/ff")}, bellTerminated: true}, want: want{mock: m("SetColor", int(NamedColorBackground), color.RGBA{0xbb, 0xee, 0xff, 0xff})}},
		{name: "OSC 11 ; ? BEL", args: args{params: [][]byte{[]byte("11"), []byte("?")}, bellTerminated: true}, want: want{mock: m("SetDynamicColor", "11", int(NamedColorBackground), "\x07")}},

		// OSC 12 ; Ps BEL
		{name: "OSC 12 ; rgb:bb/ee/ff BEL", args: args{params: [][]byte{[]byte("12"), []byte("rgb:bb/ee/ff")}, bellTerminated: true}, want: want{mock: m("SetColor", int(NamedColorCursor), color.RGBA{0xbb, 0xee, 0xff, 0xff})}},
		{name: "OSC 12 ; ? BEL", args: args{params: [][]byte{[]byte("12"), []byte("?")}, bellTerminated: true}, want: want{mock: m("SetDynamicColor", "12", int(NamedColorCursor), "\x07")}},

		// OSC 104 ; c BEL
		{name: "OSC 104 BEL", args: args{params: [][]byte{[]byte("104")}}, want: want{mock: m("ResetColor", 255)}},
		{name: "OSC 104 ; 5 BEL", args: args{params: [][]byte{[]byte("104"), []byte("5")}}, want: want{mock: m("ResetColor", 5)}},
		{name: "OSC 104 ; 5 ; 6 BEL", args: args{params: [][]byte{[]byte("104"), []byte("5"), []byte("6")}}, want: want{mock: m("ResetColor", 6)}},

		// OSC 110 BEL
		{name: "OSC 110 BEL", args: args{params: [][]byte{[]byte("110")}, bellTerminated: true}, want: want{mock: m("ResetColor", int(NamedColorForeground))}},

		// OSC 111 BEL
		{name: "OSC 111 BEL", args: args{params: [][]byte{[]byte("111")}, bellTerminated: true}, want: want{mock: m("ResetColor", int(NamedColorBackground))}},

		// OSC 112 BEL
		{name: "OSC 112 BEL", args: args{params: [][]byte{[]byte("112")}, bellTerminated: true}, want: want{mock: m("ResetColor", int(NamedColorCursor))}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := &handlerMock{}
			logger := &loggerMock{}

			performer := NewPerformer(handler, logger)
			performer.OscDispatch(tt.args.params, tt.args.bellTerminated)

			assert.Equal(t, tt.want.mock, handler)
		})
	}
}

func TestParseXColor(t *testing.T) {
	tests := []struct {
		name  string
		bytes []byte
		want  color.Color
		ok    bool
	}{
		{
			name:  "empty bytes",
			bytes: []byte{},
			want:  color.RGBA{},
			ok:    false,
		},
		{
			name:  "invalid format",
			bytes: []byte("invalid"),
			want:  color.RGBA{},
			ok:    false,
		},
		{
			name:  "hex format 1",
			bytes: []byte("#f00"),
			want:  color.RGBA{R: 255, G: 0, B: 0, A: 255},
			ok:    true,
		},
		{
			name:  "hex format 2",
			bytes: []byte("#ff0000"),
			want:  color.RGBA{R: 255, G: 0, B: 0, A: 255},
			ok:    true,
		},
		{
			name:  "hex format 3",
			bytes: []byte("#fff000000"),
			want:  color.RGBA{R: 255, G: 0, B: 0, A: 255},
			ok:    true,
		},
		{
			name:  "hex format 4",
			bytes: []byte("#ffff00000000"),
			want:  color.RGBA{R: 255, G: 0, B: 0, A: 255},
			ok:    true,
		},
		{
			name:  "hex format invalid",
			bytes: []byte("#ffff0000zzzz"),
			want:  color.RGBA{},
			ok:    false,
		},
		{
			name:  "rgb format 1",
			bytes: []byte("rgb:0/A/F"),
			want:  color.RGBA{R: 0, G: 170, B: 255, A: 255},
			ok:    true,
		},
		{
			name:  "rgb format 1",
			bytes: []byte("rgb:FF/00/00"),
			want:  color.RGBA{R: 255, G: 0, B: 0, A: 255},
			ok:    true,
		},
		{
			name:  "rgb format 3",
			bytes: []byte("rgb:FFF/000/000"),
			want:  color.RGBA{R: 255, G: 0, B: 0, A: 255},
			ok:    true,
		},
		{
			name:  "rgb format 4",
			bytes: []byte("rgb:ffff/0000/0000"),
			want:  color.RGBA{R: 255, G: 0, B: 0, A: 255},
			ok:    true,
		},
		{
			name:  "rgb format with invalid color codes",
			bytes: []byte("rgb:ff/00/zz"),
			want:  color.RGBA{},
			ok:    false,
		},
		{
			name:  "rgb format with invalid number of color codes",
			bytes: []byte("rgb:ff/00"),
			want:  color.RGBA{},
			ok:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotOk := parseXColor(tt.bytes)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.ok, gotOk)
		})
	}
}

func TestParseNumber(t *testing.T) {
	tests := []struct {
		name  string
		bytes []byte
		want  int
		ok    bool
	}{
		{
			name:  "empty",
			bytes: []byte{},
			want:  0,
			ok:    false,
		},
		{
			name:  "invalid",
			bytes: []byte{'a', 'b', 'c'},
			want:  0,
			ok:    false,
		},
		{
			name:  "zero",
			bytes: []byte{'0'},
			want:  0,
			ok:    true,
		},
		{
			name:  "positive",
			bytes: []byte{'1', '2', '3'},
			want:  123,
			ok:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := parseNumber(tt.bytes)

			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.ok, ok)
		})
	}
}
