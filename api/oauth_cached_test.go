package api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

// rotatingTokenServer is a fake auth server whose refresh grant rotates the
// refresh token: the old one is rejected once used, as SSO may do. A second
// refresh with the same token therefore fails, which is what CachedToken's
// cross-process lock must prevent.
type rotatingTokenServer struct {
	*httptest.Server
	refreshes atomic.Int32 // successful refresh grants
	rejected  atomic.Int32 // refresh grants with a stale token
	other     atomic.Int32 // any non-refresh request (register, authorize, ...)

	mu      sync.Mutex
	current string // the refresh token the server accepts next
}

func newRotatingTokenServer(t *testing.T, refreshToken string, delay time.Duration) *rotatingTokenServer {
	t.Helper()
	s := &rotatingTokenServer{current: refreshToken}
	s.Server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != tokenPath || r.ParseForm() != nil || r.Form.Get("grant_type") != "refresh_token" {
			s.other.Add(1)
			http.Error(w, "unexpected", http.StatusInternalServerError)
			return
		}
		time.Sleep(delay) // widen the window in which a second refresh could race
		s.mu.Lock()
		defer s.mu.Unlock()
		if r.Form.Get("refresh_token") != s.current {
			s.rejected.Add(1)
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid_grant"})
			return
		}
		n := s.refreshes.Add(1)
		s.current = fmt.Sprintf("rt-rotated-%d", n)
		writeJSON(w, http.StatusOK, map[string]any{
			"access_token":  fmt.Sprintf("at-fresh-%d", n),
			"refresh_token": s.current,
			"expires_in":    3600,
		})
	}))
	t.Cleanup(s.Close)
	return s
}

func writeCreds(t *testing.T, path string, c credentials) []byte {
	t.Helper()
	if err := os.MkdirAll(filepath.Dir(path), 0o700); err != nil {
		t.Fatal(err)
	}
	b, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, b, 0o600); err != nil {
		t.Fatal(err)
	}
	return b
}

func expiredCreds() credentials {
	return credentials{
		ClientID:     "client-1",
		AccessToken:  "at-expired",
		RefreshToken: "rt-1",
		ExpiresAt:    time.Now().Add(-time.Minute).UnixMilli(),
	}
}

// failOnLogin makes any interactive sign-in attempt fail the test.
func failOnLogin(t *testing.T) {
	t.Helper()
	prev := openBrowser
	openBrowser = func(string) error {
		t.Error("CachedToken attempted an interactive login")
		return errors.New("login not allowed")
	}
	t.Cleanup(func() { openBrowser = prev })
}

func TestCachedToken_ValidTokenNoNetworkNoLock(t *testing.T) {
	failOnLogin(t)
	srv := newRotatingTokenServer(t, "rt-1", 0)
	svc := newTestService(t, srv.URL)
	writeCreds(t, svc.credPath, credentials{
		ClientID: "client-1", AccessToken: "at-valid", RefreshToken: "rt-1",
		ExpiresAt: time.Now().Add(time.Hour).UnixMilli(),
	})

	tok, err := svc.CachedToken(context.Background())
	if err != nil || tok != "at-valid" {
		t.Fatalf("CachedToken = %q, %v; want at-valid", tok, err)
	}
	if n := srv.refreshes.Load() + srv.rejected.Load() + srv.other.Load(); n != 0 {
		t.Fatalf("made %d HTTP calls, want 0", n)
	}
	if _, err := os.Stat(svc.credPath + ".lock"); !os.IsNotExist(err) {
		t.Fatalf("fast path created the lock file (stat err=%v)", err)
	}
}

func TestCachedToken_NoCredentials(t *testing.T) {
	failOnLogin(t)
	srv := newRotatingTokenServer(t, "rt-1", 0)
	svc := newTestService(t, srv.URL)

	if _, err := svc.CachedToken(context.Background()); !errors.Is(err, ErrNotSignedIn) {
		t.Fatalf("err = %v, want ErrNotSignedIn", err)
	}
	if _, err := os.Stat(svc.credPath + ".lock"); !os.IsNotExist(err) {
		t.Fatalf("lock file created without credentials (stat err=%v)", err)
	}
	if srv.other.Load() != 0 {
		t.Fatal("unexpected auth-server call")
	}
}

