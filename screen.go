package tea

import "github.com/charmbracelet/x/exp/term/input"

// WindowSizeMsg is used to report the terminal size. It's sent to Update once
// initially and then on every terminal resize. Note that Windows does not
// have support for reporting when resizes occur as it does not support the
// SIGWINCH signal.
type WindowSizeMsg = input.WindowSizeEvent

// BackgroundColorMsg is used to report the terminal's background color. It's
// sent to Update once initially and whenever the background color is queried
// using RequestBackgroundColor Cmd.
type BackgroundColorMsg = input.BackgroundColorEvent

// ForegroundColorMsg is used to report the terminal's foreground color.
type ForegroundColorMsg = input.ForegroundColorEvent

// CursorColorMsg is used to report the terminal's cursor color.
type CursorColorMsg = input.CursorColorEvent

// EnhancedKeyboardMsg is used to report the terminal's keyboard enhancement mode
// flags.
type EnhancedKeyboardMsg = input.KittyKeyboardEvent

// PrimaryDeviceAttributesMsg is used to report the terminal's primary device
// attributes.
type PrimaryDeviceAttributesMsg = input.PrimaryDeviceAttributesEvent

// ClearScreen is a special command that tells the program to clear the screen
// before the next update. This can be used to move the cursor to the top left
// of the screen and clear visual clutter when the alt screen is not in use.
//
// Note that it should never be necessary to call ClearScreen() for regular
// redraws.
func ClearScreen() Msg {
	return clearScreenMsg{}
}

// clearScreenMsg is an internal message that signals to clear the screen.
// You can send a clearScreenMsg with ClearScreen.
type clearScreenMsg struct{}

// EnterAltScreen is a special command that tells the Bubble Tea program to
// enter the alternate screen buffer.
//
// Because commands run asynchronously, this command should not be used in your
// model's Init function. To initialize your program with the altscreen enabled
// use the WithAltScreen ProgramOption instead.
func EnterAltScreen() Msg {
	return enterAltScreenMsg{}
}

// enterAltScreenMsg in an internal message signals that the program should
// enter alternate screen buffer. You can send a enterAltScreenMsg with
// EnterAltScreen.
type enterAltScreenMsg struct{}

// ExitAltScreen is a special command that tells the Bubble Tea program to exit
// the alternate screen buffer. This command should be used to exit the
// alternate screen buffer while the program is running.
//
// Note that the alternate screen buffer will be automatically exited when the
// program quits.
func ExitAltScreen() Msg {
	return exitAltScreenMsg{}
}

// exitAltScreenMsg in an internal message signals that the program should exit
// alternate screen buffer. You can send a exitAltScreenMsg with ExitAltScreen.
type exitAltScreenMsg struct{}

// EnableMouseCellMotion is a special command that enables mouse click,
// release, and wheel events. Mouse movement events are also captured if
// a mouse button is pressed (i.e., drag events).
//
// Because commands run asynchronously, this command should not be used in your
// model's Init function. Use the WithMouseCellMotion ProgramOption instead.
func EnableMouseCellMotion() Msg {
	return enableMouseCellMotionMsg{}
}

// enableMouseCellMotionMsg is a special command that signals to start
// listening for "cell motion" type mouse events (ESC[?1002l). To send an
// enableMouseCellMotionMsg, use the EnableMouseCellMotion command.
type enableMouseCellMotionMsg struct{}

// EnableMouseAllMotion is a special command that enables mouse click, release,
// wheel, and motion events, which are delivered regardless of whether a mouse
// button is pressed, effectively enabling support for hover interactions.
//
// Many modern terminals support this, but not all. If in doubt, use
// EnableMouseCellMotion instead.
//
// Because commands run asynchronously, this command should not be used in your
// model's Init function. Use the WithMouseAllMotion ProgramOption instead.
func EnableMouseAllMotion() Msg {
	return enableMouseAllMotionMsg{}
}

// enableMouseAllMotionMsg is a special command that signals to start listening
// for "all motion" type mouse events (ESC[?1003l). To send an
// enableMouseAllMotionMsg, use the EnableMouseAllMotion command.
type enableMouseAllMotionMsg struct{}

