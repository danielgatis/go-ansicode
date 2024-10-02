package ansicode

import "image/color"

// Handler is the interface that handles ANSI escape sequences.
type Handler interface {
	// Backspace moves the cursor one position to the left.
	Backspace()

	// Bell rings the bell.
	Bell()

	// CarriageReturn moves the cursor to the beginning of the line.
	CarriageReturn()

	// ClearLine clears the line.
	ClearLine(mode LineClearMode)

	// ClearScreen clears the screen.
	ClearScreen(mode ClearMode)

	// ClearTabs clears the tab stops.
	ClearTabs(mode TabulationClearMode)

	// ClipboardLoad loads data from the clipboard.
	ClipboardLoad(clipboard byte, terminator string)

	// ClipboardStore stores data in the clipboard.
	ClipboardStore(clipboard byte, data []byte)

	// ConfigureCharset configures the charset.
	ConfigureCharset(index CharsetIndex, charset Charset)

	// Dcaln runs the DECALN command.
	Decaln()

	// DeleteChars deletes n characters.
	DeleteChars(n int)

	// DeleteLines deletes n lines.
	DeleteLines(n int)

	// DeviceStatus reports the device status.
	DeviceStatus(n int)

	// EraseChars erases n characters.
	// erase means reset to the default state. (default colors, no content, no flags)
	EraseChars(n int)

	// Goto moves the cursor to the specified position.
	Goto(y int, x int)

	// GotoCol moves the cursor to the specified column.
	GotoCol(n int)

	// GotoLine moves the cursor to the specified line.
	GotoLine(n int)

	// HorizontalTab sets the current position as a tab stop.
	HorizontalTabSet()

	// IdentifyTerminal identifies the terminal.
	IdentifyTerminal(b byte)

	// Input inputs a rune to be displayed.
	Input(r rune)

	// InsertBlank inserts n blank characters.
	InsertBlank(n int)

	// InsertBlankLines inserts n blank lines.
	InsertBlankLines(n int)

	// LineFeed moves the cursor to the next line.
	LineFeed()

	// MoveBackward moves the cursor backward n columns.
	MoveBackward(n int)

	// MoveBackwardTabs moves the cursor backward n tab stops.
	MoveBackwardTabs(n int)

	// MoveDown moves the cursor down n lines.
	MoveDown(n int)

	// MoveDownCr moves the cursor down n lines and to the beginning of the line.
	MoveDownCr(n int)

	// MoveForward moves the cursor forward n columns.
	MoveForward(n int)

	// MoveForwardTabs moves the cursor forward n tab stops.
	MoveForwardTabs(n int)

	// MoveUp moves the cursor up n lines.
	MoveUp(n int)

	// MoveUpCr moves the cursor up n lines and to the beginning of the line.
	MoveUpCr(n int)

	// PopKeyboardMode pops the given amount n of keyboard modes from the stack.
	PopKeyboardMode(n int)

	// PopTitle pops the title from the stack.
	PopTitle()

	// PushKeyboardMode pushes the given keyboard mode to the stack.
	PushKeyboardMode(mode KeyboardMode)

	// PushTitle pushes the given title to the stack.
	PushTitle()

	// ReportKeyboardMode reports the keyboard mode.
	ReportKeyboardMode()

	// ReportModifyOtherKeys reports the modify other keys mode. (XTERM)
	ReportModifyOtherKeys()

	// ResetColor resets the color at the given index.
	ResetColor(i int)

	// ResetState resets the termnial state.
	ResetState()

	// RestoreCursorPosition restores the cursor position.
	RestoreCursorPosition()

	// ReverseIndex moves the active position to the same horizontal position on the preceding line.
	// If the active position is at the top margin, a scroll down is performed.
	ReverseIndex()

	// SaveCursorPosition saves the cursor position.
	SaveCursorPosition()

	// ScrollDown scrolls the screen down n lines.
	ScrollDown(n int)

	// ScrollUp scrolls the screen up n lines.
	ScrollUp(n int)

	// SetActiveCharset sets the active charset.
	SetActiveCharset(n int)

	// SetColor sets the color at the given index.
	SetColor(index int, color color.Color)

	// SetCursorStyle sets the cursor style.
	SetCursorStyle(style CursorStyle)

	// SetDynamicColor sets the dynamic color at the given index.
	SetDynamicColor(prefix string, index int, terminator string)

	// SetHyperlink sets the hyperlink.
	SetHyperlink(hyperlink *Hyperlink)

	// SetKeyboardMode sets the keyboard mode.
	SetKeyboardMode(mode KeyboardMode, behavior KeyboardModeBehavior)

	// SetKeypadApplicationMode sets keypad to applications mode.
	SetKeypadApplicationMode()

	// SetMode sets the given mode.
	SetMode(mode TerminalMode)

	// SetModifyOtherKeys sets the modify other keys mode. (XTERM)
	SetModifyOtherKeys(modify ModifyOtherKeys)

	// SetScrollingRegion sets the scrolling region.
	SetScrollingRegion(top int, bottom int)

	// SetTerminalCharAttribute sets the terminal char attribute.
	SetTerminalCharAttribute(attr TerminalCharAttribute)

	// SetTitle sets the window title.
	SetTitle(title string)

	// Substitue replaces the character under the cursor.
	Substitute()

	// Tab moves the cursor to the next tab stop.
	Tab(n int)

	// TextAreaSizeChars reports the text area size in characters.
	TextAreaSizeChars()

	// TextAreaSizePixels reports the text area size in pixels.
	TextAreaSizePixels()

	// UnsetKeypadApplicationMode sets the keypad to numeric mode.
	UnsetKeypadApplicationMode()

	// UnsetMode unsets the given mode.
	UnsetMode(mode TerminalMode)
}
