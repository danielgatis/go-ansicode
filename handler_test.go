package ansicode

import (
	"image/color"
)

var _ Handler = (*handlerMock)(nil)

type handlerMock struct {
	called []string
	args   []interface{}
}

func (h *handlerMock) Backspace() {
	h.called = append(h.called, "Backspace")
}

func (h *handlerMock) Bell() {
	h.called = append(h.called, "Bell")
}

func (h *handlerMock) CarriageReturn() {
	h.called = append(h.called, "CarriageReturn")
}

func (h *handlerMock) ClearLine(mode LineClearMode) {
	h.called = append(h.called, "ClearLine")
	h.args = append(h.args, mode)
}

func (h *handlerMock) ClearScreen(mode ClearMode) {
	h.called = append(h.called, "ClearScreen")
	h.args = append(h.args, mode)
}

func (h *handlerMock) ClearTabs(mode TabulationClearMode) {
	h.called = append(h.called, "ClearTabs")
	h.args = append(h.args, mode)
}

func (h *handlerMock) ClipboardLoad(clp byte, t string) {
	h.called = append(h.called, "ClipboardLoad")
	h.args = append(h.args, clp)
	h.args = append(h.args, t)
}

func (h *handlerMock) ClipboardStore(clp byte, d []byte) {
	h.called = append(h.called, "ClipboardStore")
	h.args = append(h.args, clp)
	h.args = append(h.args, d)
}

func (h *handlerMock) ConfigureCharset(index CharsetIndex, charset Charset) {
	h.called = append(h.called, "ConfigureCharset")
	h.args = append(h.args, index)
	h.args = append(h.args, charset)
}

func (h *handlerMock) DeleteChars(n int) {
	h.called = append(h.called, "DeleteChars")
	h.args = append(h.args, n)
}

func (h *handlerMock) DeleteLines(n int) {
	h.called = append(h.called, "DeleteLines")
	h.args = append(h.args, n)
}

func (h *handlerMock) DeviceStatus(n int) {
	h.called = append(h.called, "DeviceStatus")
	h.args = append(h.args, n)
}

func (h *handlerMock) EraseChars(n int) {
	h.called = append(h.called, "EraseChars")
	h.args = append(h.args, n)
}

func (h *handlerMock) Goto(y int, x int) {
	h.called = append(h.called, "Goto")
	h.args = append(h.args, y)
	h.args = append(h.args, x)
}

func (h *handlerMock) GotoCol(n int) {
	h.called = append(h.called, "GotoCol")
	h.args = append(h.args, n)
}

func (h *handlerMock) GotoLine(n int) {
	h.called = append(h.called, "GotoLine")
	h.args = append(h.args, n)
}

func (h *handlerMock) IdentifyTerminal(b byte) {
	h.called = append(h.called, "IdentifyTerminal")
	h.args = append(h.args, b)
}

func (h *handlerMock) Input(r rune) {
	h.called = append(h.called, "Input")
	h.args = append(h.args, r)
}

func (h *handlerMock) InsertBlank(n int) {
	h.called = append(h.called, "InsertBlank")
	h.args = append(h.args, n)
}

func (h *handlerMock) InsertBlankLines(n int) {
	h.called = append(h.called, "InsertBlankLines")
	h.args = append(h.args, n)
}

func (h *handlerMock) LineFeed() {
	h.called = append(h.called, "LineFeed")
}

func (h *handlerMock) MoveBackward(n int) {
	h.called = append(h.called, "MoveBackward")
	h.args = append(h.args, n)
}

func (h *handlerMock) MoveBackwardTabs(n int) {
	h.called = append(h.called, "MoveBackwardTabs")
	h.args = append(h.args, n)
}

func (h *handlerMock) MoveDown(n int) {
	h.called = append(h.called, "MoveDown")
	h.args = append(h.args, n)
}

func (h *handlerMock) MoveDownCr(n int) {
	h.called = append(h.called, "MoveDownCr")
	h.args = append(h.args, n)
}

func (h *handlerMock) MoveForward(n int) {
	h.called = append(h.called, "MoveForward")
	h.args = append(h.args, n)
}

func (h *handlerMock) MoveForwardTabs(n int) {
	h.called = append(h.called, "MoveForwardTabs")
	h.args = append(h.args, n)
}

func (h *handlerMock) MoveUp(n int) {
	h.called = append(h.called, "MoveUp")
	h.args = append(h.args, n)
}

func (h *handlerMock) MoveUpCr(n int) {
	h.called = append(h.called, "MoveUpCr")
	h.args = append(h.args, n)
}

func (h *handlerMock) PopKeyboardMode(n int) {
	h.called = append(h.called, "PopKeyboardMode")
	h.args = append(h.args, n)
}

func (h *handlerMock) PopTitle() {
	h.called = append(h.called, "PopTitle")
}

func (h *handlerMock) PushKeyboardMode(mode KeyboardMode) {
	h.called = append(h.called, "PushKeyboardMode")
	h.args = append(h.args, mode)
}

func (h *handlerMock) PushTitle() {
	h.called = append(h.called, "PushTitle")
}

func (h *handlerMock) ReportKeyboardMode() {
	h.called = append(h.called, "ReportKeyboardMode")
}

func (h *handlerMock) ReportModifyOtherKeys() {
	h.called = append(h.called, "ReportModifyOtherKeys")
}

func (h *handlerMock) ResetColor(i int) {
	h.called = append([]string{}, "ResetColor")
	h.args = append([]interface{}{}, i)
}

func (h *handlerMock) RestoreCursorPosition() {
	h.called = append(h.called, "RestoreCursorPosition")
}

