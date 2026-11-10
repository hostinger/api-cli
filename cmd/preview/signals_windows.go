package preview

import (
	"os"
	"syscall"
)

// stopSignals end a Preview. Windows delivers Ctrl+C and Ctrl+Break as
// os.Interrupt, and closing the console window, logging off and shutting
// down as SIGTERM.
var stopSignals = []os.Signal{os.Interrupt, syscall.SIGTERM}
