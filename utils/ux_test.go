package utils

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// newCmd returns a command with the global flags the runtime helpers read.
func newCmd() *cobra.Command {
	root := &cobra.Command{Use: "tool"}
	root.PersistentFlags().Bool("no-input", false, "")
	root.PersistentFlags().Bool("dry-run", false, "")
	cmd := &cobra.Command{Use: "get", RunE: func(*cobra.Command, []string) error { return nil }}
	root.AddCommand(cmd)
	return cmd
}

// fakeTTY makes the session look interactive and feeds input to prompts.
func fakeTTY(t *testing.T, input string) *bytes.Buffer {
	t.Helper()
	oldIn, oldErr, oldReader, oldOut := stdinIsTerminal, stderrIsTerminal, promptIn, promptOut
	var out bytes.Buffer
	stdinIsTerminal = func() bool { return true }
	stderrIsTerminal = func() bool { return true }
	promptIn = bufio.NewReader(strings.NewReader(input))
	promptOut = &out
	t.Cleanup(func() { stdinIsTerminal, stderrIsTerminal, promptIn, promptOut = oldIn, oldErr, oldReader, oldOut })
	return &out
}

func TestRegisterArgsValidates(t *testing.T) {
	cmd := newCmd()
	RegisterArgs(cmd, IntArg("vm-id"), UUIDArg("key"))

	if err := cmd.Args(cmd, []string{"12", "9b1deb4d-3b7d-4bad-9bdd-2b0d7b3dcb6d"}); err != nil {
		t.Fatalf("valid args rejected: %v", err)
	}
	for _, args := range [][]string{
		{"abc", "9b1deb4d-3b7d-4bad-9bdd-2b0d7b3dcb6d"},
		{"12", "not-a-uuid"},
		{"12", "9b1deb4d-3b7d-4bad-9bdd-2b0d7b3dcb6d", "extra"},
		{"12"}, // missing, and not interactive
	} {
		err := cmd.Args(cmd, args)
		var fe *FlagError
		if !errors.As(err, &fe) {
			t.Errorf("args %q: want FlagError, got %v", args, err)
		}
	}

	none := newCmd()
	RegisterArgs(none)
	if err := none.Args(none, []string{"x"}); err == nil || !strings.Contains(err.Error(), "takes no arguments") {
		t.Errorf("no-arg command accepted an argument: %v", err)
	}
}

func TestMissingArgsAllowedWhenInteractiveAndPrompted(t *testing.T) {
	out := fakeTTY(t, "abc\n2\n")
	cmd := newCmd()
	RegisterArgs(cmd, IntArg("vm-id"))
	cmd.ValidArgsFunction = func(*cobra.Command, []string, string) ([]cobra.Completion, cobra.ShellCompDirective) {
		return []cobra.Completion{"101\tweb", "102\tdb"}, cobra.ShellCompDirectiveNoFileComp
	}
	if err := cmd.Args(cmd, nil); err != nil {
		t.Fatalf("missing arg should be allowed interactively: %v", err)
	}
	if err := PromptMissing(cmd, nil); err != nil {
		t.Fatal(err)
	}
	// "abc" is rejected as not an integer, then "2" picks the second choice.
	if got := PromptedArgs(cmd, nil); len(got) != 1 || got[0] != "102" {
		t.Fatalf("prompted args = %q, want [102]", got)
	}
	if !strings.Contains(out.String(), "expected an integer") || !strings.Contains(out.String(), "2) 102  db") {
		t.Errorf("prompt output:\n%s", out.String())
	}
}

func TestPromptMissingRequiredFlag(t *testing.T) {
	fakeTTY(t, "302\n")
	cmd := newCmd()
	RegisterArgs(cmd)
	cmd.Flags().String("redirect-type", "", "Redirect type")
	_ = cmd.MarkFlagRequired("redirect-type")
	_ = cmd.RegisterFlagCompletionFunc("redirect-type", cobra.FixedCompletions([]cobra.Completion{"301", "302"}, cobra.ShellCompDirectiveNoFileComp))

	if err := PromptMissing(cmd, nil); err != nil {
		t.Fatal(err)
	}
	if v, _ := cmd.Flags().GetString("redirect-type"); v != "302" {
		t.Errorf("flag = %q, want 302", v)
	}
	if err := cmd.ValidateRequiredFlags(); err != nil {
		t.Errorf("required flag still missing after prompt: %v", err)
	}
}

