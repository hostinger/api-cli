package preview

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"net/netip"
	"strings"

	"github.com/hostinger/api-cli/utils"
	"github.com/hostinger/hcdn-connector/pkg/localguard"
)

// checkTarget refuses an origin that reaches a sensitive port on this
// computer: SSH, databases, debuggers, cluster agents and the like. The rule
// is the connector library's, which also applies it to every connection the
// Preview makes, so a redirect or a changed DNS answer cannot get around it.
// Any form that reaches this computer counts: 127.1, 0, a name that
// resolves to loopback, or one of its own network addresses. A DNS lookup
// ends with ctx.
func checkTarget(ctx context.Context, target string) error {
	err := localguard.CheckContext(ctx, target)
	if err == nil {
		return nil
	}
	var be *localguard.BlockedError
	if errors.As(err, &be) {
		return utils.FlagErrorf("%s", be.Error())
	}
	return utils.FlagErrorf("--url %q: %v", target, err)
}

// warnSharedPort warns about a port that is shared but usually serves a
// sensitive tool, such as Jupyter's 8888, so it is never shared silently.
func warnSharedPort(ctx context.Context, w io.Writer, target string) {
	p, err := localguard.Match(ctx, target)
	if err != nil || p == nil || !p.WarnOnly {
		return
	}
	fmt.Fprintf(w, "Warning: port %d is the usual port of %s; anyone with the URL can run code on this computer unless it asks for a token or password\n",
		p.Number, p.Service)
}

// warnHostHeader warns when --host-header is a value that dev servers such
// as Vite always accept (localhost, *.localhost or any IP address): their
// host check then lets every visitor through.
func warnHostHeader(w io.Writer, value string) {
	if !bypassesHostCheck(value) {
		return
	}
	fmt.Fprintf(w, "Warning: --host-header %s bypasses the dev server's host check, and its debug consoles become public\n",
		utils.SafeText(value))
}

// bypassesHostCheck reports a Host header value that a dev server's host
// check always allows: localhost, *.localhost or an IP address, with or
// without a port.
func bypassesHostCheck(value string) bool {
	h := strings.ToLower(strings.TrimSpace(value))
	if host, _, err := net.SplitHostPort(h); err == nil {
		h = host
	}
	h = strings.TrimSuffix(strings.TrimSuffix(strings.TrimPrefix(h, "["), "]"), ".")
	if h == "localhost" || strings.HasSuffix(h, ".localhost") {
		return true
	}
	_, err := netip.ParseAddr(h)
	return err == nil
}
