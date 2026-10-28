//go:build darwin || dragonfly || freebsd || linux || netbsd || openbsd

package api

import (
	"errors"
	"os"
	"syscall"
)

// tryLockFile takes a non-blocking exclusive flock(2) on f. It reports false
// when another open file holds the lock.
func tryLockFile(f *os.File) (bool, error) {
	err := syscall.Flock(int(f.Fd()), syscall.LOCK_EX|syscall.LOCK_NB)
	switch {
	case err == nil:
		return true, nil
	case errors.Is(err, syscall.EWOULDBLOCK), errors.Is(err, syscall.EINTR):
		return false, nil
	default:
		return false, err
	}
}

func unlockFile(f *os.File) error {
	return syscall.Flock(int(f.Fd()), syscall.LOCK_UN)
}
