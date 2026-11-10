package preview

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math"
	"sync"
	"time"

	"github.com/hostinger/api-cli/utils"
	"github.com/hostinger/hcdn-connector/pkg/connector"
)

// Exit codes of a Preview that the service ended. They follow sysexits.h,
// as the connector's own: EX_NOPERM and EX_CONFIG.
const (
	ExitSignInRequired = 77
	ExitTakenDown      = 78
)

// Reasons of the "ended" JSON event.
const (
	endStopped        = "stopped"
	endExpired        = "expired"
	endTakenDown      = "taken_down"
	endSignInRequired = "sign_in_required"
	endUnreachable    = "unreachable"
	endError          = "error"
)

// ending is how a run finished: the JSON reason, the exit code and the
// message for a person.
type ending struct {
	reason  string
	code    int
	message string
}

// classify maps the result of RunQuick to an ending. binary and tokenEnv
// name the CLI and its token variable in the sign-in hint.
func classify(expired bool, err error, binary, tokenEnv string) ending {
	switch {
	case err == nil && expired:
		return ending{endExpired, utils.ExitOK, "Preview ended"}
	case err == nil, errors.Is(err, context.Canceled):
		return ending{endStopped, utils.ExitOK, "Preview ended"}
	}
	var qe *connector.QuickError
	kind := err
	if errors.As(err, &qe) {
		kind = qe.Kind
	}
	switch {
	case errors.Is(kind, connector.ErrTakenDown):
		return ending{endTakenDown, ExitTakenDown, "this Preview was taken down"}
	case errors.Is(kind, connector.ErrOwnerRequired):
		msg := fmt.Sprintf(`Sign in to start a Preview: run "%s auth login" or set %s`, binary, tokenEnv)
		if qe != nil && qe.Owner.OtherEnvironment {
			msg += "; the owner ticket comes from another environment"
		}
		return ending{endSignInRequired, ExitSignInRequired, msg}
	case errors.Is(kind, connector.ErrEdgeUnreachable):
		return ending{endUnreachable, utils.ExitError, "outbound TCP 7844 is blocked on this network"}
	case errors.Is(kind, connector.ErrQuickRefused):
		return ending{endError, utils.ExitError, "the Preview service refused the tunnel: " + cause(qe, err)}
	case errors.Is(kind, connector.ErrQuickNotOffered):
		return ending{endError, utils.ExitError, "the Preview service does not offer Previews to this client: " + cause(qe, err)}
	case errors.Is(kind, connector.ErrCertRejected):
		return ending{endError, utils.ExitError, "the Preview service no longer accepts this tunnel; run the command again for a new URL"}
	}
	return ending{endError, utils.ExitError, "Preview failed: " + err.Error()}
}

// cause is the underlying error of a QuickError, or err itself.
func cause(qe *connector.QuickError, err error) string {
	if qe != nil && qe.Err != nil {
		return qe.Err.Error()
	}
	return err.Error()
}

// owner is the account link of a Preview, as printed.
type owner struct {
	Linked   bool   `json:"linked"`
	ClientID string `json:"client_id,omitempty"`
	Reason   string `json:"reason,omitempty"`
	text     string // the reason for a person
}

// ownerOf describes the link, or returns nil when the service did not ask
// for an owner.
func ownerOf(l connector.OwnerLink) *owner {
	switch {
	case l.Mode == "":
		return nil
	case l.Linked():
		return &owner{Linked: true, ClientID: l.ClientID}
	case l.OtherEnvironment:
		return &owner{Reason: "other_environment", text: "ticket from another environment"}
	case l.Ticket == "":
		switch l.Absent {
		case connector.OwnerAbsentNoToken:
			return &owner{Reason: l.Absent, text: "no token"}
		case connector.OwnerAbsentOptedOut:
			return &owner{Reason: l.Absent, text: "--no-owner"}
		case connector.OwnerAbsentFetchFailed:
			return &owner{Reason: l.Absent, text: "ticket request failed"}
		}
		return &owner{Reason: l.Absent, text: l.Absent}
	case l.Reason != "":
		return &owner{Reason: l.Reason, text: l.Reason}
	case l.Status != "":
		return &owner{Reason: l.Status, text: l.Status}
	}
	return &owner{Reason: "not_verified", text: "ticket not verified"}
}

// printer writes the Preview's progress: text for a person, or one JSON
// object per line on stdout. OnPause calls it from another goroutine.
type printer struct {
	mu        sync.Mutex
	reachDone bool // the "reachable" line is out
	stdout    io.Writer
	stderr    io.Writer
	json      bool
	bold      bool // emphasise the URL (a terminal with styling on)
	target    string
	now       func() time.Time
	loc       *time.Location
}

