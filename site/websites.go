package site

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/hostinger/api-cli/api"
)

type Website struct {
	Domain      string `json:"domain"`
	Username    string `json:"username"`
	OrderID     int    `json:"order_id"`
	WebsiteType string `json:"website_type"`
	VhostType   string `json:"vhost_type"`
	IsEnabled   bool   `json:"is_enabled"`
}

// Tests replace Client.
var Client = func() *api.RawClient { return api.Raw() }

// ListWebsites filters by domain as a substring; empty lists all.
func ListWebsites(ctx context.Context, domain string) ([]Website, error) {
	return ListWebsitesWith(ctx, Client(), domain)
}

func ListWebsitesWith(ctx context.Context, c *api.RawClient, domain string) ([]Website, error) {
	var all []Website
	for page := 1; ; page++ {
		q := url.Values{"page": {strconv.Itoa(page)}, "per_page": {"100"}}
		if domain != "" {
			q.Set("domain", domain)
		}
		var resp struct {
			Data []Website `json:"data"`
			Meta struct {
				Total   int `json:"total"`
				PerPage int `json:"per_page"`
			} `json:"meta"`
		}
		if err := c.JSON(ctx, http.MethodGet, "/api/hosting/v1/websites", q, nil, &resp); err != nil {
			return nil, err
		}
		all = append(all, resp.Data...)
		if len(resp.Data) == 0 || len(all) >= resp.Meta.Total || resp.Meta.PerPage == 0 {
			return all, nil
		}
	}
}

func FindWebsite(ctx context.Context, c *api.RawClient, domain string) (*Website, error) {
	return FindWebsiteWith(ctx, func(ctx context.Context, d string) ([]Website, error) { return ListWebsitesWith(ctx, c, d) }, domain)
}

// Username comes from the link when it names domain, else from the API.
func Username(ctx context.Context, domain string) (string, error) {
	if p, _, err := Find("."); err == nil && p != nil && strings.EqualFold(p.Domain, domain) && p.Username != "" {
		return p.Username, nil
	}
	w, err := FindWebsite(ctx, Client(), domain)
	if err != nil {
		return "", err
	}
	if w == nil {
		return "", fmt.Errorf("no website named %q on this account (hostinger hosting websites list shows them)", domain)
	}
	if w.Username == "" {
		return "", fmt.Errorf("website %q has no hosting username yet; pass --username", domain)
	}
	return w.Username, nil
}

// NormalizeDomain turns a pasted URL into a domain (no scheme, path, port,
// trailing dot or www.).
func NormalizeDomain(s string) string {
	s = strings.ToLower(strings.TrimSpace(s))
	if i := strings.Index(s, "://"); i >= 0 {
		s = s[i+3:]
	}
	if i := strings.IndexAny(s, "/?#"); i >= 0 {
		s = s[:i]
	}
	if i := strings.LastIndex(s, ":"); i >= 0 {
		s = s[:i]
	}
	return strings.TrimPrefix(strings.TrimSuffix(s, "."), "www.")
}

func FindWebsiteWith(ctx context.Context, list func(context.Context, string) ([]Website, error), domain string) (*Website, error) {
	domain = NormalizeDomain(domain)
	sites, err := list(ctx, domain)
	if err != nil {
		return nil, err
	}
	for i := range sites {
		if strings.EqualFold(sites[i].Domain, domain) {
			return &sites[i], nil
		}
	}
	return nil, nil
}
