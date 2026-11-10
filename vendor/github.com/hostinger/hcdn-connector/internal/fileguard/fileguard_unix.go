//go:build !windows

package fileguard

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"syscall"
)

// maxLinks matches the Linux limit on symlinks followed in one lookup.
const maxLinks = 40

// Check refuses f, opened from path, when it is world-writable, or when the
// directory that holds it is world-writable without the sticky bit. In a
// sticky directory (/tmp) another account can still create the file first, so
// there the file and the directory must belong to root or to this user.
//
// The directory is the one the file is really in, found by following path as
// the open did. Each symlink on the way can be replaced or planted like the
// file, so the same rules apply to it in its own directory. If path no longer
// leads to f, f is refused.
//
// Owners are not checked otherwise: a bind mount or a user namespace shows a
// host uid that is neither root nor the container user. Group write is allowed:
// with user private groups (umask 002) the group is often the owner alone.
func Check(path string, f *os.File) error {
	st, err := f.Stat()
	if err != nil {
		return err
	}
	if perm := st.Mode().Perm(); perm&0o002 != 0 {
		return fmt.Errorf("mode %04o is world-writable; run chmod o-w %s", perm, path)
	}
	resolved, err := resolve(path)
	if err != nil {
		return err
	}
	rst, err := os.Lstat(resolved)
	if err != nil {
		return err
	}
	if !os.SameFile(st, rst) {
		return fmt.Errorf("changed while checking: %s is not the opened file", resolved)
	}
	return checkIn(filepath.Dir(resolved), st)
}

// resolve follows path component by component, as open does, and returns the
// path it leads to with no symlink left in it. Every symlink it follows must
// pass checkIn in the directory that holds the link.
func resolve(path string) (string, error) {
	var dir string // the part walked so far, free of symlinks; "" is "."
	if filepath.IsAbs(path) {
		dir = "/"
	}
	rest := path
	for links := 0; rest != ""; {
		var name string
		name, rest, _ = strings.Cut(rest, "/")
		switch name {
		case "", ".":
			continue
		case "..":
			// dir has no symlinks, so its parent is the lexical one.
			if dir == "" || filepath.Base(dir) == ".." {
				dir = filepath.Join(dir, "..")
			} else if dir = filepath.Dir(dir); dir == "." {
				dir = ""
			}
			continue
		}
		p := filepath.Join(dir, name)
		fi, err := os.Lstat(p)
		if err != nil {
			return "", err
		}
		if fi.Mode()&fs.ModeSymlink == 0 {
			dir = p
			continue
		}
		if links++; links > maxLinks {
			return "", &fs.PathError{Op: "open", Path: path, Err: syscall.ELOOP}
		}
		if err := checkIn(filepath.Dir(p), fi); err != nil {
			return "", fmt.Errorf("symlink %s: %w", p, err)
		}
		target, err := os.Readlink(p)
		if err != nil {
			return "", err
		}
		if filepath.IsAbs(target) {
			dir = "/"
		}
		rest = target + "/" + rest
	}
	if dir == "" {
		dir = "."
	}
	return dir, nil
}

// checkIn refuses an object (a file or a symlink) with info st in dir when dir
// is world-writable without the sticky bit, or when dir is sticky and
// world-writable and dir or the object belongs to another account.
func checkIn(dir string, st fs.FileInfo) error {
	dst, err := os.Stat(dir)
	if err != nil {
		return err
	}
	perm := dst.Mode().Perm()
	if perm&0o002 == 0 {
		return nil
	}
	if dst.Mode()&fs.ModeSticky == 0 {
		return fmt.Errorf("directory %s: mode %04o is world-writable; run chmod o-w %s", dir, perm, dir)
	}
	if uid, ok := untrustedOwner(dst); ok {
		return fmt.Errorf("directory %s: world-writable and owned by uid %d", dir, uid)
	}
	if uid, ok := untrustedOwner(st); ok {
		return fmt.Errorf("owned by uid %d in the world-writable directory %s", uid, dir)
	}
	return nil
}

func untrustedOwner(st fs.FileInfo) (uint32, bool) {
	s, ok := st.Sys().(*syscall.Stat_t)
	if !ok {
		return 0, false
	}
	return s.Uid, s.Uid != 0 && int(s.Uid) != os.Geteuid()
}