// DisableMouse is a special command that stops listening for mouse events.
func DisableMouse() Msg {
	return disableMouseMsg{}
}

// disableMouseMsg is an internal message that signals to stop listening
// for mouse events. To send a disableMouseMsg, use the DisableMouse command.
type disableMouseMsg struct{}

// HideCursor is a special command for manually instructing Bubble Tea to hide
// the cursor. In some rare cases, certain operations will cause the terminal
// to show the cursor, which is normally hidden for the duration of a Bubble
// Tea program's lifetime. You will most likely not need to use this command.
func HideCursor() Msg {
	return hideCursorMsg{}
}

// hideCursorMsg is an internal command used to hide the cursor. You can send
// this message with HideCursor.
type hideCursorMsg struct{}

// ShowCursor is a special command for manually instructing Bubble Tea to show
// the cursor.
func ShowCursor() Msg {
	return showCursorMsg{}
}

// showCursorMsg is an internal command used to show the cursor. You can send
// this message with ShowCursor.
type showCursorMsg struct{}

// EnableBracketedPaste is a special command that tells the Bubble Tea program
// to accept bracketed paste input.
//
// Note that bracketed paste will be automatically disabled when the
// program quits.
func EnableBracketedPaste() Msg {
	return enableBracketedPasteMsg{}
}

// enableBracketedPasteMsg in an internal message signals that
// bracketed paste should be enabled. You can send an
// enableBracketedPasteMsg with EnableBracketedPaste.
type enableBracketedPasteMsg struct{}

// DisableBracketedPaste is a special command that tells the Bubble Tea program
// to accept bracketed paste input.
//
// Note that bracketed paste will be automatically disabled when the
// program quits.
func DisableBracketedPaste() Msg {
	return disableBracketedPasteMsg{}
}

// disableBracketedPasteMsg in an internal message signals that
// bracketed paste should be disabled. You can send an
// disableBracketedPasteMsg with DisableBracketedPaste.
type disableBracketedPasteMsg struct{}

// RequestBackgroundColor is a special command that requests the terminal's
// background color. The background color will be sent to the program in a
// BackgroundColorMsg.
func RequestBackgroundColor() Msg {
	return requestBackgroundColorMsg{}
}

// requestBackgroundColorMsg is an internal message that requests the terminal's
// background color. You can send a requestBackgroundColorMsg with
// RequestBackgroundColor.
type requestBackgroundColorMsg struct{}

// RequestForegroundColor is a special command that requests the terminal's
// foreground color. The foreground color will be sent to the program in a
// ForegroundColorMsg.
func RequestForegroundColor() Msg {
	return requestForegroundColorMsg{}
}

// requestForegroundColorMsg is an internal message that requests the terminal's
// foreground color. You can send a requestForegroundColorMsg with
// RequestForegroundColor.
type requestForegroundColorMsg struct{}

// RequestCursorColor is a special command that requests the terminal's cursor
// color. The cursor color will be sent to the program in a CursorColorMsg.
// This command is not supported by all terminals.
func RequestCursorColor() Msg {
	return requestCursorColorMsg{}
}

// requestCursorColorMsg is an internal message that requests the terminal's
// cursor color. You can send a requestCursorColorMsg with
// RequestCursorColor.
type requestCursorColorMsg struct{}

// EnableEnhancedKeyboard is a special command that enables keyboard
// enhancement mode. This mode enables the Kitty Keyboard protocol, which
// provides more accurate keyboard input. This protocol is *not)( supported
// by all terminals.
//
// https://sw.kovidgoyal.net/kitty/keyboard-protocol
func EnableEnhancedKeyboard() Msg {
	return enableEnhancedKeyboardMsg{}
}

// EnableCustomEnhancedKeyboard is a special command that enables keyboard
// enhancement mode with custom flags. This mode enables the Kitty Keyboard
// protocol, which provides more accurate keyboard input. This protocol is
// *not* supported by all terminals.
func EnableCustomEnhancedKeyboard(flags int) Msg {
	return enableEnhancedKeyboardMsg{flags}
}

