package site

import (
	"errors"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/viper"
)

func TestResolveBuild(t *testing.T) {
	fakeAPI(t)
	viper.Set("no_input", true)
	t.Cleanup(func() { viper.Set("no_input", false) })
	const id = "123e4567-e89b-12d3-a456-426614174000"

	// Website and build given, with one extra argument after them.
	u, d, b, rest, err := ResolveBuild(nil, []string{"example.com", id, "extra"}, 1)
	if err != nil || d != "example.com" || u == "" || b != id || len(rest) != 1 || rest[0] != "extra" {
		t.Errorf("explicit: %s %s %s %v %v", u, d, b, rest, err)
	}

	// A mistyped build where only a build fits is reported, not taken for the website.
	for _, args := range [][]string{{"u1", "example.com", "not-a-uuid"}, {"example.com", "123e4567"}} {
		var fe *utils.FlagError
		if _, _, _, _, err = ResolveBuild(nil, args, 0); !errors.As(err, &fe) || !strings.Contains(err.Error(), "invalid build UUID") {
			t.Errorf("%v: %v", args, err)
		}
	}

	// Only the build given: the website comes from the environment.
	t.Setenv("HOSTINGER_DOMAIN", "example.com")
	if _, d, b, _, err = ResolveBuild(nil, []string{id}, 0); err != nil || d != "example.com" || b != id {
		t.Errorf("build only: %s %s %v", d, b, err)
	}

	// Nothing given and no link: the build cannot be known without a terminal.
	dir := t.TempDir()
	t.Chdir(dir)
	if _, _, _, _, err = ResolveBuild(nil, nil, 0); !errors.Is(err, ErrNoBuild) {
		t.Errorf("no build: %v", err)
	}

	// The link remembers the build `up` last started for that website.
	const linked = "9b2f4c1e-0d3a-4e5f-8a6b-7c8d9e0f1a2b"
	if err := Save(dir, &Project{Kind: "website", Domain: "example.com", Username: "u1", Build: &BuildRef{UUID: linked, State: "running", At: time.Now()}}); err != nil {
		t.Fatal(err)
	}
	os.Unsetenv("HOSTINGER_DOMAIN")
	if _, d, b, _, err = ResolveBuild(nil, nil, 0); err != nil || d != "example.com" || b != linked {
		t.Errorf("from link: %s %s %v", d, b, err)
	}
	// A malformed build in the link is ignored, not passed on to crash uuid.MustParse.
	if err := Save(dir, &Project{Kind: "website", Domain: "example.com", Username: "u1", Build: &BuildRef{UUID: "b-9"}}); err != nil {
		t.Fatal(err)
	}
	if _, _, _, _, err = ResolveBuild(nil, nil, 0); !errors.Is(err, ErrNoBuild) {
		t.Errorf("malformed link build: %v", err)
	}
	// A different website's build is not borrowed from the link.
	t.Setenv("HOSTINGER_DOMAIN", "blog.example.com")
	if _, _, _, _, err = ResolveBuild(nil, nil, 0); !errors.Is(err, ErrNoBuild) {
		t.Errorf("other website: %v", err)
	}
}
