package api

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync/atomic"
	"testing"
	"time"
)

// newTestService builds a white-box service pointed at the given issuer with an
// isolated temp credentials file, so tests never touch the real ~/.config path.
func newTestService(t *testing.T, issuer string) *OAuthService {
	t.Helper()
	return &OAuthService{
		issuer:     strings.TrimRight(issuer, "/"),
		clientName: defaultClientName,
		credPath:   filepath.Join(t.TempDir(), "credentials.json"),
		httpClient: &http.Client{},
	}
}

func TestPKCE_S256Relationship(t *testing.T) {
	verifier, challenge, err := pkce()
	if err != nil {
		t.Fatalf("pkce: %v", err)
	}
	// 32 raw bytes base64url-encoded is 43 chars.
	if len(verifier) != 43 {
		t.Errorf("verifier length = %d, want 43", len(verifier))
	}
	sum := sha256.Sum256([]byte(verifier))
	want := base64.RawURLEncoding.EncodeToString(sum[:])
	if challenge != want {
		t.Errorf("challenge = %q, want sha256(verifier) = %q", challenge, want)
	}
	// Distinct each call.
	v2, _, _ := pkce()
	if verifier == v2 {
		t.Error("two pkce verifiers were identical")
	}
}

func TestRandomState(t *testing.T) {
	s1, err := randomState()
	if err != nil {
		t.Fatalf("randomState: %v", err)
	}
	if len(s1) != 32 { // 16 bytes hex
		t.Errorf("state length = %d, want 32", len(s1))
	}
	s2, _ := randomState()
	if s1 == s2 {
		t.Error("two states were identical")
	}
}

func TestExpiresAtMillis(t *testing.T) {
	before := time.Now().Add((3600 - expiryBufferSeconds) * time.Second).UnixMilli()
	got := expiresAtMillis(3600)
	after := time.Now().Add((3600 - expiryBufferSeconds) * time.Second).UnixMilli()
	if got < before || got > after {
		t.Errorf("expiresAtMillis(3600) = %d, want within [%d,%d]", got, before, after)
	}

	// A lifetime shorter than the buffer must clamp to ~now, never to the past.
	lower := time.Now().UnixMilli()
	clamped := expiresAtMillis(10)
	upper := time.Now().Add(time.Second).UnixMilli()
	if clamped < lower || clamped > upper {
		t.Errorf("expiresAtMillis(10) = %d, want clamped to ~now [%d,%d]", clamped, lower, upper)
	}
}

func TestSaveAndLoadCredentials(t *testing.T) {
	s := newTestService(t, "http://unused")
	in := credentials{ClientID: "cid", AccessToken: "at", RefreshToken: "rt", ExpiresAt: 123}

	if err := s.saveCredentials(in); err != nil {
		t.Fatalf("saveCredentials: %v", err)
	}
	got, err := s.loadCredentials()
	if err != nil {
		t.Fatalf("loadCredentials: %v", err)
	}
	if got != in {
		t.Errorf("round-trip = %+v, want %+v", got, in)
	}

	if runtime.GOOS != "windows" {
		info, err := os.Stat(s.credPath)
		if err != nil {
			t.Fatalf("stat: %v", err)
		}
		if perm := info.Mode().Perm(); perm != 0600 {
			t.Errorf("file mode = %o, want 0600", perm)
		}
	}
}

func TestLoadCredentials_MissingAndCorrupt(t *testing.T) {
	s := newTestService(t, "http://unused")

	// Missing file -> empty, no error.
	got, err := s.loadCredentials()
	if err != nil || got != (credentials{}) {
		t.Fatalf("missing file: got (%+v, %v), want (empty, nil)", got, err)
	}

	// Corrupt JSON -> empty, no error.
	if err := os.WriteFile(s.credPath, []byte("{not json"), 0600); err != nil {
		t.Fatal(err)
	}
	got, err = s.loadCredentials()
	if err != nil || got != (credentials{}) {
		t.Fatalf("corrupt file: got (%+v, %v), want (empty, nil)", got, err)
	}
}

func TestAuthorizeURL(t *testing.T) {
	s := newTestService(t, "https://auth.example.com")
	raw := s.authorizeURL("cid", "http://127.0.0.1:5000/oauth/callback", "state123", "chal")

	u, err := url.Parse(raw)
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	if u.Host != "auth.example.com" || u.Path != authorizePath {
		t.Errorf("host/path = %s%s", u.Host, u.Path)
	}
	q := u.Query()
	checks := map[string]string{
		"client_id":             "cid",
		"redirect_uri":          "http://127.0.0.1:5000/oauth/callback",
		"state":                 "state123",
		"code_challenge":        "chal",
		"code_challenge_method": "S256",
		"response_type":         "code",
	}
	for k, want := range checks {
		if got := q.Get(k); got != want {
			t.Errorf("query %q = %q, want %q", k, got, want)
		}
	}
}

