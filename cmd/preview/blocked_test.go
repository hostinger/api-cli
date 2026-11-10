package preview

import (
	"bytes"
	"context"
	"errors"
	"strings"
	"testing"

	"github.com/hostinger/api-cli/utils"
	"github.com/hostinger/hcdn-connector/pkg/localguard"
)

func TestCheckTarget(t *testing.T) {
	cases := []struct {
		url     string
		blocked bool
	}{
		{"http://localhost:6379", true},
		{"http://LOCALHOST.:3306", true},
		{"http://app.localhost:5432", true},
		{"http://127.0.0.1:22", true},
		{"http://127.1:9229", true},
		{"http://0:22", true},
		{"http://0x7f000001:9222", true},
		{"http://2130706433:6060", true},
		{"http://127.1.2.3:27017", true},
		{"http://[::1]:9200", true},
		{"http://[::1%25lo0]:11211", true},
		{"http://0.0.0.0:2375", true},
		{"http://[::]:8500", true},
		{"http://[::ffff:127.0.0.1]:6443", true},
		{"https://localhost:25", true},
		{"http://localhost:5173", false},
		{"http://localhost:8888", false}, // Jupyter: shared with a warning
		{"http://localhost", false},
		{"https://localhost", false},
		{"http://192.0.2.20:3306", false}, // another machine: its ports mean nothing here
		{"http://[2001:db8::1]:22", false},
	}
	for _, c := range cases {
		err := checkTarget(context.Background(), c.url)
		if got := err != nil; got != c.blocked {
			t.Errorf("checkTarget(%s) = %v, want blocked %v", c.url, err, c.blocked)
			continue
		}
		var fe *utils.FlagError
		if c.blocked && (!errors.As(err, &fe) || !strings.Contains(err.Error(), "is not shared")) {
			t.Errorf("checkTarget(%s) = %v, want a usage error", c.url, err)
		}
	}
}

func TestCheckTargetMessage(t *testing.T) {
	want := "port 9229 is not shared: it is the usual port of the Node inspector, which must not be reachable from the internet"
	if err := checkTarget(context.Background(), "http://0:9229"); err == nil || err.Error() != want {
		t.Fatalf("err = %v, want %q", err, want)
	}
}

// Every port that detection probes is allowed; only Jupyter's warns.
func TestProbePortsAreShared(t *testing.T) {
	for _, port := range probePorts {
		p, listed := localguard.LookupPort(port)
		if listed && !p.WarnOnly {
			t.Errorf("probe port %d is blocked (%s)", port, p.Service)
		}
		if listed != (port == 8888) {
			t.Errorf("probe port %d: listed %v", port, listed)
		}
	}
}

func TestWarnSharedPort(t *testing.T) {
	var w bytes.Buffer
	warnSharedPort(context.Background(), &w, "http://127.0.0.1:8888")
	want := "Warning: port 8888 is the usual port of Jupyter; anyone with the URL can run code on this computer unless it asks for a token or password\n"
	if w.String() != want {
		t.Fatalf("got %q, want %q", w.String(), want)
	}
	for _, u := range []string{"http://127.0.0.1:5173", "http://192.0.2.20:8888"} {
		w.Reset()
		if warnSharedPort(context.Background(), &w, u); w.Len() > 0 {
			t.Errorf("%s: unexpected warning %q", u, w.String())
		}
	}
}

func TestWarnHostHeader(t *testing.T) {
	for _, v := range []string{"localhost", "LOCALHOST", "localhost:5173", "app.localhost", "127.0.0.1", "127.0.0.1:3000", "[::1]", "[::1]:5173", "::1",
		"0.0.0.0", "192.0.2.1", "192.168.1.5:5173", "[::]", "[2001:db8::1]:8080", "fe80::1%en0"} {
		var w bytes.Buffer
		warnHostHeader(&w, v)
		want := "Warning: --host-header " + v + " bypasses the dev server's host check, and its debug consoles become public\n"
		if w.String() != want {
			t.Errorf("%q: got %q, want %q", v, w.String(), want)
		}
	}
	for _, v := range []string{"", "shop.test", "localhost.example.com", "192.0.2.1.example.com", "[shop.test]"} {
		var w bytes.Buffer
		if warnHostHeader(&w, v); w.Len() > 0 {
			t.Errorf("%q: unexpected warning %q", v, w.String())
		}
	}
}
