package getsize

import (
	"os"
)

var (
	hStdin = os.Stdin.Fd()
)

// IsTerminal returns whether the given file descriptor is a terminal.
func IsTerminal() bool {
	return IsTerminalBy(hStdin)
}

// IsTerminalBy returns whether the given file descriptor is a terminal.
func IsTerminalBy(hStdin uintptr) bool {
	return isTerminal(hStdin)
}

// GetSize returns the visible dimensions of the given terminal.
//
// These dimensions don't include any scrollback buffer height.
func GetSize() (width, height int, err error) {
	return GetSizeBy(hStdin)
}

// GetSizeBy returns the visible dimensions of the given terminal.
//
// These dimensions don't include any scrollback buffer height.
func GetSizeBy(hStdin uintptr) (width, height int, err error) {
	return getSize(hStdin)
}
