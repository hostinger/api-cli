package preview

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"net"
	"net/url"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"time"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/brand"
	"github.com/hostinger/api-cli/utils"
	"github.com/hostinger/hcdn-connector/pkg/connector"
	"github.com/hostinger/hcdn-connector/pkg/creds"
	"github.com/hostinger/hcdn-connector/pkg/quickview"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/term"
)

// Cmd is the `preview` command.
var Cmd = &cobra.Command{
	Use:   "preview [port]",
	Short: "Share a local dev server at a temporary public URL",
	Long: `Share a web server running on this computer at a temporary public https URL, for
example to show work in progress or to test a webhook. Anyone with the URL can open it.

The Preview runs in the foreground until Ctrl+C, or until its time runs out. Requests go
through an encrypted tunnel to the server while the command runs; nothing is uploaded
unless hPanel asks to deploy this folder (see below).

Without a port, preview looks for a dev server on localhost ports 5173, 3000, 4321,
8080, 8000, 5000, 4200 and 8888 (Jupyter's port, shared with a warning). With several, it
asks which one in a terminal and fails otherwise. --url shares another origin, such as
https://localhost:8443 or a server on your network. Ports of databases, remote shells,
debuggers and similar services (for example 22, 3306, 5432, 6379 and 9229) are never
shared from this computer, whatever address or name reaches it.

Requests reach the server with the Preview's host name in the Host header. A dev server
that only answers to its own names needs to allow .hostingerpreview.com; in Vite, set
server.allowedHosts: ['.hostingerpreview.com'] in vite.config.js. --host-header replaces
the Host header instead. Avoid localhost or an IP address there: it bypasses the dev
server's host check, and its debug consoles become public.

--ttl sets how long the Preview lasts: 24h by default, from 1h to 7d, such as 90m, 12h,
3d or 1d12h. It ends a few minutes before that, and the expiry is printed with the URL.

A Preview needs a Hostinger account: sign in with auth login, or set an API token. In a
terminal, preview opens the browser to sign in when you are not signed in. Elsewhere,
and with --json, it exits with code 77 instead. The Preview is linked to that account.

The URL is printed once the Preview answers. If that takes longer than a few seconds,
the URL is printed early with a note that it may answer 530 for a few seconds, and
"Preview ready." follows once it answers. In a terminal, a live view shows the same
steps and redraws in place. Piped output, --json, --no-input, CI, AI agents (AI_AGENT or
CLAUDECODE set) and NO_COLOR keep plain lines.

hPanel's Previews page can deploy the folder a Preview runs in, as ` + "`hostinger up`" + ` would:
a folder linked to a website (by an earlier up or ` + "`hostinger link`" + `) goes to that website,
and a Node.js project not linked to one goes to the website picked in hPanel, and is
linked to it. hPanel shows the linked website, the project type and the folder's name.
The terminal shows the deploy, and Ctrl+C cancels it while the Preview keeps running. A
WordPress import and static files onto a Node.js app are only deployed from the
terminal. Remote deploys are off with --no-remote-deploy or --url, and in a folder that
is neither linked to a website nor a Node.js project.

--json prints one JSON object per line on stdout: {"event":"ready",...} with the URL and
"reachable" (false while it may still answer 530), {"event":"reachable"} when such a
Preview starts to answer, {"event":"paused",...}, {"event":"deploy",...} with "state"
accepted, succeeded or failed (with a "reason") for a deploy from hPanel, and
{"event":"ended",...}, which carries a "message" when the Preview failed. Logs go to
stderr.

Exit codes: 0 when stopped or expired, 2 for a usage error such as a refused port, 77
when no account is signed in or the service refuses it, 78 when the Preview was taken
down, and 1 for any other failure.`,
	Example: `  hostinger preview
  hostinger preview 3000
  hostinger preview 3000 --ttl 3d
  hostinger preview 8080 --host-header shop.test
  hostinger preview --url https://localhost:8443
  hostinger preview --json`,
	Args: cobra.MatchAll(cobra.MaximumNArgs(1), func(_ *cobra.Command, args []string) error {
		if len(args) == 1 {
			_, err := parsePort(args[0])
			return err
		}
		return nil
	}),
	// The root prints errors once, with the exit code they map to.
	SilenceErrors: true,
	SilenceUsage:  true,
	RunE:          runPreview,
}

