package site

import (
	"os"

	"github.com/pterm/pterm"
)

// Spin shows a spinner on stderr in a terminal; call the result to end it.
// Tests replace it.
var Spin = func(msg string) func(done string) {
	if !Interactive() {
		return func(string) {}
	}
	sp, err := pterm.DefaultSpinner.WithWriter(os.Stderr).Start(msg)
	if err != nil {
		return func(string) {}
	}
	return func(done string) { sp.Success(done) }
}
