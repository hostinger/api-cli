package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/mail"
	"os"
	"slices"
	"strings"
	"time"

	"github.com/google/uuid"
	openapi_types "github.com/oapi-codegen/runtime/types"
	"github.com/spf13/cobra"
)

// FlagError is an input mistake (bad flag value, missing argument). The root
// command prints it with a pointer to --help and exits with ExitUsage.
type FlagError struct{ Err error }

func (e *FlagError) Error() string { return e.Err.Error() }
func (e *FlagError) Unwrap() error { return e.Err }

// FlagErrorf formats a FlagError.
func FlagErrorf(format string, a ...any) error {
	return &FlagError{Err: fmt.Errorf(format, a...)}
}

// EnumCheck validates that a changed string flag holds one of the allowed values.
func EnumCheck(cmd *cobra.Command, flag string, allowed []string) error {
	if !cmd.Flags().Changed(flag) {
		return nil
	}
	v, _ := cmd.Flags().GetString(flag)
	if !slices.Contains(allowed, v) {
		return FlagErrorf("invalid value %q for --%s (one of: %s)", v, flag, strings.Join(allowed, ", "))
	}
	return nil
}

// FormatCheck validates that a changed string flag parses as format
// ("date-time", "date", "uuid" or "email") before it is converted.
func FormatCheck(cmd *cobra.Command, flag, format string) error {
	if !cmd.Flags().Changed(flag) {
		return nil
	}
	v, _ := cmd.Flags().GetString(flag)
	if err := checkFormat(v, format); err != nil {
		return FlagErrorf("invalid value %q for --%s: %v", v, flag, err)
	}
	return nil
}

func checkFormat(v, format string) error {
	switch format {
	case "date-time":
		if _, err := time.Parse(time.RFC3339, v); err != nil {
			return fmt.Errorf("expected an RFC 3339 date-time such as 2025-01-31T10:00:00Z")
		}
	case "date":
		if _, err := time.Parse(openapi_types.DateFormat, v); err != nil {
			return fmt.Errorf("expected a date such as 2025-01-31")
		}
	case "uuid":
		if _, err := uuid.Parse(v); err != nil {
			return fmt.Errorf("expected a UUID")
		}
	case "email":
		if _, err := mail.ParseAddress(v); err != nil {
			return fmt.Errorf("expected an email address")
		}
	}
	return nil
}

// stdin is the reader "-" values are read from; a package var for tests.
var stdin io.Reader = os.Stdin

// JSONValue parses a JSON flag value into a generic value. The value is either
// inline JSON, @path to read it from a file, or - to read it from stdin.
func JSONValue(raw, flag string) (any, error) {
	src := []byte(raw)
	switch {
	case raw == "-":
		b, err := io.ReadAll(stdin)
		if err != nil {
			return nil, FlagErrorf("reading --%s from stdin: %v", flag, err)
		}
		src = b
	case strings.HasPrefix(raw, "@"):
		b, err := os.ReadFile(strings.TrimPrefix(raw, "@"))
		if err != nil {
			return nil, FlagErrorf("reading --%s: %v", flag, err)
		}
		src = b
	}
	var v any
	if err := json.Unmarshal(src, &v); err != nil {
		return nil, FlagErrorf("invalid JSON in --%s: %v (see the command's --help for the expected shape)", flag, err)
	}
	return v, nil
}
