package connector

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/hostinger/hcdn-connector/internal/proto"
	"github.com/hostinger/hcdn-connector/internal/session"
)

// OwnerTicketFunc is called only when the first whoami reports an owner mode
// (log or require), after the quick credentials exist and before any
// session. The connector sets the deadline: log mode calls it once (2 s);
// require mode calls it up to twice (3 s each), the second time only after
// ErrOwnerTemporary. A ticket is sent on a second whoami and on every hello.
// ErrOwnerNoToken, ErrOwnerOptedOut, ErrOwnerTemporary or any other error
// (fetch_failed) runs the Preview without an owner in log mode and returns
// ErrOwnerRequired in require mode. A nil func means the caller has no way to
// get a ticket (unsupported). A call that ignores its ctx is abandoned at the
// deadline, not waited for: it can still run during the retry or after
// LinkOwner returns, so the func must be safe for concurrent use.
type OwnerTicketFunc func(ctx context.Context, tunnelID string, notAfter time.Time) (string, error)

var (
	// ErrOwnerNoToken: the caller has no credentials to ask for a ticket (no_token).
	ErrOwnerNoToken = errors.New("connector: no token for an owner ticket")
	// ErrOwnerOptedOut: the user asked for an anonymous tunnel (opted_out).
	ErrOwnerOptedOut = errors.New("connector: owner ticket opted out")
	// ErrOwnerTemporary: a timeout, a network error or a 5xx. Require mode
	// retries once after it; it counts as fetch_failed.
	ErrOwnerTemporary = errors.New("connector: owner ticket temporarily unavailable")
	// ErrOwnerRequired: the edge runs this tunnel only with a valid owner
	// ticket. LinkOwner returns it in require mode, and Run returns an
	// *ExitError (Code 77) that matches it on goaway owner_required.
	ErrOwnerRequired = session.ErrOwnerRequired
)

// Owner modes of WhoamiReply.OwnerMode and OwnerLink.Mode.
const (
	OwnerModeLog     = proto.OwnerModeLog
	OwnerModeRequire = proto.OwnerModeRequire
)

// No-ticket reasons of OwnerLink.Absent, sent as hcdn-owner-absent.
const (
	OwnerAbsentNoToken      = proto.OwnerAbsentNoToken
	OwnerAbsentOptedOut     = proto.OwnerAbsentOptedOut
	OwnerAbsentFetchFailed  = proto.OwnerAbsentFetchFailed
	OwnerAbsentNotRequested = proto.OwnerAbsentNotRequested
	OwnerAbsentUnsupported  = proto.OwnerAbsentUnsupported
)

// WhoamiOwner is the owner object of a whoami reply.
type WhoamiOwner = session.WhoamiOwner

// OwnerLink is the owner outcome of one quick tunnel run. LinkOwner makes it
// once; pass it to Options.Owner so every session's hello carries it.
// Reconnects re-send it and never ask the issuer again.
type OwnerLink struct {
	Mode   string // OwnerModeLog, OwnerModeRequire, or "" when the edge did not ask
	Ticket string // sent on the second whoami and on every hello; "" without one
	Absent string // the no-ticket reason when Ticket is "" (OwnerAbsent*)
	// The edge's verdict on Ticket from the second whoami: Status ok, missing or
	// invalid ("" when it was not checked), Reason its label (for example
	// unknown_kid) and ClientID the linked account when Status is ok.
	Status, Reason, ClientID string
	// OtherEnvironment: the edge does not know the ticket's key or claims
	// (unknown_kid, bad_claims) and the ticket's iss is not the issuer the edge
	// trusts, for example a production api_url against a staging edge.
	OtherEnvironment bool
}

// Linked reports an owner the edge accepted.
func (l OwnerLink) Linked() bool { return l.Status == proto.OwnerStatusOK }

func (l OwnerLink) validate() error {
	if l.Ticket != "" && !validTicket(l.Ticket) {
		return errors.New("connector: Options.Owner.Ticket is not a compact JWS")
	}
	if l.Absent != "" && !proto.OwnerAbsentReasons[l.Absent] {
		return fmt.Errorf("connector: Options.Owner.Absent %q is not a known reason", l.Absent)
	}
	return nil
}

// why says why a require-mode run has no owner.
func (l OwnerLink) why() string {
	switch {
	case l.OtherEnvironment:
		return "the owner ticket comes from another environment"
	case l.Ticket == "":
		return "no owner ticket (" + l.Absent + ")"
	case l.Status == "":
		return "the edge gave no verdict on the owner ticket"
	}
	return "owner ticket " + l.Status + " (" + l.Reason + ")"
}

// Budgets per OwnerTicketFunc call (the sign-in URL must show within 5 s); vars so tests can shorten them.
var (
	ownerBudgetLog     = 2 * time.Second
	ownerBudgetRequire = 3 * time.Second
)

