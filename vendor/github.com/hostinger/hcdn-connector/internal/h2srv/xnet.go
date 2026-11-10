//lint:file-ignore SA1019 http2.Server.ServeConn on a dialed conn has no net/http equivalent.

package h2srv

import (
	"crypto/tls"
	"io"
	"net"
	"net/http"
	"sync/atomic"

	"golang.org/x/net/http2"
)

// h2server wraps the x/net server; all HTTP/2 settings come from hs.HTTP2.
type h2server struct{ srv *http2.Server }

func newH2Server(hs *http.Server) h2server {
	s := &http2.Server{}
	_ = http2.ConfigureServer(hs, s)
	return h2server{srv: s}
}

// No Context: x/net uses it only as the parent of every stream's request
// context, and a supervisor cancel must not abort in-flight streams mid-drain.
func (h h2server) serveConn(conn net.Conn, hs *http.Server, handler http.Handler) {
	h.srv.ServeConn(conn, serveConnOpts(hs, handler))
}

type connectionStater interface{ ConnectionState() tls.ConnectionState }

// readErrConn keeps the first read error: ServeConn swallows it, and with TLS 1.3
// a client-certificate rejection only shows up there (post-handshake alert).
type readErrConn struct {
	net.Conn
	first atomic.Pointer[error]
}

func (c *readErrConn) Read(b []byte) (int, error) {
	n, err := c.Conn.Read(b)
	if err != nil && err != io.EOF {
		c.first.CompareAndSwap(nil, &err)
	}
	return n, err
}

func (c *readErrConn) Err() error {
	if p := c.first.Load(); p != nil {
		return *p
	}
	return nil
}

// readErrTLSConn keeps ConnectionState reachable through the wrapper, so x/net
// still runs its TLS checks and fills r.TLS (session.Conn forwards it).
type readErrTLSConn struct {
	*readErrConn
	cs connectionStater
}

func (c *readErrTLSConn) ConnectionState() tls.ConnectionState { return c.cs.ConnectionState() }

func wrapConn(c net.Conn) (net.Conn, *readErrConn) {
	rc := &readErrConn{Conn: c}
	if cs, ok := c.(connectionStater); ok {
		return &readErrTLSConn{readErrConn: rc, cs: cs}, rc
	}
	return rc, rc
}
