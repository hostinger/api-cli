package site

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/hostinger/api-cli/api"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// fakeAPI answers the websites list with one site.
func fakeAPI(t *testing.T) *httptest.Server {
	t.Helper()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/hosting/v1/websites" {
			http.NotFound(w, r)
			return
		}
		filter := r.URL.Query().Get("domain")
		sites := []Website{{Domain: "example.com", Username: "u1"}, {Domain: "blog.example.com", Username: "u1"}, {Domain: "other.net", Username: "u2"}}
		var data []Website
		for _, s := range sites {
			if filter == "" || contains(s.Domain, filter) {
				data = append(data, s)
			}
		}
		_ = json.NewEncoder(w).Encode(map[string]any{"data": data, "meta": map[string]int{"total": len(data), "per_page": 100}})
	}))
	t.Cleanup(srv.Close)
	old := Client
	Client = func() *api.RawClient { return api.NewRawClient(srv.URL, nil) }
	t.Cleanup(func() { Client = old })
	return srv
}

func contains(s, sub string) bool {
	for i := 0; i+len(sub) <= len(s); i++ {
		if s[i:i+len(sub)] == sub {
			return true
		}
	}
	return false
}

func newCmd() *cobra.Command {
	c := &cobra.Command{Use: "x"}
	c.Flags().String("domain", "", "")
	return c
}

func chdir(t *testing.T, dir string) {
	t.Helper()
	old, _ := os.Getwd()
	if err := os.Chdir(dir); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = os.Chdir(old) })
}

func TestResolvePositional(t *testing.T) {
	fakeAPI(t)
	u, d, rest, err := Resolve(newCmd(), []string{"u9", "example.com", "extra"}, 1)
	if err != nil || u != "u9" || d != "example.com" || len(rest) != 1 || rest[0] != "extra" {
		t.Errorf("both given: %q %q %v %v", u, d, rest, err)
	}
	u, d, rest, err = Resolve(newCmd(), []string{"example.com", "extra"}, 1)
	if err != nil || u != "u1" || d != "example.com" || len(rest) != 1 {
		t.Errorf("domain only: %q %q %v %v", u, d, rest, err)
	}
	// Exact match: "example.com" must not resolve to blog.example.com's row first.
	u, d, _, err = Resolve(newCmd(), []string{"other.net"}, 0)
	if err != nil || u != "u2" || d != "other.net" {
		t.Errorf("other.net: %q %q %v", u, d, err)
	}
	if _, _, _, err := Resolve(newCmd(), []string{"nope.com"}, 0); err == nil {
		t.Error("unknown domain should fail")
	}
}

func TestResolveFromFlagEnvProjectConfig(t *testing.T) {
	fakeAPI(t)
	chdir(t, t.TempDir())
	noPrompt(t)
	viper.Reset()
	t.Setenv("HOSTINGER_DOMAIN", "")

	if _, _, _, err := Resolve(newCmd(), nil, 0); err != ErrNoWebsite {
		t.Errorf("nothing set: err = %v", err)
	}

	viper.Set("default_domain", "other.net")
	if u, d, _, err := Resolve(newCmd(), nil, 0); err != nil || d != "other.net" || u != "u2" {
		t.Errorf("config: %q %q %v", u, d, err)
	}

	dir, _ := os.Getwd()
	if err := Save(dir, &Project{Domain: "blog.example.com", Username: "cached"}); err != nil {
		t.Fatal(err)
	}
	if u, d, _, err := Resolve(newCmd(), nil, 0); err != nil || d != "blog.example.com" || u != "cached" {
		t.Errorf("project file: %q %q %v", u, d, err)
	}

	t.Setenv("HOSTINGER_DOMAIN", "example.com")
	if u, d, _, err := Resolve(newCmd(), nil, 0); err != nil || d != "example.com" || u != "u1" {
		t.Errorf("env: %q %q %v", u, d, err)
	}

	cmd := newCmd()
	_ = cmd.Flags().Set("domain", "other.net")
	if u, d, _, err := Resolve(cmd, nil, 0); err != nil || d != "other.net" || u != "u2" {
		t.Errorf("flag: %q %q %v", u, d, err)
	}
}

func TestBrokenLinkOnlyMattersWhenRead(t *testing.T) {
	fakeAPI(t)
	viper.Set("no_input", true)
	t.Cleanup(func() { viper.Set("no_input", false) })
	dir := t.TempDir()
	t.Chdir(dir)
	if err := os.MkdirAll(dir+"/"+Dir, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(Path(dir), []byte("{broken"), 0o600); err != nil {
		t.Fatal(err)
	}
	t.Setenv("HOSTINGER_DOMAIN", "https://Example.com/")
	if u, d, _, err := Resolve(nil, nil, 0); err != nil || d != "example.com" || u != "u1" {
		t.Errorf("env given: %q %q %v", u, d, err)
	}
	t.Setenv("HOSTINGER_DOMAIN", "")
	if _, _, _, err := Resolve(nil, nil, 0); err == nil || !strings.Contains(err.Error(), "hostinger unlink") {
		t.Errorf("link needed: %v", err)
	}
}

// noPrompt stubs Interactive for one test and puts the real one back.
func noPrompt(t *testing.T) {
	t.Helper()
	old := Interactive
	Interactive = func() bool { return false }
	t.Cleanup(func() { Interactive = old })
}
