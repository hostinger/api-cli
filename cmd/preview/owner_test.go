package preview

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync/atomic"
	"testing"
	"time"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/viper"
)

const testTunnelID = "0192f3a4-5b6c-8d7e-9f80-a1b2c3d4e5f6"

var testNotAfter = time.Date(2026, 10, 23, 14, 5, 0, 0, time.FixedZone("EEST", 3*3600))

// fakePublicAPI answers the owner ticket route with a fixed status and body,
// after an optional delay, and records what it received.
type fakePublicAPI struct {
	*httptest.Server
	calls atomic.Int32
	req   atomic.Pointer[capturedRequest]
}

type capturedRequest struct {
	method, path, auth, contentType string
	body                            map[string]any
}

func newFakePublicAPI(t *testing.T, status int, body string, delay time.Duration) *fakePublicAPI {
	t.Helper()
	f := &fakePublicAPI{}
	f.Server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		f.calls.Add(1)
		c := &capturedRequest{method: r.Method, path: r.URL.Path,
			auth: r.Header.Get("Authorization"), contentType: r.Header.Get("Content-Type")}
		_ = json.NewDecoder(r.Body).Decode(&c.body)
		f.req.Store(c)
		if delay > 0 {
			select {
			case <-time.After(delay):
			case <-r.Context().Done():
				return
			}
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		_, _ = w.Write([]byte(body))
	}))
	t.Cleanup(f.Close)
	return f
}

func staticSource(tok string) TokenSource {
	return func(context.Context, string) (string, error) { return tok, nil }
}

func fetchWith(t *testing.T, srv *fakePublicAPI, timeout time.Duration) (string, error) {
	t.Helper()
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	fn := NewOwnerTicketFunc(OwnerOptions{APIURL: srv.URL + "/", Token: staticSource("tok-123"), UserAgent: "test"})
	return fn(ctx, testTunnelID, testNotAfter)
}

func TestOwnerTicket_Created(t *testing.T) {
	srv := newFakePublicAPI(t, http.StatusCreated,
		`{"ticket":"eyJ.a.b","expires_at":"2026-10-23T11:05:00Z","client_id":"123456"}`, 0)

	ticket, err := fetchWith(t, srv, 2*time.Second)
	if err != nil || ticket != "eyJ.a.b" {
		t.Fatalf("ticket = %q, err = %v", ticket, err)
	}
	got := srv.req.Load()
	if got.method != http.MethodPost || got.path != OwnerTicketPath {
		t.Errorf("request = %s %s, want POST %s", got.method, got.path, OwnerTicketPath)
	}
	if got.auth != "Bearer tok-123" {
		t.Errorf("Authorization = %q", got.auth)
	}
	if got.contentType != "application/json" {
		t.Errorf("Content-Type = %q", got.contentType)
	}
	want := map[string]any{"tunnel_id": testTunnelID, "expires_at": "2026-10-23T11:05:00Z"}
	if fmt.Sprint(got.body) != fmt.Sprint(want) {
		t.Errorf("body = %v, want %v", got.body, want)
	}
}

func TestOwnerTicket_ErrorMapping(t *testing.T) {
	cases := []struct {
		name      string
		status    int
		body      string
		delay     time.Duration
		temporary bool
		contains  string
	}{
		{"401 unauthenticated", 401, `{"message":"Unauthenticated."}`, 0, false, "(HTTP 401): Unauthenticated."},
		{"403 forbidden", 403, `{"message":"Forbidden"}`, 0, false, "HTTP 403"},
		{"422 validation", 422, `{"message":"The expires at field is invalid."}`, 0, false, "HTTP 422"},
		{"429 rate limited", 429, `{"message":"Too Many Attempts."}`, 0, false, "HTTP 429"},
		{"500", 500, `{"message":"Server Error"}`, 0, true, "HTTP 500"},
		{"502 non-JSON body", 502, `<html>bad gateway</html>`, 0, true, "HTTP 502"},
		{"503", 503, ``, 0, true, "HTTP 503"},
		{"timeout", 201, `{"ticket":"late"}`, 2 * time.Second, true, "deadline exceeded"},
		{"201 not JSON", 201, `not json`, 0, false, "invalid response"},
		{"201 without ticket", 201, `{"client_id":"1"}`, 0, false, "no ticket"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			srv := newFakePublicAPI(t, tc.status, tc.body, tc.delay)
			start := time.Now()
			ticket, err := fetchWith(t, srv, 200*time.Millisecond)
			if err == nil {
				t.Fatalf("ticket = %q, want an error", ticket)
			}
			if got := errors.Is(err, ErrOwnerTemporary); got != tc.temporary {
				t.Errorf("errors.Is(err, ErrOwnerTemporary) = %v, want %v (err: %v)", got, tc.temporary, err)
			}
			if errors.Is(err, ErrOwnerNoToken) || errors.Is(err, ErrOwnerOptedOut) {
				t.Errorf("err = %v must not be ErrOwnerNoToken/ErrOwnerOptedOut", err)
			}
			if !strings.Contains(err.Error(), tc.contains) {
				t.Errorf("err = %q, want it to contain %q", err, tc.contains)
			}
			if d := time.Since(start); d > time.Second {
				t.Errorf("call took %s; the deadline was 200ms", d)
			}
		})
	}
}

