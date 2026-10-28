package hosting

import (
	"context"
	"errors"
	"fmt"
	"slices"
	"time"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/site"
)

// suggested is the headless kind the folder suggests (SuggestHeadless).
func resolveWebsite(ctx context.Context, a API, ui UI, opts Options, project *site.Project, hint Type, suggested string) (*site.Website, bool, error) {
	// Creating a website from the picker reuses the lists the picker loaded.
	a = &listed{API: a}
	domain := opts.Domain
	if domain == "" && !opts.Create && project != nil {
		domain = project.Domain
	}
	// Only a Node.js app is offered a headless website.
	nodeApp := hint == NodeJS || opts.Type == NodeJS
	// --headless names a new website, unless the website given already is one.
	if opts.Create || (opts.Headless != "" && domain == "") {
		w, err := createWebsite(ctx, a, ui, opts, domain, nodeApp, suggested)
		return w, true, err
	}
	if domain != "" {
		w, err := site.FindWebsiteWith(ctx, a.ListWebsites, domain)
		if err != nil {
			return nil, false, err
		}
		if w != nil {
			return w, false, nil
		}
		switch {
		case opts.Headless != "":
		case !ui.Interactive():
			return nil, false, fmt.Errorf("no website named %q on this account: pass --create to create it", domain)
		default:
			ok, err := ui.Confirm(fmt.Sprintf("No website named %s on your account. Create it?", domain), true)
			if err != nil {
				return nil, false, err
			}
			if !ok {
				return nil, false, errors.New("no website to deploy to")
			}
		}
		w, err = createWebsite(ctx, a, ui, opts, domain, nodeApp, suggested)
		return w, true, err
	}
	if !ui.Interactive() {
		return nil, false, errors.New("no website given: pass --domain, or --create for a new one")
	}
	step := ui.Start("Looking up your websites")
	sites, err := a.ListWebsites(ctx, "")
	if err != nil {
		step.Fail("Could not list your websites")
		return nil, false, err
	}
	if len(sites) == 0 {
		step.Done("No websites yet; creating one")
		w, err := createWebsite(ctx, a, ui, opts, "", nodeApp, suggested)
		return w, true, err
	}
	plans := planNames(ctx, a, sites)
	site.SortByPlan(sites, plans)
	step.Done(websitesSummary(sites))
	// Preselect the first site matching the folder's type, else "Create".
	options := make([]string, 0, len(sites)+1)
	options = append(options, "Create a new website")
	def := 0
	for i, s := range sites {
		options = append(options, site.PlanOption(s, plans))
		if def == 0 && ((hint == NodeJS && (s.WebsiteType == "nodejs" || s.IsHeadless())) || (hint == Static && (s.WebsiteType == "other" || s.WebsiteType == ""))) {
			def = i + 1
		}
	}
	idx, err := ui.Select("Which website do you want to deploy to?", options, def)
	if err != nil {
		return nil, false, err
	}
	if idx == 0 {
		w, err := createWebsite(ctx, a, ui, opts, "", nodeApp, suggested)
		return w, true, err
	}
	return &sites[idx-1], false, nil
}

// listed lists the orders and all websites once; a failed list is not kept.
type listed struct {
	API
	orders      []Order
	sites       []site.Website
	sitesListed bool
}

func (l *listed) ListOrders(ctx context.Context) ([]Order, error) {
	if l.orders != nil {
		return l.orders, nil
	}
	orders, err := l.API.ListOrders(ctx)
	if err == nil {
		l.orders = orders
	}
	return orders, err
}

func (l *listed) ListWebsites(ctx context.Context, domain string) ([]site.Website, error) {
	if domain != "" {
		return l.API.ListWebsites(ctx, domain)
	}
	if l.sitesListed {
		return l.sites, nil
	}
	sites, err := l.API.ListWebsites(ctx, "")
	if err == nil {
		l.sites, l.sitesListed = sites, true
	}
	return sites, err
}

