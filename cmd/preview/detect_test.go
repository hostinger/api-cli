package preview

import (
	"context"
	"errors"
	"net"
	"reflect"
	"strconv"
	"sync"
	"testing"
	"time"
)

// fakeDialer accepts connections to the listed host:port addresses and
// refuses the rest, recording every dial.
type fakeDialer struct {
	mu      sync.Mutex
	open    map[string]bool
	dialed  []string
	noLimit []string // dials without a deadline within probeTimeout
}

func newFakeDialer(open ...string) *fakeDialer {
	f := &fakeDialer{open: map[string]bool{}}
	for _, a := range open {
		f.open[a] = true
	}
	return f
}

func (f *fakeDialer) dial(ctx context.Context, network, addr string) (net.Conn, error) {
	f.mu.Lock()
	f.dialed = append(f.dialed, addr)
	if d, ok := ctx.Deadline(); !ok || time.Until(d) > probeTimeout {
		f.noLimit = append(f.noLimit, addr)
	}
	f.mu.Unlock()
	if network != "tcp" {
		return nil, errors.New("unexpected network " + network)
	}
	if f.open[addr] {
		c1, c2 := net.Pipe()
		_ = c2.Close()
		return c1, nil
	}
	return nil, errors.New("connection refused")
}

func testDetector(f *fakeDialer, interactive bool, choose func([]string) (int, error)) detector {
	return detector{ports: probePorts, dial: f.dial, interactive: interactive, choose: choose, binary: "hostinger"}
}

func TestDetectNone(t *testing.T) {
	f := newFakeDialer()
	_, err := testDetector(f, true, nil).detect(context.Background())
	want := `no server found on ports 5173, 3000, 4321, 8080, 8000, 5000, 4200, 8888; run "hostinger preview <port>"`
	if err == nil || err.Error() != want {
		t.Fatalf("err = %v, want %q", err, want)
	}
	if len(f.dialed) != len(probePorts)*2 {
		t.Errorf("dialed %d addresses, want %d (every port on 127.0.0.1 and ::1)", len(f.dialed), len(probePorts)*2)
	}
	if len(f.noLimit) > 0 {
		t.Errorf("dials without a %s deadline: %v", probeTimeout, f.noLimit)
	}
}

func TestDetectOne(t *testing.T) {
	for _, interactive := range []bool{false, true} {
		got, err := testDetector(newFakeDialer("127.0.0.1:4321"), interactive, func([]string) (int, error) {
			t.Error("asked with a single server")
			return 0, nil
		}).detect(context.Background())
		if err != nil || got != "http://127.0.0.1:4321" {
			t.Errorf("interactive=%v: got %q, %v", interactive, got, err)
		}
	}
}

func TestDetectIPv6Only(t *testing.T) {
	got, err := testDetector(newFakeDialer("[::1]:3000"), false, nil).detect(context.Background())
	if err != nil || got != "http://[::1]:3000" {
		t.Fatalf("got %q, %v", got, err)
	}
}

func TestDetectBothFamiliesCountOnce(t *testing.T) {
	got, err := testDetector(newFakeDialer("127.0.0.1:8080", "[::1]:8080"), false, nil).detect(context.Background())
	if err != nil || got != "http://127.0.0.1:8080" {
		t.Fatalf("got %q, %v", got, err)
	}
}

func TestDetectManyWithoutTerminal(t *testing.T) {
	f := newFakeDialer("127.0.0.1:8888", "[::1]:3000", "127.0.0.1:5173")
	_, err := testDetector(f, false, func([]string) (int, error) {
		t.Error("asked without a terminal")
		return 0, nil
	}).detect(context.Background())
	want := `servers found on ports 5173, 3000, 8888; run "hostinger preview <port>" with the one to share`
	if err == nil || err.Error() != want {
		t.Fatalf("err = %v, want %q", err, want)
	}
}

func TestDetectManyAsksOnTerminal(t *testing.T) {
	f := newFakeDialer("127.0.0.1:8888", "[::1]:3000", "127.0.0.1:5173")
	var offered []string
	got, err := testDetector(f, true, func(labels []string) (int, error) {
		offered = labels
		return 1, nil
	}).detect(context.Background())
	if err != nil || got != "http://[::1]:3000" {
		t.Fatalf("got %q, %v", got, err)
	}
	want := []string{"5173 (http://127.0.0.1:5173)", "3000 (http://[::1]:3000)", "8888 (http://127.0.0.1:8888)"}
	if !reflect.DeepEqual(offered, want) {
		t.Errorf("offered %q, want %q", offered, want)
	}

	aborted := errors.New("aborted")
	if _, err := testDetector(f, true, func([]string) (int, error) { return 0, aborted }).detect(context.Background()); !errors.Is(err, aborted) {
		t.Errorf("a failed prompt: err = %v", err)
	}
}

func TestDetectCancelled(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if _, err := testDetector(newFakeDialer("127.0.0.1:5173"), false, nil).detect(ctx); !errors.Is(err, context.Canceled) {
		t.Fatalf("err = %v, want context.Canceled", err)
	}
}

func TestDetectRealListener(t *testing.T) {
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Skip("no loopback listener:", err)
	}
	defer ln.Close()
	port := ln.Addr().(*net.TCPAddr).Port
	d := detector{ports: []int{port}, dial: (&net.Dialer{}).DialContext, binary: "hostinger"}
	got, err := d.detect(context.Background())
	if err != nil || got != "http://127.0.0.1:"+strconv.Itoa(port) {
		t.Fatalf("got %q, %v", got, err)
	}
}

func TestParsePort(t *testing.T) {
	for in, want := range map[string]int{"1": 1, "80": 80, "5173": 5173, "65535": 65535, "08080": 8080} {
		if got, err := parsePort(in); err != nil || got != want {
			t.Errorf("parsePort(%q) = %d, %v; want %d", in, got, err, want)
		}
	}
	for _, in := range []string{"", "0", "65536", "99999", "123456", "-1", "+80", " 80", "80 ", "8o", "0x50", "stop", "1e3", "٣٠٠٠"} {
		if got, err := parsePort(in); err == nil {
			t.Errorf("parsePort(%q) = %d, want an error", in, got)
		} else if want := `invalid port "` + in + `": want a number from 1 to 65535`; err.Error() != want {
			t.Errorf("parsePort(%q) error %q, want %q", in, err, want)
		}
	}
}