func TestNoInputDisablesPrompts(t *testing.T) {
	fakeTTY(t, "")
	cmd := newCmd()
	RegisterArgs(cmd, IntArg("vm-id"))
	if err := cmd.Root().PersistentFlags().Set("no-input", "true"); err != nil {
		t.Fatal(err)
	}
	if Interactive(cmd) {
		t.Fatal("--no-input should make the session non-interactive")
	}
	if err := cmd.Args(cmd, nil); err == nil {
		t.Error("missing arg accepted with --no-input")
	}
}

func TestConfirm(t *testing.T) {
	// Non-interactive without --yes refuses.
	cmd := newCmd()
	cmd.Flags().BoolP("yes", "y", false, "")
	if err := Confirm(cmd, "Stop virtual machine", []string{"1"}); err == nil || !strings.Contains(err.Error(), "--yes") {
		t.Errorf("non-interactive confirm: %v", err)
	}
	// --yes skips the question.
	_ = cmd.Flags().Set("yes", "true")
	if err := Confirm(cmd, "Stop virtual machine", nil); err != nil {
		t.Errorf("--yes: %v", err)
	}

	for input, want := range map[string]error{"y\n": nil, "YES\n": nil, "n\n": ErrAborted, "\n": ErrAborted} {
		out := fakeTTY(t, input)
		c := newCmd()
		c.Flags().BoolP("yes", "y", false, "")
		if err := Confirm(c, "Stop virtual machine.", []string{"42"}); !errors.Is(err, want) && err != want {
			t.Errorf("input %q: got %v, want %v", input, err, want)
		}
		if !strings.Contains(out.String(), "Stop virtual machine 42? [y/N]") {
			t.Errorf("prompt = %q", out.String())
		}
	}
}

func TestJSONValueSources(t *testing.T) {
	if v, err := JSONValue(`{"a":1}`, "setup"); err != nil || v.(map[string]any)["a"] != float64(1) {
		t.Errorf("inline: %v %v", v, err)
	}
	path := filepath.Join(t.TempDir(), "setup.json")
	if err := os.WriteFile(path, []byte(`["x"]`), 0o600); err != nil {
		t.Fatal(err)
	}
	if v, err := JSONValue("@"+path, "setup"); err != nil || v.([]any)[0] != "x" {
		t.Errorf("@file: %v %v", v, err)
	}
	old := stdin
	stdin = strings.NewReader(`true`)
	defer func() { stdin = old }()
	if v, err := JSONValue("-", "setup"); err != nil || v != true {
		t.Errorf("stdin: %v %v", v, err)
	}
	var fe *FlagError
	if _, err := JSONValue("{bad", "setup"); !errors.As(err, &fe) || !strings.Contains(err.Error(), "--setup") {
		t.Errorf("bad json: %v", err)
	}
	if _, err := JSONValue("@/does/not/exist.json", "setup"); !errors.As(err, &fe) {
		t.Errorf("missing file: %v", err)
	}
}

func TestEnumAndFormatChecks(t *testing.T) {
	cmd := newCmd()
	cmd.Flags().String("type", "", "")
	cmd.Flags().String("from", "", "")
	if err := EnumCheck(cmd, "type", []string{"a"}); err != nil {
		t.Errorf("unchanged flag checked: %v", err)
	}
	_ = cmd.Flags().Set("type", "b")
	if err := EnumCheck(cmd, "type", []string{"a", "c"}); err == nil || !strings.Contains(err.Error(), "one of: a, c") {
		t.Errorf("enum: %v", err)
	}
	for format, value := range map[string]string{"date-time": "2025-01-31", "date": "31/01/2025", "uuid": "x", "email": "nobody"} {
		_ = cmd.Flags().Set("from", value)
		if err := FormatCheck(cmd, "from", format); err == nil {
			t.Errorf("%s accepted %q", format, value)
		}
	}
	_ = cmd.Flags().Set("from", "2025-01-31T10:00:00Z")
	if err := FormatCheck(cmd, "from", "date-time"); err != nil {
		t.Errorf("valid date-time: %v", err)
	}
}

