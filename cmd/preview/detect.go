package preview

import (
	"context"
	"errors"
	"fmt"
	"net"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/hostinger/api-cli/utils"
)

// probePorts are the default ports of common dev servers, in the order a
// match is preferred: Vite, Next.js/Express/Rails, Astro, then generic ones,
// Angular and Jupyter.
var probePorts = []int{5173, 3000, 4321, 8080, 8000, 5000, 4200, 8888}

// probeHosts are tried for every port: a dev server may listen on IPv4 or
// IPv6 loopback only.
var probeHosts = []string{"127.0.0.1", "::1"}

// probeTimeout bounds one TCP connect.
const probeTimeout = 200 * time.Millisecond

// found is a port that accepted a connection, on the first host that did.
type found struct {
	port int
	host string
}

// url is the origin the Preview forwards to: the address that answered, so
// a server on only one loopback family is reached.
func (f found) url() string {
	return "http://" + net.JoinHostPort(f.host, strconv.Itoa(f.port))
}

func (f found) label() string {
	return fmt.Sprintf("%d (%s)", f.port, f.url())
}

// detector finds a local dev server. Its dependencies are fields so tests
// can fake the network and the terminal.
type detector struct {
	ports       []int
	dial        func(ctx context.Context, network, addr string) (net.Conn, error)
	interactive bool
	// choose asks which server to share and returns the index picked.
	choose func(labels []string) (int, error)
	// binary is the CLI's name, for the hints.
	binary string
}

// probe connects to every port on every loopback host in parallel and
// returns the ports that answered, in the order of d.ports.
func (d detector) probe(ctx context.Context) []found {
	hits := make([][]bool, len(d.ports))
	var wg sync.WaitGroup
	for i, port := range d.ports {
		hits[i] = make([]bool, len(probeHosts))
		for j, host := range probeHosts {
			wg.Add(1)
			go func() {
				defer wg.Done()
				pctx, cancel := context.WithTimeout(ctx, probeTimeout)
				defer cancel()
				conn, err := d.dial(pctx, "tcp", net.JoinHostPort(host, strconv.Itoa(port)))
				if err == nil {
					_ = conn.Close()
					hits[i][j] = true
				}
			}()
		}
	}
	wg.Wait()
	var out []found
	for i, port := range d.ports {
		for j, host := range probeHosts {
			if hits[i][j] {
				out = append(out, found{port: port, host: host})
				break
			}
		}
	}
	return out
}

// detect returns the origin URL of the one dev server found. Several are
// offered in a prompt on a terminal and are an error otherwise.
func (d detector) detect(ctx context.Context) (string, error) {
	hits := d.probe(ctx)
	if err := ctx.Err(); err != nil {
		return "", err
	}
	hint := fmt.Sprintf(`run "%s preview <port>"`, d.binary)
	switch {
	case len(hits) == 0:
		return "", fmt.Errorf("no server found on ports %s; %s", joinPorts(d.ports), hint)
	case len(hits) == 1:
		return hits[0].url(), nil
	case !d.interactive || d.choose == nil:
		ports := make([]int, len(hits))
		for i, h := range hits {
			ports[i] = h.port
		}
		return "", fmt.Errorf("servers found on ports %s; %s with the one to share", joinPorts(ports), hint)
	}
	labels := make([]string, len(hits))
	for i, h := range hits {
		labels[i] = h.label()
	}
	i, err := d.choose(labels)
	if err != nil {
		return "", err
	}
	if i < 0 || i >= len(hits) {
		return "", errors.New("no server picked")
	}
	return hits[i].url(), nil
}

func joinPorts(ports []int) string {
	s := make([]string, len(ports))
	for i, p := range ports {
		s[i] = strconv.Itoa(p)
	}
	return strings.Join(s, ", ")
}

// parsePort reads a port argument: decimal digits only, 1 to 65535.
func parsePort(s string) (int, error) {
	invalid := utils.FlagErrorf("invalid port %q: want a number from 1 to 65535", s)
	if s == "" || len(s) > 5 {
		return 0, invalid
	}
	for _, c := range s {
		if c < '0' || c > '9' {
			return 0, invalid
		}
	}
	p, err := strconv.Atoi(s)
	if err != nil || p < 1 || p > 65535 {
		return 0, invalid
	}
	return p, nil
}
