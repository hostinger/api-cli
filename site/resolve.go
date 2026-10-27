package site

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/hostinger/api-cli/utils"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var ErrNoWebsite error = &utils.FlagError{Err: errors.New("no website given: pass the domain, or run `hostinger link`")}

// Resolve returns a site-scoped command's website and the remaining args.
// One arg is the domain; none means --domain, HOSTINGER_DOMAIN, the link,
// default_domain, then a prompt.
func Resolve(cmd *cobra.Command, args []string, extras int) (username, domain string, rest []string, err error) {
	ctx := context.Background()
	if cmd != nil && cmd.Context() != nil {
		ctx = cmd.Context()
	}
	switch len(args) - extras {
	case 2:
		return args[0], NormalizeDomain(args[1]), args[2:], nil
	case 1:
		domain, rest = NormalizeDomain(args[0]), args[1:]
	default:
		rest = args
		if domain, err = DefaultDomain(cmd); err != nil {
			return "", "", nil, err
		}
	}
	if domain == "" {
		if !Interactive() {
			return "", "", nil, ErrNoWebsite
		}
		w, err := ChooseWebsite(ctx, "")
		if err != nil {
			return "", "", nil, err
		}
		return w.Username, w.Domain, rest, nil
	}
	username, err = Username(ctx, domain)
	if err != nil {
		return "", "", nil, err
	}
	return username, domain, rest, nil
}

// DefaultDomain is the website named by flags, env, link or config, not args.
// A link that cannot be read is an error only when it is consulted.
func DefaultDomain(cmd *cobra.Command) (string, error) {
	if cmd != nil {
		if f := contextFlag(cmd, "domain"); f != nil && f.Changed {
			return NormalizeDomain(f.Value.String()), nil
		}
	}
	if v := os.Getenv("HOSTINGER_DOMAIN"); v != "" {
		return NormalizeDomain(v), nil
	}
	p, _, err := Find(".")
	if err != nil {
		return "", linkError(err)
	}
	if p != nil {
		return p.Domain, nil
	}
	return NormalizeDomain(viper.GetString("default_domain")), nil
}

// contextFlag is the root's --domain, --username or --order-id: a command's
// own flag of the same name (hosting wordpress install --domain) means
// something else. Commands without the root flag use their own.
func contextFlag(cmd *cobra.Command, name string) *pflag.Flag {
	if f := cmd.Root().PersistentFlags().Lookup(name); f != nil {
		return f
	}
	return cmd.Flags().Lookup(name)
}

// linkError says how to recover from a link that cannot be read.
func linkError(err error) error {
	return fmt.Errorf("the project link could not be read: %v (fix the file, or run `hostinger unlink`)", err)
}

// ChooseWebsite returns the only website without asking; current is preselected.
func ChooseWebsite(ctx context.Context, current string) (*Website, error) {
	done := Spin("Looking up your websites")
	sites, err := ListWebsites(ctx, "")
	if err != nil {
		done("Could not list your websites")
		return nil, err
	}
	done(fmt.Sprintf("%d websites", len(sites)))
	if len(sites) == 0 {
		return nil, errors.New("this account has no websites yet: run `hostinger up` to create one")
	}
	if len(sites) == 1 {
		return &sites[0], nil
	}
	names, _ := HostingPlanNames(ctx, Client()) // a failure leaves the list flat
	labels := PlanLabels(sites, names)
	SortByPlan(sites, labels)
	options := make([]string, len(sites))
	sel := pterm.DefaultInteractiveSelect.WithFilter(len(options) > 8).WithMaxHeight(SelectHeight(len(options)))
	for i, s := range sites {
		options[i] = PlanOption(s, labels)
		if current != "" && strings.EqualFold(current, s.Domain) {
			options[i] += "  (current)"
			sel = sel.WithDefaultOption(options[i])
		}
	}
	picked, err := pick(sel.WithOptions(options), "Which website?")
	if err != nil {
		return nil, err
	}
	for i, o := range options {
		if o == picked {
			return &sites[i], nil
		}
	}
	return nil, fmt.Errorf("no website selected")
}

func Label(w Website) string {
	label := w.Domain
	if w.WebsiteType != "" && w.WebsiteType != "other" {
		label += " (" + w.WebsiteType + ")"
	}
	return label
}
