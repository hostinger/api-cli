package api

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/hostinger/api-cli/client"
	"github.com/hostinger/api-cli/utils"
)

// Streams for --dry-run (the command's output) and --verbose (diagnostics);
// package vars for tests. nil means os.Stdout / os.Stderr at write time.
var (
	dryRunOut  io.Writer
	verboseOut io.Writer
)

func writerOr(w, fallback io.Writer) io.Writer {
	if w != nil {
		return w
	}
	return fallback
}

// dryRunDoer prints each request as an equivalent curl command and returns
// utils.ErrDryRun instead of sending it. The token is written as an
// environment variable reference, never its value.
type dryRunDoer struct{}

func (dryRunDoer) Do(req *http.Request) (*http.Response, error) {
	var b strings.Builder
	fmt.Fprintf(&b, "curl -X %s %s", req.Method, shellQuote(req.URL.String()))
	fmt.Fprintf(&b, " \\\n  -H \"Authorization: Bearer $%sAPI_TOKEN\"", utils.EnvPrefix)
	if ct := req.Header.Get("Content-Type"); ct != "" {
		fmt.Fprintf(&b, " \\\n  -H %s", shellQuote("Content-Type: "+ct))
	}
	if req.Body != nil && req.Body != http.NoBody {
		body, err := io.ReadAll(req.Body)
		if err != nil {
			return nil, err
		}
		if len(body) > 0 {
			fmt.Fprintf(&b, " \\\n  -d %s", shellQuote(string(body)))
		}
	}
	fmt.Fprintln(writerOr(dryRunOut, os.Stdout), b.String())
	return nil, utils.ErrDryRun
}

// noTokenDoer fails every request of a build without OAuth when no API token
// is configured.
type noTokenDoer struct{}

func (noTokenDoer) Do(*http.Request) (*http.Response, error) {
	return nil, fmt.Errorf("%w: set %sAPI_TOKEN or run '%s config set api_token'",
		utils.ErrNoCredentials, utils.EnvPrefix, filepath.Base(os.Args[0]))
}

// verboseDoer logs each request line and the response status with timing to
// stderr. Headers are not logged, so credentials never reach the log.
type verboseDoer struct {
	base client.HttpRequestDoer
}

func (d *verboseDoer) Do(req *http.Request) (*http.Response, error) {
	fmt.Fprintf(writerOr(verboseOut, os.Stderr), "> %s %s\n", req.Method, req.URL.String())
	start := time.Now()
	resp, err := d.base.Do(req)
	elapsed := time.Since(start).Round(time.Millisecond)
	if err != nil {
		fmt.Fprintf(writerOr(verboseOut, os.Stderr), "< error after %s: %v\n", elapsed, err)
		return resp, err
	}
	fmt.Fprintf(writerOr(verboseOut, os.Stderr), "< %s (%s)\n", resp.Status, elapsed)
	return resp, nil
}

func shellQuote(s string) string {
	return "'" + strings.ReplaceAll(s, "'", `'\''`) + "'"
}
