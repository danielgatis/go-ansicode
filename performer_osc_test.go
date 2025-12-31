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

		// OSC 7 ; URI BEL - Set working directory
		{name: "OSC 7 ; file://host/path BEL", args: args{params: [][]byte{[]byte("7"), []byte("file://localhost/home/user")}, bellTerminated: true}, want: want{mock: m("SetWorkingDirectory", "file://localhost/home/user")}},
		{name: "OSC 7 ; file://host/path ST", args: args{params: [][]byte{[]byte("7"), []byte("file://myhost/var/log")}}, want: want{mock: m("SetWorkingDirectory", "file://myhost/var/log")}},
		{name: "OSC 7 ; URI with semicolons BEL", args: args{params: [][]byte{[]byte("7"), []byte("file://host/path"), []byte("extra")}, bellTerminated: true}, want: want{mock: m("SetWorkingDirectory", "file://host/path;extra")}},
		{name: "OSC 7 (no URI)", args: args{params: [][]byte{[]byte("7")}, bellTerminated: true}, want: want{mock: &handlerMock{}}},

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

		// OSC 133 - Shell Integration
		{name: "OSC 133 ; A BEL - Prompt start", args: args{params: [][]byte{[]byte("133"), []byte("A")}, bellTerminated: true}, want: want{mock: m("ShellIntegrationMark", PromptStart, -1)}},
		{name: "OSC 133 ; B BEL - Command start", args: args{params: [][]byte{[]byte("133"), []byte("B")}, bellTerminated: true}, want: want{mock: m("ShellIntegrationMark", CommandStart, -1)}},
		{name: "OSC 133 ; C BEL - Command executed", args: args{params: [][]byte{[]byte("133"), []byte("C")}, bellTerminated: true}, want: want{mock: m("ShellIntegrationMark", CommandExecuted, -1)}},
		{name: "OSC 133 ; D BEL - Command finished (no exit code)", args: args{params: [][]byte{[]byte("133"), []byte("D")}, bellTerminated: true}, want: want{mock: m("ShellIntegrationMark", CommandFinished, -1)}},
		{name: "OSC 133 ; D ; 0 BEL - Command finished (exit code 0)", args: args{params: [][]byte{[]byte("133"), []byte("D"), []byte("0")}, bellTerminated: true}, want: want{mock: m("ShellIntegrationMark", CommandFinished, 0)}},
		{name: "OSC 133 ; D ; 1 BEL - Command finished (exit code 1)", args: args{params: [][]byte{[]byte("133"), []byte("D"), []byte("1")}, bellTerminated: true}, want: want{mock: m("ShellIntegrationMark", CommandFinished, 1)}},
		{name: "OSC 133 ; D ; 127 BEL - Command finished (exit code 127)", args: args{params: [][]byte{[]byte("133"), []byte("D"), []byte("127")}, bellTerminated: true}, want: want{mock: m("ShellIntegrationMark", CommandFinished, 127)}},
		{name: "OSC 133 (no params)", args: args{params: [][]byte{[]byte("133")}, bellTerminated: true}, want: want{mock: &handlerMock{}}},

		// OSC 99 - Desktop Notifications (Kitty protocol)
		{name: "OSC 99 ; ; body BEL - simple notification", args: args{params: [][]byte{[]byte("99"), []byte(""), []byte("Hello World")}, bellTerminated: true}, want: want{mock: m("DesktopNotification", &NotificationPayload{Done: true, Urgency: 1, Timeout: -1, Data: []byte("Hello World")})}},
		{name: "OSC 99 ; i=test ; body BEL - with ID", args: args{params: [][]byte{[]byte("99"), []byte("i=test"), []byte("Hello")}, bellTerminated: true}, want: want{mock: m("DesktopNotification", &NotificationPayload{ID: "test", Done: true, Urgency: 1, Timeout: -1, Data: []byte("Hello")})}},
		{name: "OSC 99 ; i=1:p=title ; My Title BEL", args: args{params: [][]byte{[]byte("99"), []byte("i=1:p=title"), []byte("My Title")}, bellTerminated: true}, want: want{mock: m("DesktopNotification", &NotificationPayload{ID: "1", PayloadType: "title", Done: true, Urgency: 1, Timeout: -1, Data: []byte("My Title")})}},
		{name: "OSC 99 ; i=1:p=body:d=1 ; body text BEL", args: args{params: [][]byte{[]byte("99"), []byte("i=1:p=body:d=1"), []byte("body text")}, bellTerminated: true}, want: want{mock: m("DesktopNotification", &NotificationPayload{ID: "1", PayloadType: "body", Done: true, Urgency: 1, Timeout: -1, Data: []byte("body text")})}},
		{name: "OSC 99 ; i=1:d=0 ; chunk1 BEL - incomplete chunk", args: args{params: [][]byte{[]byte("99"), []byte("i=1:d=0"), []byte("chunk1")}, bellTerminated: true}, want: want{mock: m("DesktopNotification", &NotificationPayload{ID: "1", Done: false, Urgency: 1, Timeout: -1, Data: []byte("chunk1")})}},
		{name: "OSC 99 ; u=2 ; critical BEL - urgency", args: args{params: [][]byte{[]byte("99"), []byte("u=2"), []byte("critical")}, bellTerminated: true}, want: want{mock: m("DesktopNotification", &NotificationPayload{Done: true, Urgency: 2, Timeout: -1, Data: []byte("critical")})}},
		{name: "OSC 99 ; n=error ; error msg BEL - icon name", args: args{params: [][]byte{[]byte("99"), []byte("n=error"), []byte("error msg")}, bellTerminated: true}, want: want{mock: m("DesktopNotification", &NotificationPayload{Done: true, Urgency: 1, Timeout: -1, IconName: "error", Data: []byte("error msg")})}},
		{name: "OSC 99 ; s=silent ; quiet BEL - sound", args: args{params: [][]byte{[]byte("99"), []byte("s=silent"), []byte("quiet")}, bellTerminated: true}, want: want{mock: m("DesktopNotification", &NotificationPayload{Done: true, Urgency: 1, Timeout: -1, Sound: "silent", Data: []byte("quiet")})}},
		{name: "OSC 99 ; p=? ; BEL - query", args: args{params: [][]byte{[]byte("99"), []byte("i=q1:p=?"), []byte("")}, bellTerminated: true}, want: want{mock: m("DesktopNotification", &NotificationPayload{ID: "q1", PayloadType: "?", Done: true, Urgency: 1, Timeout: -1, Data: nil})}},
		{name: "OSC 99 ; c=1 ; track close BEL", args: args{params: [][]byte{[]byte("99"), []byte("c=1"), []byte("track me")}, bellTerminated: true}, want: want{mock: m("DesktopNotification", &NotificationPayload{Done: true, Urgency: 1, Timeout: -1, TrackClose: true, Data: []byte("track me")})}},
		{name: "OSC 99 ; w=5000 ; timeout BEL", args: args{params: [][]byte{[]byte("99"), []byte("w=5000"), []byte("expires")}, bellTerminated: true}, want: want{mock: m("DesktopNotification", &NotificationPayload{Done: true, Urgency: 1, Timeout: 5000, Data: []byte("expires")})}},
		{name: "OSC 99 ; a=focus,report ; actions BEL", args: args{params: [][]byte{[]byte("99"), []byte("a=focus,report"), []byte("click me")}, bellTerminated: true}, want: want{mock: m("DesktopNotification", &NotificationPayload{Done: true, Urgency: 1, Timeout: -1, Actions: []string{"focus", "report"}, Data: []byte("click me")})}},
		{name: "OSC 99 ; o=unfocused ; occasion BEL", args: args{params: [][]byte{[]byte("99"), []byte("o=unfocused"), []byte("when away")}, bellTerminated: true}, want: want{mock: m("DesktopNotification", &NotificationPayload{Done: true, Urgency: 1, Timeout: -1, Occasion: "unfocused", Data: []byte("when away")})}},

		// OSC 99 - Base64 encoding tests
		{name: "OSC 99 ; e=1 ; base64 payload BEL", args: args{params: [][]byte{[]byte("99"), []byte("e=1"), []byte("SGVsbG8gV29ybGQ=")}, bellTerminated: true}, want: want{mock: m("DesktopNotification", &NotificationPayload{Done: true, Urgency: 1, Timeout: -1, Encoding: "1", Data: []byte("Hello World")})}},
		{name: "OSC 99 ; e=1 ; base64 without padding BEL", args: args{params: [][]byte{[]byte("99"), []byte("e=1"), []byte("SGVsbG8")}, bellTerminated: true}, want: want{mock: m("DesktopNotification", &NotificationPayload{Done: true, Urgency: 1, Timeout: -1, Encoding: "1", Data: []byte("Hello")})}},
		{name: "OSC 99 ; f=base64appname ; appname BEL", args: args{params: [][]byte{[]byte("99"), []byte("f=TXlBcHA="), []byte("test")}, bellTerminated: true}, want: want{mock: m("DesktopNotification", &NotificationPayload{Done: true, Urgency: 1, Timeout: -1, AppName: "MyApp", Data: []byte("test")})}},
		{name: "OSC 99 ; t=base64type ; type BEL", args: args{params: [][]byte{[]byte("99"), []byte("t=YnVpbGQ="), []byte("test")}, bellTerminated: true}, want: want{mock: m("DesktopNotification", &NotificationPayload{Done: true, Urgency: 1, Timeout: -1, Type: "build", Data: []byte("test")})}},
		{name: "OSC 99 ; g=uuid ; icon cache BEL", args: args{params: [][]byte{[]byte("99"), []byte("g=550e8400-e29b-41d4-a716-446655440000"), []byte("test")}, bellTerminated: true}, want: want{mock: m("DesktopNotification", &NotificationPayload{Done: true, Urgency: 1, Timeout: -1, IconCacheID: "550e8400-e29b-41d4-a716-446655440000", Data: []byte("test")})}},

		// OSC 99 - Edge cases
		{name: "OSC 99 ; u=5 ; invalid urgency BEL", args: args{params: [][]byte{[]byte("99"), []byte("u=5"), []byte("test")}, bellTerminated: true}, want: want{mock: m("DesktopNotification", &NotificationPayload{Done: true, Urgency: 1, Timeout: -1, Data: []byte("test")})}},
		{name: "OSC 99 ; u=0 ; low urgency BEL", args: args{params: [][]byte{[]byte("99"), []byte("u=0"), []byte("low")}, bellTerminated: true}, want: want{mock: m("DesktopNotification", &NotificationPayload{Done: true, Urgency: 0, Timeout: -1, Data: []byte("low")})}},
		{name: "OSC 99 ; w=0 ; no timeout BEL", args: args{params: [][]byte{[]byte("99"), []byte("w=0"), []byte("never")}, bellTerminated: true}, want: want{mock: m("DesktopNotification", &NotificationPayload{Done: true, Urgency: 1, Timeout: 0, Data: []byte("never")})}},
		{name: "OSC 99 ; w=-1 ; OS default timeout BEL", args: args{params: [][]byte{[]byte("99"), []byte("w=-1"), []byte("os default")}, bellTerminated: true}, want: want{mock: m("DesktopNotification", &NotificationPayload{Done: true, Urgency: 1, Timeout: -1, Data: []byte("os default")})}},
		{name: "OSC 99 ; ; empty body BEL", args: args{params: [][]byte{[]byte("99"), []byte(""), []byte("")}, bellTerminated: true}, want: want{mock: m("DesktopNotification", &NotificationPayload{Done: true, Urgency: 1, Timeout: -1, Data: nil})}},
		{name: "OSC 99 no metadata or payload", args: args{params: [][]byte{[]byte("99")}, bellTerminated: true}, want: want{mock: m("DesktopNotification", &NotificationPayload{Done: true, Urgency: 1, Timeout: -1, Data: nil})}},

		// OSC 99 - Combined metadata
		{name: "OSC 99 full notification", args: args{params: [][]byte{[]byte("99"), []byte("i=notif1:p=body:u=2:n=warning:s=warn:a=focus:o=unfocused"), []byte("Important message")}, bellTerminated: true}, want: want{mock: m("DesktopNotification", &NotificationPayload{ID: "notif1", PayloadType: "body", Done: true, Urgency: 2, Timeout: -1, IconName: "warning", Sound: "warn", Actions: []string{"focus"}, Occasion: "unfocused", Data: []byte("Important message")})}},

		// OSC 99 - Close notification
		{name: "OSC 99 ; i=1:p=close ; close notification BEL", args: args{params: [][]byte{[]byte("99"), []byte("i=1:p=close"), []byte("")}, bellTerminated: true}, want: want{mock: m("DesktopNotification", &NotificationPayload{ID: "1", PayloadType: "close", Done: true, Urgency: 1, Timeout: -1, Data: nil})}},

		// OSC 99 - Alive query
		{name: "OSC 99 ; i=1:p=alive ; alive query BEL", args: args{params: [][]byte{[]byte("99"), []byte("i=1:p=alive"), []byte("")}, bellTerminated: true}, want: want{mock: m("DesktopNotification", &NotificationPayload{ID: "1", PayloadType: "alive", Done: true, Urgency: 1, Timeout: -1, Data: nil})}},

		// OSC 1337 - User variables (iTerm2/WezTerm)
		// dGVzdA== is base64 for "test"
		{name: "OSC 1337 ; SetUserVar=NAME=dGVzdA== BEL", args: args{params: [][]byte{[]byte("1337"), []byte("SetUserVar=MY_VAR=dGVzdA==")}, bellTerminated: true}, want: want{mock: m("SetUserVar", "MY_VAR", "test")}},
		{name: "OSC 1337 ; SetUserVar=NAME=dGVzdA== ST", args: args{params: [][]byte{[]byte("1337"), []byte("SetUserVar=MY_VAR=dGVzdA==")}}, want: want{mock: m("SetUserVar", "MY_VAR", "test")}},
		// aGVsbG8gd29ybGQ= is base64 for "hello world"
		{name: "OSC 1337 ; SetUserVar with spaces in value", args: args{params: [][]byte{[]byte("1337"), []byte("SetUserVar=GREETING=aGVsbG8gd29ybGQ=")}, bellTerminated: true}, want: want{mock: m("SetUserVar", "GREETING", "hello world")}},
		// Empty base64 value
		{name: "OSC 1337 ; SetUserVar with empty value", args: args{params: [][]byte{[]byte("1337"), []byte("SetUserVar=EMPTY=")}, bellTerminated: true}, want: want{mock: m("SetUserVar", "EMPTY", "")}},
		// Invalid base64 should not call handler
		{name: "OSC 1337 ; SetUserVar with invalid base64", args: args{params: [][]byte{[]byte("1337"), []byte("SetUserVar=TEST=!!invalid!!")}, bellTerminated: true}, want: want{mock: &handlerMock{}}},
		// Missing value part should not call handler
		{name: "OSC 1337 ; SetUserVar without value", args: args{params: [][]byte{[]byte("1337"), []byte("SetUserVar=NOVALUE")}, bellTerminated: true}, want: want{mock: &handlerMock{}}},
		// Unknown 1337 command should be ignored
		{name: "OSC 1337 ; UnknownCommand", args: args{params: [][]byte{[]byte("1337"), []byte("UnknownCommand=foo")}, bellTerminated: true}, want: want{mock: &handlerMock{}}},
		// Only 1337 code without data
		{name: "OSC 1337 (no params)", args: args{params: [][]byte{[]byte("1337")}, bellTerminated: true}, want: want{mock: &handlerMock{}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := &handlerMock{}

			performer := NewPerformer(handler)
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

func TestParseOSC99(t *testing.T) {
	tests := []struct {
		name   string
		params [][]byte
		want   *NotificationPayload
	}{
		{
			name:   "empty params",
			params: [][]byte{[]byte("99")},
			want:   &NotificationPayload{Done: true, Urgency: 1, Timeout: -1},
		},
		{
			name:   "simple body",
			params: [][]byte{[]byte("99"), []byte(""), []byte("Hello")},
			want:   &NotificationPayload{Done: true, Urgency: 1, Timeout: -1, Data: []byte("Hello")},
		},
		{
			name:   "with ID and title",
			params: [][]byte{[]byte("99"), []byte("i=test:p=title"), []byte("My Title")},
			want:   &NotificationPayload{ID: "test", PayloadType: "title", Done: true, Urgency: 1, Timeout: -1, Data: []byte("My Title")},
		},
		{
			name:   "base64 encoded payload",
			params: [][]byte{[]byte("99"), []byte("e=1"), []byte("SGVsbG8gV29ybGQ=")},
			want:   &NotificationPayload{Done: true, Encoding: "1", Urgency: 1, Timeout: -1, Data: []byte("Hello World")},
		},
		{
			name:   "base64 without padding",
			params: [][]byte{[]byte("99"), []byte("e=1"), []byte("SGVsbG8")},
			want:   &NotificationPayload{Done: true, Encoding: "1", Urgency: 1, Timeout: -1, Data: []byte("Hello")},
		},
		{
			name:   "invalid base64 fallback to raw",
			params: [][]byte{[]byte("99"), []byte("e=1"), []byte("not-valid-base64!!!")},
			want:   &NotificationPayload{Done: true, Encoding: "1", Urgency: 1, Timeout: -1, Data: []byte("not-valid-base64!!!")},
		},
		{
			name:   "chunked notification",
			params: [][]byte{[]byte("99"), []byte("i=chunk1:d=0"), []byte("partial")},
			want:   &NotificationPayload{ID: "chunk1", Done: false, Urgency: 1, Timeout: -1, Data: []byte("partial")},
		},
		{
			name:   "all metadata fields",
			params: [][]byte{[]byte("99"), []byte("i=id1:d=1:p=body:e=0:a=focus,report:c=1:w=5000:f=TXlBcHA=:t=YnVpbGQ=:n=info:g=uuid123:s=system:u=2:o=unfocused"), []byte("test")},
			want: &NotificationPayload{
				ID:          "id1",
				Done:        true,
				PayloadType: "body",
				Encoding:    "0",
				Actions:     []string{"focus", "report"},
				TrackClose:  true,
				Timeout:     5000,
				AppName:     "MyApp",
				Type:        "build",
				IconName:    "info",
				IconCacheID: "uuid123",
				Sound:       "system",
				Urgency:     2,
				Occasion:    "unfocused",
				Data:        []byte("test"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := parseOSC99(tt.params)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestParseOSC99Metadata(t *testing.T) {
	tests := []struct {
		name string
		meta []byte
		want *NotificationPayload
	}{
		{
			name: "empty metadata",
			meta: []byte(""),
			want: &NotificationPayload{},
		},
		{
			name: "single key-value",
			meta: []byte("i=test"),
			want: &NotificationPayload{ID: "test"},
		},
		{
			name: "multiple key-values",
			meta: []byte("i=id1:p=title:u=2"),
			want: &NotificationPayload{ID: "id1", PayloadType: "title", Urgency: 2},
		},
		{
			name: "done flag true",
			meta: []byte("d=1"),
			want: &NotificationPayload{Done: true},
		},
		{
			name: "done flag false",
			meta: []byte("d=0"),
			want: &NotificationPayload{Done: false},
		},
		{
			name: "actions list",
			meta: []byte("a=focus,report,other"),
			want: &NotificationPayload{Actions: []string{"focus", "report", "other"}},
		},
		{
			name: "track close enabled",
			meta: []byte("c=1"),
			want: &NotificationPayload{TrackClose: true},
		},
		{
			name: "track close disabled",
			meta: []byte("c=0"),
			want: &NotificationPayload{TrackClose: false},
		},
		{
			name: "timeout value",
			meta: []byte("w=10000"),
			want: &NotificationPayload{Timeout: 10000},
		},
		{
			name: "invalid timeout",
			meta: []byte("w=invalid"),
			want: &NotificationPayload{},
		},
		{
			name: "base64 appname",
			meta: []byte("f=TXlBcHA="),
			want: &NotificationPayload{AppName: "MyApp"},
		},
		{
			name: "invalid base64 appname fallback",
			meta: []byte("f=not-base64!!!"),
			want: &NotificationPayload{AppName: "not-base64!!!"},
		},
		{
			name: "base64 type",
			meta: []byte("t=YnVpbGQ="),
			want: &NotificationPayload{Type: "build"},
		},
		{
			name: "urgency values",
			meta: []byte("u=0"),
			want: &NotificationPayload{Urgency: 0},
		},
		{
			name: "urgency out of range high",
			meta: []byte("u=5"),
			want: &NotificationPayload{},
		},
		{
			name: "urgency out of range negative",
			meta: []byte("u=-1"),
			want: &NotificationPayload{},
		},
		{
			name: "icon name",
			meta: []byte("n=warning"),
			want: &NotificationPayload{IconName: "warning"},
		},
		{
			name: "icon cache ID",
			meta: []byte("g=550e8400-e29b-41d4-a716-446655440000"),
			want: &NotificationPayload{IconCacheID: "550e8400-e29b-41d4-a716-446655440000"},
		},
		{
			name: "sound",
			meta: []byte("s=silent"),
			want: &NotificationPayload{Sound: "silent"},
		},
		{
			name: "occasion",
			meta: []byte("o=invisible"),
			want: &NotificationPayload{Occasion: "invisible"},
		},
		{
			name: "encoding",
			meta: []byte("e=1"),
			want: &NotificationPayload{Encoding: "1"},
		},
		{
			name: "payload type query",
			meta: []byte("p=?"),
			want: &NotificationPayload{PayloadType: "?"},
		},
		{
			name: "payload type close",
			meta: []byte("p=close"),
			want: &NotificationPayload{PayloadType: "close"},
		},
		{
			name: "payload type alive",
			meta: []byte("p=alive"),
			want: &NotificationPayload{PayloadType: "alive"},
		},
		{
			name: "payload type icon",
			meta: []byte("p=icon"),
			want: &NotificationPayload{PayloadType: "icon"},
		},
		{
			name: "payload type buttons",
			meta: []byte("p=buttons"),
			want: &NotificationPayload{PayloadType: "buttons"},
		},
		{
			name: "invalid key without value",
			meta: []byte("invalid"),
			want: &NotificationPayload{},
		},
		{
			name: "unknown key ignored",
			meta: []byte("unknown=value"),
			want: &NotificationPayload{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := &NotificationPayload{}
			parseOSC99Metadata(tt.meta, got)
			assert.Equal(t, tt.want, got)
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
