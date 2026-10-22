package cmd

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"text/tabwriter"

	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.yaml.in/yaml/v3"
	"golang.org/x/term"
)

// configKeys are the settings `config` manages; each can also be set through
// the matching environment variable (EnvPrefix + upper-case key).
var configKeys = []struct{ key, desc string }{
	{"api_token", "API token (takes precedence over any sign-in)"},
	{"api_url", "API base URL"},
	{"format", "Default output format (" + strings.Join(output.Formats, "|") + ")"},
	{"no_input", "Never prompt for missing values or confirmations (true|false)"},
}

var ConfigCmd = &cobra.Command{
	Use:     "config",
	Short:   "Read and write CLI settings",
	GroupID: groupSetup,
	Args:    cobra.ArbitraryArgs,
	RunE:    utils.GroupRunE,
}

var configSetCmd = &cobra.Command{
	Use:   "set <key> [value]",
	Short: "Set a setting in the config file",
	Long: `Set a setting in the config file. For api_token the value may be left out:
it is then read from a hidden prompt so it stays out of the shell history.`,
	Example: "  config set format json\n  config set api_token",
	Args:    cobra.RangeArgs(1, 2),
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]cobra.Completion, cobra.ShellCompDirective) {
		if len(args) == 0 {
			return configKeyCompletions(), cobra.ShellCompDirectiveNoFileComp
		}
		if len(args) == 1 && args[0] == "format" {
			return output.Formats, cobra.ShellCompDirectiveNoFileComp
		}
		return nil, cobra.ShellCompDirectiveNoFileComp
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		key := args[0]
		if err := checkConfigKey(key); err != nil {
			return err
		}
		var value string
		if len(args) == 2 {
			value = args[1]
		} else if key == "api_token" && term.IsTerminal(int(os.Stdin.Fd())) {
			fmt.Fprint(os.Stderr, "API token: ")
			b, err := term.ReadPassword(int(os.Stdin.Fd()))
			fmt.Fprintln(os.Stderr)
			if err != nil {
				return err
			}
			value = strings.TrimSpace(string(b))
		} else {
			return utils.FlagErrorf("missing value for %s", key)
		}
		if err := checkConfigValue(key, value); err != nil {
			return err
		}
		path := configPath()
		if err := writeConfig(path, func(m map[string]any) { m[key] = value }); err != nil {
			return err
		}
		fmt.Fprintf(cmd.OutOrStdout(), "Set %s in %s\n", key, path)
		return nil
	},
}

var configGetCmd = &cobra.Command{
	Use:               "get <key>",
	Short:             "Print the effective value of a setting",
	Args:              cobra.ExactArgs(1),
	ValidArgsFunction: completeConfigKey,
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := checkConfigKey(args[0]); err != nil {
			return err
		}
		value := viper.GetString(args[0])
		if args[0] == "api_token" && value != "" {
			value = maskSecret(value)
		}
		fmt.Fprintln(cmd.OutOrStdout(), value)
		return nil
	},
}

var configUnsetCmd = &cobra.Command{
	Use:               "unset <key>",
	Short:             "Remove a setting from the config file",
	Args:              cobra.ExactArgs(1),
	ValidArgsFunction: completeConfigKey,
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := checkConfigKey(args[0]); err != nil {
			return err
		}
		path := configPath()
		if err := writeConfig(path, func(m map[string]any) { delete(m, args[0]) }); err != nil {
			return err
		}
		fmt.Fprintf(cmd.OutOrStdout(), "Removed %s from %s\n", args[0], path)
		return nil
	},
}

var configListCmd = &cobra.Command{
	Use:   "list",
	Short: "List the settings and their effective values",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		out := cmd.OutOrStdout()
		fmt.Fprintf(out, "Config file: %s\n\n", configPath())
		w := tabwriter.NewWriter(out, 0, 0, 2, ' ', 0)
		for _, k := range configKeys {
			value := viper.GetString(k.key)
			if k.key == "api_token" && value != "" {
				value = maskSecret(value)
			}
			if value == "" {
				value = "(not set)"
			}
			fmt.Fprintf(w, "%s\t%s\t%s\n", k.key, value, k.desc)
		}
		if err := w.Flush(); err != nil {
			return err
		}
		fmt.Fprintf(out, "\nEnvironment variables override the file: %sAPI_TOKEN, %sFORMAT, …\n", utils.EnvPrefix, utils.EnvPrefix)
		return nil
	},
}

var configPathCmd = &cobra.Command{
	Use:   "path",
	Short: "Print the config file path",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Fprintln(cmd.OutOrStdout(), configPath())
		return nil
	},
}

func init() {
	ConfigCmd.AddCommand(configSetCmd, configGetCmd, configUnsetCmd, configListCmd, configPathCmd)
}

// configPath is the file `config set` writes: --config, else the file that was
// read, else the default in the home directory.
func configPath() string {
	if cfgFile != "" {
		return cfgFile
	}
	if used := viper.ConfigFileUsed(); used != "" {
		return used
	}
	home, err := os.UserHomeDir()
	if err != nil {
		return utils.ConfigFileName
	}
	return filepath.Join(home, utils.ConfigFileName)
}

// writeConfig applies change to the file's own settings only (never values
// that came from the environment or flags) and writes it back with mode 0600.
func writeConfig(path string, change func(m map[string]any)) error {
	m := map[string]any{}
	raw, err := os.ReadFile(path)
	if err != nil && !errors.Is(err, fs.ErrNotExist) {
		return fmt.Errorf("reading %s: %w", path, err)
	}
	if err := yaml.Unmarshal(raw, &m); err != nil {
		return fmt.Errorf("reading %s: %w", path, err)
	}
	if m == nil {
		m = map[string]any{}
	}
	change(m)
	out, err := yaml.Marshal(m)
	if err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o700); err != nil {
		return err
	}
	if err := os.WriteFile(path, out, 0o600); err != nil {
		return fmt.Errorf("writing %s: %w", path, err)
	}
	return os.Chmod(path, 0o600)
}

func checkConfigKey(key string) error {
	for _, k := range configKeys {
		if k.key == key {
			return nil
		}
	}
	names := make([]string, len(configKeys))
	for i, k := range configKeys {
		names[i] = k.key
	}
	return utils.FlagErrorf("unknown setting %q (one of: %s)", key, strings.Join(names, ", "))
}

func checkConfigValue(key, value string) error {
	switch key {
	case "format":
		if !slices.Contains(output.Formats, value) {
			return utils.FlagErrorf("invalid format %q (one of: %s)", value, strings.Join(output.Formats, ", "))
		}
	case "no_input":
		if value != "true" && value != "false" {
			return utils.FlagErrorf("no_input must be true or false")
		}
	}
	return nil
}

func configKeyCompletions() []cobra.Completion {
	out := make([]cobra.Completion, len(configKeys))
	for i, k := range configKeys {
		out[i] = cobra.CompletionWithDesc(k.key, k.desc)
	}
	return out
}

func completeConfigKey(cmd *cobra.Command, args []string, toComplete string) ([]cobra.Completion, cobra.ShellCompDirective) {
	if len(args) > 0 {
		return nil, cobra.ShellCompDirectiveNoFileComp
	}
	return configKeyCompletions(), cobra.ShellCompDirectiveNoFileComp
}
