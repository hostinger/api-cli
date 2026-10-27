package api

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/viper"
)

// --dry-run prints the command's own request; the lookups that build it still go out.
func TestRawSendsUnderDryRun(t *testing.T) {
	var got string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		got = r.Header.Get("Authorization")
		_, _ = w.Write([]byte(`{"ok":true}`))
	}))
	defer srv.Close()
	for key, v := range map[string]any{"api_url": srv.URL, "api_token": "t", "dry_run": true} {
		old := viper.Get(key)
		viper.Set(key, v)
		t.Cleanup(func() { viper.Set(key, old) })
	}

	var out struct{ OK bool }
	if err := Raw().JSON(context.Background(), http.MethodGet, "/api/hosting/v1/websites", nil, nil, &out); err != nil {
		t.Fatal(err)
	}
	if !out.OK || got != "Bearer t" {
		t.Errorf("out = %+v, Authorization = %q", out, got)
	}
}

// Without credentials a dry run's lookup fails at once: it never starts a login.
func TestRawDryRunNeverLogsIn(t *testing.T) {
	t.Setenv("HOME", t.TempDir()) // no stored sign-in
	for key, v := range map[string]any{"api_url": "http://127.0.0.1:1", "api_token": "", "dry_run": true} {
		old := viper.Get(key)
		viper.Set(key, v)
		t.Cleanup(func() { viper.Set(key, old) })
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := Raw().JSON(ctx, http.MethodGet, "/api/hosting/v1/websites", nil, nil, nil); err == nil || errors.Is(err, context.DeadlineExceeded) {
		t.Errorf("err = %v, want a quick failure", err)
	}
}

func TestCommandErrorMapsAPIErrors(t *testing.T) {
	err := CommandError(&Error{Method: "GET", Path: "/x", Status: http.StatusUnauthorized, Body: []byte(`{"message":"Unauthenticated."}`)})
	var apiErr *utils.APIError
	if !errors.As(err, &apiErr) || apiErr.Status != http.StatusUnauthorized {
		t.Errorf("err = %#v", err)
	}
	other := errors.New("boom")
	if CommandError(other) != other {
		t.Error("other errors should pass through")
	}
}
