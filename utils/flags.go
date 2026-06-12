package utils

import (
	"encoding/json"
	"log"
	"slices"

	"github.com/spf13/cobra"
)

// EnumCheck validates that a changed string flag holds one of the allowed values.
func EnumCheck(cmd *cobra.Command, flag string, allowed []string) {
	if !cmd.Flags().Changed(flag) {
		return
	}
	v, _ := cmd.Flags().GetString(flag)
	if !slices.Contains(allowed, v) {
		log.Fatalf("invalid value %q for --%s (one of: %v)", v, flag, allowed)
	}
}

// JSONValue parses a JSON-string flag value into a generic value.
func JSONValue(raw, flag string) any {
	var v any
	if err := json.Unmarshal([]byte(raw), &v); err != nil {
		log.Fatalf("invalid JSON in --%s: %v", flag, err)
	}
	return v
}