// options are the command's flags.
type options struct {
	url          string
	hostHeader   string
	json         bool
	bootstrap    []string
	serverCAFile string
	ports        []int
	noOwner      bool
	ttl          string
	noRemote     bool
}

var flags options

func init() {
	f := Cmd.Flags()
	f.StringVar(&flags.url, "url", "", "Share this origin instead of a local port, e.g. http://192.168.1.20:8080 (http or https)")
	f.StringVar(&flags.hostHeader, "host-header", "", "Host header sent to the server (default: the Preview's host name; prefer allowing .hostingerpreview.com in the dev server)")
	f.BoolVar(&flags.json, "json", false, "Print one JSON object per event on stdout")
	f.StringVar(&flags.ttl, "ttl", "24h", "How long the Preview lasts, from 1h to 7d (e.g. 90m, 12h, 3d)")
	f.BoolVar(&flags.noRemote, "no-remote-deploy", false, "Never deploy this folder when asked from hPanel")

	f.StringSliceVar(&flags.bootstrap, "bootstrap", nil, "Tunnel endpoints to connect to (comma separated)")
	f.StringVar(&flags.serverCAFile, "server-ca-file", "", "CA bundle that signs the tunnel endpoints' certificates")
	f.IntSliceVar(&flags.ports, "ports", nil, "Tunnel ports to try (comma separated)")
	f.BoolVar(&flags.noOwner, "no-owner", false, "Never link the Preview to an account")
	for _, name := range []string{"bootstrap", "server-ca-file", "ports", "no-owner"} {
		_ = f.MarkHidden(name)
	}
}

// tokenEnv is the environment variable of the API token, named in the sign-in
// messages.
var tokenEnv = utils.EnvPrefix + "API_TOKEN"

// runQuick is connector.RunQuick; tests replace it.
var runQuick = connector.RunQuick

// interactive reports whether a person can answer prompts and the browser
// sign-in; tests replace it.
var interactive = utils.Interactive

// newView returns the live view when w is a terminal that can show it, or
// nil for plain lines; tests replace it.
var newView = func(w io.Writer, o quickview.Options) *quickview.View {
	f, ok := w.(*os.File)
	if !ok || !quickview.Supported(f) {
		return nil
	}
	return quickview.New(f, o)
}

// newDetector builds the dev server detector for cmd; tests replace it.
var newDetector = func(cmd *cobra.Command, jsonOut bool) detector {
	d := &net.Dialer{}
	return detector{
		ports:       probePorts,
		dial:        d.DialContext,
		interactive: !jsonOut && utils.Interactive(cmd),
		choose:      chooseServer,
		binary:      cmd.Root().Name(),
	}
}

func chooseServer(labels []string) (int, error) {
	pterm.SetDefaultOutput(os.Stderr)
	picked, err := pterm.DefaultInteractiveSelect.WithOptions(labels).Show("Which server?")
	if err != nil {
		return -1, err
	}
	for i, l := range labels {
		if l == picked {
			return i, nil
		}
	}
	return -1, nil
}