func TestCachedToken_RefreshesExpiredToken(t *testing.T) {
	failOnLogin(t)
	srv := newRotatingTokenServer(t, "rt-1", 0)
	svc := newTestService(t, srv.URL)
	writeCreds(t, svc.credPath, expiredCreds())

	tok, err := svc.CachedToken(context.Background())
	if err != nil || tok != "at-fresh-1" {
		t.Fatalf("CachedToken = %q, %v; want at-fresh-1", tok, err)
	}
	creds, _ := svc.loadCredentials()
	if creds.AccessToken != "at-fresh-1" || creds.RefreshToken != "rt-rotated-1" || !accessTokenValid(creds) {
		t.Fatalf("stored creds not updated: %+v", creds)
	}
	// The second call is served from the file.
	if tok, err := svc.CachedToken(context.Background()); err != nil || tok != "at-fresh-1" {
		t.Fatalf("second CachedToken = %q, %v", tok, err)
	}
	if n := srv.refreshes.Load(); n != 1 {
		t.Fatalf("refreshes = %d, want 1", n)
	}
}

func TestCachedToken_DeadRefreshTokenLeavesCredentials(t *testing.T) {
	failOnLogin(t)
	srv := newRotatingTokenServer(t, "rt-other", 0) // rejects rt-1
	svc := newTestService(t, srv.URL)
	before := writeCreds(t, svc.credPath, expiredCreds())

	_, err := svc.CachedToken(context.Background())
	if !errors.Is(err, ErrNotSignedIn) {
		t.Fatalf("err = %v, want ErrNotSignedIn", err)
	}
	after, _ := os.ReadFile(svc.credPath)
	if !bytes.Equal(before, after) {
		t.Fatalf("credentials changed on failure:\n%s", after)
	}
	if srv.other.Load() != 0 {
		t.Fatal("unexpected non-refresh auth-server call (login?)")
	}
}

func TestCachedToken_TransientFailureLeavesCredentials(t *testing.T) {
	failOnLogin(t)
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "down", http.StatusServiceUnavailable)
	}))
	defer srv.Close()
	svc := newTestService(t, srv.URL)
	before := writeCreds(t, svc.credPath, expiredCreds())

	_, err := svc.CachedToken(context.Background())
	if err == nil || errors.Is(err, ErrNotSignedIn) {
		t.Fatalf("err = %v, want a transient (non-ErrNotSignedIn) error", err)
	}
	after, _ := os.ReadFile(svc.credPath)
	if !bytes.Equal(before, after) {
		t.Fatalf("credentials changed on failure:\n%s", after)
	}
}

func TestCachedToken_LockWaitHonoursDeadline(t *testing.T) {
	failOnLogin(t)
	srv := newRotatingTokenServer(t, "rt-1", 0)
	svc := newTestService(t, srv.URL)
	before := writeCreds(t, svc.credPath, expiredCreds())

	unlock, err := lockFile(context.Background(), svc.credPath+".lock")
	if err != nil {
		t.Fatal(err)
	}
	defer unlock()

	ctx, cancel := context.WithTimeout(context.Background(), 150*time.Millisecond)
	defer cancel()
	start := time.Now()
	_, err = svc.CachedToken(ctx)
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("err = %v, want context.DeadlineExceeded", err)
	}
	if d := time.Since(start); d > 2*time.Second {
		t.Fatalf("lock wait took %s, deadline was 150ms", d)
	}
	if n := srv.refreshes.Load() + srv.rejected.Load(); n != 0 {
		t.Fatalf("refreshed %d times while another holder had the lock", n)
	}
	after, _ := os.ReadFile(svc.credPath)
	if !bytes.Equal(before, after) {
		t.Fatal("credentials changed on failure")
	}
}

