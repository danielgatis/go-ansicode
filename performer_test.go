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

func TestPerformer_SosPmApcDispatch(t *testing.T) {
	tests := []struct {
		name           string
		kind           byte
		data           []byte
		bellTerminated bool
	}{
		{
			name:           "SOS sequence",
			kind:           0, // SosKind
			data:           []byte("test sos data"),
			bellTerminated: false,
		},
		{
			name:           "PM sequence",
			kind:           1, // PmKind
			data:           []byte("test pm data"),
			bellTerminated: false,
		},
		{
			name:           "APC sequence",
			kind:           2, // ApcKind
			data:           []byte("test apc data"),
			bellTerminated: false,
		},
		{
			name:           "APC sequence bell terminated",
			kind:           2, // ApcKind
			data:           []byte("kitty graphics data"),
			bellTerminated: true,
		},
		{
			name:           "empty data",
			kind:           0,
			data:           []byte{},
			bellTerminated: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := NewPerformer(nil)
			// Should not panic even with nil handler
			assert.NotPanics(t, func() {
				p.SosPmApcDispatch(tt.kind, tt.data, tt.bellTerminated)
			})
		})
	}
}