func TestOwnerTicket_NetworkErrorIsTemporary(t *testing.T) {
	srv := newFakePublicAPI(t, 201, `{}`, 0)
	srv.Close() // nothing listens any more
	_, err := fetchWith(t, srv, time.Second)
	if !errors.Is(err, ErrOwnerTemporary) {
		t.Fatalf("err = %v, want ErrOwnerTemporary", err)
	}
}

func TestOwnerTicket_OptedOutMakesNoRequest(t *testing.T) {
	srv := newFakePublicAPI(t, 201, `{"ticket":"t"}`, 0)
	tokenCalled := false
	fn := NewOwnerTicketFunc(OwnerOptions{APIURL: srv.URL, NoOwner: true,
		Token: func(context.Context, string) (string, error) { tokenCalled = true; return "x", nil }})
	if _, err := fn(context.Background(), testTunnelID, testNotAfter); !errors.Is(err, ErrOwnerOptedOut) {
		t.Fatalf("err = %v, want ErrOwnerOptedOut", err)
	}
	if srv.calls.Load() != 0 || tokenCalled {
		t.Fatal("--no-owner must not look up a token or call public-api")
	}
}

func TestOwnerTicket_NoTokenMakesNoRequest(t *testing.T) {
	srv := newFakePublicAPI(t, 201, `{"ticket":"t"}`, 0)
	for name, src := range map[string]TokenSource{
		"nil source":   nil,
		"empty token":  staticSource(""),
		"no token err": func(context.Context, string) (string, error) { return "", ErrOwnerNoToken },
	} {
		t.Run(name, func(t *testing.T) {
			fn := NewOwnerTicketFunc(OwnerOptions{APIURL: srv.URL, Token: src})
			if _, err := fn(context.Background(), testTunnelID, testNotAfter); !errors.Is(err, ErrOwnerNoToken) {
				t.Fatalf("err = %v, want ErrOwnerNoToken", err)
			}
		})
	}
	if srv.calls.Load() != 0 {
		t.Fatal("public-api called without a token")
	}
}

func TestOwnerTicket_TokenSourceErrorPassesThrough(t *testing.T) {
	srv := newFakePublicAPI(t, 201, `{"ticket":"t"}`, 0)
	want := fmt.Errorf("%w: refresh: 503", ErrOwnerTemporary)
	fn := NewOwnerTicketFunc(OwnerOptions{APIURL: srv.URL,
		Token: func(context.Context, string) (string, error) { return "", want }})
	if _, err := fn(context.Background(), testTunnelID, testNotAfter); !errors.Is(err, ErrOwnerTemporary) {
		t.Fatalf("err = %v, want ErrOwnerTemporary", err)
	}
	if srv.calls.Load() != 0 {
		t.Fatal("public-api called after the token source failed")
	}
}

// fakeOAuth stands in for api.OAuthService's CachedToken and ReplaceToken.
type fakeOAuth struct {
	tok      string
	err      error
	called   bool
	replaced string // the token ReplaceToken was given
}

func (f *fakeOAuth) CachedToken(context.Context) (string, error) {
	f.called = true
	return f.tok, f.err
}