func websitesSummary(sites []site.Website) string {
	orders := map[int]bool{}
	for _, s := range sites {
		orders[s.OrderID] = true
	}
	msg := fmt.Sprintf("%d website", len(sites))
	if len(sites) != 1 {
		msg += "s"
	}
	if len(orders) > 1 {
		msg += fmt.Sprintf(" on %d plans", len(orders))
	}
	return msg
}

// planNames labels the websites' plans for a grouped picker; nil keeps it
// flat (one plan, or the orders cannot be listed).
func planNames(ctx context.Context, a API, sites []site.Website) map[int]string {
	ids := map[int]bool{}
	for _, s := range sites {
		if s.OrderID != 0 {
			ids[s.OrderID] = true
		}
	}
	if len(ids) < 2 {
		return nil
	}
	orders, err := a.ListOrders(ctx)
	if err != nil {
		return nil
	}
	return site.PlanLabels(sites, site.PlanNames(orders))
}

// createWebsite sets up a website; an empty domain gets a free subdomain. A
// Node.js app is asked whether it should be headless.
func createWebsite(ctx context.Context, a API, ui UI, opts Options, domain string, nodeApp bool, suggested string) (*site.Website, error) {
	websiteType := ""
	if opts.Headless != "" {
		websiteType = "headless"
	}
	if opts.DryRun {
		if opts.OrderID != 0 {
			orders, err := a.ListOrders(ctx)
			if err != nil {
				return nil, err
			}
			if !slices.ContainsFunc(orders, func(o Order) bool { return o.ID == opts.OrderID }) {
				return nil, fmt.Errorf("hosting order %d was not found on this account (`hostinger hosting orders list` shows them)", opts.OrderID)
			}
		}
		if domain == "" {
			domain = "<free subdomain>"
		}
		return &site.Website{Domain: domain, OrderID: opts.OrderID, WebsiteType: websiteType}, nil
	}
	orderID := opts.OrderID
	if orderID == 0 {
		orders, err := a.ListOrders(ctx)
		if err != nil {
			return nil, err
		}
		var active []Order
		for _, o := range orders {
			if o.Status == "" || o.Status == "active" {
				active = append(active, o)
			}
		}
		switch {
		case len(active) == 0:
			return nil, errors.New("this account has no Web or Cloud hosting plan; buy one at https://hpanel.hostinger.com, then run this again")
		case len(active) == 1:
			orderID = active[0].ID
			ui.Info("Using " + active[0].Label())
		case !ui.Interactive():
			return nil, fmt.Errorf("several hosting plans: pass --order-id (hostinger hosting orders list shows them)")
		default:
			options := make([]string, len(active))
			for i, o := range active {
				options[i] = o.Label()
			}
			idx, err := ui.Select("Which hosting plan should host the new website?", options, 0)
			if err != nil {
				return nil, err
			}
			orderID = active[idx].ID
		}
	}
	kind, err := chooseKind(ui, opts, nodeApp, suggested)
	if err != nil {
		return nil, err
	}
	if domain == "" && ui.Interactive() {
		v, err := ui.Input("Domain for the new website (leave empty for a free subdomain)", "")
		if err != nil {
			return nil, err
		}
		domain = site.NormalizeDomain(v)
		if domain != "" {
			owned, err := a.OwnsDomain(ctx, domain)
			if err == nil && !owned {
				ok, err := ui.Confirm("This account does not own "+domain+"; the website will not be reachable until the domain's DNS points to it. Create it anyway?", false)
				if err != nil {
					return nil, err
				}
				if !ok {
					return nil, errors.New("no website created: use a domain this account owns, or leave it empty for a free subdomain")
				}
			}
		}
	}
	datacenter, err := chooseDatacenter(ctx, a, ui, opts, orderID)
	if err != nil {
		return nil, err
	}
	step := ui.Start("Creating website")
	setupType := ""
	if kind != "" {
		setupType, websiteType = "headless_"+kind, "headless"
	}
	setup, err := a.StartSetup(ctx, orderID, domain, datacenter, setupType)
	if err != nil {
		step.Fail("Could not create website")
		var apiErr *api.Error
		if errors.As(err, &apiErr) && apiErr.Status == 404 {
			return nil, fmt.Errorf("hosting order %d was not found on this account (`hostinger hosting orders list` shows them)", orderID)
		}
		return nil, err
	}
	step.Done("Created " + setup.Domain)
	return &site.Website{Domain: setup.Domain, Username: setup.Username, OrderID: orderID, WebsiteType: websiteType}, nil
}