func TestAllPages(t *testing.T) {
	items := []int{1, 2, 3, 4, 5, 6, 7}
	var pages []int
	page := 0
	call := func(ctx context.Context) ([]byte, int, error) {
		pages = append(pages, page)
		start, end := (page-1)*3, min(page*3, len(items))
		b, _ := json.Marshal(map[string]any{
			"data": items[start:end],
			"meta": map[string]int{"current_page": page, "per_page": 3, "total": len(items)},
		})
		return b, 200, nil
	}
	body, status, err := AllPages(call, func(p int) { page = p })(context.Background())
	if err != nil || status != 200 {
		t.Fatalf("status %d err %v", status, err)
	}
	var got struct {
		Data []int    `json:"data"`
		Meta pageMeta `json:"meta"`
	}
	if err := json.Unmarshal(body, &got); err != nil {
		t.Fatal(err)
	}
	if fmt.Sprint(got.Data) != fmt.Sprint(items) || got.Meta.Total != 7 || fmt.Sprint(pages) != "[1 2 3]" {
		t.Errorf("data %v meta %+v pages %v", got.Data, got.Meta, pages)
	}

	// Non-envelope responses and errors pass through unchanged.
	plain := func(ctx context.Context) ([]byte, int, error) { return []byte(`[1]`), 200, nil }
	if b, _, _ := AllPages(plain, func(int) {})(context.Background()); string(b) != `[1]` {
		t.Errorf("plain body changed: %s", b)
	}
	failing := func(ctx context.Context) ([]byte, int, error) { return []byte(`{"message":"x"}`), 500, nil }
	if _, st, _ := AllPages(failing, func(int) {})(context.Background()); st != 500 {
		t.Errorf("error status lost: %d", st)
	}
}

func TestWait(t *testing.T) {
	oldFirst, oldMax, oldOut := waitFirstInterval, waitMaxInterval, promptOut
	waitFirstInterval, waitMaxInterval = time.Millisecond, time.Millisecond
	var out bytes.Buffer
	promptOut = &out
	defer func() { waitFirstInterval, waitMaxInterval, promptOut = oldFirst, oldMax, oldOut }()

	spec := WaitSpec{IDField: "id", StateField: "state", Done: []string{"success"}, Failed: []string{"error"}}
	run := func(states ...string) ([]byte, error) {
		cmd := newCmd()
		RegisterWaitFlags(cmd)
		cmd.SetContext(context.Background())
		polls := 0
		body, _, err := Wait(cmd, []byte(`{"id": 8123712, "state": "created"}`), spec, func(ctx context.Context, id string) ([]byte, int, error) {
			if id != "8123712" {
				t.Errorf("polled id %q", id)
			}
			s := states[min(polls, len(states)-1)]
			polls++
			return []byte(`{"id": 8123712, "state": "` + s + `"}`), 200, nil
		})
		return body, err
	}
	if body, err := run("sent", "success"); err != nil || !strings.Contains(string(body), "success") {
		t.Errorf("success: %s %v", body, err)
	}
	if _, err := run("error"); !errors.Is(err, ErrActionFailed) {
		t.Errorf("failure: %v", err)
	}
	if !strings.Contains(out.String(), "Action 8123712: sent") {
		t.Errorf("progress output:\n%s", out.String())
	}

	cmd := newCmd()
	RegisterWaitFlags(cmd)
	cmd.SetContext(context.Background())
	_ = cmd.Flags().Set("wait-timeout", "5ms")
	_, _, err := Wait(cmd, []byte(`{"id": 1, "state": "sent"}`), spec, func(ctx context.Context, id string) ([]byte, int, error) {
		return []byte(`{"id": 1, "state": "sent"}`), 200, nil
	})
	if err == nil || !strings.Contains(err.Error(), "gave up") {
		t.Errorf("timeout: %v", err)
	}
}

