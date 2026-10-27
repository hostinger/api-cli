package utils

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"golang.org/x/term"
)

// Terminal detection and the prompt streams are package vars so tests can
// simulate an interactive session.
var (
	stdinIsTerminal            = func() bool { return term.IsTerminal(int(os.Stdin.Fd())) }
	stderrIsTerminal           = func() bool { return term.IsTerminal(int(os.Stderr.Fd())) }
	promptIn                   = bufio.NewReader(os.Stdin)
	promptOut        io.Writer = stderrWriter{}
)

// stderrWriter writes to os.Stderr as it is at write time.
type stderrWriter struct{}

func (stderrWriter) Write(p []byte) (int, error) { return os.Stderr.Write(p) }

// maxChoices caps how many options a numbered prompt lists.
const maxChoices = 30

// Interactive reports whether the CLI may prompt: stdin and stderr are
// terminals and neither --no-input nor the no_input config/env key is set.
func Interactive(cmd *cobra.Command) bool {
	if flagBool(cmd, "no-input") || viper.GetBool("no_input") || os.Getenv(EnvPrefix+"NONINTERACTIVE") != "" {
		return false
	}
	// Agents and CI runners often have a pseudo-terminal.
	if ci := strings.ToLower(os.Getenv("CI")); ci != "" && ci != "false" && ci != "0" {
		return false
	}
	return stdinIsTerminal() && stderrIsTerminal()
}

// PromptMissing asks for any missing positional arguments and required flags
// when the session is interactive. It runs from the root's PersistentPreRunE:
// after argument validation (which lets missing arguments through only when
// interactive) and before cobra's required-flag check.
func PromptMissing(cmd *cobra.Command, args []string) error {
	setPromptedArgs(cmd, nil)
	if !Interactive(cmd) {
		return nil
	}
	specs, scope := ArgSpecs(cmd), ArgScope(cmd)
	if len(args) < len(specs)-scope {
		full := append([]string{}, args...)
		for i := scope + len(args); i < len(specs); i++ {
			v, err := promptArg(cmd, specs[i], full)
			if err != nil {
				return err
			}
			full = append(full, v)
		}
		setPromptedArgs(cmd, full)
	}

	var missing []*pflag.Flag
	cmd.LocalFlags().VisitAll(func(f *pflag.Flag) {
		if IsRequired(f) && !f.Changed {
			missing = append(missing, f)
		}
	})
	for _, f := range missing {
		v, err := promptFlag(cmd, f)
		if err != nil {
			return err
		}
		if err := cmd.Flags().Set(f.Name, v); err != nil {
			return FlagErrorf("invalid value for --%s: %v", f.Name, err)
		}
	}
	return nil
}

// IsRequired reports whether a flag was marked required with MarkFlagRequired.
func IsRequired(f *pflag.Flag) bool {
	ann := f.Annotations[cobra.BashCompOneRequiredFlag]
	return len(ann) > 0 && ann[0] == "true"
}

func promptArg(cmd *cobra.Command, spec ArgSpec, before []string) (string, error) {
	var choices []string
	if cmd.ValidArgsFunction != nil {
		comps, _ := cmd.ValidArgsFunction(cmd, before, "")
		for _, c := range comps {
			choices = append(choices, string(c))
		}
	}
	for {
		v, err := ask("<"+spec.Name+">", "", choices)
		if err != nil {
			return "", err
		}
		if err := checkArg(spec, v); err != nil {
			fmt.Fprintln(promptOut, err)
			continue
		}
		return v, nil
	}
}

func promptFlag(cmd *cobra.Command, f *pflag.Flag) (string, error) {
	var choices []string
	if fn, ok := cmd.GetFlagCompletionFunc(f.Name); ok {
		comps, _ := fn(cmd, nil, "")
		for _, c := range comps {
			choices = append(choices, string(c))
		}
	}
	if f.Value.Type() == "bool" {
		choices = []string{"true", "false"}
	}
	return ask("--"+f.Name, f.Usage, choices)
}

