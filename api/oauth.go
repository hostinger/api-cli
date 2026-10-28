package api

import (
	"bytes"
	"context"
	"crypto/rand"
	"crypto/sha256"
	"embed"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"io"
	"net"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"time"
)

// OAuth 2.0 + PKCE (S256) client for Hostinger SSO. It runs an interactive
// browser sign-in against a local loopback callback, persists the resulting
// tokens to disk, and refreshes them on demand so callers can obtain a bearer
// token without a statically configured API token.
const (
	defaultIssuer     = "https://auth.hostinger.com"
	defaultClientName = "hostinger-cli"

	registerPath  = "/api/external/v1/oauth-server/register"
	authorizePath = "/api/external/v1/oauth-server/authorize"
	tokenPath     = "/api/external/v1/oauth-server/token"
	revokePath    = "/api/external/v1/oauth-server/token/revoke"

	callbackPath = "/oauth/callback"

	// httpTimeout bounds a single auth-server round trip; loginTimeout bounds
	// the interactive wait for the user to finish in their browser. Both are
	// applied only when the caller's context carries no deadline of its own.
	httpTimeout  = 30 * time.Second
	loginTimeout = 5 * time.Minute

	// expiryBufferSeconds is subtracted from the server-reported lifetime so we
	// treat a token as expired slightly early and avoid a mid-request expiry.
	expiryBufferSeconds = 60
)

// errRefreshTokenDead marks a refresh grant rejected with a 4xx — the refresh
// token is definitively dead and the caller must fall back to a full login.
// A 5xx or network failure is returned as a plain error instead, so transient
// outages never spuriously launch the browser.
var errRefreshTokenDead = errors.New("refresh token rejected")

// openBrowser launches the user's default browser. It is a package var so tests
// can substitute it; in production it is fire-and-forget (see Login).
var openBrowser = func(rawURL string) error {
	switch runtime.GOOS {
	case "darwin":
		return exec.Command("open", rawURL).Start()
	case "windows":
		// The empty quoted string is cmd.exe's required first "title" arg so
		// that it interprets the URL as the target rather than the title.
		return exec.Command("cmd", "/c", "start", "", rawURL).Start()
	default:
		return exec.Command("xdg-open", rawURL).Start()
	}
}

//go:embed templates/*.html
var templateFS embed.FS

// Callback pages, parsed once at load and reused. The error page interpolates
// the untrusted auth-server error code via {{ . }}, which html/template escapes.
var (
	successTemplate = template.Must(template.ParseFS(templateFS, "templates/success.html"))
	errorTemplate   = template.Must(template.ParseFS(templateFS, "templates/error.html"))
)

// OAuthService obtains and maintains a bearer token via the OAuth flow.
// The zero value is not usable; construct one with NewOAuthService.
type OAuthService struct {
	issuer     string
	clientName string
	credPath   string
	httpClient *http.Client

	// mu serializes the public methods so concurrent callers in one process
	// share a single auth-server round trip rather than racing on disk.
	mu sync.Mutex
}