func runPreview(cmd *cobra.Command, args []string) error {
	o := flags
	jsonOut := o.json || formatJSON(cmd)
	lifetime, err := parseTTL(o.ttl)
	if err != nil {
		return err
	}

	base := cmd.Context()
	if base == nil {
		base = context.Background()
	}
	// The root command's context ends on the first Ctrl+C. notifyStop
	// decides what each stop signal does here: the first one cancels a
	// deploy from hPanel, when one runs, and leaves the Preview running.
	base = context.WithoutCancel(base)
	stderr := cmd.ErrOrStderr()
	p := &printer{
		stdout: cmd.OutOrStdout(),
		stderr: stderr,
		log:    stderr,
		json:   jsonOut,
		bold:   styled(cmd.OutOrStdout()),
		now:    time.Now,
		loc:    time.Local,
	}
	rd := newRemote(base, o, p)
	var interrupt func() bool
	if rd != nil {
		interrupt = rd.cancelDeploy
	}
	ctx, stop := notifyStop(base, interrupt)
	defer stop()

	target, err := resolveTarget(ctx, args, o, func() detector { return newDetector(cmd, jsonOut) })
	if err != nil {
		return err
	}
	p.target = target
	detected := len(args) == 0 && o.url == ""
	// A person at a terminal gets the live view. CI, --no-input and AI agents
	// keep plain lines: agents often have a pseudo-terminal, and a view that
	// redraws for the whole Preview only fills their transcript.
	var view *quickview.View
	if !jsonOut && interactive(cmd) && !brand.AgentShell() {
		view = newView(cmd.OutOrStdout(), quickview.Options{Target: target, Detected: detected})
	}
	if detected && !jsonOut && view == nil {
		fmt.Fprintf(stderr, "Found a server at %s\n", target)
	}
	warn := func(w io.Writer) {
		warnSharedPort(ctx, w, target)
		warnHostHeader(w, o.hostHeader)
	}
	if view == nil {
		warn(stderr)
	}

	// A Preview needs an account. Sign in first: the connector's owner ticket
	// callback has no time to wait for a browser.
	if !o.noOwner {
		browser := !jsonOut && interactive(cmd)
		if err := ensureSignedIn(ctx, stderr, browser); err != nil {
			return p.ended(notSignedIn(ctx, err, cmd.Root().Name(), tokenEnv))
		}
	}
	level := slog.LevelWarn
	if viper.GetBool("verbose") {
		level = slog.LevelDebug
	}
	logOut := stderr
	if view != nil {
		view.Start()
		p.view, logOut = view, view.LogWriter(stderr)
		p.log = logOut
		// --verbose request lines print above the view too.
		api.SetDiagnostics(logOut)
		defer api.SetDiagnostics(nil)
		// Below the logo, where the view cannot push them off screen.
		warn(logOut)
	}
	qo := connector.QuickOptions{
		URL:          target,
		HostHeader:   o.hostHeader,
		Bootstrap:    o.bootstrap,
		Ports:        o.ports,
		ServerCAFile: o.serverCAFile,
		Lifetime:     lifetime,
		Version:      api.UserAgent(),
		Logger:       slog.New(slog.NewTextHandler(logOut, &slog.HandlerOptions{Level: level})),
		Owner:        NewOwnerTicketFunc(DefaultOwnerOptions(o.noOwner)),
		OnReady:      p.ready,
		OnReachable:  p.reachable,
		OnPause:      p.paused,
	}
	if rd != nil {
		rd.attach(&qo, DefaultOwnerOptions(o.noOwner))
	}
	expired, err := runQuick(ctx, qo)
	if rd != nil {
		rd.close()
	}
	return p.ended(classify(expired, err, cmd.Root().Name(), tokenEnv))
}

// notifyStop returns a context that the first stop signal ends, as
// signal.NotifyContext does, and a second one kills the process. Ctrl+C
// first goes to interrupt, when set: it returns true when it took the signal,
// as a running deploy from hPanel does, and the Preview keeps running.
func notifyStop(base context.Context, interrupt func() bool) (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithCancel(base)
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, stopSignals...)
	go func() {
		defer signal.Stop(sigs)
		for {
			select {
			case s := <-sigs:
				if s == os.Interrupt && interrupt != nil && interrupt() {
					continue
				}
				cancel()
				return
			case <-ctx.Done():
				return
			}
		}
	}()
	return ctx, cancel
}

