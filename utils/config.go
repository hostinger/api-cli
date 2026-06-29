package utils

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

// MigrateLegacyConfig returns the YAML config path the CLI should read and
// whether a one-time migration (or legacy fallback) happened. When the new
// ~/.hostinger.yaml is absent but a legacy ~/.hapi.yaml exists, it moves the old
// file to the new path. If the move fails, it falls back to the old path in
// place. Notices are written to w. With no legacy file, it is a no-op.
func MigrateLegacyConfig(home string, w io.Writer) (configPath string, migrated bool) {
	newPath := filepath.Join(home, ".hostinger.yaml")
	oldPath := filepath.Join(home, ".hapi.yaml")

	if _, err := os.Stat(newPath); err == nil {
		return newPath, false
	}
	if _, err := os.Stat(oldPath); os.IsNotExist(err) {
		return newPath, false
	}
	if err := os.Rename(oldPath, newPath); err != nil {
		fmt.Fprintf(w, "warning: reading %s; could not migrate to %s: %v\n", oldPath, newPath, err)
		// migrated=true means "legacy fallback in effect"; configPath is the old path.
		return oldPath, true
	}
	fmt.Fprintf(w, "migrated %s -> %s\n", oldPath, newPath)
	return newPath, true
}

// BindLegacyEnv binds config keys to their HOSTINGER_* env var with a deprecated
// HAPI_* fallback on the global viper, emitting a one-time notice to w per key
// whose value comes from the legacy variable.
func BindLegacyEnv(w io.Writer) {
	for _, key := range []string{"api_token", "api_url"} {
		newEnv := "HOSTINGER_" + strings.ToUpper(key)
		oldEnv := "HAPI_" + strings.ToUpper(key)
		_ = viper.BindEnv(key, newEnv, oldEnv)
		if msg := legacyEnvWarning(os.Getenv(newEnv), os.Getenv(oldEnv), newEnv, oldEnv); msg != "" {
			fmt.Fprintln(w, msg)
		}
	}
}

// legacyEnvWarning returns a deprecation message when only the legacy variable
// is set, or "" otherwise.
func legacyEnvWarning(newVal, oldVal, newEnv, oldEnv string) string {
	if newVal == "" && oldVal != "" {
		return fmt.Sprintf("deprecated: %s is deprecated, use %s instead", oldEnv, newEnv)
	}
	return ""
}
