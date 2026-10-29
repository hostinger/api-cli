package cmd

import (
	"errors"
	"fmt"
	"os"

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
}

func newTermUI(jsonOutput bool) *termUI {
	pterm.SetDefaultOutput(os.Stderr)
	return &termUI{interactive: site.Interactive() && !jsonOutput}
}

var errNotInteractive = errors.New("a prompt was needed but the session is not interactive")

func (u *termUI) Interactive() bool { return u.interactive }

func (u *termUI) Select(title string, options []string, def int) (int, error) {
	if !u.interactive {
		return 0, errNotInteractive
	}
	options = safeAll(options)
	sel := pterm.DefaultInteractiveSelect.WithOptions(options).WithFilter(len(options) > 8).WithMaxHeight(site.SelectHeight(len(options)))
	if def >= 0 && def < len(options) {
		sel = sel.WithDefaultOption(options[def])
	}
	picked, err := sel.Show(utils.SafeText(title))
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
	v, err := pterm.DefaultInteractiveTextInput.WithDefaultValue(utils.SafeText(def)).Show(utils.SafeText(title))
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
	return pterm.DefaultInteractiveConfirm.WithDefaultValue(def).Show(utils.SafeText(title))
}

func (u *termUI) Info(msg string) {
	pterm.Info.Println(utils.SafeText(msg))
}

func (u *termUI) Start(msg string) hosting.Step {
	msg = utils.SafeText(msg)
	if !u.interactive {
		fmt.Fprintln(os.Stderr, msg)
		return &plainStep{}
	}
	sp, err := pterm.DefaultSpinner.WithWriter(os.Stderr).Start(spinnerText(msg))
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
func (s *spinnerStep) Done(msg string)   { s.sp.Success(utils.SafeText(msg)) }
func (s *spinnerStep) Fail(msg string)   { s.sp.Fail(utils.SafeText(msg)) }

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
