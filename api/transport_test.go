package api

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/hostinger/api-cli/utils"
)

func TestDryRunDoerPrintsCurlWithoutToken(t *testing.T) {
	var out bytes.Buffer
	old := dryRunOut
	dryRunOut = &out
	defer func() { dryRunOut = old }()

	req, _ := http.NewRequest(http.MethodPost, "https://api.example.com/v1/items?x=1", strings.NewReader(`{"name":"it's"}`))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer secret-token")

	resp, err := dryRunDoer{}.Do(req)
	if resp != nil || !errors.Is(err, utils.ErrDryRun) {
		t.Fatalf("resp %v err %v, want ErrDryRun", resp, err)
	}
	got := out.String()
	for _, want := range []string{
		`curl -X POST 'https://api.example.com/v1/items?x=1'`,
		`-H "Authorization: Bearer $` + utils.EnvPrefix + `API_TOKEN"`,
		`-H 'Content-Type: application/json'`,
		`-d '{"name":"it'\''s"}'`,
	} {
		if !strings.Contains(got, want) {
			t.Errorf("dry run output missing %q:\n%s", want, got)
		}
	}
	if strings.Contains(got, "secret-token") {
		t.Errorf("dry run leaked the token:\n%s", got)
	}
}

func TestVerboseDoerLogsRequestAndStatus(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusTeapot)
	}))
	defer srv.Close()

	var out bytes.Buffer
	old := verboseOut
	verboseOut = &out
	defer func() { verboseOut = old }()

	req, _ := http.NewRequest(http.MethodGet, srv.URL+"/v1/items", nil)
	req.Header.Set("Authorization", "Bearer secret-token")
	resp, err := (&verboseDoer{base: srv.Client()}).Do(req)
	if err != nil {
		t.Fatal(err)
	}
	resp.Body.Close()
	got := out.String()
	if !strings.Contains(got, "> GET "+srv.URL+"/v1/items") || !strings.Contains(got, "< 418 I'm a teapot") {
		t.Errorf("verbose output:\n%s", got)
	}
	if strings.Contains(got, "secret-token") {
		t.Errorf("verbose output leaked the token:\n%s", got)
	}
}

func TestNoTokenDoerReportsMissingCredentials(t *testing.T) {
	_, err := noTokenDoer{}.Do(&http.Request{})
	if !errors.Is(err, utils.ErrNoCredentials) || !strings.Contains(err.Error(), utils.EnvPrefix+"API_TOKEN") {
		t.Errorf("err = %v", err)
	}
}
