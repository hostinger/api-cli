package connector

import (
	"context"

	"github.com/hostinger/hcdn-connector/internal/h2srv"
)

type (
	// Command is a command that the edge relays to the connector (POST
	// /_hcdn/v1/command, docs/protocol.md section 11). Fields: Body, the
	// request body as the edge relayed it, at most 8 KiB (the connector
	// answers a larger one 413 itself) and not parsed by the connector:
	// verify it before you act on it; Pop, the PoP of the session that
	// carried it; and EdgeNode, the edge node of that session.
	Command = h2srv.Command
	// CommandError refuses a command. Status is a 4xx code (any other
	// value is sent as 400). Code goes out in hcdn-error: lower-case
	// letters, digits and underscores, at most 64 bytes (any other value is
	// sent as "refused").
	CommandError = h2srv.CommandError
)

// CommandFunc handles a Command. Return nil to accept it (202), or a
// *CommandError to refuse it with a status and a reason; any other error
// refuses it with 400 and the reason "refused".
//
// It runs on the request goroutine of the session, and the edge waits for
// the answer: return quickly and do long work on another goroutine. ctx ends
// when the call returns, the edge drops the request, or the session closes,
// so do not pass it to work that outlives the call. Calls can run
// concurrently, and the edge can relay one command on several sessions. A panic resets the command's stream
// with INTERNAL_ERROR, so the edge gets no verdict; the session stays up.
type CommandFunc func(ctx context.Context, c Command) error
