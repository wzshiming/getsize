// +build windows
package getsize

import (
	"golang.org/x/sys/windows"
)

// isTerminal returns whether the given file descriptor is a terminal.
func isTerminal(hStdin uintptr) bool {
	var st uint32
	err := windows.GetConsoleMode(windows.Handle(hStdin), &st)
	return err == nil
}

// getSize returns the visible dimensions of the given terminal.
//
// These dimensions don't include any scrollback buffer height.
func getSize(hStdin uintptr) (width, height int, err error) {
	var info windows.ConsoleScreenBufferInfo
	if err := windows.GetConsoleScreenBufferInfo(windows.Handle(hStdin), &info); err != nil {
		return 0, 0, err
	}
	return int(info.Window.Right - info.Window.Left + 1), int(info.Window.Bottom - info.Window.Top + 1), nil
}
