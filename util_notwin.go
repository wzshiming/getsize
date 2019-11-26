// +build !windows

package getsize

import (
	"golang.org/x/sys/unix"
)

// isTerminal returns whether the given file descriptor is a terminal.
func isTerminal(hStdin uintptr) bool {
	_, err := unix.IoctlGetTermios(int(hStdin), ioctlReadTermios)
	return err == nil
}

// getSize returns the dimensions of the given terminal.
func getSize(hStdin uintptr) (width, height int, err error) {
	ws, err := unix.IoctlGetWinsize(int(hStdin), unix.TIOCGWINSZ)
	if err != nil {
		return -1, -1, err
	}
	return int(ws.Col), int(ws.Row), nil
}
