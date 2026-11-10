//go:build http2legacy

//lint:file-ignore SA1019 http2.Server.ServeConn on a dialed conn has no net/http equivalent.

package h2srv

import (
	"net/http"

	"golang.org/x/net/http2"
)

// The legacy x/net server reads the HTTP/2 settings from BaseConfig only.
func serveConnOpts(hs *http.Server, h http.Handler) *http2.ServeConnOpts {
	return &http2.ServeConnOpts{BaseConfig: hs, Handler: h}
}
