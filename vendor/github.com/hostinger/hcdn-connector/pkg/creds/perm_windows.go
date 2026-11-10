//go:build windows

package creds

import (
	"fmt"
	"io/fs"
	"os"
	"os/user"
	"slices"
	"strings"

	"github.com/hostinger/hcdn-connector/internal/winacl"
)

const permHint = "read access for your user only"

func openCredentials(path string) (*os.File, error) { return os.Open(path) }

// Every Windows file reports a group/world mode, so the DACL decides.
func checkPerm(path string, f *os.File, _ fs.FileInfo) error {
	who, err := winacl.BroadAccess(f)
	if err != nil {
		return fmt.Errorf("credentials %s: read the ACL: %w", path, err)
	}
	if len(who) == 0 {
		return nil
	}
	var names, sids []string
	for _, r := range who {
		names = append(names, r.Name)
		if !slices.Contains(sids, "*"+r.SID) {
			sids = append(sids, "*"+r.SID)
		}
	}
	// /inheritance:r drops only inherited ACEs; /remove:g drops the explicit ones.
	return fmt.Errorf(`credentials %s: accessible by %s; run icacls "%s" /inheritance:r /grant:r "%s:F", then icacls "%s" /remove:g %s`,
		path, strings.Join(names, ", "), path, currentUser(), path, strings.Join(sids, " "))
}

func currentUser() string {
	if u, err := user.Current(); err == nil {
		return u.Username
	}
	return os.Getenv("USERNAME")
}

func createPrivate(path string) (*os.File, error) { return winacl.CreatePrivate(path) }
