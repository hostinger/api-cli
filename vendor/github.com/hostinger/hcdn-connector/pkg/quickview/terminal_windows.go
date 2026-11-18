//go:build windows

package quickview

import (
	"os"

	"golang.org/x/sys/windows"
)

// enableEscapes turns on escape sequences in the console f writes to. It
// fails on a console that has no such mode, such as an old Windows 10.
func enableEscapes(f *os.File) bool {
	h := windows.Handle(f.Fd())
	var mode uint32
	if err := windows.GetConsoleMode(h, &mode); err != nil {
		return false
	}
	if mode&windows.ENABLE_VIRTUAL_TERMINAL_PROCESSING != 0 {
		return true
	}
	return windows.SetConsoleMode(h, mode|windows.ENABLE_VIRTUAL_TERMINAL_PROCESSING) == nil
}
