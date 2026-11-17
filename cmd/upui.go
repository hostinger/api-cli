package cmd

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/brand"
	"github.com/hostinger/api-cli/deploy/hosting"
	"github.com/hostinger/api-cli/site"
	"github.com/hostinger/api-cli/utils"
	"github.com/pterm/pterm"
)

// termUI writes progress to stderr so stdout stays clean for JSON. Everything
// it shows goes through utils.SafeText: file names from the deployed folder,
// build settings from its package.json and API errors all reach it.
type termUI struct {
	interactive bool
	brand       *brandView // nil: pterm's plain look
}

func newTermUI(jsonOutput bool) *termUI {
	pterm.SetDefaultOutput(os.Stderr)
	u := &termUI{interactive: site.Interactive() && !jsonOutput}
	if u.interactive && brand.Supported(os.Stderr) {
		u.brand = newBrandView(os.Stderr)
		api.SetDiagnostics(u.brand.logWriter())
	}
	return u
}

// Close must come before printing anything else (JSON); a second call does nothing.
func (u *termUI) Close() {
	if u.brand != nil {
		u.brand.stop()
		api.SetDiagnostics(nil)
	}
}

func (u *termUI) Finish(out *hosting.Outcome) {
	if u.brand != nil {
		u.brand.finish(out)
		return
	}
	switch {
	case out.Note != "":
		pterm.Info.Printfln("Not finished yet: %s", utils.SafeText(out.Note))
	case out.Reachable:
		pterm.Success.Printfln("Live at %s", utils.SafeText(out.URL))
	default:
		pterm.Success.Printfln("Deployed to %s", utils.SafeText(out.Target))
		pterm.Info.Println(utils.SafeText(out.ReachableNote))
	}
	if out.NewLink {
		pterm.Info.Printfln("Linked this directory to %s; `hostinger up` now redeploys there. %s", utils.SafeText(out.Target), hosting.RelinkHint)
	}
}

func (u *termUI) Fail(err error) {
	if u.brand != nil {
		u.brand.fail(errors.Is(err, context.Canceled))
	}
}

func (u *termUI) prompt() func() {
	if u.brand == nil {
		return func() {}
	}
	u.brand.hide()
	return u.brand.show
}

func (u *termUI) question(title string) string {
	title = utils.SafeText(title)
	if u.brand != nil {
		return "  " + title
	}
	return title
}

var errNotInteractive = errors.New("a prompt was needed but the session is not interactive")

func (u *termUI) Interactive() bool { return u.interactive }

func (u *termUI) Select(title string, options []string, def int) (int, error) {
	if !u.interactive {
		return 0, errNotInteractive
	}
	defer u.prompt()()
	options = safeAll(options)
	sel := pterm.DefaultInteractiveSelect.WithOptions(options).WithFilter(len(options) > 8).WithMaxHeight(site.SelectHeight(len(options)))
	if def >= 0 && def < len(options) {
		sel = sel.WithDefaultOption(options[def])
	}
	picked, err := sel.Show(u.question(title))
	if err != nil {
		return 0, err
	}
	for i, o := range options {
		if o == picked {
			return i, nil
		}
	}
	return 0, fmt.Errorf("nothing selected")
}

func (u *termUI) Input(title, def string) (string, error) {
	if !u.interactive {
		return "", errNotInteractive
	}
	defer u.prompt()()
	v, err := pterm.DefaultInteractiveTextInput.WithDefaultValue(utils.SafeText(def)).Show(u.question(title))
	if err != nil {
		return "", err
	}
	if v == "" {
		return def, nil
	}
	return v, nil
}

func (u *termUI) Confirm(title string, def bool) (bool, error) {
	if !u.interactive {
		return false, errNotInteractive
	}
	defer u.prompt()()
	return pterm.DefaultInteractiveConfirm.WithDefaultValue(def).Show(u.question(title))
}

func (u *termUI) Info(msg string) {
	if u.brand != nil {
		u.brand.info(msg)
		return
	}
	pterm.Info.Println(utils.SafeText(msg))
}

func (u *termUI) Start(msg string) hosting.Step {
	msg = utils.SafeText(msg)
	if !u.interactive {
		fmt.Fprintln(os.Stderr, msg)
		return &plainStep{}
	}
	if u.brand != nil {
		return u.brand.start(msg)
	}
	sp, err := pterm.DefaultSpinner.WithWriter(os.Stderr).WithRemoveWhenDone(true).Start(spinnerText(msg))
	if err != nil {
		return &plainStep{}
	}
	return &spinnerStep{sp: sp}
}

// spinnerText fits a spinner line in the terminal, leaving room for the
// spinner and its timer; read on every update, so a resize is followed.
func spinnerText(msg string) string {
	if w := utils.StderrWidth(); w > 0 {
		return utils.FitWidth(msg, max(w-12, 10))
	}
	return msg
}

type spinnerStep struct{ sp *pterm.SpinnerPrinter }

func (s *spinnerStep) Update(msg string) { s.sp.UpdateText(spinnerText(utils.SafeText(msg))) }

// Done and Fail stop the spinner first: pterm's Success and Fail stop it after
// printing, so a late frame can overwrite the result line.
func (s *spinnerStep) Done(msg string) {
	_ = s.sp.Stop()
	pterm.Success.WithWriter(s.sp.Writer).Println(utils.SafeText(msg))
}

func (s *spinnerStep) Fail(msg string) {
	_ = s.sp.Stop()
	pterm.Error.WithWriter(s.sp.Writer).Println(utils.SafeText(msg))
}

// plainStep prints one line per state, for logs and pipes.
type plainStep struct{}

func (plainStep) Update(string)   {}
func (plainStep) Done(msg string) { fmt.Fprintln(os.Stderr, utils.SafeText(msg)) }
func (plainStep) Fail(msg string) { fmt.Fprintln(os.Stderr, utils.SafeText(msg)) }

func safeAll(ss []string) []string {
	out := make([]string, len(ss))
	for i, s := range ss {
		out[i] = utils.SafeText(s)
	}
	return out
}
