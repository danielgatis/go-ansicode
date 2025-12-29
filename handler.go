package ansicode

import "image/color"

// ShellIntegrationMark represents a shell integration mark type (OSC 133).
type ShellIntegrationMark int

const (
	// PromptStart marks the beginning of the shell prompt (OSC 133 ; A).
	PromptStart ShellIntegrationMark = iota
	// CommandStart marks the end of prompt/start of user input (OSC 133 ; B).
	CommandStart
	// CommandExecuted marks when command execution begins (OSC 133 ; C).
	CommandExecuted
	// CommandFinished marks when command execution ends (OSC 133 ; D).
	CommandFinished
)

// NotificationPayload represents a desktop notification from OSC 99 (Kitty protocol).
// See: https://sw.kovidgoyal.net/kitty/desktop-notifications/
type NotificationPayload struct {
	// ID is the unique identifier for chunking and tracking notifications.
	ID string

	// Done indicates if this is the final chunk (default true).
	Done bool

	// PayloadType specifies what kind of data is in Data.
	// Values: "title", "body", "icon", "buttons", "close", "alive", "?" (query).
	PayloadType string

	// Encoding specifies the payload encoding ("1" = base64).
	Encoding string

	// Actions specifies click behavior: "focus", "report".
	Actions []string

	// TrackClose indicates if close events should be reported.
	TrackClose bool

	// Timeout is auto-expiry in milliseconds (-1 = OS default, 0 = never).
	Timeout int

	// AppName is the application name.
	AppName string

	// Type is the notification type for filtering.
	Type string

	// IconName is a standard icon: "error", "warning", "info", "question".
	IconName string

	// IconCacheID is a UUID for caching custom icons.
	IconCacheID string

	// Sound is the notification sound: "system", "silent", "error", "warn", "info".
	Sound string

	// Urgency is 0 (low), 1 (normal), or 2 (critical).
	Urgency int

	// Occasion is when to show: "always", "unfocused", "invisible".
	Occasion string

	// Data is the payload content (decoded if base64).
	Data []byte
}

// Handler is the interface that handles ANSI escape sequences.
type Handler interface {
	// ApplicationCommandReceived handles Application Program Command (APC) sequences (ESC _ ... ST).
	// Used by protocols like Kitty Graphics.
	ApplicationCommandReceived(data []byte)

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

	// DesktopNotification handles OSC 99 desktop notifications (Kitty protocol).
	// The payload contains all parsed notification data. The provider is responsible
	// for handling queries (PayloadType="?") and returning its supported capabilities.
	DesktopNotification(payload *NotificationPayload)

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

	// PrivacyMessageReceived handles Privacy Message (PM) sequences (ESC ^ ... ST).
	PrivacyMessageReceived(data []byte)

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

	// StartOfStringReceived handles Start of String (SOS) sequences (ESC X ... ST).
	StartOfStringReceived(data []byte)

	// SetTerminalCharAttribute sets the terminal char attribute.
	SetTerminalCharAttribute(attr TerminalCharAttribute)

	// SetTitle sets the window title.
	SetTitle(title string)

	// SetWorkingDirectory sets the current working directory (OSC 7).
	// uri is in the format "file://hostname/path/to/dir".
	SetWorkingDirectory(uri string)

	// ShellIntegrationMark handles shell integration marks (OSC 133).
	// mark indicates the type of mark (PromptStart, CommandStart, CommandExecuted, CommandFinished).
	// exitCode is only valid for CommandFinished marks (-1 if not provided).
	ShellIntegrationMark(mark ShellIntegrationMark, exitCode int)

	// SixelReceived is called when a complete Sixel image sequence is received.
	// params contains the DCS parameters before 'q' (e.g., aspect ratio settings).
	// data contains the Sixel image data.
	SixelReceived(params [][]uint16, data []byte)

	// Substitute replaces the character under the cursor.
	Substitute()

	// Tab moves the cursor to the next tab stop.
	Tab(n int)

	// TextAreaSizeChars reports the text area size in characters.
	TextAreaSizeChars()

	// TextAreaSizePixels reports the text area size in pixels.
	TextAreaSizePixels()

	// CellSizePixels reports the cell size in pixels.
	CellSizePixels()

	// UnsetKeypadApplicationMode sets the keypad to numeric mode.
	UnsetKeypadApplicationMode()

	// UnsetMode unsets the given mode.
	UnsetMode(mode TerminalMode)
}
