package cmd

import (
	"context"
	"os"
	"strings"

	"github.com/hostinger/api-cli/site"
	"github.com/hostinger/api-cli/utils"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"golang.org/x/term"
)

// A flag typed without its value is asked for in a terminal and the command
// re-run with the answer: cobra cannot resume a parse, but nothing has run yet.
func init() {
	if _, noColor := os.LookupEnv("NO_COLOR"); noColor || !term.IsTerminal(int(os.Stderr.Fd())) {
		pterm.DisableColor()
		pterm.DisableStyling()
	}
	// pterm's light text colours vanish on light backgrounds; keep colour on badges only.
	plain := &pterm.Style{}
	pterm.Info.MessageStyle = plain
	pterm.Success.MessageStyle = plain
	pterm.Warning.MessageStyle = plain
	pterm.Error.MessageStyle = plain
	pterm.DefaultSpinner.MessageStyle = plain
	pterm.DefaultInteractiveConfirm.TextStyle = plain
	pterm.DefaultInteractiveSelect.TextStyle = plain
	pterm.DefaultInteractiveTextInput.TextStyle = plain
	RootCmd.SetFlagErrorFunc(flagError)
}

// AskFlagValue returns ok=false when it cannot ask. Tests replace it.
var AskFlagValue = askFlagValue

// jsonOutput reports whether the command line asks for JSON, or output is piped
// with no format set: then nothing is asked, as for `up`. The flags failed to
// parse, so the arguments are read directly.
func jsonOutput(args []string) bool {
	format := ""
	for i, a := range args {
		switch {
		case a == "--format" || a == "-o":
			if i+1 < len(args) {
				format = args[i+1]
			}
		case strings.HasPrefix(a, "--format="):
			format = strings.TrimPrefix(a, "--format=")
		case strings.HasPrefix(a, "-o") && len(a) > 2:
			format = strings.TrimPrefix(strings.TrimPrefix(a, "-o"), "=")
		}
	}
	if format == "" {
		format = viper.GetString("format")
	}
	return format == "json" || (format == "" && !term.IsTerminal(int(os.Stdout.Fd())))
}

func flagError(c *cobra.Command, err error) error {
	if name, value, ok := badValueFlag(err.Error()); ok {
		if f := c.Flags().Lookup(name); f != nil {
			return utils.FlagErrorf("--%s: %q is not a valid %s\n  --%s %s   %s", name, value, f.Value.Type(), f.Name, f.Value.Type(), f.Usage)
		}
	}
	name, ok := missingValueFlag(err.Error())
	if !ok {
		return err
	}
	f := c.Flags().Lookup(name)
	if f == nil {
		f = c.Flags().ShorthandLookup(name)
	}
	if f == nil {
		return err
	}
	// Flag errors come before cobra's OnInitialize: load the config now, so
	// no_input and the API token are seen.
	initConfig()
	if site.Interactive() && !jsonOutput(os.Args[1:]) {
		value, ok, aerr := AskFlagValue(c, f)
		if aerr != nil {
			return aerr
		}
		if ok {
			// pflag only reports a missing value for the last argument. The
			// second parse starts from clean flags, or list flags given
			// before it would be appended twice.
			resetFlags(c)
			root := c.Root()
			root.SetArgs(append(os.Args[1:], value))
			return root.Execute()
		}
	}
	return utils.FlagErrorf("%s\n  --%s %s   %s", err, f.Name, f.Value.Type(), f.Usage)
}

// resetFlags puts c's flags, and the persistent ones it inherits, back to
// their defaults.
func resetFlags(c *cobra.Command) {
	reset := func(f *pflag.Flag) {
		if f.Changed {
			if sv, ok := f.Value.(pflag.SliceValue); ok {
				_ = sv.Replace(nil)
			} else {
				_ = f.Value.Set(f.DefValue)
			}
			f.Changed = false
		}
	}
	c.Flags().VisitAll(reset)
	c.InheritedFlags().VisitAll(reset)
}

