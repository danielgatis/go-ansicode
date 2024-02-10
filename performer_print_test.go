package ansicode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPerformer_Print(t *testing.T) {
	type args struct {
		r rune
	}

	type want struct {
		mock             *handlerMock
		precedingRune    rune
		hasPrecedingRune bool
	}

	tests := []struct {
		name string
		args args
		want want
	}{
		{name: "a", args: args{r: 'a'}, want: want{mock: m("Input", 'a'), precedingRune: 'a', hasPrecedingRune: true}},
		{name: "A", args: args{r: 'A'}, want: want{mock: m("Input", 'A'), precedingRune: 'A', hasPrecedingRune: true}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := &handlerMock{}

			performer := NewPerformer(handler)
			performer.Print(tt.args.r)

			assert.Equal(t, tt.want.mock, handler)
			assert.Equal(t, tt.want.precedingRune, performer.precedingRune)
			assert.Equal(t, tt.want.hasPrecedingRune, performer.hasPrecedingRune)
		})
	}
}