// LinkOwner runs the owner step after the first whoami (first, made with o):
//   - the edge did not ask (no owner object, mode off, an older tunneld): no
//     call, Absent not_requested;
//   - fn is nil: no call, Absent unsupported;
//   - else fn runs with the tunnel id and NotAfter of o.Credentials, under
//     the deadlines of OwnerTicketFunc. With a ticket, a second whoami
//     carries it in hcdn-owner-ticket and returns the edge's verdict.
//
// In require mode a run without a ticket, or a verdict other than ok,
// returns an error that matches ErrOwnerRequired, with the link for the
// message; a failed second whoami returns its error. In log mode only a ctx
// error is returned: the run goes on without an owner, and a failed second
// whoami keeps the ticket for the hellos.
func LinkOwner(ctx context.Context, o WhoamiOptions, first *WhoamiReply, fn OwnerTicketFunc) (OwnerLink, error) {
	link := OwnerLink{Mode: first.OwnerMode()}
	if link.Mode == "" {
		link.Absent = OwnerAbsentNotRequested
		return link, nil
	}
	if o.Credentials == nil {
		return link, errors.New("connector: WhoamiOptions.Credentials is required")
	}
	log := o.logger()
	if fn == nil {
		link.Absent = OwnerAbsentUnsupported
	} else {
		cur := o.Credentials.Current()
		link.Ticket, link.Absent = fetchOwnerTicket(ctx, fn, link.Mode, cur.TunnelID, cur.Leaf.NotAfter, log.Warn)
		if err := ctx.Err(); err != nil {
			return link, err
		}
	}
	if link.Ticket != "" {
		w, err := whoami(ctx, o, link.Ticket)
		switch {
		case err != nil && (ctx.Err() != nil || link.Mode == OwnerModeRequire):
			return link, fmt.Errorf("whoami with the owner ticket: %w", err)
		case err != nil:
			log.Warn("whoami with the owner ticket failed; the sessions still send it", "err", err)
		case w.Owner != nil:
			link.Status, link.Reason, link.ClientID = w.Owner.Status, w.Owner.Reason, w.Owner.ClientID
			issuer := w.Owner.Issuer
			if issuer == "" {
				issuer = first.Owner.Issuer
			}
			if !link.Linked() && (link.Reason == "unknown_kid" || link.Reason == "bad_claims") {
				if iss := ticketIssuer(link.Ticket); iss != "" && iss != issuer {
					link.OtherEnvironment = true
				}
			}
		}
	}
	switch {
	case link.Linked():
		log.Info("owner linked", "client_id", link.ClientID)
	case link.Mode == OwnerModeRequire:
		return link, fmt.Errorf("%w: %s", ErrOwnerRequired, link.why())
	case link.Ticket != "":
		log.Warn("owner not linked", "status", link.Status, "reason", link.Reason, "other_environment", link.OtherEnvironment)
	default:
		log.Info("owner not linked", "reason", link.Absent)
	}
	return link, nil
}

// fetchOwnerTicket calls fn under the mode's budget: log once (2 s), require
// up to twice (3 s each), the second time only after a temporary failure.
// It returns a ticket, or "" and the no-ticket reason.
func fetchOwnerTicket(ctx context.Context, fn OwnerTicketFunc, mode, tunnelID string, notAfter time.Time,
	warn func(string, ...any)) (ticket, absent string) {
	budget, tries := ownerBudgetLog, 1
	if mode == OwnerModeRequire {
		budget, tries = ownerBudgetRequire, 2
	}
	for i := 0; i < tries; i++ {
		t, err := callOwnerTicket(ctx, fn, budget, tunnelID, notAfter)
		if err == nil && !validTicket(t) {
			err = errors.New("the owner ticket is not a compact JWS")
		}
		if err == nil {
			return t, ""
		}
		absent = ownerAbsentReason(err)
		if ctx.Err() != nil {
			return "", absent
		}
		warn("owner ticket unavailable", "reason", absent, "attempt", i+1, "err", err)
		if !errors.Is(err, ErrOwnerTemporary) {
			break
		}
	}
	return "", absent
}

// callOwnerTicket bounds fn by budget even when fn ignores its ctx; the
// deadline counts as ErrOwnerTemporary (a timeout).
func callOwnerTicket(ctx context.Context, fn OwnerTicketFunc, budget time.Duration, tunnelID string, notAfter time.Time) (string, error) {
	cctx, cancel := context.WithTimeout(ctx, budget)
	defer cancel()
	type result struct {
		ticket string
		err    error
	}
	ch := make(chan result, 1)
	go func() {
		t, err := fn(cctx, tunnelID, notAfter)
		ch <- result{t, err}
	}()
	var r result
	select {
	case r = <-ch:
	case <-cctx.Done():
		r.err = cctx.Err()
	}
	if r.err != nil && ctx.Err() == nil && errors.Is(cctx.Err(), context.DeadlineExceeded) {
		r.err = fmt.Errorf("%w: no answer within %s", ErrOwnerTemporary, budget)
	}
	return r.ticket, r.err
}

func ownerAbsentReason(err error) string {
	switch {
	case errors.Is(err, ErrOwnerNoToken):
		return OwnerAbsentNoToken
	case errors.Is(err, ErrOwnerOptedOut):
		return OwnerAbsentOptedOut
	}
	return OwnerAbsentFetchFailed
}

// maxOwnerTicketLen matches the edge's malformed limit (protocol.md section 9).
const maxOwnerTicketLen = 2048

// validTicket accepts three non-empty dot-separated base64url parts within the
// edge's size limit: the value goes into a header as is. The edge does the
// real checks.
func validTicket(t string) bool {
	if len(t) > maxOwnerTicketLen || strings.Count(t, ".") != 2 {
		return false
	}
	if t[0] == '.' || t[len(t)-1] == '.' || strings.Contains(t, "..") {
		return false
	}
	for i := 0; i < len(t); i++ {
		c := t[i]
		if !(c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' || c >= '0' && c <= '9' || c == '-' || c == '_' || c == '.') {
			return false
		}
	}
	return true
}

// ticketIssuer reads the unverified iss claim, or "".
func ticketIssuer(ticket string) string {
	parts := strings.Split(ticket, ".")
	if len(parts) != 3 {
		return ""
	}
	b, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return ""
	}
	var c struct {
		Iss any `json:"iss"`
	}
	if json.Unmarshal(b, &c) != nil {
		return ""
	}
	s, _ := c.Iss.(string)
	return s
}