// credentials is the on-disk token state. Absent fields imply "not yet known".
type credentials struct {
	ClientID     string `json:"client_id,omitempty"`
	AccessToken  string `json:"access_token,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
	ExpiresAt    int64  `json:"expires_at,omitempty"` // local expiry, ms epoch
}

type tokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int64  `json:"expires_in"`
}

// NewOAuthService builds a service with the default issuer (overridable via the
// HOSTINGER_OAUTH_ISSUER env var) and the standard on-disk credentials path.
func NewOAuthService() (*OAuthService, error) {
	credPath, err := defaultCredPath()
	if err != nil {
		return nil, err
	}

	issuer := defaultIssuer
	if v := strings.TrimSpace(os.Getenv("HOSTINGER_OAUTH_ISSUER")); v != "" {
		issuer = v
	}

	return &OAuthService{
		issuer:     strings.TrimRight(issuer, "/"),
		clientName: defaultClientName,
		credPath:   credPath,
		httpClient: &http.Client{},
	}, nil
}

// Token returns a usable access token: the cached one if still valid, otherwise
// a proactive refresh, otherwise a full interactive login.
func (s *OAuthService) Token(ctx context.Context) (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.resolve(ctx)
}

// ErrNotSignedIn is returned by TokenNoLogin when only an interactive sign-in
// could produce a token.
var ErrNotSignedIn = errors.New("not signed in")

// TokenNoLogin is Token without the interactive fallback: the cached token, or
// a refresh, or ErrNotSignedIn. Shell completion uses it so a TAB press never
// opens a browser.
func (s *OAuthService) TokenNoLogin(ctx context.Context) (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	creds, _ := s.loadCredentials()
	if creds.AccessToken != "" && creds.ExpiresAt != 0 && nowMillis() < creds.ExpiresAt {
		return creds.AccessToken, nil
	}
	if creds.RefreshToken == "" || creds.ClientID == "" {
		return "", ErrNotSignedIn
	}
	token, err := s.lockedRefresh(ctx, creds.AccessToken)
	if errors.Is(err, errNoRefreshToken) {
		return "", ErrNotSignedIn
	}
	return token, err
}

// CachedToken is the cross-process safe, non-interactive way to get a token,
// for callers that may run in parallel (several `preview` runs started by an
// agent at once). It returns the stored access token while it is valid, and
// otherwise refreshes it under an exclusive lock on a sibling lock file, so
// that only one process uses a refresh token the server may rotate. Every
// refresh path (Token, TokenNoLogin, Refresh, Reauthenticate) and the sign-in
// and logout writes take the same lock. It never signs in.
//
//  1. The stored token is returned without locking while it is valid. Writes
//     are atomic renames, so a read never sees a half-written file.
//  2. Otherwise it locks <credentials>.lock, waiting no longer than ctx allows,
//     and re-reads the file: another process may have refreshed meanwhile.
//  3. Still expired: it refreshes, writes the file atomically, and unlocks.
//     The refresh is not cut short by ctx (see lockedRefresh), so a call can
//     return after its deadline, by at most httpTimeout.
//
// "Valid" means the stored expiry has not passed. That expiry is already
// expiryBufferSeconds (60 s) before the server's, so a returned token is good
// for at least another 60 s.
//
// No stored refresh token, or one the server rejects (4xx), returns an error
// matching ErrNotSignedIn. Any other failure (lock wait past the deadline,
// network, 5xx) returns a plain error. On every failure the stored
// credentials are left as they were.
func (s *OAuthService) CachedToken(ctx context.Context) (string, error) {
	return s.cachedToken(ctx, "")
}

// ReplaceToken is CachedToken for a token the server rejected (401) while its
// stored expiry says it is valid: it returns a different stored token, or
// refreshes past the rejected one. It never signs in.
func (s *OAuthService) ReplaceToken(ctx context.Context, rejected string) (string, error) {
	return s.cachedToken(ctx, rejected)
}

func (s *OAuthService) cachedToken(ctx context.Context, rejected string) (string, error) {
	ctx, cancel := withTimeout(ctx, httpTimeout)
	defer cancel()

	s.mu.Lock()
	defer s.mu.Unlock()

	creds, err := s.loadCredentials()
	if err != nil {
		return "", err
	}
	if accessTokenValid(creds) && creds.AccessToken != rejected {
		return creds.AccessToken, nil
	}
	if creds.RefreshToken == "" || creds.ClientID == "" {
		return "", ErrNotSignedIn
	}

	if rejected == "" {
		rejected = creds.AccessToken
	}
	token, err := s.lockedRefresh(ctx, rejected)
	switch {
	case errors.Is(err, errNoRefreshToken):
		return "", ErrNotSignedIn
	case errors.Is(err, errRefreshTokenDead):
		return "", fmt.Errorf("%w: %w", ErrNotSignedIn, err)
	}
	return token, err
}

// errNoRefreshToken is lockedRefresh's answer when the file, re-read under the
// lock, holds no refresh token (another process logged out meanwhile).
var errNoRefreshToken = errors.New("no refresh token stored")

// lockedRefresh refreshes under the cross-process lock on <credentials>.lock.
// Once the lock is held it re-reads the file and returns the stored access
// token instead when another process refreshed meanwhile: valid and different
// from stale, the token the caller already holds (for Reauthenticate, the one
// the server just rejected).
//
// Only the lock wait follows ctx. The refresh grant and the save run under
// httpTimeout even when ctx ends first: the server may rotate the refresh
// token, and dropping its answer would leave a dead one on disk. Callers must
// hold s.mu.
func (s *OAuthService) lockedRefresh(ctx context.Context, stale string) (string, error) {
	unlock, err := s.lockCredentials(ctx)
	if err != nil {
		return "", err
	}
	defer unlock()

	creds, err := s.loadCredentials()
	if err != nil {
		return "", err
	}
	if accessTokenValid(creds) && creds.AccessToken != stale {
		return creds.AccessToken, nil
	}
	if creds.RefreshToken == "" || creds.ClientID == "" {
		return "", errNoRefreshToken
	}
	rctx, cancel := context.WithTimeout(context.WithoutCancel(ctx), httpTimeout)
	defer cancel()
	return s.refresh(rctx, creds)
}

// lockCredentials takes the cross-process lock on <credentials>.lock that
// every credentials write except the first sign-in's client registration
// happens under. The wait follows ctx, bounded by httpTimeout when ctx has no
// deadline.
func (s *OAuthService) lockCredentials(ctx context.Context) (unlock func(), err error) {
	ctx, cancel := withTimeout(ctx, httpTimeout)
	defer cancel()
	return lockFile(ctx, s.credPath+".lock")
}

// accessTokenValid reports whether creds hold an access token whose stored
// (already buffered) expiry is still ahead.
func accessTokenValid(creds credentials) bool {
	return creds.AccessToken != "" && creds.ExpiresAt != 0 && nowMillis() < creds.ExpiresAt
}

// Status describes the stored sign-in without touching the network.
type Status struct {
	SignedIn   bool      // an access or refresh token is stored
	Expired    bool      // the access token has expired (a refresh may still work)
	CanRefresh bool      // a refresh token is stored
	ExpiresAt  time.Time // access token expiry; zero when unknown
	CredPath   string
}

// Status reports the stored sign-in state.
func (s *OAuthService) Status() Status {
	s.mu.Lock()
	defer s.mu.Unlock()

	creds, _ := s.loadCredentials()
	st := Status{
		SignedIn:   creds.AccessToken != "" || creds.RefreshToken != "",
		CanRefresh: creds.RefreshToken != "" && creds.ClientID != "",
		CredPath:   s.credPath,
	}
	if creds.ExpiresAt != 0 {
		st.ExpiresAt = time.UnixMilli(creds.ExpiresAt)
		st.Expired = nowMillis() >= creds.ExpiresAt
	}
	return st
}

// Login runs the full interactive PKCE flow regardless of cache state.
func (s *OAuthService) Login(ctx context.Context) (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.login(ctx)
}

// Refresh exchanges the stored refresh token for a fresh access token.
func (s *OAuthService) Refresh(ctx context.Context) (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	creds, err := s.loadCredentials()
	if err != nil {
		return "", err
	}
	if creds.RefreshToken == "" || creds.ClientID == "" {
		return "", errors.New("no refresh token available; run an interactive login first")
	}
	token, err := s.lockedRefresh(ctx, creds.AccessToken)
	if errors.Is(err, errNoRefreshToken) {
		return "", errors.New("no refresh token available; run an interactive login first")
	}
	return token, err
}

// Reauthenticate forces a fresh token, bypassing the cached-token fast path. It
// mirrors Token's refresh-then-login fallback and is meant for the reactive 401
// recovery path, where the cached token is already known to be dead.
func (s *OAuthService) Reauthenticate(ctx context.Context) (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	creds, _ := s.loadCredentials()
	return s.refreshOrLogin(ctx, creds)
}

// Logout best-effort revokes the access token and wipes local tokens, keeping
// the registered client_id so a later login can skip dynamic registration.
func (s *OAuthService) Logout(ctx context.Context) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Under the lock, so a refresh in flight in another process cannot write
	// its tokens back after the wipe, and the token revoked is the latest.
	unlock, err := s.lockCredentials(ctx)
	if err != nil {
		return err
	}
	defer unlock()

	creds, _ := s.loadCredentials()
	if creds.ClientID == "" && creds.AccessToken == "" {
		return nil // nothing stored
	}

	if creds.AccessToken != "" {
		form := url.Values{}
		form.Set("token", creds.AccessToken)
		form.Set("client_id", creds.ClientID)
		// Revocation is best-effort: logout must succeed locally even if the
		// auth server is unreachable.
		_, _, _ = s.postForm(ctx, revokePath, form)
	}

	return s.saveCredentials(credentials{ClientID: creds.ClientID})
}

// resolve is the token decision tree: cached -> refresh -> login. Callers must
// hold s.mu.
func (s *OAuthService) resolve(ctx context.Context) (string, error) {
	creds, _ := s.loadCredentials()

	if creds.AccessToken != "" && creds.ExpiresAt != 0 && nowMillis() < creds.ExpiresAt {
		return creds.AccessToken, nil
	}
	return s.refreshOrLogin(ctx, creds)
}

// refreshOrLogin attempts a proactive refresh and falls back to a full
// interactive login when there is no usable refresh token or the refresh token
// is dead (4xx). A transient refresh failure (5xx/network) is surfaced as-is so
// it never spuriously launches a browser. The refresh runs under the
// cross-process lock (lockedRefresh); the login does not, since it waits on
// the user. Callers must hold s.mu.
func (s *OAuthService) refreshOrLogin(ctx context.Context, creds credentials) (string, error) {
	if creds.RefreshToken != "" && creds.ClientID != "" {
		token, err := s.lockedRefresh(ctx, creds.AccessToken)
		if err == nil {
			return token, nil
		}
		if !errors.Is(err, errRefreshTokenDead) && !errors.Is(err, errNoRefreshToken) {
			return "", err
		}
		// 4xx, or signed out meanwhile: fall through to a full login.
	}
	return s.login(ctx)
}

// login runs the PKCE flow end to end and returns the issued access token.
// Callers must hold s.mu.
func (s *OAuthService) login(ctx context.Context) (string, error) {
	// Bind the loopback callback before launching the browser so the redirect
	// can never race ahead of the listener.
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		return "", fmt.Errorf("failed to start callback listener: %w", err)
	}
	defer listener.Close()

	port := listener.Addr().(*net.TCPAddr).Port
	redirectURI := fmt.Sprintf("http://127.0.0.1:%d%s", port, callbackPath)

	creds, _ := s.loadCredentials()
	if creds.ClientID == "" {
		clientID, err := s.register(ctx, redirectURI)
		if err != nil {
			return "", err
		}
		creds.ClientID = clientID
		if err := s.saveCredentials(creds); err != nil {
			return "", err
		}
	}

	verifier, challenge, err := pkce()
	if err != nil {
		return "", err
	}
	state, err := randomState()
	if err != nil {
		return "", err
	}

	authURL := s.authorizeURL(creds.ClientID, redirectURI, state, challenge)
	if err := openBrowser(authURL); err != nil {
		fmt.Fprintf(os.Stderr, "Could not open a browser automatically.\nOpen this URL to continue signing in:\n\n%s\n\n", authURL)
	}

	waitCtx, cancel := withTimeout(ctx, loginTimeout)
	defer cancel()
	code, err := s.awaitCallback(waitCtx, listener, state)
	if err != nil {
		return "", err
	}

	return s.exchangeCode(ctx, creds.ClientID, code, verifier, redirectURI)
}

// register performs RFC 7591 dynamic client registration and returns the new
// client_id. It is a public (PKCE-only) client, so no client_secret is expected.
func (s *OAuthService) register(ctx context.Context, redirectURI string) (string, error) {
	payload, err := json.Marshal(map[string]any{
		"client_name":   s.clientName,
		"redirect_uris": []string{redirectURI},
	})
	if err != nil {
		return "", err
	}

	status, body, err := s.postJSON(ctx, registerPath, payload)
	if err != nil {
		return "", err
	}
	if !is2xx(status) {
		return "", fmt.Errorf("client registration failed (status %d): %s", status, snippet(body))
	}

	var out struct {
		ClientID string `json:"client_id"`
	}
	if err := json.Unmarshal(body, &out); err != nil {
		return "", fmt.Errorf("client registration: invalid response: %w", err)
	}
	if out.ClientID == "" {
		return "", fmt.Errorf("client registration response missing client_id: %s", snippet(body))
	}
	return out.ClientID, nil
}

// exchangeCode swaps the authorization code for tokens and persists them.
func (s *OAuthService) exchangeCode(ctx context.Context, clientID, code, verifier, redirectURI string) (string, error) {
	form := url.Values{}
	form.Set("grant_type", "authorization_code")
	form.Set("code", code)
	form.Set("code_verifier", verifier)
	form.Set("redirect_uri", redirectURI)
	form.Set("client_id", clientID)

	status, body, err := s.postForm(ctx, tokenPath, form)
	if err != nil {
		return "", err
	}
	if !is2xx(status) {
		return "", fmt.Errorf("token exchange failed (status %d): %s", status, snippet(body))
	}

	tr, err := parseTokenResponse(body)
	if err != nil {
		return "", err
	}

	// Under the lock, so a refresh of the previous token chain in another
	// process cannot overwrite the new sign-in.
	unlock, err := s.lockCredentials(ctx)
	if err != nil {
		return "", err
	}
	defer unlock()

	creds, _ := s.loadCredentials()
	creds.ClientID = clientID
	creds.AccessToken = tr.AccessToken
	creds.RefreshToken = tr.RefreshToken
	creds.ExpiresAt = expiresAtMillis(tr.ExpiresIn)
	if err := s.saveCredentials(creds); err != nil {
		return "", err
	}
	return creds.AccessToken, nil
}

// refresh runs a refresh-token grant. A 4xx is reported as errRefreshTokenDead;
// a 5xx or network error is returned verbatim so the caller can treat it as
// transient.
func (s *OAuthService) refresh(ctx context.Context, creds credentials) (string, error) {
	form := url.Values{}
	form.Set("grant_type", "refresh_token")
	form.Set("refresh_token", creds.RefreshToken)
	form.Set("client_id", creds.ClientID)

	status, body, err := s.postForm(ctx, tokenPath, form)
	if err != nil {
		return "", err
	}
	switch {
	case is2xx(status):
		// handled below
	case status >= 400 && status < 500:
		return "", fmt.Errorf("%w (status %d)", errRefreshTokenDead, status)
	default:
		return "", fmt.Errorf("token refresh failed (status %d): %s", status, snippet(body))
	}

	tr, err := parseTokenResponse(body)
	if err != nil {
		return "", err
	}

	creds.AccessToken = tr.AccessToken
	if tr.RefreshToken != "" {
		creds.RefreshToken = tr.RefreshToken // server may rotate; keep old otherwise
	}
	creds.ExpiresAt = expiresAtMillis(tr.ExpiresIn)
	if err := s.saveCredentials(creds); err != nil {
		return "", err
	}
	return creds.AccessToken, nil
}

// awaitCallback serves the loopback listener until the OAuth redirect arrives,
// then returns the authorization code. Only GET /oauth/callback is honored;
// everything else is 404 and does not affect the flow.
func (s *OAuthService) awaitCallback(ctx context.Context, listener net.Listener, expectedState string) (string, error) {
	type result struct {
		code string
		err  error
	}
	// Buffered + non-blocking send: the first callback resolves the flow; any
	// duplicate hit (browser retry, replay) is answered but its result dropped,
	// so the handler goroutine never blocks on a full channel.
	done := make(chan result, 1)
	reply := func(res result) {
		select {
		case done <- res:
		default:
		}
	}

	mux := http.NewServeMux()
	mux.HandleFunc(callbackPath, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.NotFound(w, r)
			return
		}
		q := r.URL.Query()

		if e := q.Get("error"); e != "" {
			renderPage(w, http.StatusBadRequest, errorTemplate, e)
			reply(result{err: fmt.Errorf("authorization failed: %s", e)})
			return
		}
		// CSRF guard: the state must match byte-for-byte what we generated.
		if q.Get("state") != expectedState {
			renderPage(w, http.StatusBadRequest, errorTemplate, "state mismatch")
			reply(result{err: errors.New("oauth state mismatch")})
			return
		}
		code := q.Get("code")
		if code == "" {
			renderPage(w, http.StatusBadRequest, errorTemplate, "missing authorization code")
			reply(result{err: errors.New("oauth callback missing authorization code")})
			return
		}
		renderPage(w, http.StatusOK, successTemplate, nil)
		reply(result{code: code})
	})

	srv := &http.Server{Handler: mux}
	go srv.Serve(listener) //nolint:errcheck // returns ErrServerClosed on shutdown
	defer func() {
		shutdownCtx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		_ = srv.Shutdown(shutdownCtx) // graceful: lets the final response flush
	}()

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case res := <-done:
		return res.code, res.err
	}
}

func (s *OAuthService) authorizeURL(clientID, redirectURI, state, challenge string) string {
	q := url.Values{}
	q.Set("client_id", clientID)
	q.Set("redirect_uri", redirectURI)
	q.Set("state", state)
	q.Set("code_challenge", challenge)
	q.Set("code_challenge_method", "S256")
	q.Set("response_type", "code")
	return s.issuer + authorizePath + "?" + q.Encode()
}

// --- HTTP plumbing ---------------------------------------------------------

func (s *OAuthService) postForm(ctx context.Context, path string, form url.Values) (int, []byte, error) {
	ctx, cancel := withTimeout(ctx, httpTimeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, s.issuer+path, strings.NewReader(form.Encode()))
	if err != nil {
		return 0, nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	return s.do(req)
}

func (s *OAuthService) postJSON(ctx context.Context, path string, payload []byte) (int, []byte, error) {
	ctx, cancel := withTimeout(ctx, httpTimeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, s.issuer+path, bytes.NewReader(payload))
	if err != nil {
		return 0, nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	return s.do(req)
}

func (s *OAuthService) do(req *http.Request) (int, []byte, error) {
	resp, err := s.httpClient.Do(req)
	if err != nil {
		return 0, nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return resp.StatusCode, nil, err
	}
	return resp.StatusCode, body, nil
}

// --- credentials persistence -----------------------------------------------

// loadCredentials reads the on-disk state. A missing or corrupt file is treated
// as "no credentials" (zero value, nil error).
func (s *OAuthService) loadCredentials() (credentials, error) {
	data, err := os.ReadFile(s.credPath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return credentials{}, nil
		}
		return credentials{}, err
	}

	var creds credentials
	if err := json.Unmarshal(data, &creds); err != nil {
		return credentials{}, nil // corrupt: start fresh
	}
	return creds, nil
}

// saveCredentials writes the state atomically (temp file + rename) with mode
// 0600, creating the parent directory on demand.
func (s *OAuthService) saveCredentials(creds credentials) error {
	data, err := json.MarshalIndent(creds, "", "  ")
	if err != nil {
		return err
	}

	dir := filepath.Dir(s.credPath)
	if err := os.MkdirAll(dir, 0700); err != nil {
		return err
	}

	tmp, err := os.CreateTemp(dir, "credentials-*.tmp")
	if err != nil {
		return err
	}
	tmpName := tmp.Name()
	defer os.Remove(tmpName) // no-op once the rename succeeds

	if err := tmp.Chmod(0600); err != nil {
		tmp.Close()
		return err
	}
	if _, err := tmp.Write(data); err != nil {
		tmp.Close()
		return err
	}
	if err := tmp.Close(); err != nil {
		return err
	}
	return replaceFile(tmpName, s.credPath)
}

// replaceFile renames from over to. On Windows the rename fails while another
// process has the target open, such as a parallel run's lock-free read in
// CachedToken, so it retries briefly there: giving up would drop a refresh
// token the server has just rotated.
func replaceFile(from, to string) error {
	err := os.Rename(from, to)
	for i := 0; err != nil && runtime.GOOS == "windows" && i < 10; i++ {
		time.Sleep(10 * time.Millisecond)
		err = os.Rename(from, to)
	}
	return err
}

// --- helpers ---------------------------------------------------------------

// defaultCredPath returns ~/.config/hostinger/api-cli/credentials.json on POSIX
// and %APPDATA%\hostinger\api-cli\credentials.json on Windows.
func defaultCredPath() (string, error) {
	if runtime.GOOS == "windows" {
		base := os.Getenv("APPDATA")
		if base == "" {
			home, err := os.UserHomeDir()
			if err != nil {
				return "", err
			}
			base = filepath.Join(home, "AppData", "Roaming")
		}
		return filepath.Join(base, "hostinger", "api-cli", "credentials.json"), nil
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".config", "hostinger", "api-cli", "credentials.json"), nil
}

// withTimeout applies d only when ctx carries no deadline of its own.
func withTimeout(ctx context.Context, d time.Duration) (context.Context, context.CancelFunc) {
	if _, ok := ctx.Deadline(); ok {
		return ctx, func() {}
	}
	return context.WithTimeout(ctx, d)
}

// pkce generates an S256 verifier/challenge pair from 32 bytes of CSPRNG output.
func pkce() (verifier, challenge string, err error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return "", "", err
	}
	verifier = base64.RawURLEncoding.EncodeToString(b)
	sum := sha256.Sum256([]byte(verifier))
	challenge = base64.RawURLEncoding.EncodeToString(sum[:])
	return verifier, challenge, nil
}

// randomState returns 16 bytes of CSPRNG output as lowercase hex (the CSRF token).
func randomState() (string, error) {
	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

func parseTokenResponse(body []byte) (tokenResponse, error) {
	var tr tokenResponse
	if err := json.Unmarshal(body, &tr); err != nil {
		return tokenResponse{}, fmt.Errorf("invalid token response: %w", err)
	}
	if tr.AccessToken == "" {
		return tokenResponse{}, errors.New("token response missing access_token")
	}
	return tr, nil
}

func expiresAtMillis(expiresIn int64) int64 {
	// Clamp so a short-lived (or zero) server lifetime can't underflow into the
	// past; the token is simply treated as already expired on the next call.
	effective := expiresIn - expiryBufferSeconds
	if effective < 0 {
		effective = 0
	}
	return time.Now().Add(time.Duration(effective) * time.Second).UnixMilli()
}

func nowMillis() int64 {
	return time.Now().UnixMilli()
}

func is2xx(status int) bool {
	return status >= 200 && status < 300
}

// renderPage executes a callback page template into the response. It renders to
// a buffer first so a template failure can't leave a half-written body, and
// flushes so the page reaches the browser before the listener is shut down.
func renderPage(w http.ResponseWriter, status int, tmpl *template.Template, data any) {
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(status)
	_, _ = buf.WriteTo(w)
	if f, ok := w.(http.Flusher); ok {
		f.Flush()
	}
}

// snippet trims and bounds an untrusted response body for inclusion in errors.
func snippet(body []byte) string {
	s := strings.TrimSpace(string(body))
	if len(s) > 512 {
		s = s[:512] + "…"
	}
	return s
}
