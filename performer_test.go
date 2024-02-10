package ansicode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPerformer(t *testing.T) {
	type args struct {
		handler Handler
	}

	type want struct {
		performer *Performer
	}

	tests := []struct {
		name string
		args args
		want want
	}{
		{name: "new", args: args{handler: nil}, want: want{performer: &Performer{handler: nil}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewPerformer(tt.args.handler)
			assert.Equal(t, tt.want.performer, got)
		})
	}
}