func TestRefresh_Success_RotatesToken(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != tokenPath {
			t.Errorf("unexpected path %s", r.URL.Path)
		}
		_ = r.ParseForm()
		if r.Form.Get("grant_type") != "refresh_token" {
			t.Errorf("grant_type = %q", r.Form.Get("grant_type"))
		}
		writeJSON(w, 200, tokenResponse{AccessToken: "new-at", RefreshToken: "new-rt", ExpiresIn: 3600})
	}))
	defer srv.Close()

	s := newTestService(t, srv.URL)
	token, err := s.refresh(context.Background(), credentials{ClientID: "cid", RefreshToken: "old-rt"})
	if err != nil {
		t.Fatalf("refresh: %v", err)
	}
	if token != "new-at" {
		t.Errorf("token = %q, want new-at", token)
	}

	got, _ := s.loadCredentials()
	if got.AccessToken != "new-at" || got.RefreshToken != "new-rt" {
		t.Errorf("persisted = %+v, want at=new-at rt=new-rt", got)
	}
	if got.ExpiresAt <= nowMillis() {
		t.Errorf("expires_at = %d not in the future", got.ExpiresAt)
	}
}

func TestRefresh_KeepsOldRefreshTokenWhenOmitted(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, 200, tokenResponse{AccessToken: "new-at", ExpiresIn: 3600}) // no refresh_token
	}))
	defer srv.Close()

	s := newTestService(t, srv.URL)
	if _, err := s.refresh(context.Background(), credentials{ClientID: "cid", RefreshToken: "keep-me"}); err != nil {
		t.Fatalf("refresh: %v", err)
	}
	got, _ := s.loadCredentials()
	if got.RefreshToken != "keep-me" {
		t.Errorf("refresh token = %q, want keep-me", got.RefreshToken)
	}
}

func TestRefresh_4xxIsDead(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid_grant"})
	}))
	defer srv.Close()

	s := newTestService(t, srv.URL)
	_, err := s.refresh(context.Background(), credentials{ClientID: "cid", RefreshToken: "rt"})
	if !errors.Is(err, errRefreshTokenDead) {
		t.Fatalf("err = %v, want errRefreshTokenDead", err)
	}
}

func TestRefresh_5xxIsTransient(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "boom", http.StatusInternalServerError)
	}))
	defer srv.Close()

	s := newTestService(t, srv.URL)
	_, err := s.refresh(context.Background(), credentials{ClientID: "cid", RefreshToken: "rt"})
	if err == nil {
		t.Fatal("expected error on 5xx")
	}
	if errors.Is(err, errRefreshTokenDead) {
		t.Error("5xx must not be classified as dead refresh token")
	}
}

func TestRegister(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != registerPath {
			t.Errorf("unexpected path %s", r.URL.Path)
		}
		var body map[string]any
		_ = json.NewDecoder(r.Body).Decode(&body)
		if body["client_name"] != defaultClientName {
			t.Errorf("client_name = %v", body["client_name"])
		}
		writeJSON(w, 200, map[string]any{"client_id": "client-xyz"})
	}))
	defer srv.Close()

	s := newTestService(t, srv.URL)
	id, err := s.register(context.Background(), "http://127.0.0.1:1/oauth/callback")
	if err != nil {
		t.Fatalf("register: %v", err)
	}
	if id != "client-xyz" {
		t.Errorf("client_id = %q, want client-xyz", id)
	}
}

func TestRegister_MissingClientID(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, 200, map[string]any{}) // no client_id
	}))
	defer srv.Close()

	s := newTestService(t, srv.URL)
	if _, err := s.register(context.Background(), "http://127.0.0.1:1/oauth/callback"); err == nil {
		t.Fatal("expected error when client_id is missing")
	}
}

func TestRegister_4xx(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid_redirect_uri"})
	}))
	defer srv.Close()

	s := newTestService(t, srv.URL)
	if _, err := s.register(context.Background(), "http://127.0.0.1:1/oauth/callback"); err == nil {
		t.Fatal("expected error on 4xx registration")
	}
}

func TestExchangeCode_PersistsCreds(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_ = r.ParseForm()
		if r.Form.Get("grant_type") != "authorization_code" {
			t.Errorf("grant_type = %q", r.Form.Get("grant_type"))
		}
		if r.Form.Get("code") != "the-code" || r.Form.Get("code_verifier") == "" {
			t.Errorf("missing code/verifier: %v", r.Form)
		}
		writeJSON(w, 200, tokenResponse{AccessToken: "at", RefreshToken: "rt", ExpiresIn: 3600})
	}))
	defer srv.Close()

	s := newTestService(t, srv.URL)
	token, err := s.exchangeCode(context.Background(), "cid", "the-code", "verifier", "http://127.0.0.1:1/oauth/callback")
	if err != nil {
		t.Fatalf("exchangeCode: %v", err)
	}
	if token != "at" {
		t.Errorf("token = %q, want at", token)
	}
	got, _ := s.loadCredentials()
	if got.ClientID != "cid" || got.AccessToken != "at" || got.RefreshToken != "rt" {
		t.Errorf("persisted creds = %+v", got)
	}
}

