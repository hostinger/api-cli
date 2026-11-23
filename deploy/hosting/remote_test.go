package hosting

import (
	"context"
	"errors"
	"fmt"
	"path/filepath"
	"strings"
	"testing"

	"github.com/hostinger/api-cli/site"
)

func refusal(t *testing.T, err error) string {
	t.Helper()
	var r *RefusedError
	if !errors.As(err, &r) {
		t.Fatalf("err = %v, want a *RefusedError", err)
	}
	return r.Reason
}

func TestRemoteFirstDeployRefusesToOverwriteContent(t *testing.T) {
	fastPolling(t)
	api := newFake(t)
	api.settings = NodeSettings{AppType: "express", NodeVersion: 22}
	api.content = true
	dir := nodeProject(t)

	_, err := Run(context.Background(), api, &fakeUI{t: t}, Options{Path: dir, Domain: "example.com", Remote: true})
	if got := refusal(t, err); got != ReasonWouldOverwrite || !strings.HasPrefix(err.Error(), "example.com already has ") {
		t.Errorf("reason = %q (%v)", got, err)
	}
	if strings.Contains(strings.Join(api.calls, "\n"), "creds") {
		t.Error("uploaded after a refused overwrite")
	}
	if p, _ := site.Load(dir); p != nil {
		t.Errorf("refused deploy linked: %+v", p)
	}

	// Yes is hPanel's confirmation: the deploy runs and links the folder.
	res, err := Run(context.Background(), api, &fakeUI{t: t}, Options{Path: dir, Domain: "example.com", Remote: true, Yes: true, Message: "From hPanel"})
	if err != nil {
		t.Fatal(err)
	}
	if res.Type != NodeJS || api.message != "From hPanel" {
		t.Errorf("result = %+v, message %q", res, api.message)
	}
	if p, _ := site.Load(dir); p == nil || p.Domain != "example.com" {
		t.Errorf("project file = %+v", p)
	}
}

func TestRemoteFirstDeployToAnEmptyWebsiteNeedsNoConfirmation(t *testing.T) {
	fastPolling(t)
	api := newFake(t)
	api.settings = NodeSettings{AppType: "express", NodeVersion: 22}
	if _, err := Run(context.Background(), api, &fakeUI{t: t}, Options{Path: nodeProject(t), Domain: "example.com", Remote: true}); err != nil {
		t.Fatal(err)
	}
}

func TestRemoteRefusesStaticFilesOntoANodeApp(t *testing.T) {
	fastPolling(t)
	api := newFake(t)
	api.websites[0].WebsiteType = "nodejs"
	dir := staticProject(t)
	if err := site.Save(dir, &site.Project{Domain: "example.com", Username: "u1", Type: "static"}); err != nil {
		t.Fatal(err)
	}
	_, err := Run(context.Background(), api, &fakeUI{t: t}, Options{Path: dir, Remote: true, Yes: true})
	if got := refusal(t, err); got != ReasonUnsupportedType {
		t.Errorf("reason = %q", got)
	}
}

func TestRemoteRefusesAWordPressImport(t *testing.T) {
	fastPolling(t)
	api := newFake(t)
	dir := t.TempDir()
	linkHere(t, dir)
	write(t, dir, "wp-config.php", "<?php")
	write(t, dir, "database.sql", "CREATE TABLE x (id int);")
	if err := site.Save(dir, &site.Project{Domain: "example.com", Username: "u1", Type: "wordpress-site"}); err != nil {
		t.Fatal(err)
	}
	_, err := Run(context.Background(), api, &fakeUI{t: t}, Options{Path: dir, Remote: true, Yes: true})
	if got := refusal(t, err); got != ReasonUnsupportedType {
		t.Errorf("reason = %q", got)
	}
	if strings.Contains(strings.Join(api.calls, "\n"), "creds") {
		t.Error("uploaded a refused import")
	}
}

func TestFailureReason(t *testing.T) {
	fastPolling(t)
	api := newFake(t)
	api.upload.Close()
	_, err := Run(context.Background(), api, &fakeUI{t: t}, Options{Path: staticProject(t), Domain: "example.com", Yes: true})
	// Its message is the cause's.
	if got := FailureReason(err); got != ReasonUploadFailed || !strings.HasPrefix(err.Error(), "starting upload: ") {
		t.Errorf("a failed upload: %q (%v)", got, err)
	}
	for err, want := range map[error]string{
		&RefusedError{Reason: ReasonWouldOverwrite}: ReasonWouldOverwrite,
		&BuildError{UUID: "b"}:                      ReasonBuildFailed,
		fmt.Errorf("packing: %w", context.Canceled): ReasonCancelled,
		errors.New("HTTP 500"):                      ReasonDeployFailed,
	} {
		if got := FailureReason(err); got != want {
			t.Errorf("%v: %q, want %q", err, got, want)
		}
	}
}

func TestLinkedType(t *testing.T) {
	static, node, wp := t.TempDir(), t.TempDir(), t.TempDir()
	write(t, static, "index.html", "")
	write(t, wp, "wp-config.php", "<?php")
	write(t, node, "package.json", `{"scripts":{"build":"vite build"}}`)
	for _, tc := range []struct {
		link *site.Project
		src  string
		want Type
	}{
		{&site.Project{Type: "static"}, static, Static},
		{&site.Project{Type: "static"}, node, NodeJS}, // it became a Node.js app
		{&site.Project{Type: "nodejs"}, static, Static},
		{&site.Project{Type: "wordpress-site"}, wp, WordPressSite},
		{&site.Project{}, node, NodeJS},
		{&site.Project{Type: "static", Headless: true}, static, NodeJS},
	} {
		if got, err := LinkedType(tc.link, tc.src); err != nil || got != tc.want {
			t.Errorf("%+v in %s: %q, %v; want %q", tc.link, filepath.Base(tc.src), got, err, tc.want)
		}
	}
}
