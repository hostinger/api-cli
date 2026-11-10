//go:build !windows

package preview

import (
	"os"
	"syscall"
)

// stopSignals end a Preview: Ctrl+C, a service manager's stop, and a closed
// terminal.
var stopSignals = []os.Signal{os.Interrupt, syscall.SIGTERM, syscall.SIGHUP}
