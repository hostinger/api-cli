package quickview

import (
	"os"

	"golang.org/x/term"
)

// Supported reports whether f is a terminal that can show the view: a
// terminal, NO_COLOR unset, TERM not "dumb", and on Windows a console that
// takes escape sequences (the view turns them on when it can).
func Supported(f *os.File) bool {
	if f == nil || !term.IsTerminal(int(f.Fd())) {
		return false
	}
	if _, noColor := os.LookupEnv("NO_COLOR"); noColor || os.Getenv("TERM") == "dumb" {
		return false
	}
	return enableEscapes(f)
}

// terminalSize is the size of f, or 80x24 when it has none.
func terminalSize(f *os.File) func() (int, int) {
	return func() (int, int) {
		w, h, err := term.GetSize(int(f.Fd()))
		if err != nil || w <= 0 {
			return 80, 24
		}
		return w, h
	}
}
