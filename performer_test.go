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
		wantCalled     string
	}{
		{
			name:           "SOS sequence",
			kind:           0,
			data:           []byte("test sos data"),
			bellTerminated: false,
			wantCalled:     "StartOfStringReceived",
		},
		{
			name:           "PM sequence",
			kind:           1,
			data:           []byte("test pm data"),
			bellTerminated: false,
			wantCalled:     "PrivacyMessageReceived",
		},
		{
			name:           "APC sequence",
			kind:           2,
			data:           []byte("test apc data"),
			bellTerminated: false,
			wantCalled:     "ApplicationCommandReceived",
		},
		{
			name:           "APC sequence bell terminated",
			kind:           2,
			data:           []byte("kitty graphics data"),
			bellTerminated: true,
			wantCalled:     "ApplicationCommandReceived",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &handlerMock{}
			p := NewPerformer(h)
			p.SosPmApcDispatch(tt.kind, tt.data, tt.bellTerminated)
			assert.Equal(t, []string{tt.wantCalled}, h.called)
			assert.Equal(t, tt.data, h.args[0])
		})
	}
}