// chooseKind confirms the kind of a new website for a Node.js app: the headless
// website it looks like a frontend for, else a regular one. Declining asks the
// full question. Without a terminal it never guesses: --headless chooses.
func chooseKind(ui UI, opts Options, nodeApp bool, suggested string) (string, error) {
	if opts.Headless != "" || !nodeApp {
		return opts.Headless, nil
	}
	s, known := suggestions[suggested]
	if !ui.Interactive() {
		if known {
			ui.Info("This looks like " + s.looks + "; pass --headless " + suggested + " to create a headless website for it")
		}
		return "", nil
	}
	question, def := "Create a regular website?", 1 // declined: offer headless first
	if known {
		question, def = "This looks like "+s.looks+". "+s.create+"?", 0
	}
	ok, err := ui.Confirm(question, true)
	if err != nil || ok {
		return suggested, err
	}
	idx, err := ui.Select("What kind of website?", headlessChoices, def)
	if err != nil || idx == 0 {
		return "", err
	}
	return HeadlessKinds[idx-1], nil
}

// suggestions word chooseKind's question for what SuggestHeadless finds.
var suggestions = map[string]struct{ looks, create string }{
	"empty":      {"a web frontend", "Create a headless website (frontend only, add a backend later)"},
	"wordpress":  {"a headless WordPress frontend", "Create a headless website with WordPress in /cms"},
	"pocketbase": {"a frontend for PocketBase", "Create a headless website for a PocketBase backend"},
}

// headlessChoices follow HeadlessKinds, after a regular website.
var headlessChoices = []string{
	"Regular website",
	"Headless, frontend only (add a backend later)",
	"Headless with WordPress (installed in /cms)",
	"Headless for a Hostinger Ecommerce store",
	"Headless for a PocketBase backend",
}

// chooseDatacenter matters only for a plan's first website; later ones share its server.
func chooseDatacenter(ctx context.Context, a API, ui UI, opts Options, orderID int) (string, error) {
	if opts.Datacenter != "" {
		return opts.Datacenter, nil
	}
	step := ui.Start("Checking the plan")
	sites, err := a.ListWebsites(ctx, "")
	if err != nil {
		step.Fail("Could not list your websites")
		return "", err
	}
	for _, s := range sites {
		if s.OrderID == orderID {
			step.Done("The plan already has a server")
			return "", nil
		}
	}
	centers, err := a.ListDatacenters(ctx, orderID)
	if err != nil || len(centers) == 0 {
		step.Done("The platform picks the datacenter")
		return "", nil // let the platform choose
	}
	step.Done("First website on this plan: pick where it lives")
	if !ui.Interactive() {
		ui.Info("Hosting the plan in " + centers[0].Title + " (the best match; --datacenter to choose)")
		return centers[0].Code, nil
	}
	options := make([]string, len(centers))
	for i, c := range centers {
		options[i] = c.Title + " (" + c.Code + ")"
	}
	idx, err := ui.Select("Where should the website be hosted? (first is the best match)", options, 0)
	if err != nil {
		return "", err
	}
	return centers[idx].Code, nil
}