// Environment of the helper process used by TestCachedToken_CrossProcess.
const (
	helperEnv       = "API_CACHED_TOKEN_HELPER"
	helperCredsEnv  = "API_CACHED_TOKEN_CREDS"
	helperIssuerEnv = "API_CACHED_TOKEN_ISSUER"
	helperGoEnv     = "API_CACHED_TOKEN_GO"
	helperMethodEnv = "API_CACHED_TOKEN_METHOD"
)

// helperMethods are the refresh paths the cross-process test runs side by
// side: they must all take the same lock.
var helperMethods = []string{"CachedToken", "Token", "TokenNoLogin"}

// TestCachedToken_CrossProcess starts three separate processes (this test
// binary, re-executed) that call CachedToken, Token and TokenNoLogin at the
// same moment against one credentials file with an expired access token. The server rotates the
// refresh token, so without the cross-process lock all but one would get a
// rejected refresh. Exactly one refresh must happen, every process must get
// the fresh token, the file must stay valid, and nothing may try to sign in.
func TestCachedToken_CrossProcess(t *testing.T) {
	if os.Getenv(helperEnv) == "1" {
		t.Skip("running as helper")
	}
	srv := newRotatingTokenServer(t, "rt-1", 300*time.Millisecond)
	dir := t.TempDir()
	credPath := filepath.Join(dir, "api-cli", "credentials.json")
	writeCreds(t, credPath, expiredCreds())
	goFile := filepath.Join(dir, "go")

	const procs = 3
	cmds := make([]*exec.Cmd, procs)
	outs := make([]*bytes.Buffer, procs)
	for i := range cmds {
		cmd := exec.Command(os.Args[0], "-test.run=^TestCachedTokenHelperProcess$", "-test.count=1")
		cmd.Env = append(os.Environ(),
			helperEnv+"=1",
			helperCredsEnv+"="+credPath,
			helperIssuerEnv+"="+srv.URL,
			helperGoEnv+"="+goFile,
			helperMethodEnv+"="+helperMethods[i%len(helperMethods)],
		)
		outs[i] = new(bytes.Buffer)
		cmd.Stdout = outs[i]
		cmd.Stderr = outs[i]
		if err := cmd.Start(); err != nil {
			t.Fatal(err)
		}
		cmds[i] = cmd
	}
	// Release all helpers at once.
	if err := os.WriteFile(goFile, nil, 0o600); err != nil {
		t.Fatal(err)
	}

	for i, cmd := range cmds {
		if err := cmd.Wait(); err != nil {
			t.Errorf("helper %d failed: %v\n%s", i, err, outs[i])
			continue
		}
		var tok string
		for _, line := range strings.Split(outs[i].String(), "\n") {
			if v, ok := strings.CutPrefix(line, "TOKEN="); ok {
				tok = v
			}
		}
		if tok != "at-fresh-1" {
			t.Errorf("helper %d token = %q, want at-fresh-1\n%s", i, tok, outs[i])
		}
	}

	if n := srv.refreshes.Load(); n != 1 {
		t.Errorf("successful refreshes = %d, want exactly 1", n)
	}
	if n := srv.rejected.Load(); n != 0 {
		t.Errorf("refreshes with a stale (rotated) token = %d, want 0", n)
	}
	if n := srv.other.Load(); n != 0 {
		t.Errorf("non-refresh auth-server calls (login/register) = %d, want 0", n)
	}

	raw, err := os.ReadFile(credPath)
	if err != nil {
		t.Fatal(err)
	}
	var creds credentials
	if err := json.Unmarshal(raw, &creds); err != nil {
		t.Fatalf("credentials file is not valid JSON: %v\n%s", err, raw)
	}
	if creds.ClientID != "client-1" || creds.AccessToken != "at-fresh-1" ||
		creds.RefreshToken != "rt-rotated-1" || !accessTokenValid(creds) {
		t.Errorf("credentials after refresh = %+v", creds)
	}
}