// enableEnhancedKeyboardMsg is an internal message that enables keyboard
// enhancement mode. You can send an enableEnhancedKeyboardMsg with
// EnableKeyboardEnhancement.
type enableEnhancedKeyboardMsg struct {
	flags int // zero means the default flags
}

// DisableEnhancedKeyboard is a special command that disables keyboard
// enhancement mode. This command should be used to disable the keyboard
// enhancement mode while the program is running.
//
// Note that the keyboard enhancement mode will be automatically disabled when
// the program quits.
func DisableEnhancedKeyboard() Msg {
	return disableEnhancedKeyboardMsg{}
}

// disableEnhancedKeyboardMsg is an internal message that disables keyboard
// enhancement mode. You can send a disableEnhancedKeyboardMsg with
// DisableKeyboardEnhancement.
type disableEnhancedKeyboardMsg struct{}

// RequestEnhancedKeyboard is a special command that requests the terminal's
// keyboard enhancement mode flags. The keyboard enhancement mode will be sent
// to the program in a EnhancedKeyboardMsg.
func RequestEnhancedKeyboard() Msg {
	return requestEnhancedKeyboardMsg{}
}

// requestEnhancedKeyboardMsg is an internal message that requests the terminal's
// keyboard enhancement mode flags. You can send a requestEnhancedKeyboardMsg with
// RequestEnhancedKeyboard.
type requestEnhancedKeyboardMsg struct{}

// RequestPrimaryDeviceAttributes is a special command that requests the
// terminal's primary device attributes. The primary device attributes will be
// sent to the program in a PrimaryDeviceAttributesMsg.
func RequestPrimaryDeviceAttributes() Msg {
	return requestPrimaryDeviceAttributesMsg{}
}

// requestPrimaryDeviceAttributesMsg is an internal message that requests the
// terminal's primary device attributes. You can send a
// requestPrimaryDeviceAttributesMsg with RequestPrimaryDeviceAttributes.
type requestPrimaryDeviceAttributesMsg struct{}

// EnterAltScreen enters the alternate screen buffer, which consumes the entire
// terminal window. ExitAltScreen will return the terminal to its former state.
//
// Deprecated: Use the WithAltScreen ProgramOption instead.
func (p *Program) EnterAltScreen() {
	if p.renderer != nil {
		p.renderer.enterAltScreen()
	}
}

// ExitAltScreen exits the alternate screen buffer.
//
// Deprecated: The altscreen will exited automatically when the program exits.
func (p *Program) ExitAltScreen() {
	if p.renderer != nil {
		p.renderer.exitAltScreen()
	}
}

// EnableMouseCellMotion enables mouse click, release, wheel and motion events
// if a mouse button is pressed (i.e., drag events).
//
// Deprecated: Use the WithMouseCellMotion ProgramOption instead.
func (p *Program) EnableMouseCellMotion() {
	p.renderer.enableMouseCellMotion()
}

// DisableMouseCellMotion disables Mouse Cell Motion tracking. This will be
// called automatically when exiting a Bubble Tea program.
//
// Deprecated: The mouse will automatically be disabled when the program exits.
func (p *Program) DisableMouseCellMotion() {
	p.renderer.disableMouseCellMotion()
}

// EnableMouseAllMotion enables mouse click, release, wheel and motion events,
// regardless of whether a mouse button is pressed. Many modern terminals
// support this, but not all.
//
// Deprecated: Use the WithMouseAllMotion ProgramOption instead.
func (p *Program) EnableMouseAllMotion() {
	p.renderer.enableMouseAllMotion()
}

// DisableMouseAllMotion disables All Motion mouse tracking. This will be
// called automatically when exiting a Bubble Tea program.
//
// Deprecated: The mouse will automatically be disabled when the program exits.
func (p *Program) DisableMouseAllMotion() {
	p.renderer.disableMouseAllMotion()
}

// SetWindowTitle sets the terminal window title.
func (p *Program) SetWindowTitle(title string) {
	p.renderer.setWindowTitle(title)
}