type readyEvent struct {
	Event     string `json:"event"`
	URL       string `json:"url"`
	ExpiresAt string `json:"expires_at"`
	Reachable bool   `json:"reachable"` // false: still connecting; a "reachable" event follows
	Owner     *owner `json:"owner,omitempty"`
}

type reachableEvent struct {
	Event string `json:"event"`
}

type pausedEvent struct {
	Event      string `json:"event"`
	Reason     string `json:"reason"`
	RetryAfter int64  `json:"retry_after"`
}

type endedEvent struct {
	Event   string `json:"event"`
	Reason  string `json:"reason"`
	Message string `json:"message,omitempty"` // why it failed; absent for a normal end
}

func (p *printer) emit(v any) {
	enc := json.NewEncoder(p.stdout)
	enc.SetEscapeHTML(false)
	_ = enc.Encode(v)
}

func (p *printer) ready(t connector.QuickTunnel) {
	p.mu.Lock()
	defer p.mu.Unlock()
	o := ownerOf(t.Owner)
	if p.json {
		p.emit(readyEvent{Event: "ready", URL: t.URL,
			ExpiresAt: t.ExpiresAt.UTC().Format(time.RFC3339), Reachable: t.Reachable, Owner: o})
		return
	}
	url := utils.SafeText(t.URL)
	if p.bold {
		url = "\x1b[1m" + url + "\x1b[0m"
	}
	w := p.stdout
	if t.Reachable {
		fmt.Fprintf(w, "\nPreview ready: %s\n\n", url)
	} else {
		fmt.Fprintf(w, "\nPreview URL: %s (still connecting; it may answer 530 for a few seconds)\n\n", url)
	}
	fmt.Fprintf(w, "  Forwarding to  %s\n", utils.SafeText(p.target))
	fmt.Fprintf(w, "  Expires        %s (in %s)\n", t.ExpiresAt.In(p.loc).Format("2006-01-02 15:04 MST"),
		roughDuration(t.ExpiresAt.Sub(p.now())))
	switch {
	case o == nil:
	case o.Linked:
		fmt.Fprintln(w, "  Linked to your Hostinger account")
	default:
		fmt.Fprintf(w, "  Not linked to an account (%s)\n", utils.SafeText(o.text))
	}
	fmt.Fprintln(w, "\nAnyone with the URL can reach this server, including debug consoles and dev tools.")
	fmt.Fprintln(w, "Press Ctrl+C to stop.")
}

// reachable reports that a Preview announced while still connecting now
// answers. It prints once.
func (p *printer) reachable(connector.QuickTunnel) {
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.reachDone {
		return
	}
	p.reachDone = true
	if p.json {
		p.emit(reachableEvent{Event: "reachable"})
		return
	}
	fmt.Fprintln(p.stdout, "Preview ready.")
}

func (p *printer) paused(q connector.QuickPause) {
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.json {
		wait := math.Ceil(q.RetryAt.Sub(p.now()).Seconds())
		p.emit(pausedEvent{Event: "paused", Reason: q.Reason, RetryAfter: int64(max(wait, 0))})
		return
	}
	at := q.RetryAt.In(p.loc).Format("15:04")
	if q.Reason == connector.PauseCapExceeded {
		fmt.Fprintf(p.stderr, "too many Previews from your network; retrying at %s\n", at)
		return
	}
	fmt.Fprintf(p.stderr, "Preview paused (%s); retrying at %s\n", utils.SafeText(q.Reason), at)
}

// ended reports the end of the run and returns the error that sets the exit
// code: nil for a normal end.
func (p *printer) ended(e ending) error {
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.json {
		if e.code == utils.ExitOK {
			p.emit(endedEvent{Event: "ended", Reason: e.reason})
			return nil
		}
		// stdout stays JSON only: the reason travels in the event, not on stderr.
		p.emit(endedEvent{Event: "ended", Reason: e.reason, Message: e.message})
		return &utils.ExitCodeError{Code: e.code}
	}
	if e.code == utils.ExitOK {
		fmt.Fprintln(p.stderr, e.message)
		return nil
	}
	return &utils.ExitCodeError{Code: e.code, Message: e.message}
}

// roughDuration is d to the minute: "23h59m", "45m", "under a minute".
func roughDuration(d time.Duration) string {
	d = d.Round(time.Minute)
	switch {
	case d < time.Minute:
		return "under a minute"
	case d < time.Hour:
		return fmt.Sprintf("%dm", int(d.Minutes()))
	case d%time.Hour == 0:
		return fmt.Sprintf("%dh", int(d.Hours()))
	}
	return fmt.Sprintf("%dh%dm", int(d.Hours()), int(d.Minutes())%60)
}