// TestCachedTokenHelperProcess is the child side of TestCachedToken_CrossProcess.
func TestCachedTokenHelperProcess(t *testing.T) {
	if os.Getenv(helperEnv) != "1" {
		t.Skip("helper process only")
	}
	openBrowser = func(string) error {
		fmt.Println("LOGIN ATTEMPTED")
		os.Exit(3)
		return nil
	}
	goFile := os.Getenv(helperGoEnv)
	for deadline := time.Now().Add(10 * time.Second); ; time.Sleep(5 * time.Millisecond) {
		if _, err := os.Stat(goFile); err == nil {
			break
		}
		if time.Now().After(deadline) {
			t.Fatal("start signal never came")
		}
	}

	svc := &OAuthService{
		issuer:     os.Getenv(helperIssuerEnv),
		clientName: defaultClientName,
		credPath:   os.Getenv(helperCredsEnv),
		httpClient: &http.Client{},
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	method := os.Getenv(helperMethodEnv)
	var tok string
	var err error
	switch method {
	case "Token":
		tok, err = svc.Token(ctx)
	case "TokenNoLogin":
		tok, err = svc.TokenNoLogin(ctx)
	default:
		tok, err = svc.CachedToken(ctx)
	}
	if err != nil {
		t.Fatalf("%s: %v", method, err)
	}
	fmt.Println("TOKEN=" + tok)
}

// Reauthenticate runs after the server rejected the stored token, which may
// still look valid by its expiry: it must refresh instead of returning it, and
// it must take the lock like every other refresh path.
func TestReauthenticate_RefreshesRejectedToken(t *testing.T) {
	failOnLogin(t)
	srv := newRotatingTokenServer(t, "rt-1", 0)
	svc := newTestService(t, srv.URL)
	writeCreds(t, svc.credPath, credentials{
		ClientID: "client-1", AccessToken: "at-rejected", RefreshToken: "rt-1",
		ExpiresAt: time.Now().Add(time.Hour).UnixMilli(),
	})

	tok, err := svc.Reauthenticate(context.Background())
	if err != nil || tok != "at-fresh-1" {
		t.Fatalf("Reauthenticate = %q, %v; want at-fresh-1", tok, err)
	}
	if n := srv.refreshes.Load(); n != 1 {
		t.Fatalf("refreshes = %d, want 1", n)
	}
	if _, err := os.Stat(svc.credPath + ".lock"); err != nil {
		t.Fatalf("refresh did not take the lock: %v", err)
	}
}

// A process that waits on the lock picks up the token another process wrote
// meanwhile instead of refreshing with a refresh token that is now rotated.
func TestToken_UsesTokenRefreshedWhileWaiting(t *testing.T) {
	failOnLogin(t)
	srv := newRotatingTokenServer(t, "rt-1", 0)
	svc := newTestService(t, srv.URL)
	writeCreds(t, svc.credPath, expiredCreds())

	unlock, err := lockFile(context.Background(), svc.credPath+".lock")
	if err != nil {
		t.Fatal(err)
	}
	fresh, err := json.Marshal(credentials{
		ClientID: "client-1", AccessToken: "at-other", RefreshToken: "rt-other",
		ExpiresAt: time.Now().Add(time.Hour).UnixMilli(),
	})
	if err != nil {
		t.Fatal(err)
	}
	go func() {
		time.Sleep(100 * time.Millisecond)
		_ = os.WriteFile(svc.credPath, fresh, 0o600)
		unlock()
	}()

	tok, err := svc.Token(context.Background())
	if err != nil || tok != "at-other" {
		t.Fatalf("Token = %q, %v; want at-other", tok, err)
	}
	if n := srv.refreshes.Load() + srv.rejected.Load() + srv.other.Load(); n != 0 {
		t.Fatalf("made %d auth-server calls, want 0", n)
	}
}

// The refresh outlives the caller's deadline: when the server rotates the
// refresh token, its answer must still be saved, or the next call finds a dead
// refresh token on disk.
func TestCachedToken_RefreshOutlivesCallerDeadline(t *testing.T) {
	failOnLogin(t)
	srv := newRotatingTokenServer(t, "rt-1", 300*time.Millisecond)
	svc := newTestService(t, srv.URL)
	writeCreds(t, svc.credPath, expiredCreds())

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	_, _ = svc.CachedToken(ctx)

	tok, err := svc.CachedToken(context.Background())
	if err != nil || tok != "at-fresh-1" {
		t.Fatalf("next CachedToken = %q, %v; want at-fresh-1", tok, err)
	}
	if r, n := srv.rejected.Load(), srv.refreshes.Load(); r != 0 || n != 1 {
		t.Fatalf("refreshes = %d, rejected = %d; want 1 and 0", n, r)
	}
}

func TestReplaceToken(t *testing.T) {
	failOnLogin(t)
	srv := newRotatingTokenServer(t, "rt-1", 0)
	svc := newTestService(t, srv.URL)
	writeCreds(t, svc.credPath, credentials{
		ClientID: "client-1", AccessToken: "at-revoked", RefreshToken: "rt-1",
		ExpiresAt: time.Now().Add(time.Hour).UnixMilli(),
	})

	tok, err := svc.ReplaceToken(context.Background(), "at-revoked")
	if err != nil || tok != "at-fresh-1" {
		t.Fatalf("ReplaceToken = %q, %v; want at-fresh-1", tok, err)
	}
	// Another stored token than the rejected one is returned as is.
	tok, err = svc.ReplaceToken(context.Background(), "at-revoked")
	if err != nil || tok != "at-fresh-1" || srv.refreshes.Load() != 1 {
		t.Fatalf("second ReplaceToken = %q, %v after %d refreshes", tok, err, srv.refreshes.Load())
	}
}

func TestRefresh(t *testing.T) {
	failOnLogin(t)
	srv := newRotatingTokenServer(t, "rt-1", 0)
	svc := newTestService(t, srv.URL)

	if _, err := svc.Refresh(context.Background()); err == nil || !strings.Contains(err.Error(), "no refresh token") {
		t.Fatalf("Refresh without credentials: err = %v", err)
	}

	writeCreds(t, svc.credPath, expiredCreds())
	tok, err := svc.Refresh(context.Background())
	if err != nil || tok != "at-fresh-1" {
		t.Fatalf("Refresh = %q, %v; want at-fresh-1", tok, err)
	}
	if _, err := os.Stat(svc.credPath + ".lock"); err != nil {
		t.Fatalf("Refresh did not take the lock: %v", err)
	}
}

func TestTokenNoLogin(t *testing.T) {
	failOnLogin(t)

	t.Run("valid token, no network", func(t *testing.T) {
		srv := newRotatingTokenServer(t, "rt-1", 0)
		svc := newTestService(t, srv.URL)
		writeCreds(t, svc.credPath, credentials{
			ClientID: "client-1", AccessToken: "at-valid", RefreshToken: "rt-1",
			ExpiresAt: time.Now().Add(time.Hour).UnixMilli(),
		})
		if tok, err := svc.TokenNoLogin(context.Background()); err != nil || tok != "at-valid" {
			t.Fatalf("TokenNoLogin = %q, %v", tok, err)
		}
		if n := srv.refreshes.Load() + srv.rejected.Load() + srv.other.Load(); n != 0 {
			t.Fatalf("made %d auth-server calls, want 0", n)
		}
	})
	t.Run("expired token refreshes under the lock", func(t *testing.T) {
		srv := newRotatingTokenServer(t, "rt-1", 0)
		svc := newTestService(t, srv.URL)
		writeCreds(t, svc.credPath, expiredCreds())
		if tok, err := svc.TokenNoLogin(context.Background()); err != nil || tok != "at-fresh-1" {
			t.Fatalf("TokenNoLogin = %q, %v", tok, err)
		}
		if _, err := os.Stat(svc.credPath + ".lock"); err != nil {
			t.Fatalf("no lock taken: %v", err)
		}
	})
	t.Run("not signed in", func(t *testing.T) {
		svc := newTestService(t, "http://127.0.0.1:1")
		if _, err := svc.TokenNoLogin(context.Background()); !errors.Is(err, ErrNotSignedIn) {
			t.Fatalf("err = %v, want ErrNotSignedIn", err)
		}
	})
	t.Run("dead refresh token is an error, not a login", func(t *testing.T) {
		srv := newRotatingTokenServer(t, "rt-other", 0)
		svc := newTestService(t, srv.URL)
		writeCreds(t, svc.credPath, expiredCreds())
		if _, err := svc.TokenNoLogin(context.Background()); !errors.Is(err, errRefreshTokenDead) {
			t.Fatalf("err = %v, want errRefreshTokenDead", err)
		}
	})
}

// holdLockWhile takes the credentials lock, runs change while holding it after
// a short delay, then releases it, as another process would.
func holdLockWhile(t *testing.T, svc *OAuthService, change func()) {
	t.Helper()
	unlock, err := lockFile(context.Background(), svc.credPath+".lock")
	if err != nil {
		t.Fatal(err)
	}
	go func() {
		time.Sleep(100 * time.Millisecond)
		change()
		unlock()
	}()
}

// Signed out by another process while waiting on the lock: Token falls
// through to the interactive sign-in instead of failing.
func TestToken_SignedOutWhileWaitingFallsThroughToLogin(t *testing.T) {
	srv := newRotatingTokenServer(t, "rt-1", 0)
	svc := newTestService(t, srv.URL)
	writeCreds(t, svc.credPath, expiredCreds())
	wiped, err := json.Marshal(credentials{ClientID: "client-1"})
	if err != nil {
		t.Fatal(err)
	}
	holdLockWhile(t, svc, func() { _ = os.WriteFile(svc.credPath, wiped, 0o600) })

	var opened atomic.Bool
	prev := openBrowser
	openBrowser = func(string) error { opened.Store(true); return nil }
	t.Cleanup(func() { openBrowser = prev })

	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	if _, err := svc.Token(ctx); err == nil {
		t.Fatal("Token succeeded without a sign-in")
	}
	if !opened.Load() {
		t.Fatal("Token did not fall through to the interactive sign-in")
	}
	if srv.refreshes.Load()+srv.rejected.Load() != 0 {
		t.Fatal("refreshed with a refresh token that was wiped")
	}
}

func TestLockedRefresh_LoadErrorUnderLock(t *testing.T) {
	failOnLogin(t)
	srv := newRotatingTokenServer(t, "rt-1", 0)
	svc := newTestService(t, srv.URL)
	writeCreds(t, svc.credPath, expiredCreds())
	holdLockWhile(t, svc, func() {
		_ = os.Remove(svc.credPath)
		_ = os.Mkdir(svc.credPath, 0o700) // unreadable as a file
	})

	if _, err := svc.TokenNoLogin(context.Background()); err == nil || errors.Is(err, ErrNotSignedIn) {
		t.Fatalf("err = %v, want the read error", err)
	}
	if srv.refreshes.Load() != 0 {
		t.Fatal("refreshed after a failed read")
	}
}

// A logout waits for a refresh in flight in another process, so the refresh
// cannot write its tokens back after the wipe.
func TestLogout_WaitsForRefreshInFlight(t *testing.T) {
	failOnLogin(t)
	srv := newRotatingTokenServer(t, "rt-1", 300*time.Millisecond)
	refresher := newTestService(t, srv.URL)
	writeCreds(t, refresher.credPath, expiredCreds())
	other := &OAuthService{issuer: refresher.issuer, clientName: defaultClientName,
		credPath: refresher.credPath, httpClient: &http.Client{}}

	done := make(chan struct{})
	go func() {
		defer close(done)
		_, _ = refresher.CachedToken(context.Background())
	}()
	time.Sleep(50 * time.Millisecond) // the refresh holds the lock now
	if err := other.Logout(context.Background()); err != nil {
		t.Fatal(err)
	}
	<-done

	raw, err := os.ReadFile(refresher.credPath)
	if err != nil {
		t.Fatal(err)
	}
	var creds credentials
	if err := json.Unmarshal(raw, &creds); err != nil {
		t.Fatal(err)
	}
	if creds.AccessToken != "" || creds.RefreshToken != "" || creds.ClientID != "client-1" {
		t.Fatalf("credentials after logout = %+v, want only the client id", creds)
	}
}
