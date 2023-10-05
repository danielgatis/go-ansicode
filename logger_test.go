package ansicode

var _ Logger = (*loggerMock)(nil)

type loggerMock struct {
	called string
	args   []interface{}
}

func (m *loggerMock) Tracef(format string, args ...interface{}) {
	m.called = "Tracef"
	m.args = args
}
