//go:build !windows

package brand

import "os"

// enableEscapes: terminals outside Windows take escape sequences.
func enableEscapes(*os.File) bool { return true }
