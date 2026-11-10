package config

import (
	"os"
	"path/filepath"
)

// The socket gets its own directory: a config directory that an administrator
// created inherits Users read access from %ProgramData%, and admin refuses such
// a socket directory. The connector creates a missing config directory for
// SYSTEM, Administrators and the owner only.
var (
	DefaultConfigPath  = filepath.Join(programData(), "hcdn-connector", "config.yml")
	DefaultAdminSocket = filepath.Join(programData(), "hcdn-connector", "run", "admin.sock")
)

func programData() string {
	if p := os.Getenv("ProgramData"); p != "" {
		return p
	}
	return `C:\ProgramData`
}
