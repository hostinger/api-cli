// Package proto holds the wire contract shared with tunneld (docs/protocol.md).
package proto

import "time"

const (
	Version = "1"

	PathHello   = "/_hcdn/v1/hello"
	PathGoaway  = "/_hcdn/v1/goaway"
	PathVisible = "/_hcdn/v1/visible"
	PathPops    = "/_hcdn/v1/pops"
	PathPing    = "/_hcdn/v1/ping"
	PathWhoami  = "/_hcdn/v1/whoami"
	PathConnect = "/_tunnel/v1/connect"

	ALPNSession = "hcdn-tunnel/1"
	ALPNAPI     = "hcdn-tunnel-api/1"
	ALPNHTTP1   = "http/1.1"

	UpgradeToken = "hcdn-tunnel"

	PortDirect   = 7844
	PortFallback = 443

	TransportDirect   = "7844"
	TransportFallback = "443-fallback"

	HdrProtocol       = "Hcdn-Protocol"
	HdrEdgeNode       = "Hcdn-Edge-Node"
	HdrEdgePop        = "Hcdn-Edge-Pop"
	HdrCertDaysLeft   = "Hcdn-Cert-Days-Left"
	HdrMaxInflight    = "Hcdn-Max-Inflight"
	HdrMaxStreams     = "Hcdn-Max-Streams"
	HdrPopsGen        = "Hcdn-Pops-Gen"
	HdrConnVersion    = "Hcdn-Connector-Version"
	HdrConnID         = "Hcdn-Connector-Id"
	HdrFeatures       = "Hcdn-Features"
	HdrConnectedPops  = "Hcdn-Connected-Pops"
	HdrTransport      = "Hcdn-Transport"
	HdrReason         = "Hcdn-Reason"
	HdrRetryAfter     = "Hcdn-Retry-After"
	HdrRequestID      = "Hcdn-Request-Id"
	HdrVhost          = "Hcdn-Vhost"
	HdrError          = "Hcdn-Error"
	HdrConnectorRT    = "Hcdn-Connector-Rt"
	HdrTunnelID       = "Hcdn-Tunnel-Id"
	HdrUpgrade        = "Hcdn-Upgrade"
	HdrPath           = "Hcdn-Path"
	HdrWSAccept       = "Hcdn-Ws-Accept"
	HdrWSProtocol     = "Hcdn-Ws-Protocol"
	HdrOwnerTicket    = "Hcdn-Owner-Ticket"
	HdrOwnerAbsent    = "Hcdn-Owner-Absent"
	HdrEdgeFeatures   = "Hcdn-Edge-Features"
	HdrVisibleScope   = "Hcdn-Visible-Scope"
	HdrControlPrefix  = "Hcdn-"
	HdrXControlPrefix = "X-Hcdn-"

	Features = "ws,owner,visible"

	// EdgeFeatureVisible in hcdn-edge-features: the edge sends a visible
	// control request once the tunnel is in its presence.
	EdgeFeatureVisible = "visible"

	// hcdn-visible-scope values: the edge confirmed every node of its PoP,
	// or only itself.
	VisibleScopePop  = "pop"
	VisibleScopeNode = "node"

	// hcdn-error values. tunneld renders a 530 with X-Hcdn-Tunnel-Error set
	// to the value. overloaded extends the original set (per-session inflight cap).
	ErrServiceUnreachable = "service_unreachable"
	ErrServiceTimeout     = "service_timeout"
	ErrBadGateway         = "bad_gateway"
	ErrOverloaded         = "overloaded"

	ReasonDrain           = "drain"
	ReasonRevoked         = "revoked"
	ReasonDisabled        = "disabled"
	ReasonUnknownSerial   = "unknown_serial"
	ReasonCapExceeded     = "cap_exceeded"
	ReasonAuthUnavailable = "auth_unavailable"
	ReasonReselect        = "reselect"
	ReasonProtocolError   = "protocol_error"
	ReasonOwnerRequired   = "owner_required"

	DefaultRetryAfterDisabled        = 300 * time.Second
	DefaultRetryAfterCapExceeded     = 60 * time.Second
	DefaultRetryAfterAuthUnavailable = 15 * time.Second

	HelloTimeout = 10 * time.Second

	// Exit code for revoked / unknown_serial; RestartPreventExitStatus=78 in the unit.
	ExitRevoked = 78
	// Exit code for owner_required (EX_NOPERM): the edge wants a linked owner.
	ExitOwnerRequired = 77

	MaxConcurrentStreams = 256
	MaxFrameSize         = 64 << 10
	MaxHeaderBytes       = 64 << 10
	DefaultMaxInflight   = 512
)

// KnownReasons lists the goaway reasons the connector acts on.
var KnownReasons = map[string]bool{
	ReasonDrain:           true,
	ReasonRevoked:         true,
	ReasonDisabled:        true,
	ReasonUnknownSerial:   true,
	ReasonCapExceeded:     true,
	ReasonAuthUnavailable: true,
	ReasonReselect:        true,
	ReasonProtocolError:   true,
	ReasonOwnerRequired:   true,
}

// Owner ticket (docs/protocol.md section 9). The whoami owner.mode values that
// ask for a ticket; tunneld omits the object in mode off.
const (
	OwnerModeLog     = "log"
	OwnerModeRequire = "require"

	OwnerStatusOK = "ok"

	// hcdn-owner-absent values: why a hello answer carries no ticket.
	OwnerAbsentNoToken      = "no_token"
	OwnerAbsentOptedOut     = "opted_out"
	OwnerAbsentFetchFailed  = "fetch_failed"
	OwnerAbsentNotRequested = "not_requested"
	OwnerAbsentUnsupported  = "unsupported"
)

// OwnerAbsentReasons lists the hcdn-owner-absent values the edge counts.
var OwnerAbsentReasons = map[string]bool{
	OwnerAbsentNoToken:      true,
	OwnerAbsentOptedOut:     true,
	OwnerAbsentFetchFailed:  true,
	OwnerAbsentNotRequested: true,
	OwnerAbsentUnsupported:  true,
}

// DefaultRetryAfter returns the default for reasons that carry one.
func DefaultRetryAfter(reason string) time.Duration {
	switch reason {
	case ReasonDisabled:
		return DefaultRetryAfterDisabled
	case ReasonCapExceeded:
		return DefaultRetryAfterCapExceeded
	case ReasonAuthUnavailable:
		return DefaultRetryAfterAuthUnavailable
	}
	return 0
}

const (
	MaxPause           = time.Hour
	MaxRetryAfterShort = 15 * time.Minute
)

// MaxRetryAfter is the ceiling for an edge-supplied hcdn-retry-after.
func MaxRetryAfter(reason string) time.Duration {
	switch reason {
	case ReasonCapExceeded, ReasonAuthUnavailable:
		return MaxRetryAfterShort
	}
	return MaxPause
}
