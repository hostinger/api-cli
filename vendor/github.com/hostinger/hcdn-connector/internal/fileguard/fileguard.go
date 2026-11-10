// Package fileguard refuses a file that changes what the connector trusts
// (config.yml, credentials.json, a CA file) when another account can replace
// it: through the file itself or through the directory that holds it.
package fileguard

import (
	"fmt"
	"io"
	"os"
)

// ReadFile reads path after Check, from the handle that Check inspected.
func ReadFile(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	if err := Check(path, f); err != nil {
		return nil, fmt.Errorf("%s: %w", path, err)
	}
	return io.ReadAll(f)
}