func TestCompletionItems(t *testing.T) {
	body := []byte(`{"data": [{"id": 8123712, "hostname": "web.example.com"}, {"id": 2, "name": "db"}, {"other": 1}]}`)
	got := completionItems(body, "id")
	want := []completionItem{{"8123712", "web.example.com"}, {"2", "db"}}
	if fmt.Sprint(got) != fmt.Sprint(want) {
		t.Errorf("items = %v, want %v", got, want)
	}
	if got := completionItems([]byte(`[{"domain": "example.com", "status": "active"}]`), "domain"); len(got) != 1 || got[0].label != "active" {
		t.Errorf("array items = %v", got)
	}
}

func TestCompleteArgsUsesSourceAndCache(t *testing.T) {
	t.Setenv("HOME", t.TempDir())
	t.Setenv("XDG_CACHE_HOME", t.TempDir())
	viper.Set("api_token", "cache-test") // only static-token results are cached
	defer viper.Set("api_token", "")
	calls := 0
	src := &ArgSource{Value: "id", Fetch: func(ctx context.Context, args []string) ([]byte, int, error) {
		calls++
		return []byte(`[{"id": 11, "name": "a"}, {"id": 22, "name": "b"}]`), 200, nil
	}}
	cmd := newCmd()
	fn := CompleteArgs(src, nil)

	comps, dir := fn(cmd, nil, "2")
	if len(comps) != 1 || comps[0] != "22\tb" || dir != cobra.ShellCompDirectiveNoFileComp {
		t.Errorf("completions %q directive %v", comps, dir)
	}
	fn(cmd, nil, "")
	if calls != 1 {
		t.Errorf("source called %d times, want 1 (cached)", calls)
	}
	if comps, dir := fn(cmd, []string{"11"}, ""); comps != nil || dir != cobra.ShellCompDirectiveNoFileComp {
		t.Errorf("second position should have no completions and no files: %q %v", comps, dir)
	}
}

func TestHandleError(t *testing.T) {
	cmd := newCmd()
	cmd.Flags().String("redirect-url", "", "")
	cmd.Flags().String("setup", "", "")
	cases := []struct {
		err  error
		code int
		want []string
	}{
		{nil, ExitOK, nil},
		{ErrDryRun, ExitOK, nil},
		{ErrAborted, ExitError, []string{"Aborted."}},
		{FlagErrorf("bad --x"), ExitUsage, []string{"Error: bad --x", "Run 'tool get --help' for usage."}},
		{errors.New(`unknown flag: --nope`), ExitUsage, []string{"--help"}},
		{errors.New("dial tcp: refused"), ExitError, []string{"Error: dial tcp: refused"}},
		{fmt.Errorf("%w: set a token", ErrNoCredentials), ExitAuth, []string{"no credentials configured: set a token"}},
		{NewAPIError(404, []byte(`{"message":"Not found","correlation_id":"c-1"}`)), ExitNotFound, []string{"Error: Not found (HTTP 404)", "Correlation ID: c-1"}},
		{NewAPIError(401, []byte(`{"message":"Unauthenticated"}`)), ExitAuth, []string{"Hint:"}},
		{NewAPIError(422, []byte(`{"message":"Invalid","errors":{"redirect_url":["must be a URL"],"setup.hostname":["is taken"],"other":["x"]}}`)), ExitValidation,
			[]string{"  --redirect-url: must be a URL", "  --setup (hostname): is taken", "  other: x"}},
		{NewAPIError(500, []byte(`oops`)), ExitError, []string{"Error: oops (HTTP 500)"}},
	}
	for _, c := range cases {
		var w bytes.Buffer
		if got := HandleError(&w, cmd, c.err); got != c.code {
			t.Errorf("%v: exit %d, want %d", c.err, got, c.code)
		}
		for _, want := range c.want {
			if !strings.Contains(w.String(), want) {
				t.Errorf("%v: output missing %q:\n%s", c.err, want, w.String())
			}
		}
	}
}