func (h *handlerMock) SaveCursorPosition() {
	h.called = append(h.called, "SaveCursorPosition")
}

func (h *handlerMock) ScrollDown(n int) {
	h.called = append(h.called, "ScrollDown")
	h.args = append(h.args, n)
}

func (h *handlerMock) ScrollUp(n int) {
	h.called = append(h.called, "ScrollUp")
	h.args = append(h.args, n)
}

func (h *handlerMock) SetActiveCharset(cs int) {
	h.called = append(h.called, "SetActiveCharset")
	h.args = append(h.args, cs)
}

func (h *handlerMock) SetColor(i int, c color.Color) {
	h.called = append(h.called, "SetColor")
	h.args = append(h.args, i)
	h.args = append(h.args, c)
}

func (h *handlerMock) SetCursorStyle(style CursorStyle) {
	h.called = append(h.called, "SetCursorStyle")
	h.args = append(h.args, style)
}

func (h *handlerMock) SetDynamicColor(p string, i int, t string) {
	h.called = append(h.called, "SetDynamicColor")
	h.args = append(h.args, p)
	h.args = append(h.args, i)
	h.args = append(h.args, t)
}

func (h *handlerMock) SetHyperlink(hyperlink *Hyperlink) {
	h.called = append(h.called, "SetHyperlink")
	h.args = append(h.args, hyperlink)
}

func (h *handlerMock) SetKeyboardMode(mode KeyboardMode, behavior KeyboardModeBehavior) {
	h.called = append(h.called, "SetKeyboardMode")
	h.args = append(h.args, mode)
	h.args = append(h.args, behavior)
}

func (h *handlerMock) SetMode(mode TerminalMode) {
	h.called = append(h.called, "SetMode")
	h.args = append(h.args, mode)
}

func (h *handlerMock) SetModifyOtherKeys(modify ModifyOtherKeys) {
	h.called = append(h.called, "SetModifyOtherKeys")
	h.args = append(h.args, modify)
}

func (h *handlerMock) SetScrollingRegion(top int, bottom int) {
	h.called = append(h.called, "SetScrollingRegion")
	h.args = append(h.args, top)
	h.args = append(h.args, bottom)
}

func (h *handlerMock) StartOfStringReceived(data []byte) {
	h.called = append(h.called, "StartOfStringReceived")
	h.args = append(h.args, data)
}

func (h *handlerMock) PrivacyMessageReceived(data []byte) {
	h.called = append(h.called, "PrivacyMessageReceived")
	h.args = append(h.args, data)
}

func (h *handlerMock) ApplicationCommandReceived(data []byte) {
	h.called = append(h.called, "ApplicationCommandReceived")
	h.args = append(h.args, data)
}

func (h *handlerMock) SetTerminalCharAttribute(attr TerminalCharAttribute) {
	h.called = append(h.called, "SetTerminalCharAttribute")
	h.args = append(h.args, attr)
}

func (h *handlerMock) SetTitle(title string) {
	h.called = append(h.called, "SetTitle")
	h.args = append(h.args, title)
}

func (h *handlerMock) Substitute() {
	h.called = append(h.called, "Substitute")
}

func (h *handlerMock) Tab(n int) {
	h.called = append(h.called, "Tab")
	h.args = append(h.args, n)
}

func (h *handlerMock) TextAreaSizeChars() {
	h.called = append(h.called, "TextAreaSizeChars")
}

func (h *handlerMock) TextAreaSizePixels() {
	h.called = append(h.called, "TextAreaSizePixels")
}

func (h *handlerMock) CellSizePixels() {
	h.called = append(h.called, "CellSizePixels")
}

func (h *handlerMock) UnsetMode(mode TerminalMode) {
	h.called = append(h.called, "UnsetMode")
	h.args = append(h.args, mode)
}

func (h *handlerMock) UnsetKeypadApplicationMode() {
	h.called = append(h.called, "UnsetKeypadApplicationMode")
}

func (h *handlerMock) SetKeypadApplicationMode() {
	h.called = append(h.called, "SetKeypadApplicationMode")
}

func (h *handlerMock) ReverseIndex() {
	h.called = append(h.called, "ReverseIndex")
}

func (h *handlerMock) ResetState() {
	h.called = append(h.called, "ResetState")
}

func (h *handlerMock) HorizontalTabSet() {
	h.called = append(h.called, "HorizontalTabSet")
}

func (h *handlerMock) Decaln() {
	h.called = append(h.called, "Decaln")
}

func (h *handlerMock) ShellIntegrationMark(mark ShellIntegrationMark, exitCode int) {
	h.called = append(h.called, "ShellIntegrationMark")
	h.args = append(h.args, mark)
	h.args = append(h.args, exitCode)
}

func (h *handlerMock) SetWorkingDirectory(uri string) {
	h.called = append(h.called, "SetWorkingDirectory")
	h.args = append(h.args, uri)
}

func (h *handlerMock) SixelReceived(params [][]uint16, data []byte) {
	h.called = append(h.called, "SixelReceived")
	h.args = append(h.args, params)
	h.args = append(h.args, data)
}

func (h *handlerMock) DesktopNotification(payload *NotificationPayload) {
	h.called = append(h.called, "DesktopNotification")
	h.args = append(h.args, payload)
}

func ms(called []string, args ...interface{}) *handlerMock {
	return &handlerMock{
		called: called,
		args:   args,
	}
}

func m(called string, args ...interface{}) *handlerMock {
	return &handlerMock{
		called: []string{called},
		args:   args,
	}
}
