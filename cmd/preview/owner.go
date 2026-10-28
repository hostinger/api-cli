// Package preview holds the `preview` command's pieces. It is shipped only to
// CLIs whose manifest sets [generators.cli] preview = true; the generator
// deletes this directory for every other CLI.
package preview

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/hostinger/api-cli/api"
	"github.com/spf13/viper"
)

// The owner ticket errors. They map 1:1 to hcdn-connector's pkg/connector
// errors of the same names; replace them with those when the connector
// dependency lands.
var (
	ErrOwnerNoToken   = errors.New("no Hostinger API token or sign-in")
	ErrOwnerOptedOut  = errors.New("owner link turned off with --no-owner")
	ErrOwnerTemporary = errors.New("owner ticket request failed temporarily")
)

// OwnerTicketPath is public-api's internal owner ticket operation. Released
// CLIs call it, so public-api keeps it backward compatible.
const OwnerTicketPath = "/api/hosting/v1/previews/owner-tickets"

const (
	// defaultOwnerTimeout bounds a call whose context has no deadline. The
	// connector always sets one (2 s in log mode, 3 s in require mode).
	defaultOwnerTimeout = 5 * time.Second
	// maxOwnerResponse bounds the response body read.
	maxOwnerResponse = 64 << 10
)

// OwnerTicketFunc is the shape of hcdn-connector's QuickOptions.OwnerTicket.
type OwnerTicketFunc func(ctx context.Context, tunnelID string, notAfter time.Time) (string, error)

// TokenSource returns the bearer token for the owner ticket call. rejected is
// "" on the first call; after a 401 it is the token public-api refused, and
// the source returns a different token to retry with, or that same token when
// it has no other (a static API token). It returns an error matching
// ErrOwnerNoToken when there are no usable credentials, and one matching
// ErrOwnerTemporary for a failure worth one retry.
type TokenSource func(ctx context.Context, rejected string) (string, error)

// OwnerOptions configures NewOwnerTicketFunc.
type OwnerOptions struct {
	APIURL     string       // public-api base URL, e.g. https://developers.hostinger.com
	NoOwner    bool         // --no-owner: never ask, return ErrOwnerOptedOut
	Token      TokenSource  // credentials; see CLIToken
	HTTPClient *http.Client // nil means a plain http.Client
	UserAgent  string
}

// DefaultOwnerOptions builds the options from the CLI configuration: the
// configured api_url and CLIToken.
func DefaultOwnerOptions(noOwner bool) OwnerOptions {
	return OwnerOptions{
		APIURL:    api.BaseURL(),
		NoOwner:   noOwner,
		Token:     CLIToken,
		UserAgent: api.UserAgent(),
	}
}

// NewOwnerTicketFunc returns the callback the connector calls when the edge
// asks for an owner. It POSTs {tunnel_id, expires_at} to OwnerTicketPath and
// returns the ticket, or:
//   - ErrOwnerOptedOut for --no-owner, without a request;
//   - ErrOwnerNoToken when there are no credentials or the refresh token is dead;
//   - an error wrapping ErrOwnerTemporary for a timeout, a network error or a 5xx;
//   - a plain error for any 4xx or an unusable response (counted as fetch_failed).
//
// A 401 is retried once when the token source has another token: a sign-in
// whose access token the server revoked while its stored expiry says it is
// valid. A 429 is not retried: the limit is per hour, so a retry within the
// connector's budget would only be refused again.
//
// It never opens a browser and never uses api.Request(), whose OAuth fallback
// signs in interactively.
func NewOwnerTicketFunc(o OwnerOptions) OwnerTicketFunc {
	return func(ctx context.Context, tunnelID string, notAfter time.Time) (string, error) {
		if o.NoOwner {
			return "", ErrOwnerOptedOut
		}
		if _, ok := ctx.Deadline(); !ok {
			var cancel context.CancelFunc
			ctx, cancel = context.WithTimeout(ctx, defaultOwnerTimeout)
			defer cancel()
		}
		if o.Token == nil {
			return "", ErrOwnerNoToken
		}
		token, err := o.Token(ctx, "")
		if err != nil {
			return "", err
		}
		if token == "" {
			return "", ErrOwnerNoToken
		}
		ticket, err := o.fetch(ctx, token, tunnelID, notAfter)
		var se *statusError
		if !errors.As(err, &se) || se.status != http.StatusUnauthorized {
			return ticket, err
		}
		fresh, terr := o.Token(ctx, token)
		if terr != nil || fresh == "" || fresh == token {
			return "", err
		}
		return o.fetch(ctx, fresh, tunnelID, notAfter)
	}
}

// formatExpiresAt formats the ticket expiry as public-api accepts it: RFC 3339
// in UTC with whole seconds ("2026-10-23T11:05:00Z"). Its validation
// (date_format Y-m-d\TH:i:sP / Y-m-d\TH:i:sp) answers 422 to fractional
// seconds, so this must never become RFC3339Nano.
func formatExpiresAt(t time.Time) string {
	return t.UTC().Truncate(time.Second).Format(time.RFC3339)
}

