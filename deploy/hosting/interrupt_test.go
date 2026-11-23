package hosting

import (
	"context"
	"net/http"
	"net/http/httptest"
	"slices"
	"testing"
)

// failUI records what failed steps show.
type failUI struct {
	*fakeUI
	fails []string
}

func (u *failUI) Start(string) Step { return failStep{u} }

type failStep struct{ u *failUI }

func (failStep) Update(string)     {}
func (failStep) Done(string)       {}
func (s failStep) Fail(msg string) { s.u.fails = append(s.u.fails, msg) }

func TestUploadStepSaysInterruptedOnCtrlC(t *testing.T) {
	fastPolling(t)

	// A failed upload still says so.
	api := newFake(t)
	api.upload.Close()
	ui := &failUI{fakeUI: &fakeUI{t: t}}
	if _, err := Run(context.Background(), api, ui, Options{Path: staticProject(t), Domain: "example.com", Yes: true}); FailureReason(err) != ReasonUploadFailed {
		t.Fatalf("closed upload server: %v", err)
	}
	if want := []string{"Upload failed"}; !slices.Equal(ui.fails, want) {
		t.Errorf("failed upload shows %q, want %q", ui.fails, want)
	}

	// Ctrl+C during the upload is not an upload error.
	api = newFake(t)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	srv := httptest.NewServer(http.HandlerFunc(func(_ http.ResponseWriter, r *http.Request) {
		cancel()
		<-r.Context().Done()
	}))
	t.Cleanup(srv.Close)
	api.upload = srv
	ui = &failUI{fakeUI: &fakeUI{t: t}}
	if _, err := Run(ctx, api, ui, Options{Path: staticProject(t), Domain: "example.com", Yes: true}); FailureReason(err) != ReasonCancelled {
		t.Fatalf("cancelled upload: %v (reason %q)", err, FailureReason(err))
	}
	if want := []string{"Interrupted"}; !slices.Equal(ui.fails, want) {
		t.Errorf("cancelled upload shows %q, want %q", ui.fails, want)
	}
}
