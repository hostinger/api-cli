package hosting

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/spf13/pflag"
)

var flags struct {
	orderID    int
	datacenter string
	create     bool
	typ        string
	yes        bool
	settings   bool
	source     string
	headless   string
	dryRun     bool
	noWait     bool
	node       NodeOverrides
}

func AddFlags(f *pflag.FlagSet) {
	f.IntVar(&flags.orderID, "order-id", 0, "Hosting plan (order ID) a new website is created on")
	f.StringVar(&flags.datacenter, "datacenter", "", "Where the first website on a plan is hosted, a code from hostinger hosting datacenters list (default: asked, or the API's best match)")
	f.BoolVar(&flags.create, "create", false, "Create a new website instead of deploying to an existing one")
	f.StringVar(&flags.headless, "headless", "", "Create a headless website for a Node.js frontend: empty, wordpress (installed in /cms), ecommerce or pocketbase. A new website unless the linked website or --domain is already headless")
	f.StringVar(&flags.typ, "type", "", "What to deploy as: static or nodejs (default: detected, or remembered)")
	f.BoolVarP(&flags.yes, "yes", "y", false, "Accept the detected type, and replace a website's content or deploy static files onto a Node.js site without asking (required for those without a terminal)")
	f.StringVar(&flags.source, "source", "", "Where the files come from: upload (pack this directory) or git (the pushed branch, through a connected Git account); default: remembered, else asked")
	f.BoolVar(&flags.dryRun, "dry-run", false, "Show what would be deployed where, as JSON, without creating, uploading or building anything")
	f.BoolVar(&flags.noWait, "no-wait", false, "Return as soon as a Node.js build is started instead of waiting for it")
	f.BoolVar(&flags.settings, "settings", false, "Review the Node.js build settings before building (a first deploy in a terminal reviews them anyway; without a terminal the detected settings and flags are used)")
	f.IntVar(&flags.node.NodeVersion, "node-version", 0, "Node.js major version (18, 20, 22, 24)")
	f.StringVar(&flags.node.AppType, "app-type", "", "Node.js framework, e.g. express, next, vite, nest, other")
	f.StringVar(&flags.node.RootDirectory, "root-directory", "", "Directory holding package.json, relative to the archive root")
	f.StringVar(&flags.node.OutputDirectory, "output-directory", "", "Build output directory, relative to the root directory")
	f.StringVar(&flags.node.BuildScript, "build-script", "", "npm script that builds the app")
	f.StringVar(&flags.node.EntryFile, "entry-file", "", "Entry file of the server (compiled apps)")
	f.StringVar(&flags.node.PackageManager, "package-manager", "", "npm, yarn or pnpm")
}

// Deploy runs Run with the up command's flags.
func Deploy(ctx context.Context, ui UI, path, domain string) (*Outcome, error) {
	opts := Options{Path: path, Domain: domain, Create: flags.create, OrderID: flags.orderID, Datacenter: flags.datacenter, Yes: flags.yes, Settings: flags.settings, Node: flags.node, DryRun: flags.dryRun, NoWait: flags.noWait}
	if flags.typ != "" {
		t, err := ParseType(flags.typ)
		if err != nil {
			return nil, err
		}
		opts.Type = t
	}
	if flags.headless != "" {
		kind, err := ParseHeadless(flags.headless)
		if err != nil {
			return nil, err
		}
		opts.Headless = kind
	}
	if flags.source != "" {
		src, err := ParseSource(flags.source)
		if err != nil {
			return nil, err
		}
		opts.Source = src
	}
	result, err := Run(ctx, NewAPI(api.Raw()), ui, opts)
	if result == nil {
		return nil, err
	}
	return &Outcome{URL: result.URL, Target: result.Domain, ProjectFile: result.ProjectFile, NewLink: result.NewLink, DryRun: result.DryRun, Note: result.Next, Reachable: result.Reachable, ReachableNote: result.ReachableNote, Data: result}, err
}