func TestResolve_HappyPathMakesNoHTTPCalls(t *testing.T) {
	var hits int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(&hits, 1)
	}))
	defer srv.Close()

	s := newTestService(t, srv.URL)
	// Valid, unexpired cached token.
	_ = s.saveCredentials(credentials{
		ClientID:     "cid",
		AccessToken:  "cached-at",
		RefreshToken: "rt",
		ExpiresAt:    nowMillis() + 60_000,
	})

	token, err := s.resolve(context.Background())
	if err != nil {
		t.Fatalf("resolve: %v", err)
	}
	if token != "cached-at" {
		t.Errorf("token = %q, want cached-at", token)
	}
	if n := atomic.LoadInt32(&hits); n != 0 {
		t.Errorf("made %d HTTP calls on the happy path, want 0", n)
	}
}

func TestResolve_ExpiredTokenTriggersRefresh(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_ = r.ParseForm()
		if r.Form.Get("grant_type") != "refresh_token" {
			t.Errorf("expected refresh grant, got %q", r.Form.Get("grant_type"))
		}
		writeJSON(w, 200, tokenResponse{AccessToken: "refreshed-at", RefreshToken: "rt2", ExpiresIn: 3600})
	}))
	defer srv.Close()

	s := newTestService(t, srv.URL)
	_ = s.saveCredentials(credentials{
		ClientID:     "cid",
		AccessToken:  "stale-at",
		RefreshToken: "rt",
		ExpiresAt:    nowMillis() - 1000, // expired
	})

	token, err := s.resolve(context.Background())
	if err != nil {
		t.Fatalf("resolve: %v", err)
	}
	if token != "refreshed-at" {
		t.Errorf("token = %q, want refreshed-at", token)
	}
}

func TestAwaitCallback_Success(t *testing.T) {
	s := newTestService(t, "http://unused")
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	port := ln.Addr().(*net.TCPAddr).Port

	type res struct {
		code string
		err  error
	}
	out := make(chan res, 1)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	go func() {
		code, err := s.awaitCallback(ctx, ln, "state-ok")
		out <- res{code, err}
	}()

	// A non-callback path must 404 and leave the flow untouched.
	resp404, err := http.Get(fmt.Sprintf("http://127.0.0.1:%d/whatever", port))
	if err != nil {
		t.Fatalf("get /whatever: %v", err)
	}
	if resp404.StatusCode != http.StatusNotFound {
		t.Errorf("wrong-path status = %d, want 404", resp404.StatusCode)
	}
	resp404.Body.Close()

	// The real callback resolves the flow.
	resp, err := http.Get(fmt.Sprintf("http://127.0.0.1:%d/oauth/callback?code=abc&state=state-ok", port))
	if err != nil {
		t.Fatalf("get callback: %v", err)
	}
	body, _ := io.ReadAll(resp.Body)
	resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("callback status = %d, want 200", resp.StatusCode)
	}
	// The embedded success template should have rendered into the response.
	if !strings.Contains(strings.ToLower(string(body)), "signed in") {
		t.Errorf("success page body missing expected content: %q", string(body))
	}

	r := <-out
	if r.err != nil || r.code != "abc" {
		t.Errorf("awaitCallback = (%q, %v), want (abc, nil)", r.code, r.err)
	}
}

func TestAwaitCallback_StateMismatch(t *testing.T) {
	code, err := runCallback(t, "expected-state", "/oauth/callback?code=abc&state=wrong")
	if err == nil || !strings.Contains(err.Error(), "state mismatch") {
		t.Fatalf("err = %v, want state mismatch", err)
	}
	if code != "" {
		t.Errorf("code = %q, want empty", code)
	}
}

func TestAwaitCallback_AuthServerError(t *testing.T) {
	code, err := runCallback(t, "s", "/oauth/callback?error=access_denied&state=s")
	if err == nil || !strings.Contains(err.Error(), "access_denied") {
		t.Fatalf("err = %v, want access_denied", err)
	}
	if code != "" {
		t.Errorf("code = %q, want empty", code)
	}
}

