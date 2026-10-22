package output

import (
	"bytes"
	"errors"
	"testing"

	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func newCmd(t *testing.T, flags map[string]string) *cobra.Command {
	t.Helper()
	cmd := &cobra.Command{Use: "list"}
	cmd.Flags().String("format", "", "")
	cmd.Flags().String("jq", "", "")
	for k, v := range flags {
		if err := cmd.Flags().Set(k, v); err != nil {
			t.Fatal(err)
		}
	}
	return cmd
}

func capture(t *testing.T, terminal bool) *bytes.Buffer {
	t.Helper()
	var buf bytes.Buffer
	oldOut, oldTerm := stdout, stdoutIsTerminal
	stdout, stdoutIsTerminal = &buf, func() bool { return terminal }
	t.Cleanup(func() { stdout, stdoutIsTerminal = oldOut, oldTerm })
	return &buf
}

func TestSelectedFormat(t *testing.T) {
	viper.Reset()
	defer viper.Reset()

	capture(t, true)
	if f, _ := SelectedFormat(newCmd(t, nil)); f != "table" {
		t.Errorf("terminal default = %q, want table", f)
	}
	capture(t, false)
	if f, _ := SelectedFormat(newCmd(t, nil)); f != "json" {
		t.Errorf("piped default = %q, want json", f)
	}
	viper.Set("format", "yaml")
	if f, _ := SelectedFormat(newCmd(t, nil)); f != "yaml" {
		t.Errorf("config default = %q, want yaml", f)
	}
	if f, _ := SelectedFormat(newCmd(t, map[string]string{"format": "tree"})); f != "tree" {
		t.Errorf("flag = %q, want tree", f)
	}
	var fe *utils.FlagError
	if _, err := SelectedFormat(newCmd(t, map[string]string{"format": "xml"})); !errors.As(err, &fe) {
		t.Errorf("invalid format: %v", err)
	}
}

func TestFormatYAMLKeepsOrderAndQuotesNumericStrings(t *testing.T) {
	out := capture(t, false)
	body := []byte(`{"zone":"a","id":8123712,"code":"0123","tags":["x","y"],"meta":{"ok":true}}`)
	if err := Format(newCmd(t, map[string]string{"format": "yaml"}), body, 200); err != nil {
		t.Fatal(err)
	}
	want := "zone: a\nid: 8123712\ncode: \"0123\"\ntags:\n  - x\n  - y\nmeta:\n  ok: true\n"
	if out.String() != want {
		t.Errorf("yaml:\n%s\nwant:\n%s", out.String(), want)
	}
}

func TestFormatJQ(t *testing.T) {
	out := capture(t, true)
	body := []byte(`{"data":[{"id":8123712,"hostname":"web"},{"id":2,"hostname":"db"}]}`)
	if err := Format(newCmd(t, map[string]string{"jq": ".data[] | .hostname"}), body, 200); err != nil {
		t.Fatal(err)
	}
	if err := Format(newCmd(t, map[string]string{"jq": "[.data[].id]"}), body, 200); err != nil {
		t.Fatal(err)
	}
	if got, want := out.String(), "web\ndb\n[8123712,2]\n"; got != want {
		t.Errorf("jq output %q, want %q", got, want)
	}
	var fe *utils.FlagError
	if err := Format(newCmd(t, map[string]string{"jq": ".data["}), body, 200); !errors.As(err, &fe) {
		t.Errorf("invalid jq: %v", err)
	}
}

func TestFormatReturnsAPIError(t *testing.T) {
	out := capture(t, false)
	err := Format(newCmd(t, nil), []byte(`{"message":"Not found"}`), 404)
	var apiErr *utils.APIError
	if !errors.As(err, &apiErr) || apiErr.Status != 404 || apiErr.Message != "Not found" {
		t.Errorf("err = %v", err)
	}
	if out.Len() != 0 {
		t.Errorf("error body printed to stdout: %q", out.String())
	}
}

func TestFormatJQOnEmptyBodyPrintsNothing(t *testing.T) {
	out := capture(t, false)
	if err := Format(newCmd(t, map[string]string{"jq": "."}), nil, 204); err != nil {
		t.Errorf("empty 204 with --jq: %v", err)
	}
	if out.Len() != 0 {
		t.Errorf("printed %q", out.String())
	}
}

func TestFormatJQDoesNotEscapeHTML(t *testing.T) {
	out := capture(t, false)
	if err := Format(newCmd(t, map[string]string{"jq": "{url}"}), []byte(`{"url":"https://example.com/?a=1&b=<2>"}`), 200); err != nil {
		t.Fatal(err)
	}
	if got, want := out.String(), `{"url":"https://example.com/?a=1&b=<2>"}`+"\n"; got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