type ownerTicketRequest struct {
	TunnelID  string `json:"tunnel_id"`
	ExpiresAt string `json:"expires_at"`
}

// ownerTicketResponse is public-api's 201 body. public-api resources are not
// wrapped in a "data" envelope (JsonResource::withoutWrapping).
type ownerTicketResponse struct {
	Ticket    string `json:"ticket"`
	ExpiresAt string `json:"expires_at"`
	ClientID  string `json:"client_id"`
}

func (o OwnerOptions) fetch(ctx context.Context, token, tunnelID string, notAfter time.Time) (string, error) {
	body, err := json.Marshal(ownerTicketRequest{
		TunnelID:  tunnelID,
		ExpiresAt: formatExpiresAt(notAfter),
	})
	if err != nil {
		return "", err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost,
		strings.TrimRight(o.APIURL, "/")+OwnerTicketPath, bytes.NewReader(body))
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	if o.UserAgent != "" {
		req.Header.Set("User-Agent", o.UserAgent)
	}

	hc := o.HTTPClient
	if hc == nil {
		hc = &http.Client{}
	}
	resp, err := hc.Do(req)
	if err != nil {
		return "", fmt.Errorf("%w: %w", ErrOwnerTemporary, err)
	}
	defer resp.Body.Close()
	raw, err := io.ReadAll(io.LimitReader(resp.Body, maxOwnerResponse))
	if err != nil {
		return "", fmt.Errorf("%w: reading the response: %w", ErrOwnerTemporary, err)
	}

	switch {
	case resp.StatusCode >= 500:
		return "", fmt.Errorf("%w: %w", ErrOwnerTemporary, newStatusError(resp.StatusCode, raw))
	case resp.StatusCode < 200 || resp.StatusCode > 299:
		return "", newStatusError(resp.StatusCode, raw)
	}
	var out ownerTicketResponse
	if err := json.Unmarshal(raw, &out); err != nil {
		return "", fmt.Errorf("owner ticket: invalid response: %w", err)
	}
	if out.Ticket == "" {
		return "", errors.New("owner ticket: response has no ticket")
	}
	return out.Ticket, nil
}

// statusError is a failed response, with the API's message when the body
// carries one.
type statusError struct {
	status int
	msg    string
}

func (e *statusError) Error() string { return e.msg }

// newStatusError describes a failed response. The body is untrusted, so the
// message taken from it is bounded.
func newStatusError(status int, body []byte) *statusError {
	msg := fmt.Sprintf("owner ticket request failed (HTTP %d)", status)
	var apiErr struct {
		Message string `json:"message"`
	}
	if json.Unmarshal(body, &apiErr) == nil && apiErr.Message != "" {
		m := apiErr.Message
		if len(m) > 200 {
			m = m[:200] + "…"
		}
		msg += ": " + m
	}
	return &statusError{status: status, msg: msg}
}

// Seams for tests.
var (
	configToken  = func() string { return viper.GetString("api_token") }
	newOAuthFunc = func() (cachedTokener, error) { return api.NewOAuthService() }
)

type cachedTokener interface {
	CachedToken(ctx context.Context) (string, error)
	ReplaceToken(ctx context.Context, rejected string) (string, error)
}

// CLIToken is the CLI's token source for the owner ticket (OT10). The order
// matches api/request.go:
//  1. api_token as viper resolves it (HOSTINGER_API_TOKEN, then the legacy
//     HAPI_API_TOKEN, then the config file), sent as is;
//  2. else the stored OAuth sign-in through OAuthService.CachedToken, or
//     ReplaceToken after a 401, which may refresh it but never sign in;
//  3. else ErrOwnerNoToken.
func CLIToken(ctx context.Context, rejected string) (string, error) {
	if tok := StaticToken(); tok != "" {
		return tok, nil
	}
	oauth, err := newOAuthFunc()
	if err != nil {
		return "", fmt.Errorf("%w: %w", ErrOwnerNoToken, err)
	}
	var tok string
	if rejected == "" {
		tok, err = oauth.CachedToken(ctx)
	} else {
		tok, err = oauth.ReplaceToken(ctx, rejected)
	}
	switch {
	case err == nil && tok != "":
		return tok, nil
	case err == nil:
		return "", ErrOwnerNoToken
	case errors.Is(err, api.ErrNotSignedIn):
		return "", fmt.Errorf("%w: %w", ErrOwnerNoToken, err)
	default:
		return "", fmt.Errorf("%w: %w", ErrOwnerTemporary, err)
	}
}

// StaticToken returns the user-managed API token, or "".
func StaticToken() string {
	return strings.TrimSpace(configToken())
}
