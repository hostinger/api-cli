package cmd

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/deploy/hosting"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/site"
	"github.com/hostinger/api-cli/utils"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/term"
)

var UpCmd = &cobra.Command{
	Use:     "up [path]",
	Aliases: []string{"deploy"},
	Short:   "Deploy a directory",
	Long: `Deploy a directory (default: the current one) to a Web or Cloud hosting website.

The website comes from --domain, the project file written by an earlier run or by
` + "`hostinger link`" + `, or a prompt. A website that does not exist yet is created on a
hosting plan (--order-id, or a prompt), with a free subdomain when no domain is given.
The directory is deployed as a static site (HTML, CSS, JavaScript, PHP) or a Node.js
app, which is built on the server, from an uploaded copy of the directory or from its
pushed Git branch (--source git, asked once and remembered, which also redeploys on
every push). Deploying overwrites the website's current files.

A headless website (--headless creates one: empty, wordpress with WordPress in /cms,
ecommerce or pocketbase) serves a Node.js frontend, so a folder with a package.json
is deployed there as a Node.js app and its backend is left alone. In a terminal, a new
website is confirmed as a regular one, or as the headless website its package.json
points to: a PocketBase or WordPress client, else an empty one for a web frontend
(Next.js, Nuxt, Astro, Vite and similar); --headless skips that.

The directory is linked to the website in .hostinger/site.json, written in the
current directory with the deployed folder recorded, so ` + "`hostinger up`" + ` from the
project root redeploys the same folder. --domain on a linked directory deploys there
once and leaves the link alone; ` + "`hostinger link`" + ` changes it.

Without a terminal (CI, pipes, --format json) nothing is asked: pass --domain for an
existing website, or --create for a new one (--order-id when the account has more than
one hosting plan; the only plan is used by itself), and --yes to replace a website that
already has content. up never deploys to default_domain from the config file: only to the
linked website, --domain or HOSTINGER_DOMAIN.`,
	Example: `  hostinger up
  hostinger up ./dist --domain example.com
  hostinger up --headless wordpress
  hostinger up --create --order-id 123 --type nodejs --format json`,
	Args:         cobra.MaximumNArgs(1),
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		path := "."
		if len(args) == 1 {
			path = args[0]
		}
		return runUp(cmd, path)
	},
}

func init() {
	flags := UpCmd.Flags()
	hosting.AddFlags(flags)
	versions := make([]string, len(hosting.NodeVersions))
	for i, v := range hosting.NodeVersions {
		versions[i] = strconv.Itoa(v)
	}
	for name, values := range map[string][]string{
		"type":            {string(hosting.Static), string(hosting.NodeJS)},
		"source":          {string(hosting.SourceUpload), string(hosting.SourceGit)},
		"headless":        hosting.HeadlessKinds,
		"app-type":        hosting.AppTypes,
		"package-manager": hosting.PackageManagers,
		"node-version":    versions,
	} {
		flagChoices[flags.Lookup(name)] = values
	}
	UpCmd.GroupID = groupDeploy
	RootCmd.AddCommand(UpCmd)
}

func runUp(cmd *cobra.Command, path string) error {
	format, err := output.SelectedFormat(cmd)
	if err != nil {
		return err
	}
	jsonOutput := format == "json"
	domain, _ := cmd.Flags().GetString("domain")
	if domain == "" {
		domain = os.Getenv(utils.EnvPrefix + "DOMAIN")
	}
	domain = site.NormalizeDomain(domain)

	ui := newTermUI(jsonOutput)
	// default_domain is too easy to deploy over by accident: name it, don't use it.
	create, _ := cmd.Flags().GetBool("create")
	if def := viper.GetString("default_domain"); domain == "" && def != "" && !create {
		if p, _, err := site.Find("."); err == nil && p == nil {
			ui.Info("default_domain " + def + " in the config file is not used by up; pass --domain " + def + " to deploy there")
		}
	}
	base := cmd.Context()
	if base == nil {
		base = context.Background()
	}
	// Ctrl-C cancels: packing or the upload stops and the temp archive is removed.
	// A second Ctrl-C kills the process, as root.go's handler stops too.
	ctx, stop := signal.NotifyContext(base, os.Interrupt, syscall.SIGTERM)
	defer stop()
	context.AfterFunc(ctx, stop)
	out, err := hosting.Deploy(ctx, ui, path, domain)
	if err != nil {
		if errors.Is(err, context.Canceled) {
			err = errors.New("cancelled")
		}
		var berr *hosting.BuildError
		if errors.As(err, &berr) && berr.Logs != "" && !jsonOutput {
			fmt.Fprintln(os.Stderr, plainLogs(berr.Logs))
		}
		if jsonOutput {
			printJSON(withError(out, err))
		}
		return api.CommandError(err)
	}
	if jsonOutput || out.DryRun {
		printJSON(out)
		return nil
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
	return nil
}

// withError adds an error field, so a script never mistakes a failure for success.
func withError(out *hosting.Outcome, err error) map[string]any {
	m := map[string]any{}
	if out != nil {
		if raw, merr := json.Marshal(out); merr == nil {
			_ = json.Unmarshal(raw, &m)
		}
	}
	m["error"] = err.Error()
	return m
}

// plainLogs keeps build-tool colours only where they show as colours.
func plainLogs(logs string) string {
	_, noColor := os.LookupEnv("NO_COLOR")
	return utils.SafeLog(logs, !noColor && term.IsTerminal(int(os.Stderr.Fd())))
}
