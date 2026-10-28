//go:build !(darwin || dragonfly || freebsd || linux || netbsd || openbsd || windows)

package api

import (
	"errors"
	"os"
)

// Platforms without flock(2) or LockFileEx (none of the released targets) get
// no cross-process lock, so CachedToken returns an error instead of refreshing
// unguarded.
func tryLockFile(*os.File) (bool, error) {
	return false, errors.New("file locking is not supported on this platform")
}

func unlockFile(*os.File) error { return nil }
