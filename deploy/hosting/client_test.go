package hosting

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/hostinger/api-cli/api"
)

func TestUploadCredentialsRetriesUntilURL(t *testing.T) {
	CredentialRetryDelay = time.Millisecond
	calls := 0
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		calls++
		switch r.URL.Path {
		case "/api/hosting/v1/files/upload-urls":
			var body map[string]string
			_ = json.NewDecoder(r.Body).Decode(&body)
			if body["domain"] == "rejected.com" {
				w.WriteHeader(http.StatusForbidden)
				return
			}
			if calls < 3 {
				_ = json.NewEncoder(w).Encode(map[string]string{}) // not ready yet
				return
			}
			_ = json.NewEncoder(w).Encode(map[string]string{"url": "https://files.example", "auth_key": "a", "rest_auth_key": "b"})
		case "/api/hosting/v1/accounts/u1/domains/broken.com/files":
			w.WriteHeader(http.StatusInternalServerError)
		case "/api/domains/v1/portfolio":
			_ = json.NewEncoder(w).Encode([]map[string]string{{"domain": "mine.com"}})
		default:
			http.NotFound(w, r)
		}
	}))
	defer srv.Close()
	c := NewAPI(api.NewRawClient(srv.URL, nil))
	creds, err := c.UploadCredentials(context.Background(), "u1", "rejected.com")
	if err == nil || calls != 1 {
		t.Errorf("a 403 was retried: calls=%d err=%v", calls, err)
	}
	calls = 0
	creds, err = c.UploadCredentials(context.Background(), "u1", "example.com")
	if err != nil || creds.URL != "https://files.example" || calls != 3 {
		t.Errorf("creds=%+v err=%v calls=%d", creds, err, calls)
	}
	// No document root yet (404) is an empty site; any other failure is an error.
	if has, err := c.HasContent(context.Background(), "u1", "new.com"); has || err != nil {
		t.Errorf("404: %v %v", has, err)
	}
	if _, err := c.HasContent(context.Background(), "u1", "broken.com"); err == nil {
		t.Error("500 was not an error")
	}
	for domain, want := range map[string]bool{"mine.com": true, "blog.mine.com": true, "other.com": false} {
		if got, err := c.OwnsDomain(context.Background(), domain); err != nil || got != want {
			t.Errorf("OwnsDomain(%s) = %v %v", domain, got, err)
		}
	}
}

func TestStartBuildSendsTrimmedMessage(t *testing.T) {
	var body map[string]any
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost || r.URL.Path != "/api/hosting/v1/accounts/u1/websites/example.com/nodejs/builds" {
			http.NotFound(w, r)
			return
		}
		body = nil
		_ = json.NewDecoder(r.Body).Decode(&body)
		_ = json.NewEncoder(w).Encode(map[string]string{"uuid": "b-1", "state": "pending"})
	}))
	defer srv.Close()
	c := NewAPI(api.NewRawClient(srv.URL, nil))
	settings := NodeSettings{NodeVersion: 22, AppType: "express", RootDirectory: "."}

	for message, want := range map[string]map[string]any{
		"  Fix the login redirect\n": {"archive_path": "site.zip", "message": "Fix the login redirect"},
		"":                           {"archive_path": "site.zip"},
		" \t\n ":                     {"archive_path": "site.zip"},
	} {
		if _, err := c.StartBuild(context.Background(), "u1", "example.com", settings, "site.zip", message); err != nil {
			t.Fatalf("message %q: %v", message, err)
		}
		if body["source_type"] != "archive" {
			t.Errorf("message %q: source_type = %v", message, body["source_type"])
		}
		got, _ := body["source_options"].(map[string]any)
		if len(got) != len(want) {
			t.Errorf("message %q: source_options = %v, want %v", message, got, want)
			continue
		}
		for k, v := range want {
			if got[k] != v {
				t.Errorf("message %q: source_options[%s] = %v, want %v", message, k, got[k], v)
			}
		}
	}
}

func TestParseMessage(t *testing.T) {
	if m, err := ParseMessage("  Add search\n\nIndexes posts by title.\n"); err != nil || m != "Add search\n\nIndexes posts by title." {
		t.Errorf("ParseMessage = %q, %v", m, err)
	}
	if m, err := ParseMessage(" " + strings.Repeat("ž", MaxMessageLength) + " "); err != nil || m != strings.Repeat("ž", MaxMessageLength) {
		t.Errorf("500 characters were rejected: %v", err)
	}
	if _, err := ParseMessage(strings.Repeat("a", MaxMessageLength+1)); err == nil || !strings.Contains(err.Error(), "limit is 500") {
		t.Errorf("501 characters: %v", err)
	}
}
