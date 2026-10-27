package hosting

import (
	"context"
	"errors"
	"net"
	"strings"
)

// Tests replace LookupHost.
var LookupHost = func(ctx context.Context, host string) ([]string, error) {
	return net.DefaultResolver.LookupHost(ctx, host)
}

// platformSuffixes are free subdomains, which always point at the website.
var platformSuffixes = []string{".hostingersite.com"}

// reachability can only catch a domain with no DNS: the API exposes no
// server IP to compare against.
func reachability(ctx context.Context, domain string) (bool, string) {
	for _, suffix := range platformSuffixes {
		if strings.HasSuffix(strings.ToLower(domain), suffix) {
			return true, ""
		}
	}
	addrs, err := LookupHost(ctx, domain)
	if err == nil && len(addrs) > 0 {
		return true, ""
	}
	// Offline or a failing resolver says nothing about the domain.
	var dnsErr *net.DNSError
	if err != nil && !(errors.As(err, &dnsErr) && dnsErr.IsNotFound) {
		return true, ""
	}
	return false, domain + " has no DNS records yet, so https://" + domain + " will not answer until its DNS points to Hostinger. " +
		"Preview the website in hPanel (Websites, " + domain + ", Preview) in the meantime."
}
