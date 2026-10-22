package cmd

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var completionShells = []string{"bash", "zsh", "fish", "powershell"}

var completionInstallCmd = &cobra.Command{
	Use:   "install [bash|zsh|fish|powershell]",
	Short: "Install the completion script for your shell",
	Long: `Write the completion script where your shell loads it from. The shell is
taken from $SHELL unless given. Run it again after upgrading to pick up new
commands. Homebrew installs already include completions.`,
	Args:      cobra.MaximumNArgs(1),
	ValidArgs: completionShells,
	RunE: func(cmd *cobra.Command, args []string) error {
		shell := filepath.Base(os.Getenv("SHELL"))
		if len(args) == 1 {
			shell = args[0]
		}
		path, note, err := installCompletion(cmd.Root(), shell)
		if err != nil {
			return err
		}
		w := cmd.OutOrStdout()
		fmt.Fprintf(w, "Installed %s completion to %s\n", shell, path)
		if note != "" {
			fmt.Fprintln(w, note)
		}
		return nil
	},
}

// registerCompletionInstall adds `completion install` to cobra's default
// completion command, creating that command up front.
func registerCompletionInstall(root *cobra.Command) {
	root.InitDefaultCompletionCmd()
	for _, c := range root.Commands() {
		if c.Name() == "completion" {
			c.AddCommand(completionInstallCmd)
			return
		}
	}
}

// installCompletion writes the script for shell and returns its path plus any
// one-time setup the shell still needs.
func installCompletion(root *cobra.Command, shell string) (path, note string, err error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", "", err
	}
	name := root.Name()
	var script bytes.Buffer
	switch shell {
	case "bash":
		path = filepath.Join(envOr("XDG_DATA_HOME", filepath.Join(home, ".local", "share")), "bash-completion", "completions", name)
		err = root.GenBashCompletionV2(&script, true)
		note = "Loaded by bash-completion 2 in new shells (install it with your package manager if TAB does nothing)."
	case "zsh":
		dir := filepath.Join(envOr("ZDOTDIR", home), ".zsh", "completions")
		path = filepath.Join(dir, "_"+name)
		err = root.GenZshCompletion(&script)
		note = fmt.Sprintf("If %s is not in your fpath yet, add this to ~/.zshrc and open a new shell:\n  fpath=(%s $fpath)\n  autoload -U compinit && compinit", dir, dir)
	case "fish":
		path = filepath.Join(envOr("XDG_CONFIG_HOME", filepath.Join(home, ".config")), "fish", "completions", name+".fish")
		err = root.GenFishCompletion(&script, true)
		note = "Loaded automatically in new fish shells."
	case "powershell", "pwsh":
		dir, derr := os.UserConfigDir()
		if derr != nil {
			return "", "", derr
		}
		path = filepath.Join(dir, name, "completion.ps1")
		err = root.GenPowerShellCompletionWithDesc(&script)
		note = fmt.Sprintf("Add this line to your PowerShell profile ($PROFILE):\n  . '%s'", path)
	default:
		return "", "", utils.FlagErrorf("cannot detect a supported shell from %q; pass one of: %s", shell, strings.Join(completionShells, ", "))
	}
	if err != nil {
		return "", "", err
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return "", "", err
	}
	if err := os.WriteFile(path, script.Bytes(), 0o644); err != nil {
		return "", "", err
	}
	return path, note, nil
}

func envOr(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