// resolveTarget picks the origin to share: the port argument, --url, or a
// dev server that detector finds. A sensitive port on this computer is
// refused before anything connects to the Preview service.
func resolveTarget(ctx context.Context, args []string, o options, detector func() detector) (string, error) {
	var target string
	switch {
	case len(args) == 1 && o.url != "":
		return "", utils.FlagErrorf("pass a port or --url, not both")
	case len(args) == 1:
		port, err := parsePort(args[0])
		if err != nil {
			return "", err
		}
		target = "http://localhost:" + strconv.Itoa(port)
	case o.url != "":
		u, err := parseOrigin(o.url)
		if err != nil {
			return "", err
		}
		target = u.String()
	default:
		var err error
		if target, err = detector().detect(ctx); err != nil {
			return "", err
		}
	}
	if err := checkTarget(ctx, target); err != nil {
		return "", err
	}
	return target, nil
}

// parseOrigin accepts an http or https origin: a scheme, a host and an
// optional port, with no credentials, path, query or fragment.
func parseOrigin(raw string) (*url.URL, error) {
	u, err := url.Parse(raw)
	if err != nil || (u.Scheme != "http" && u.Scheme != "https") || u.Host == "" || u.Hostname() == "" {
		return nil, utils.FlagErrorf("--url %q: want an http or https origin such as http://localhost:8080", raw)
	}
	if u.User != nil || (u.Path != "" && u.Path != "/") || u.RawQuery != "" || u.Fragment != "" || u.Opaque != "" {
		return nil, utils.FlagErrorf("--url %q: want an origin only (scheme, host and port), such as http://localhost:8080", raw)
	}
	if p := u.Port(); p != "" {
		if _, err := parsePort(p); err != nil {
			return nil, utils.FlagErrorf("--url %q: invalid port %q", raw, p)
		}
	}
	u.Path = ""
	u.Host = strings.ToLower(u.Host)
	return u, nil
}

// parseTTL reads --ttl: a Go duration such as 90m or 12h, optionally after
// a whole number of days (3d, 1d12h), within the connector's
// creds.MinQuickLifetime and creds.MaxQuickLifetime. The help text names
// that range too; TestTTLHelpMatchesTheConnector keeps it in step.
func parseTTL(raw string) (time.Duration, error) {
	invalid := utils.FlagErrorf("invalid --ttl %q: want a duration from %s to %s, such as 90m, 12h or 3d",
		raw, roughDuration(creds.MinQuickLifetime), roughDuration(creds.MaxQuickLifetime))
	s := strings.ToLower(strings.TrimSpace(raw))
	var d time.Duration
	if i := strings.IndexByte(s, 'd'); i >= 0 {
		days, err := strconv.Atoi(s[:i])
		if err != nil || days < 0 || days > int(creds.MaxQuickLifetime/(24*time.Hour)) {
			return 0, invalid
		}
		d, s = time.Duration(days)*24*time.Hour, s[i+1:]
	}
	if s != "" {
		rest, err := time.ParseDuration(s)
		if err != nil || rest < 0 {
			return 0, invalid
		}
		d += rest
	}
	if d < creds.MinQuickLifetime || d > creds.MaxQuickLifetime {
		return 0, invalid
	}
	return d, nil
}

// formatJSON reports an explicit --format json, which selects JSON lines as
// --json does. A configured default format does not: it is for API output.
func formatJSON(cmd *cobra.Command) bool {
	f := cmd.Flag("format")
	return f != nil && f.Changed && f.Value.String() == "json"
}

// styled reports whether w is a terminal that takes styling.
func styled(w io.Writer) bool {
	if _, noColor := os.LookupEnv("NO_COLOR"); noColor {
		return false
	}
	f, ok := w.(*os.File)
	return ok && term.IsTerminal(int(f.Fd()))
}