func TestGroupRunESuggests(t *testing.T) {
	group := &cobra.Command{Use: "vm", RunE: GroupRunE}
	group.AddCommand(&cobra.Command{Use: "start", Run: func(*cobra.Command, []string) {}})
	err := GroupRunE(group, []string{"stat"})
	var fe *FlagError
	if !errors.As(err, &fe) || !strings.Contains(err.Error(), "Did you mean this?\n\tstart") {
		t.Errorf("err = %v", err)
	}
}

func TestUsageTemplateSplitsRequiredFlags(t *testing.T) {
	cmd := newCmd()
	cmd.Root().SetUsageTemplate(UsageTemplate)
	cmd.Flags().String("domain", "", "Domain name")
	cmd.Flags().String("note", "", "Optional note")
	_ = cmd.MarkFlagRequired("domain")
	var out bytes.Buffer
	cmd.SetOut(&out)
	if err := cmd.Usage(); err != nil {
		t.Fatal(err)
	}
	u := out.String()
	req, opt := strings.Index(u, "Required Flags:"), strings.Index(u, "\nFlags:")
	if req < 0 || opt < 0 || !(req < strings.Index(u, "--domain") && strings.Index(u, "--domain") < opt && opt < strings.Index(u, "--note")) {
		t.Errorf("usage:\n%s", u)
	}
}

func TestConfirmSkippedForDryRunFromConfig(t *testing.T) {
	viper.Set("dry_run", true)
	defer viper.Set("dry_run", false)
	cmd := newCmd()
	cmd.Flags().BoolP("yes", "y", false, "")
	if err := Confirm(cmd, "Stop virtual machine", []string{"1"}); err != nil {
		t.Errorf("dry_run from config/env should skip confirmation: %v", err)
	}
}

func TestAskPrefersListedValueOverIndex(t *testing.T) {
	for _, c := range []struct {
		choices []string
		input   string
		want    string
	}{
		{[]string{"12", "24", "3"}, "24\n", "24"}, // a listed value
		{[]string{"12", "24", "3"}, "2\n", "24"},  // not listed: menu index
		{[]string{"2", "1"}, "1\n", "1"},          // listed value wins over index 1 ("2")
		{[]string{"101\tweb", "102\tdb"}, "102\n", "102"},
	} {
		fakeTTY(t, c.input)
		got, err := ask("<id>", "", c.choices)
		if err != nil || got != c.want {
			t.Errorf("choices %q input %q: got %q, %v; want %q", c.choices, c.input, got, err, c.want)
		}
	}
}

func TestWaitStoppedEarlyKeepsLastState(t *testing.T) {
	oldFirst, oldMax, oldOut := waitFirstInterval, waitMaxInterval, promptOut
	waitFirstInterval, waitMaxInterval, promptOut = time.Millisecond, time.Millisecond, io.Discard
	defer func() { waitFirstInterval, waitMaxInterval, promptOut = oldFirst, oldMax, oldOut }()
	spec := WaitSpec{IDField: "id", StateField: "state", Done: []string{"success"}}
	start := []byte(`{"id": 1, "state": "initiated"}`)

	// Interrupted (Ctrl-C) before the first poll.
	cmd := newCmd()
	RegisterWaitFlags(cmd)
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	cmd.SetContext(ctx)
	body, _, err := Wait(cmd, start, spec, func(ctx context.Context, id string) ([]byte, int, error) {
		t.Error("polled after cancel")
		return nil, 0, nil
	})
	if err == nil || !strings.Contains(err.Error(), `interrupted; last state "initiated"`) || string(body) != string(start) {
		t.Errorf("cancel: body %s err %v", body, err)
	}

	// Deadline hits while a poll is in flight.
	cmd = newCmd()
	RegisterWaitFlags(cmd)
	cmd.SetContext(context.Background())
	_ = cmd.Flags().Set("wait-timeout", "20ms")
	body, _, err = Wait(cmd, start, spec, func(ctx context.Context, id string) ([]byte, int, error) {
		<-ctx.Done()
		return nil, 0, ctx.Err()
	})
	if err == nil || !strings.Contains(err.Error(), `gave up after 20ms; last state "initiated"`) || string(body) != string(start) {
		t.Errorf("deadline: body %s err %v", body, err)
	}
}