func (f *fakeOAuth) ReplaceToken(_ context.Context, rejected string) (string, error) {
	f.called = true
	f.replaced = rejected
	return f.tok, f.err
}

func stubSources(t *testing.T, cfgToken string, oauth *fakeOAuth, oauthErr error) {
	t.Helper()
	prevCfg, prevOAuth := configToken, newOAuthFunc
	t.Cleanup(func() { configToken, newOAuthFunc = prevCfg, prevOAuth })
	configToken = func() string { return cfgToken }
	newOAuthFunc = func() (cachedTokener, error) {
		if oauthErr != nil {
			return nil, oauthErr
		}
		return oauth, nil
	}
}

func TestCLIToken_SourceOrder(t *testing.T) {
	cases := []struct {
		name      string
		cfg       string
		oauth     fakeOAuth
		want      string
		wantOAuth bool
	}{
		{"api_token wins over OAuth", "cfg", fakeOAuth{tok: "oauth"}, "cfg", false},
		{"blank api_token is ignored", "  ", fakeOAuth{tok: "oauth"}, "oauth", true},
		{"OAuth last", "", fakeOAuth{tok: "oauth"}, "oauth", true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			oauth := tc.oauth
			stubSources(t, tc.cfg, &oauth, nil)
			got, err := CLIToken(context.Background(), "")
			if err != nil || got != tc.want {
				t.Fatalf("CLIToken = %q, %v; want %q", got, err, tc.want)
			}
			if oauth.called != tc.wantOAuth {
				t.Errorf("OAuth consulted = %v, want %v", oauth.called, tc.wantOAuth)
			}
		})
	}
}

func TestCLIToken_OAuthErrors(t *testing.T) {
	cases := []struct {
		name     string
		oauth    fakeOAuth
		buildErr error
		want     error
	}{
		{"not signed in", fakeOAuth{err: api.ErrNotSignedIn}, nil, ErrOwnerNoToken},
		{"dead refresh token", fakeOAuth{err: fmt.Errorf("%w: refresh token rejected (status 400)", api.ErrNotSignedIn)}, nil, ErrOwnerNoToken},
		{"empty token", fakeOAuth{}, nil, ErrOwnerNoToken},
		{"no credentials path", fakeOAuth{}, errors.New("no home directory"), ErrOwnerNoToken},
		{"refresh 5xx", fakeOAuth{err: errors.New("token refresh failed (status 503)")}, nil, ErrOwnerTemporary},
		{"lock wait timed out", fakeOAuth{err: fmt.Errorf("waiting for lock: %w", context.DeadlineExceeded)}, nil, ErrOwnerTemporary},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			oauth := tc.oauth
			stubSources(t, "", &oauth, tc.buildErr)
			_, err := CLIToken(context.Background(), "")
			if !errors.Is(err, tc.want) {
				t.Fatalf("err = %v, want %v", err, tc.want)
			}
		})
	}
}

func TestDefaultOwnerOptions(t *testing.T) {
	viper.Set("api_url", "https://api.example.test")
	t.Cleanup(func() { viper.Set("api_url", "") })
	o := DefaultOwnerOptions(true)
	if o.APIURL != "https://api.example.test" || !o.NoOwner || o.Token == nil || o.UserAgent == "" {
		t.Fatalf("options = %+v", o)
	}
}

// public-api rejects fractional seconds (422), so expires_at is pinned to whole
// seconds in UTC whatever the connector's NotAfter carries.
func TestOwnerTicket_ExpiresAtWholeSecondsUTC(t *testing.T) {
	notAfter := time.Date(2026, 10, 23, 14, 5, 7, 987654321, time.FixedZone("EEST", 3*3600))
	if got, want := formatExpiresAt(notAfter), "2026-10-23T11:05:07Z"; got != want {
		t.Fatalf("formatExpiresAt = %q, want %q", got, want)
	}

	srv := newFakePublicAPI(t, http.StatusCreated, `{"ticket":"t","expires_at":"2026-10-23T11:05:07Z","client_id":"1"}`, 0)
	fn := NewOwnerTicketFunc(OwnerOptions{APIURL: srv.URL, Token: staticSource("tok")})
	if _, err := fn(context.Background(), testTunnelID, notAfter); err != nil {
		t.Fatal(err)
	}
	if got := srv.req.Load().body["expires_at"]; got != "2026-10-23T11:05:07Z" {
		t.Fatalf("sent expires_at = %v, want 2026-10-23T11:05:07Z", got)
	}
}

