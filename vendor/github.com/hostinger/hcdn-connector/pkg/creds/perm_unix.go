//go:build !windows

package creds

import (
	"fmt"
	"io/fs"
	"os"
	"syscall"
)

const (
	forbiddenPerm = os.FileMode(0o077)
	permHint      = "mode 0600"
)

// O_NONBLOCK: a FIFO at the path must not block the open.
func openCredentials(path string) (*os.File, error) {
	return os.OpenFile(path, os.O_RDONLY|syscall.O_NONBLOCK, 0)
}

func checkPerm(path string, _ *os.File, st fs.FileInfo) error {
	if perm := st.Mode().Perm(); perm&forbiddenPerm != 0 {
		return fmt.Errorf("credentials %s: mode %04o is group/world readable; run chmod 0600", path, perm)
	}
	return nil
}

func createPrivate(path string) (*os.File, error) {
	return os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0o600)
}
