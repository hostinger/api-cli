package hosting

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
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
