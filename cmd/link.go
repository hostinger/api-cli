package cmd

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"strings"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/site"
	"github.com/hostinger/api-cli/utils"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

var (
	linkGlobal bool
	linkShow   bool
)

var LinkCmd = &cobra.Command{
	Use:   "link [domain]",
	Short: "Link this directory to a website",
	Long: `Remember which website the commands run from this directory act on, in
.hostinger/site.json. Commands that take [domain] then use it when the domain is
left out, and ` + "`hostinger up`" + ` redeploys to it.

With --global the website is written as default_domain in the config file instead,
for commands run outside any linked directory.

With --show nothing is changed: it prints what commands run from this directory
would act on and where each value comes from (flags, environment, the link, the
config file).`,
	Args:         cobra.MaximumNArgs(1),
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		if linkShow {
			return showContext(cmd)
		}
		ctx := cmd.Context()
		if ctx == nil {
			ctx = context.Background()
		}
		domain := ""
		if len(args) == 1 {
			domain = site.NormalizeDomain(args[0])
		} else if f := cmd.Flags().Lookup("domain"); f != nil && f.Changed {
			domain = site.NormalizeDomain(f.Value.String())
		} else {
			domain = site.NormalizeDomain(os.Getenv(utils.EnvPrefix + "DOMAIN"))
		}

		c := api.Raw()
		var w *site.Website
		if domain == "" {
			if !site.Interactive() {
				return errors.New("pass the domain to link: hostinger link example.com")
			}
			pterm.SetDefaultOutput(os.Stderr)
			cwd, _ := os.Getwd()
			currentDomain := ""
			if current, linkedDir, _ := site.Find(cwd); current != nil {
				currentDomain = current.Domain
				if linkedDir == cwd {
					pterm.Info.Printfln("This directory is linked to %s", current.Domain)
				} else {
					pterm.Info.Printfln("This directory is linked to %s by %s", current.Domain, site.Path(linkedDir))
				}
			}
			var err error
			if w, err = site.ChooseWebsite(ctx, currentDomain); err != nil {
				return api.CommandError(err)
			}
		} else {
			var err error
			w, err = site.FindWebsite(ctx, c, domain)
			if err != nil {
				return api.CommandError(err)
			}
			if w == nil {
				return fmt.Errorf("no website named %q on this account (hostinger hosting websites list shows them)", domain)
			}
		}

		if linkGlobal {
			path := configPath()
			if _, err := writeConfigKey(path, "default_domain", w.Domain); err != nil {
				return err
			}
			return printResult(cmd, "Default website set to "+w.Domain+" in "+path, map[string]any{"domain": w.Domain, "configFile": path})
		}
		dir, err := os.Getwd()
		if err != nil {
			return err
		}
		p := &site.Project{Kind: "website", OrderID: w.OrderID, Domain: w.Domain, Username: w.Username}
		// Re-linking the same website updates its link where it lives, here or in a
		// parent, and keeps what deploys remembered; another website is linked here.
		if old, oldDir, err := site.Find(dir); err == nil && old != nil && strings.EqualFold(old.Domain, w.Domain) {
			dir = oldDir
			p.Path, p.Type, p.Source, p.Node, p.Build = old.Path, old.Type, old.Source, old.Node, old.Build
		}
		if err := site.Save(dir, p); err != nil {
			return err
		}
		return printResult(cmd, fmt.Sprintf("Linked %s to %s (%s)", dir, w.Domain, site.Path(dir)), map[string]any{"domain": w.Domain, "linkFile": site.Path(dir)})
	},
}

// printResult prints text, or v as JSON when the output is JSON (asked for, or piped).
func printResult(cmd *cobra.Command, text string, v any) error {
	format, err := output.SelectedFormat(cmd)
	if err != nil {
		return err
	}
	if format == "json" {
		printJSON(v)
		return nil
	}
	fmt.Println(text)
	return nil
}

// printJSON matches the output package: indented, & < > unescaped.
func printJSON(v any) {
	enc := json.NewEncoder(os.Stdout)
	enc.SetEscapeHTML(false)
	enc.SetIndent("", "  ")
	_ = enc.Encode(v)
}

func showContext(cmd *cobra.Command) error {
	c, err := site.Explain(cmd)
	if err != nil {
		return fmt.Errorf("the project link could not be read: %v (fix the file, or run `hostinger unlink`)", err)
	}
	format, err := output.SelectedFormat(cmd)
	if err != nil {
		return err
	}
	if format == "json" {
		printJSON(c)
		return nil
	}
	for _, line := range c.Lines() {
		fmt.Println(line)
	}
	return nil
}

// writeConfigKey sets or (empty value) removes one key, unlike viper's
// WriteConfig, which also writes env and flag values. An emptied file is deleted.
func writeConfigKey(path, key, value string) (removed bool, err error) {
	left := 0
	err = writeConfig(path, func(m map[string]any) {
		if value == "" {
			delete(m, key)
		} else {
			m[key] = value
		}
		left = len(m)
	})
	if err != nil || left > 0 {
		return false, err
	}
	if err := os.Remove(path); err != nil && !errors.Is(err, fs.ErrNotExist) {
		return false, err
	}
	return true, nil
}

func init() {
	LinkCmd.Flags().BoolVar(&linkGlobal, "global", false, "Set default_domain in the config file instead of linking this directory")
	LinkCmd.Flags().BoolVar(&linkShow, "show", false, "Print what commands run from this directory act on, and where each value comes from")
	LinkCmd.GroupID = groupDeploy
	RootCmd.AddCommand(LinkCmd)
}