// ask prints a prompt (with a numbered list when choices are known) and reads
// one non-empty line. A listed value is taken as is; otherwise a number
// selects that menu entry, and any other text is the value itself.
func ask(label, help string, choices []string) (string, error) {
	if help != "" {
		fmt.Fprintf(promptOut, "%s: %s\n", label, help)
	} else if len(choices) > 0 {
		fmt.Fprintf(promptOut, "%s:\n", label)
	}
	shown := choices
	if len(shown) > maxChoices {
		shown = shown[:maxChoices]
	}
	for i, c := range shown {
		value, desc, _ := strings.Cut(c, "\t")
		if desc != "" {
			fmt.Fprintf(promptOut, "  %2d) %s  %s\n", i+1, value, desc)
		} else {
			fmt.Fprintf(promptOut, "  %2d) %s\n", i+1, value)
		}
	}
	if len(choices) > maxChoices {
		fmt.Fprintf(promptOut, "  … %d more, type a value\n", len(choices)-maxChoices)
	}
	prompt := label + ": "
	if help != "" || len(shown) > 0 {
		prompt = "> "
	}
	for {
		fmt.Fprint(promptOut, prompt)
		line, err := promptIn.ReadString('\n')
		line = strings.TrimSpace(line)
		if line == "" && err != nil {
			return "", FlagErrorf("no value given for %s", label)
		}
		if line == "" {
			continue
		}
		for _, c := range choices {
			if value, _, _ := strings.Cut(c, "\t"); value == line {
				return line, nil // a listed value, even if it also looks like a menu number
			}
		}
		if n, convErr := strconv.Atoi(line); convErr == nil && n >= 1 && n <= len(shown) {
			value, _, _ := strings.Cut(shown[n-1], "\t")
			return value, nil
		}
		return line, nil
	}
}

// ErrAborted is returned when the user declines a confirmation prompt.
var ErrAborted = fmt.Errorf("aborted")

// Confirm asks before a destructive command runs. --yes and dry-run (flag,
// config or environment) skip the question; a non-interactive session must pass --yes explicitly.
func Confirm(cmd *cobra.Command, action string, args []string) error {
	if flagBool(cmd, "yes") || flagBool(cmd, "dry-run") || viper.GetBool("dry_run") {
		return nil
	}
	if !Interactive(cmd) {
		return FlagErrorf("%q needs confirmation; pass --yes to run it non-interactively", cmd.CommandPath())
	}
	target := ""
	if len(args) > 0 {
		target = " " + strings.Join(args, " ")
	}
	fmt.Fprintf(promptOut, "%s%s? [y/N]: ", strings.TrimSuffix(action, "."), target)
	line, _ := promptIn.ReadString('\n')
	switch strings.ToLower(strings.TrimSpace(line)) {
	case "y", "yes":
		return nil
	}
	return ErrAborted
}

// ConfirmScoped is Confirm for a command whose target can come from the link,
// the environment or the config. Such a target is acted on only after a
// question; with --yes or without a terminal it must be typed.
func ConfirmScoped(cmd *cobra.Command, action string, args []string, typed bool) error {
	dryRun := flagBool(cmd, "dry-run") || viper.GetBool("dry_run")
	if !typed && !dryRun && (flagBool(cmd, "yes") || !Interactive(cmd)) && len(args) > 0 {
		return FlagErrorf("%q would act on %s, which was not typed but came from --domain, the environment, the link or the config: pass it as an argument", cmd.CommandPath(), args[0])
	}
	return Confirm(cmd, action, args)
}

// FitWidth shortens s to width columns, ending in "…". A line wider than the
// terminal wraps, and a spinner that redraws it with \r then adds a row per frame.
func FitWidth(s string, width int) string {
	if width <= 0 {
		return s
	}
	r := []rune(s)
	if len(r) <= width {
		return s
	}
	if width == 1 {
		return "…"
	}
	return string(r[:width-1]) + "…"
}

// StderrWidth is the terminal width of stderr, or 0 when unknown.
func StderrWidth() int {
	w, _, err := term.GetSize(int(os.Stderr.Fd()))
	if err != nil {
		return 0
	}
	return w
}

// flagBool reads a bool flag of cmd or one it inherits; false when absent.
func flagBool(cmd *cobra.Command, name string) bool {
	f := cmd.Flag(name)
	if f == nil {
		return false
	}
	v, err := strconv.ParseBool(f.Value.String())
	return err == nil && v
}
