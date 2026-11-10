//go:build windows

package fileguard

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/hostinger/hcdn-connector/internal/winacl"
)

const hint = "only SYSTEM, Administrators and the account that runs hcdn-connector may own it or change it"

// Check refuses f, opened from path, or the directory that holds it, when an
// account other than SYSTEM, Administrators and the current user owns it or
// can change it. Under %ProgramData% any user can create a directory first and
// then owns it, or create a junction there that it can re-point: a junction or
// symbolic link in path must pass the same check (winacl.OpenTrustedPath), and
// the directory judged is the real one, links resolved. A file on a network
// volume is refused. The file is checked on its handle; read from f, not from
// path again.
func Check(path string, f *os.File) error {
	l, err := winacl.OpenTrustedPath(path)
	if err != nil {
		return withHint(err)
	}
	defer l.Close()
	if same, err := sameFile(l, f); err != nil || !same {
		if err == nil {
			err = errors.New("changed while it was checked")
		}
		return err
	}
	if err := winacl.CheckTrusted(f); err != nil {
		return withHint(err)
	}
	dir := filepath.Dir(l.Name())
	d, err := winacl.OpenTrustedPath(dir)
	if err != nil {
		return withHint(err)
	}
	defer d.Close()
	if err := winacl.CheckTrusted(d); err != nil {
		return fmt.Errorf("directory %s: %w", dir, withHint(err))
	}
	return nil
}

// sameFile reports whether the walk reached the file that os.Open opened.
func sameFile(a, b *os.File) (bool, error) {
	sa, err := a.Stat()
	if err != nil {
		return false, err
	}
	sb, err := b.Stat()
	if err != nil {
		return false, err
	}
	return os.SameFile(sa, sb), nil
}

func withHint(err error) error {
	if errors.Is(err, winacl.ErrRemote) {
		return err
	}
	return fmt.Errorf("%w; %s", err, hint)
}