// runCallback drives awaitCallback through a single GET and returns its result.
func runCallback(t *testing.T, expectedState, requestPath string) (string, error) {
	t.Helper()
	s := newTestService(t, "http://unused")
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	port := ln.Addr().(*net.TCPAddr).Port

	type res struct {
		code string
		err  error
	}
	out := make(chan res, 1)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	go func() {
		code, err := s.awaitCallback(ctx, ln, expectedState)
		out <- res{code, err}
	}()

	resp, err := http.Get(fmt.Sprintf("http://127.0.0.1:%d%s", port, requestPath))
	if err != nil {
		t.Fatalf("get: %v", err)
	}
	io.Copy(io.Discard, resp.Body)
	resp.Body.Close()

	r := <-out
	return r.code, r.err
}

func TestErrorTemplate_EscapesUntrustedInput(t *testing.T) {
	var buf bytes.Buffer
	if err := errorTemplate.Execute(&buf, `<script>alert('x')</script>`); err != nil {
		t.Fatalf("execute error template: %v", err)
	}
	page := buf.String()
	if strings.Contains(page, "<script>") {
		t.Error("error template did not escape the reason")
	}
	if !strings.Contains(page, "&lt;script&gt;") {
		t.Error("error template missing escaped entities")
	}
}

func TestNewOAuthService_IssuerFromEnv(t *testing.T) {
	t.Setenv("HOSTINGER_OAUTH_ISSUER", "https://issuer.example.com/")
	s, err := NewOAuthService()
	if err != nil {
		t.Fatalf("NewOAuthService: %v", err)
	}
	if s.issuer != "https://issuer.example.com" { // trailing slash stripped
		t.Errorf("issuer = %q", s.issuer)
	}
	if s.clientName != defaultClientName {
		t.Errorf("clientName = %q", s.clientName)
	}
}

// TestLogin_HappyPath exercises register -> authorize -> callback -> exchange end
// to end against an httptest auth server, substituting only the browser launch
// (which a test cannot perform) to replay the redirect to the loopback callback.
func TestLogin_HappyPath(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case registerPath:
			writeJSON(w, 200, map[string]any{"client_id": "client-login"})
		case tokenPath:
			_ = r.ParseForm()
			if r.Form.Get("client_id") != "client-login" {
				t.Errorf("token client_id = %q", r.Form.Get("client_id"))
			}
			writeJSON(w, 200, tokenResponse{AccessToken: "login-at", RefreshToken: "login-rt", ExpiresIn: 3600})
		default:
			http.NotFound(w, r)
		}
	}))
	defer srv.Close()

	orig := openBrowser
	defer func() { openBrowser = orig }()
	openBrowser = func(rawURL string) error {
		u, err := url.Parse(rawURL)
		if err != nil {
			return err
		}
		q := u.Query()
		redirect := q.Get("redirect_uri")
		state := q.Get("state")
		// Simulate the auth server redirecting back to the loopback listener.
		// Fired asynchronously so login() can reach awaitCallback and serve it.
		go func() {
			resp, err := http.Get(redirect + "?code=login-code&state=" + url.QueryEscape(state))
			if err == nil {
				resp.Body.Close()
			}
		}()
		return nil
	}

	s := newTestService(t, srv.URL)
	token, err := s.login(context.Background())
	if err != nil {
		t.Fatalf("login: %v", err)
	}
	if token != "login-at" {
		t.Errorf("token = %q, want login-at", token)
	}

	got, _ := s.loadCredentials()
	if got.ClientID != "client-login" || got.AccessToken != "login-at" || got.RefreshToken != "login-rt" {
		t.Errorf("persisted creds = %+v", got)
	}
}

func TestLogout_RevokesAndWipesTokens(t *testing.T) {
	var revoked bool
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == revokePath {
			revoked = true
			_ = r.ParseForm()
			if r.Form.Get("token") != "at" {
				t.Errorf("revoke token = %q", r.Form.Get("token"))
			}
		}
		w.WriteHeader(200)
	}))
	defer srv.Close()

	s := newTestService(t, srv.URL)
	_ = s.saveCredentials(credentials{ClientID: "cid", AccessToken: "at", RefreshToken: "rt", ExpiresAt: 999})

	if err := s.Logout(context.Background()); err != nil {
		t.Fatalf("logout: %v", err)
	}
	if !revoked {
		t.Error("revoke endpoint was not called")
	}
	got, _ := s.loadCredentials()
	if got.ClientID != "cid" {
		t.Errorf("client_id = %q, want preserved cid", got.ClientID)
	}
	if got.AccessToken != "" || got.RefreshToken != "" {
		t.Errorf("tokens not wiped: %+v", got)
	}
}

func TestLogout_NoCredentialsIsNoOp(t *testing.T) {
	s := newTestService(t, "http://unused")
	if err := s.Logout(context.Background()); err != nil {
		t.Fatalf("logout with no creds: %v", err)
	}
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}
