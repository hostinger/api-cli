package site

import (
	"os"
	"strings"

	"github.com/pterm/pterm"
	"github.com/spf13/viper"
	"golang.org/x/term"
)

// Interactive also checks CI and HOSTINGER_NONINTERACTIVE: agents and CI
// runners often have a pseudo-terminal. Tests replace it.
var Interactive = func() bool {
	if viper.GetBool("no_input") || os.Getenv("HOSTINGER_NONINTERACTIVE") != "" {
		return false
	}
	if ci := strings.ToLower(os.Getenv("CI")); ci != "" && ci != "false" && ci != "0" {
		return false
	}
	return term.IsTerminal(int(os.Stdin.Fd())) && term.IsTerminal(int(os.Stderr.Fd()))
}

// pick shows a select on stderr, so stdout stays clean for the command's
// output (--format json | jq); pterm writes to stdout otherwise.
func pick(sel *pterm.InteractiveSelectPrinter, title string) (string, error) {
	pterm.SetDefaultOutput(os.Stderr)
	defer pterm.SetDefaultOutput(os.Stdout)
	return sel.Show(title)
}

// SelectHeight fits the list below the question, at least five rows.
func SelectHeight(options int) int {
	_, rows, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil || rows <= 0 {
		rows = 24
	}
	fit := rows - 4
	if fit < 5 {
		fit = 5
	}
	if options < fit {
		return max(options, 1)
	}
	return fit
}
