package api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/hostinger/api-cli/client"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/viper"
)

// RawClient calls API paths directly, configured like the typed client, so
// hand-written commands don't depend on generated method names.
type RawClient struct {
	server  string
	doer    client.HttpRequestDoer
	editors []client.RequestEditorFn
}

// Raw ignores --dry-run: hand-written commands look up the website, account or
// order with it before the request --dry-run prints, so those lookups are sent.
// Under --dry-run they use a token or a stored sign-in but never start a login,
// so a dry run without credentials fails at once instead of waiting on a browser.
func Raw() *RawClient {
	apiURL, opts := clientOptions(!viper.GetBool("dry_run"), false)
	c, err := client.NewClient(apiURL, opts...)
	if err != nil {
		log.Fatal(err)
	}
	return &RawClient{server: strings.TrimRight(c.Server, "/"), doer: c.Client, editors: c.RequestEditors}
}

// NewRawClient has no authentication (tests).
func NewRawClient(server string, doer client.HttpRequestDoer) *RawClient {
	if doer == nil {
		doer = http.DefaultClient
	}
	return &RawClient{server: strings.TrimRight(server, "/"), doer: doer}
}

// CommandError gives an API error from a raw call the exit code and hints the
// generated commands give theirs; other errors pass through.
func CommandError(err error) error {
	var raw *Error
	if errors.As(err, &raw) {
		return utils.NewAPIError(raw.Status, raw.Body)
	}
	return err
}

type Error struct {
	Method string
	Path   string
	Status int
	Body   []byte
}

func (e *Error) Error() string {
	msg := strings.TrimSpace(string(e.Body))
	var parsed struct {
		Message string `json:"message"`
	}
	if json.Unmarshal(e.Body, &parsed) == nil && parsed.Message != "" {
		msg = parsed.Message
	}
	return fmt.Sprintf("%s (HTTP %d from %s %s)", msg, e.Status, e.Method, e.Path)
}

// JSON decodes a 2xx response into out (may be nil); other statuses are *Error.
func (r *RawClient) JSON(ctx context.Context, method, path string, query url.Values, body any, out any) error {
	var reader io.Reader
	var raw []byte
	if body != nil {
		var err error
		raw, err = json.Marshal(body)
		if err != nil {
			return err
		}
		reader = bytes.NewReader(raw)
	}
	u := r.server + path
	if len(query) > 0 {
		u += "?" + query.Encode()
	}
	req, err := http.NewRequestWithContext(ctx, method, u, reader)
	if err != nil {
		return err
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
		req.GetBody = func() (io.ReadCloser, error) { return io.NopCloser(bytes.NewReader(raw)), nil }
	}
	req.Header.Set("Accept", "application/json")
	for _, edit := range r.editors {
		if err := edit(ctx, req); err != nil {
			return err
		}
	}
	resp, err := r.doer.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return &Error{Method: method, Path: path, Status: resp.StatusCode, Body: data}
	}
	if out == nil || len(bytes.TrimSpace(data)) == 0 {
		return nil
	}
	if err := json.Unmarshal(data, out); err != nil {
		return fmt.Errorf("%s %s: decoding response: %w", method, path, err)
	}
	return nil
}