// badValueFlag reads the flag and value out of pflag's
// `invalid argument "abc" for "--order-id" flag: strconv.ParseInt: ...`.
func badValueFlag(msg string) (name, value string, ok bool) {
	rest, found := strings.CutPrefix(msg, "invalid argument \"")
	if !found {
		return "", "", false
	}
	value, rest, found = strings.Cut(rest, "\" for \"")
	if !found {
		return "", "", false
	}
	name, _, found = strings.Cut(rest, "\" flag")
	if !found {
		return "", "", false
	}
	return strings.TrimLeft(name, "-"), value, true
}

// missingValueFlag reads the flag name out of pflag's two phrasings:
// "flag needs an argument: --type" and "flag needs an argument: 't' in -t".
func missingValueFlag(msg string) (string, bool) {
	if name, ok := strings.CutPrefix(msg, "flag needs an argument: --"); ok {
		return name, true
	}
	if strings.HasPrefix(msg, "flag needs an argument:") {
		if i := strings.LastIndex(msg, " in -"); i > 0 {
			return strings.TrimPrefix(msg[i+len(" in -"):], "-"), true
		}
	}
	return "", false
}

// flagChoices are value lists commands register for their own flags, keyed
// by the flag so another command's flag of the same name is not affected.
var flagChoices = map[*pflag.Flag][]string{}

func choices(c *cobra.Command, f *pflag.Flag) []string {
	if values, ok := flagChoices[f]; ok {
		return values
	}
	switch f.Name {
	case "format":
		return []string{"table", "json", "yaml", "tree"}
	}
	// Generated flags spell their values as "(one of: a, b, c)".
	if i := strings.Index(f.Usage, "(one of: "); i >= 0 {
		rest := f.Usage[i+len("(one of: "):]
		if j := strings.Index(rest, ")"); j > 0 {
			return strings.Split(strings.ReplaceAll(rest[:j], " ", ""), ",")
		}
	}
	return nil
}

// Question is a short title for common flags, else the trimmed flag help.
func Question(f *pflag.Flag) string {
	switch f.Name {
	case "type":
		return "Deploy as"
	case "source":
		return "Deploy from"
	case "app-type":
		return "Framework"
	case "node-version":
		return "Node.js version"
	case "package-manager":
		return "Package manager"
	case "format":
		return "Output format"
	}
	usage := f.Usage
	for {
		i := strings.Index(usage, " (")
		if i < 0 {
			break
		}
		j := strings.Index(usage[i:], ")")
		if j < 0 {
			break
		}
		usage = usage[:i] + usage[i+j+1:]
	}
	usage = strings.TrimSuffix(strings.TrimSpace(usage), ".")
	if usage == "" {
		return "--" + f.Name
	}
	return "--" + f.Name + ": " + usage
}

func askFlagValue(c *cobra.Command, f *pflag.Flag) (string, bool, error) {
	pterm.SetDefaultOutput(os.Stderr)
	ctx := context.Background()
	// The pickers are for the root's --domain/--order-id and up's --order-id; a
	// command's own flag of the same name (wordpress install --domain) is not.
	name := f.Name
	if f != c.Root().PersistentFlags().Lookup(name) && c.Name() != "up" {
		name = ""
	}
	switch name {
	case "domain":
		w, err := site.ChooseWebsite(ctx, "")
		if err != nil {
			return "", false, err
		}
		return w.Domain, true, nil
	case "order-id":
		group := "hosting"
		if parts := strings.Fields(c.CommandPath()); len(parts) > 1 && parts[1] != "up" && parts[1] != "link" {
			group = parts[1]
		}
		o, err := site.ChooseOrder(ctx, group)
		if err != nil {
			return "", false, err
		}
		return o.ID, true, nil
	}
	title := Question(f)
	if opts := choices(c, f); len(opts) > 0 {
		sel := pterm.DefaultInteractiveSelect.WithOptions(opts).WithFilter(len(opts) > 8).WithMaxHeight(site.SelectHeight(len(opts)))
		if f.DefValue != "" && f.DefValue != "0" {
			sel = sel.WithDefaultOption(f.DefValue)
		}
		v, err := sel.Show(title)
		return v, err == nil, err
	}
	v, err := pterm.DefaultInteractiveTextInput.WithDefaultValue(f.DefValue).Show(title)
	if err != nil {
		return "", false, err
	}
	return v, v != "", nil
}
