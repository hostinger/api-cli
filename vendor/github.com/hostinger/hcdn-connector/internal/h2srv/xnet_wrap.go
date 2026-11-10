//go:build !http2legacy

//lint:file-ignore SA1019 http2.Server.ServeConn on a dialed conn has no net/http equivalent.

package h2srv

import (
	"net/http"

	"golang.org/x/net/http2"
)

// No BaseConfig: x/net (Go 1.27+) then serves on a copy of hs, and Drain's
// hs.Shutdown never reaches the conn. ConfigureServer already bound the server to hs.
func serveConnOpts(_ *http.Server, h http.Handler) *http2.ServeConnOpts {
	return &http2.ServeConnOpts{Handler: h}
}
