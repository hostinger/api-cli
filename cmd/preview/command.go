package preview

import (
	"bytes"
	"context"
	"crypto/ed25519"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"sync"
	"time"

	"github.com/hostinger/hcdn-connector/pkg/connector"
)

// A command is a compact JWS that wh-api signs with the owner ticket's key
// when someone clicks Deploy in hPanel. The edge only relays it; the CLI
// verifies it with the keys of its owner ticket response.
const (
	commandType     = "hcdn-command+jwt"
	commandAudience = "hcdn-command"
	commandDeploy   = "deploy"
	// commandLeeway allows for clock skew past exp, as the edge does for
	// owner tickets.
	commandLeeway = 60 * time.Second
	// keyRefreshInterval spaces the owner ticket calls an unknown kid makes.
	keyRefreshInterval = 10 * time.Minute
	// keyRefreshTimeout bounds that call: the edge resets a command's stream
	// after 10 s, so the answer must come well before.
	keyRefreshTimeout = 4 * time.Second
)

// commandClaims are a command's payload. Username and Domain name the
// website hPanel showed, or the one picked for a first deploy; Overwrite is
// set once the customer confirmed replacing its content.
type commandClaims struct {
	Aud       string `json:"aud"`
	Sub       string `json:"sub"` // the owner's client id, as in the owner ticket
	Tid       string `json:"tid"`
	Cmd       string `json:"cmd"`
	Username  string `json:"username"`
	Domain    string `json:"domain"`
	Overwrite bool   `json:"overwrite"`
	Jti       string `json:"jti"`
	Exp       int64  `json:"exp"`
}

// command is a verified command.
type command struct {
	raw    string // the token, sent back with every status event
	claims commandClaims
}

// commandError is the connector's answer to a command the CLI does not run.
func commandError(status int, code string) error {
	return &connector.CommandError{Status: status, Code: code}
}

// CommandKey is a public key that verifies commands, from the owner ticket
// response.
type CommandKey struct {
	Kid       string `json:"kid"`
	PublicKey string `json:"public_key"` // base64url, no padding: the raw 32-byte Ed25519 key
}

// verifier checks commands against the latest owner ticket response: its
// command keys, and the owner's client id. A command signed with a key it
// does not know repeats the owner ticket call, at most once per
// keyRefreshInterval: the key may have rotated during a run that lasts up to
// 7 days.
type verifier struct {
	now func() time.Time

	refreshMu sync.Mutex // held through a refresh
	mu        sync.Mutex
	clientID  string
	keys      map[string]ed25519.PublicKey
	ticket    OwnerTicketFunc // nil until the connector asks for an owner
	tunnelID  string
	notAfter  time.Time
	refreshed time.Time
}

// setKeys stores an owner ticket response, as OwnerOptions.OnTicket. Keys
// that do not decode are left out.
func (v *verifier) setKeys(clientID string, keys []CommandKey) {
	m := make(map[string]ed25519.PublicKey, len(keys))
	for _, key := range keys {
		raw, err := base64.RawURLEncoding.DecodeString(key.PublicKey)
		if err == nil && len(raw) == ed25519.PublicKeySize && key.Kid != "" {
			m[key.Kid] = ed25519.PublicKey(raw)
		}
	}
	v.mu.Lock()
	defer v.mu.Unlock()
	v.clientID, v.keys = clientID, m
}

// owner wraps the owner ticket callback: it learns the tunnel, and a key
// refresh repeats the call.
func (v *verifier) owner(next OwnerTicketFunc) OwnerTicketFunc {
	return func(ctx context.Context, tunnelID string, notAfter time.Time) (string, error) {
		v.mu.Lock()
		v.ticket, v.tunnelID, v.notAfter = next, tunnelID, notAfter
		v.mu.Unlock()
		return next(ctx, tunnelID, notAfter)
	}
}

// verify checks a command's signature and claims, and returns a
// *connector.CommandError for one the CLI must not run.
func (v *verifier) verify(ctx context.Context, body []byte) (*command, error) {
	tok := bytes.TrimSpace(body)
	parts := bytes.Split(tok, []byte("."))
	if len(parts) != 3 {
		return nil, commandError(http.StatusBadRequest, "bad_command")
	}
	var header struct {
		Alg string `json:"alg"`
		Kid string `json:"kid"`
		Typ string `json:"typ"`
	}
	c := &command{raw: string(tok)}
	if !decodeJSON(parts[0], &header) || !decodeJSON(parts[1], &c.claims) || header.Alg != "EdDSA" {
		return nil, commandError(http.StatusBadRequest, "bad_command")
	}
	sig, err := base64.RawURLEncoding.DecodeString(string(parts[2]))
	if err != nil {
		return nil, commandError(http.StatusBadRequest, "bad_command")
	}
	key, clientID, tunnelID := v.lookup(ctx, header.Kid)
	cl := c.claims
	switch {
	case key == nil:
		return nil, commandError(http.StatusForbidden, "unknown_kid")
	case !ed25519.Verify(key, tok[:len(parts[0])+1+len(parts[1])], sig):
		return nil, commandError(http.StatusForbidden, "bad_signature")
	case header.Typ != commandType || cl.Aud != commandAudience:
		return nil, commandError(http.StatusForbidden, "wrong_audience")
	case cl.Cmd != commandDeploy:
		return nil, commandError(http.StatusBadRequest, "unsupported_command")
	case cl.Tid == "" || cl.Tid != tunnelID:
		return nil, commandError(http.StatusForbidden, "wrong_tunnel")
	case cl.Sub == "" || cl.Sub != clientID:
		return nil, commandError(http.StatusForbidden, "wrong_owner")
	case v.now().After(time.Unix(cl.Exp, 0).Add(commandLeeway)):
		return nil, commandError(http.StatusForbidden, "expired")
	case cl.Jti == "" || cl.Username == "" || cl.Domain == "":
		return nil, commandError(http.StatusBadRequest, "bad_command")
	}
	return c, nil
}

func decodeJSON(part []byte, v any) bool {
	raw, err := base64.RawURLEncoding.DecodeString(string(part))
	return err == nil && json.Unmarshal(raw, v) == nil
}

// lookup returns the key for kid, the owner's client id and the tunnel id,
// refreshing the keys once when kid is unknown. A nil key means kid stays
// unknown.
func (v *verifier) lookup(ctx context.Context, kid string) (ed25519.PublicKey, string, string) {
	if key, id, tid := v.get(kid); key != nil {
		return key, id, tid
	}
	// One refresh at a time: copies of the command that arrive over other
	// sessions meanwhile wait for it, then find the key.
	v.refreshMu.Lock()
	defer v.refreshMu.Unlock()
	if key, id, tid := v.get(kid); key != nil {
		return key, id, tid
	}
	v.mu.Lock()
	ticket, tunnelID, notAfter := v.ticket, v.tunnelID, v.notAfter
	due := ticket != nil && (v.refreshed.IsZero() || v.now().Sub(v.refreshed) >= keyRefreshInterval)
	if due {
		v.refreshed = v.now()
	}
	v.mu.Unlock()
	if !due {
		return nil, "", ""
	}
	ctx, cancel := context.WithTimeout(ctx, keyRefreshTimeout)
	defer cancel()
	if _, err := ticket(ctx, tunnelID, notAfter); err != nil {
		return nil, "", ""
	}
	return v.get(kid)
}

func (v *verifier) get(kid string) (ed25519.PublicKey, string, string) {
	v.mu.Lock()
	defer v.mu.Unlock()
	return v.keys[kid], v.clientID, v.tunnelID
}
