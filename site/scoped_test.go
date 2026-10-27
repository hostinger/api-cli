package site

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/hostinger/api-cli/api"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func scopedCmd() *cobra.Command {
	c := &cobra.Command{Use: "x"}
	c.Flags().String("domain", "", "")
	c.Flags().String("username", "", "")
	return c
}

func portfolioAPI(t *testing.T, domains ...string) {
	t.Helper()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/api/domains/v1/portfolio":
			var rows []map[string]string
			for _, d := range domains {
				rows = append(rows, map[string]string{"domain": d})
			}
			_ = json.NewEncoder(w).Encode(rows)
		case "/api/hosting/v1/websites":
			_ = json.NewEncoder(w).Encode(map[string]any{"data": []Website{{Domain: "example.com", Username: "u1"}}, "meta": map[string]int{"total": 1, "per_page": 100}})
		default:
			http.NotFound(w, r)
		}
	}))
	t.Cleanup(srv.Close)
	old := Client
	Client = func() *api.RawClient { return api.NewRawClient(srv.URL, nil) }
	t.Cleanup(func() { Client = old })
}

func TestResolveUsername(t *testing.T) {
	portfolioAPI(t)
	chdir(t, t.TempDir())
	noPrompt(t)
	t.Setenv("HOSTINGER_USERNAME", "")
	t.Setenv("HOSTINGER_DOMAIN", "")
	viper.Reset()

	if u, rest, err := ResolveUsername(scopedCmd(), []string{"u9", "wp"}, 1); err != nil || u != "u9" || len(rest) != 1 {
		t.Errorf("positional: %q %v %v", u, rest, err)
	}
	if _, _, err := ResolveUsername(scopedCmd(), nil, 0); err != ErrNoUsername {
		t.Errorf("nothing: err = %v", err)
	}
	dir, _ := os.Getwd()
	if err := Save(dir, &Project{Domain: "other.com", Username: "linked"}); err != nil {
		t.Fatal(err)
	}
	if u, _, err := ResolveUsername(scopedCmd(), nil, 0); err != nil || u != "linked" {
		t.Errorf("link: %q %v", u, err)
	}
	// A domain that is not the linked one is looked up.
	t.Setenv("HOSTINGER_DOMAIN", "example.com")
	if u, _, err := ResolveUsername(scopedCmd(), nil, 0); err != nil || u != "u1" {
		t.Errorf("via domain lookup: %q %v", u, err)
	}
	t.Setenv("HOSTINGER_USERNAME", "u7")
	if u, _, err := ResolveUsername(scopedCmd(), nil, 0); err != nil || u != "u7" {
		t.Errorf("env: %q %v", u, err)
	}
	c := scopedCmd()
	_ = c.Flags().Set("username", "u8")
	if u, _, err := ResolveUsername(c, nil, 0); err != nil || u != "u8" {
		t.Errorf("flag: %q %v", u, err)
	}
}

func TestResolveDomainForDNS(t *testing.T) {
	portfolioAPI(t, "mine.com")
	chdir(t, t.TempDir())
	noPrompt(t)
	t.Setenv("HOSTINGER_DOMAIN", "")
	viper.Reset()

	if d, _, err := ResolveDomain(scopedCmd(), []string{"x.com"}, 0, "dns"); err != nil || d != "x.com" {
		t.Errorf("positional: %q %v", d, err)
	}
	t.Setenv("HOSTINGER_DOMAIN", "https://Mine.com/")
	if d, _, err := ResolveDomain(scopedCmd(), nil, 0, "dns"); err != nil || d != "mine.com" {
		t.Errorf("env URL: %q %v", d, err)
	}
	t.Setenv("HOSTINGER_DOMAIN", "")
	// A linked free subdomain is not a zone the account owns: no silent pick.
	dir, _ := os.Getwd()
	if err := Save(dir, &Project{Domain: "frog.hostingersite.com", Username: "u1"}); err != nil {
		t.Fatal(err)
	}
	if _, _, err := ResolveDomain(scopedCmd(), nil, 0, "dns"); err != ErrNoDomain {
		t.Errorf("free subdomain link: err = %v", err)
	}
	// A linked owned domain is used.
	if err := Save(dir, &Project{Domain: "mine.com", Username: "u1"}); err != nil {
		t.Fatal(err)
	}
	if d, _, err := ResolveDomain(scopedCmd(), nil, 0, "dns"); err != nil || d != "mine.com" {
		t.Errorf("owned link: %q %v", d, err)
	}
	// The link comes before default_domain, as for website commands.
	viper.Set("default_domain", "cfg.com")
	if d, _, err := ResolveDomain(scopedCmd(), nil, 0, "domains"); err != nil || d != "mine.com" {
		t.Errorf("link over config: %q %v", d, err)
	}
	if err := Save(dir, &Project{Domain: "frog.hostingersite.com", Username: "u1"}); err != nil {
		t.Fatal(err)
	}
	if d, _, err := ResolveDomain(scopedCmd(), nil, 0, "domains"); err != nil || d != "cfg.com" {
		t.Errorf("config after an unowned link: %q %v", d, err)
	}
	t.Setenv("HOSTINGER_DOMAIN", "env.com")
	if d, _, err := ResolveDomain(scopedCmd(), nil, 0, "dns"); err != nil || d != "env.com" {
		t.Errorf("env: %q %v", d, err)
	}
	// hosting group: the website link is used directly.
	t.Setenv("HOSTINGER_DOMAIN", "")
	if err := Save(dir, &Project{Domain: "frog.hostingersite.com", Username: "u1"}); err != nil {
		t.Fatal(err)
	}
	viper.Reset()
	if d, _, err := ResolveDomain(scopedCmd(), nil, 0, "hosting"); err != nil || d != "frog.hostingersite.com" {
		t.Errorf("hosting link: %q %v", d, err)
	}
}

func TestCommandsOwnFlagIsNotTheContextFlag(t *testing.T) {
	fakeAPI(t)
	chdir(t, t.TempDir())
	// hosting wordpress install has its own --domain: the new site's.
	root := &cobra.Command{Use: "root"}
	root.PersistentFlags().String("domain", "", "")
	install := &cobra.Command{Use: "install"}
	install.Flags().String("domain", "", "")
	root.AddCommand(install)
	if err := install.Flags().Set("domain", "new-site.com"); err != nil {
		t.Fatal(err)
	}
	t.Setenv("HOSTINGER_DOMAIN", "example.com")
	if d, err := DefaultDomain(install); err != nil || d != "example.com" {
		t.Errorf("domain = %q, %v", d, err)
	}
}
