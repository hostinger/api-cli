//go:build !windows

package quickview

import "os"

// enableEscapes: terminals outside Windows take escape sequences.
func enableEscapes(*os.File) bool { return true }