// api_token is read through viper, which resolves HOSTINGER_API_TOKEN, then
// the legacy HAPI_API_TOKEN (utils.BindLegacyEnv), then the config file: the
// same lookup as api/request.go.
func TestStaticToken_ViperResolvesEnv(t *testing.T) {
	t.Cleanup(viper.Reset)
	viper.Reset()
	t.Setenv("HOSTINGER_API_TOKEN", "")
	t.Setenv("HAPI_API_TOKEN", " legacy ")
	utils.BindLegacyEnv(io.Discard)
	if got := StaticToken(); got != "legacy" {
		t.Fatalf("StaticToken = %q, want legacy", got)
	}
	t.Setenv("HOSTINGER_API_TOKEN", "current")
	if got := StaticToken(); got != "current" {
		t.Fatalf("StaticToken = %q, want current", got)
	}
}

func TestCLIToken_ReplacesRejectedSignIn(t *testing.T) {
	oauth := fakeOAuth{tok: "at-new"}
	stubSources(t, "", &oauth, nil)
	got, err := CLIToken(context.Background(), "at-old")
	if err != nil || got != "at-new" || oauth.replaced != "at-old" {
		t.Fatalf("CLIToken = %q, %v (replaced %q)", got, err, oauth.replaced)
	}

	static := fakeOAuth{tok: "unused"}
	stubSources(t, "api-tok", &static, nil)
	if got, _ := CLIToken(context.Background(), "api-tok"); got != "api-tok" || static.called {
		t.Fatalf("a static token must come back unchanged without OAuth, got %q", got)
	}
}

// ticketServer answers 201 for the good token and 401 for any other.
func ticketServer(t *testing.T, good string) (*httptest.Server, *atomic.Int32) {
	t.Helper()
	var calls atomic.Int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		calls.Add(1)
		w.Header().Set("Content-Type", "application/json")
		if r.Header.Get("Authorization") != "Bearer "+good {
			w.WriteHeader(http.StatusUnauthorized)
			_, _ = w.Write([]byte(`{"message":"Unauthenticated."}`))
			return
		}
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte(`{"ticket":"t.i.k","expires_at":"2026-10-23T11:05:00Z","client_id":"1"}`))
	}))
	t.Cleanup(srv.Close)
	return srv, &calls
}

func TestOwnerTicket_401RetriesOnceWithAnotherToken(t *testing.T) {
	cases := []struct {
		name       string
		source     TokenSource
		wantTicket string
		wantCalls  int32
	}{
		{"fresh sign-in token", func(_ context.Context, rejected string) (string, error) {
			if rejected == "" {
				return "at-revoked", nil
			}
			return "at-fresh", nil
		}, "t.i.k", 2},
		{"static token comes back unchanged", staticSource("at-revoked"), "", 1},
		{"replacement fails", func(_ context.Context, rejected string) (string, error) {
			if rejected == "" {
				return "at-revoked", nil
			}
			return "", ErrOwnerNoToken
		}, "", 1},
		{"replacement is rejected too", func(_ context.Context, rejected string) (string, error) {
			if rejected == "" {
				return "at-revoked", nil
			}
			return "at-also-bad", nil
		}, "", 2},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			srv, calls := ticketServer(t, "at-fresh")
			fn := NewOwnerTicketFunc(OwnerOptions{APIURL: srv.URL, Token: tc.source})
			ticket, err := fn(context.Background(), testTunnelID, testNotAfter)
			if ticket != tc.wantTicket {
				t.Fatalf("ticket = %q, %v; want %q", ticket, err, tc.wantTicket)
			}
			if tc.wantTicket == "" && (err == nil || !strings.Contains(err.Error(), "HTTP 401") ||
				errors.Is(err, ErrOwnerTemporary) || errors.Is(err, ErrOwnerNoToken)) {
				t.Fatalf("err = %v, want the plain 401 error", err)
			}
			if n := calls.Load(); n != tc.wantCalls {
				t.Fatalf("public-api calls = %d, want %d", n, tc.wantCalls)
			}
		})
	}
}