// waitForSetup waits for provisioning. An older site usually lists no setup,
// which counts as ready.
func waitForSetup(ctx context.Context, a API, ui UI, domain string, created bool) error {
	setups, err := a.ListSetups(ctx, domain)
	if err != nil {
		if !created {
			return nil
		}
		return err
	}
	if len(setups) == 0 {
		return nil
	}
	latest := setups[0]
	if latest.Status == "completed" {
		return nil
	}
	if latest.Status == "failed" {
		return fmt.Errorf("the setup of %s failed; check it in hPanel", domain)
	}
	step := ui.Start("Setting up " + domain)
	deadline := Now().Add(SetupTimeout)
	failures := 0
	for {
		select {
		case <-ctx.Done():
			step.Fail("Interrupted")
			return ctx.Err()
		case <-time.After(SetupPollInterval):
		}
		setups, err := a.ListSetups(ctx, domain)
		if err != nil {
			if failures++; failures < 3 { // one failed check does not end a long wait
				continue
			}
			step.Fail("Could not read setup status")
			return err
		}
		failures = 0
		if len(setups) > 0 {
			switch setups[0].Status {
			case "completed":
				step.Done(domain + " is ready")
				return nil
			case "failed":
				step.Fail("Setup failed")
				return fmt.Errorf("the setup of %s failed; check it in hPanel", domain)
			}
		}
		if Now().After(deadline) {
			step.Fail("Timed out")
			return fmt.Errorf("gave up waiting for the setup of %s after %s; run the command again later", domain, SetupTimeout)
		}
	}
}

// confirmOverwrite asks before the first deploy to an unlinked website, since
// a deploy replaces its content. A site with content defaults to no and, without
// a terminal, needs --yes. On a headless website only the frontend is replaced.
func confirmOverwrite(ctx context.Context, a API, ui UI, w *site.Website, opts Options, headless bool) error {
	has, checkErr := hasContent(ctx, a, w, headless)
	what, frontend, kept := "content", "", ""
	if headless {
		what, frontend, kept = "a frontend", "frontend ", " (its backend, such as WordPress in /cms, is kept)"
	}
	if opts.Yes {
		switch {
		case checkErr != nil:
			ui.Info("Could not check whether " + w.Domain + " has " + what + "; this deploy replaces whatever " + frontend + "it serves" + kept)
		case has:
			ui.Info(w.Domain + " already has " + what + "; this deploy replaces it" + kept)
		}
		return nil
	}
	if !ui.Interactive() {
		switch {
		case checkErr != nil:
			return fmt.Errorf("could not check whether %s has %s (%v): pass --yes to replace it", w.Domain, what, checkErr)
		case has:
			return errors.New(w.Domain + " already has " + what + ", which this deploy replaces" + kept + ": pass --yes to replace it")
		}
		return nil
	}
	question, def := "Deploying replaces whatever "+frontend+w.Domain+" serves now"+kept+". Continue?", true
	if has {
		question, def = w.Domain+" already has "+what+", which this deploy replaces"+kept+". Continue?", false
	}
	ok, err := ui.Confirm(question, def)
	if err != nil {
		return err
	}
	if !ok {
		return errors.New("deploy cancelled: nothing was changed")
	}
	return nil
}

// hasContent reports whether a deploy would replace something; content that
// cannot be checked counts, and checkErr says why. A headless website's
// document root holds its backend, so only a built frontend counts there.
func hasContent(ctx context.Context, a API, w *site.Website, headless bool) (has bool, checkErr error) {
	if headless {
		stored, err := a.BuildSettings(ctx, w.Username, w.Domain)
		return stored != nil || err != nil, err
	}
	if w.WebsiteType == "nodejs" { // a Node.js app lives outside the document root
		return true, nil
	}
	has, checkErr = a.HasContent(ctx, w.Username, w.Domain)
	return has || checkErr != nil, checkErr
}

// confirmStaticOverApp asks before static files go onto a Node.js site: the
// app keeps running and answers every other path.
func confirmStaticOverApp(ui UI, w *site.Website, opts Options) error {
	msg := w.Domain + " runs a Node.js app. A static deploy only adds files to its document root: " +
		"the app keeps running and still answers every path that is not one of those files"
	switch {
	case opts.Yes || opts.DryRun:
		ui.Info(msg + ".")
		return nil
	case !ui.Interactive():
		return errors.New(msg + ". Pass --yes to deploy the files anyway, or --type nodejs to redeploy the app")
	}
	ok, err := ui.Confirm(msg+". Deploy the static files anyway?", false)
	if err != nil {
		return err
	}
	if !ok {
		return errors.New("deploy cancelled: nothing was changed (deploy the app with --type nodejs, or remove it in hPanel first)")
	}
	return nil
}
