package api

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// lockPollInterval is how often lockFile retries a lock another process holds.
const lockPollInterval = 20 * time.Millisecond

// lockFile takes an exclusive advisory lock on path, creating the file (and its
// directory) when missing, and returns the function that releases it. It
// retries a held lock until ctx is done. The lock file is never removed:
// deleting it while another process waits on it would let two holders run at
// once. The lock is per open file, so it also excludes other callers in this
// process.
func lockFile(ctx context.Context, path string) (unlock func(), err error) {
	if err := os.MkdirAll(filepath.Dir(path), 0o700); err != nil {
		return nil, err
	}
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0o600)
	if err != nil {
		return nil, err
	}
	ticker := time.NewTicker(lockPollInterval)
	defer ticker.Stop()
	for {
		ok, err := tryLockFile(f)
		if err != nil {
			f.Close()
			return nil, fmt.Errorf("locking %s: %w", path, err)
		}
		if ok {
			return func() {
				_ = unlockFile(f)
				f.Close()
			}, nil
		}
		select {
		case <-ctx.Done():
			f.Close()
			return nil, fmt.Errorf("waiting for %s: %w", path, ctx.Err())
		case <-ticker.C:
		}
	}
}
