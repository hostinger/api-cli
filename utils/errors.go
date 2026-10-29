package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"sort"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Exit codes. Scripts can rely on these; keep README.md in sync.
const (
	ExitOK         = 0
	ExitError      = 1 // API, network or other runtime failure
	ExitUsage      = 2 // bad arguments, flags or input
	ExitAuth       = 3 // HTTP 401/403
	ExitNotFound   = 4 // HTTP 404
	ExitValidation = 5 // HTTP 422
)

// ErrDryRun is returned by the HTTP layer after --dry-run printed a request
// instead of sending it. It ends the command successfully.
var ErrDryRun = errors.New("dry run")

// ErrNoCredentials marks a request that could not be sent because no API
// token (or sign-in) is configured. It exits with ExitAuth.
var ErrNoCredentials = errors.New("no credentials configured")

// APIError is an HTTP error response from the API.
type APIError struct {
	Status        int
	Message       string
	Errors        map[string][]string
	CorrelationID string
	Body          []byte
}

// NewAPIError decodes the API's error body ({message, errors, correlation_id});
// an unexpected body is kept verbatim.
func NewAPIError(status int, body []byte) *APIError {
	e := &APIError{Status: status, Body: body}
	var parsed struct {
		Message       string              `json:"message"`
		Errors        map[string][]string `json:"errors"`
		CorrelationID string              `json:"correlation_id"`
	}
	if json.Unmarshal(body, &parsed) == nil {
		e.Message, e.Errors, e.CorrelationID = parsed.Message, parsed.Errors, parsed.CorrelationID
	}
	return e
}

func (e *APIError) Error() string {
	msg := e.Message
	if msg == "" {
		msg = strings.TrimSpace(string(e.Body))
	}
	if msg == "" {
		msg = http.StatusText(e.Status)
	}
	return fmt.Sprintf("%s (HTTP %d)", msg, e.Status)
}

// ExitCode maps the HTTP status to the CLI exit code.
func (e *APIError) ExitCode() int {
	switch e.Status {
	case http.StatusUnauthorized, http.StatusForbidden:
		return ExitAuth
	case http.StatusNotFound:
		return ExitNotFound
	case http.StatusUnprocessableEntity:
		return ExitValidation
	}
	return ExitError
}

// HandleError prints err for the command that failed and returns the exit
// code. Usage mistakes get a pointer to --help; API errors list each invalid
// field against the flag that sets it.
func HandleError(w io.Writer, cmd *cobra.Command, err error) int {
	if err == nil || errors.Is(err, ErrDryRun) {
		return ExitOK
	}
	w = safeWriter{w} // messages, field errors and correlation IDs come from the server
	if errors.Is(err, ErrAborted) {
		fmt.Fprintln(w, "Aborted.")
		return ExitError
	}
	if errors.Is(err, ErrNoCredentials) {
		fmt.Fprintln(w, "Error:", err)
		return ExitAuth
	}
	var apiErr *APIError
	if errors.As(err, &apiErr) {
		fmt.Fprintln(w, "Error:", apiErr.Error())
		fmt.Fprint(w, fieldErrors(cmd, apiErr.Errors))
		if apiErr.Status == http.StatusUnauthorized {
			if viper.GetString("api_token") != "" {
				fmt.Fprintf(w, "Hint: the configured API token was rejected; see '%s auth status'.\n", cmd.Root().Name())
			} else if login, _, err := cmd.Root().Find([]string{"auth", "login"}); err == nil && login.Name() == "login" {
				fmt.Fprintf(w, "Hint: sign in with '%s auth login' or set %sAPI_TOKEN.\n", cmd.Root().Name(), EnvPrefix)
			} else {
				fmt.Fprintf(w, "Hint: set %sAPI_TOKEN or run '%s config set api_token'.\n", EnvPrefix, cmd.Root().Name())
			}
		}
		if apiErr.CorrelationID != "" {
			fmt.Fprintln(w, "Correlation ID:", apiErr.CorrelationID)
		}
		return apiErr.ExitCode()
	}
	fmt.Fprintln(w, "Error:", err)
	if isUsageError(err) {
		fmt.Fprintf(w, "Run '%s --help' for usage.\n", cmd.CommandPath())
		return ExitUsage
	}
	return ExitError
}

// fieldErrors renders the API's per-field validation messages, naming the
// flag that sets each field ("data_center_id" -> --data-center-id; nested
// "setup.hostname" -> --setup (hostname)).
func fieldErrors(cmd *cobra.Command, fields map[string][]string) string {
	keys := make([]string, 0, len(fields))
	for k := range fields {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var b bytes.Buffer
	for _, k := range keys {
		top, rest, nested := strings.Cut(k, ".")
		name := strings.ReplaceAll(strings.ToLower(top), "_", "-")
		label := k
		if cmd != nil && cmd.Flags().Lookup(name) != nil {
			label = "--" + name
			if nested {
				label += " (" + rest + ")"
			}
		}
		for _, msg := range fields[k] {
			fmt.Fprintf(&b, "  %s: %s\n", label, msg)
		}
	}
	return b.String()
}

// isUsageError reports input mistakes: our FlagError plus cobra's own
// argument and flag errors, which are plain errors with known prefixes.
func isUsageError(err error) bool {
	var fe *FlagError
	if errors.As(err, &fe) {
		return true
	}
	msg := err.Error()
	for _, p := range []string{"unknown command", "unknown flag", "unknown shorthand flag", "invalid argument", "flag needs an argument", "required flag(s)", "accepts ", "requires at least", "if any flags in the group", "bad flag syntax"} {
		if strings.HasPrefix(msg, p) {
			return true
		}
	}
	return false
}
