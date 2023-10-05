package ansicode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPerformer_Execute(t *testing.T) {
	type args struct {
		r byte
	}

	type want struct {
		mock *handlerMock
	}

	tests := []struct {
		name string
		args args
		want want
	}{
		{name: "HT", args: args{r: byte(c0HT)}, want: want{mock: m("Tab", 1)}},
		{name: "BS", args: args{r: byte(c0BS)}, want: want{mock: m("Backspace")}},
		{name: "CR", args: args{r: byte(c0CR)}, want: want{mock: m("CarriageReturn")}},
		{name: "LF", args: args{r: byte(c0LF)}, want: want{mock: m("LineFeed")}},
		{name: "VT", args: args{r: byte(c0VT)}, want: want{mock: m("LineFeed")}},
		{name: "FF", args: args{r: byte(c0FF)}, want: want{mock: m("LineFeed")}},
		{name: "BEL", args: args{r: byte(c0BEL)}, want: want{mock: m("Bell")}},
		{name: "SUB", args: args{r: byte(c0SUB)}, want: want{mock: m("Substitute")}},
		{name: "SO", args: args{r: byte(c0SO)}, want: want{mock: m("SetActiveCharset", 1)}},
		{name: "SI", args: args{r: byte(c0SI)}, want: want{mock: m("SetActiveCharset", 0)}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := &handlerMock{}
			logger := &loggerMock{}

			performer := NewPerformer(handler, logger)
			performer.Execute(tt.args.r)

			assert.Equal(t, tt.want.mock, handler)
		})
	}
}
