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
	"github.com/hostinger/api-cli/utils"
	"github.com/hostinger/hcdn-connector/pkg/connector"
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

The Preview runs in the foreground until Ctrl+C, or until its time runs out. Nothing is
uploaded: requests go through an encrypted tunnel to the server while the command runs.

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

When the Preview service asks for an owner, the Preview is linked to the account you
signed in with (auth login, or an API token); it works without one when the service
allows it.

The URL is printed once the Preview answers. If that takes longer than a few seconds,
the URL is printed early with a note that it may answer 530 for a few seconds, and
"Preview ready." follows once it answers.

--json prints one JSON object per line on stdout: {"event":"ready",...} with the URL and
"reachable" (false while it may still answer 530), {"event":"reachable"} when such a
Preview starts to answer, {"event":"paused",...} and {"event":"ended",...}, which
carries a "message" when the Preview failed. Logs go to stderr.

Exit codes: 0 when stopped or expired, 2 for a usage error such as a refused port, 77
when the service requires a signed-in account, 78 when the Preview was taken down, and 1
for any other failure.`,
	Example: `  hostinger preview
  hostinger preview 3000
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
}

var flags options

func init() {
	f := Cmd.Flags()
	f.StringVar(&flags.url, "url", "", "Share this origin instead of a local port, e.g. http://192.168.1.20:8080 (http or https)")
	f.StringVar(&flags.hostHeader, "host-header", "", "Host header sent to the server (default: the Preview's host name; prefer allowing .hostingerpreview.com in the dev server)")
	f.BoolVar(&flags.json, "json", false, "Print one JSON object per event on stdout")

	f.StringSliceVar(&flags.bootstrap, "bootstrap", nil, "Tunnel endpoints to connect to (comma separated)")
	f.StringVar(&flags.serverCAFile, "server-ca-file", "", "CA bundle that signs the tunnel endpoints' certificates")
	f.IntSliceVar(&flags.ports, "ports", nil, "Tunnel ports to try (comma separated)")
	f.BoolVar(&flags.noOwner, "no-owner", false, "Never link the Preview to an account")
	for _, name := range []string{"bootstrap", "server-ca-file", "ports", "no-owner"} {
		_ = f.MarkHidden(name)
	}
}

// runQuick is connector.RunQuick; tests replace it.
var runQuick = connector.RunQuick

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

	base := cmd.Context()
	if base == nil {
		base = context.Background()
	}
	// The first signal ends the Preview cleanly; a second one kills the process.
	ctx, stop := signal.NotifyContext(base, stopSignals...)
	defer stop()
	context.AfterFunc(ctx, stop)

	target, err := resolveTarget(ctx, args, o, func() detector { return newDetector(cmd, jsonOut) })
	if err != nil {
		return err
	}
	stderr := cmd.ErrOrStderr()
	if len(args) == 0 && o.url == "" && !jsonOut {
		fmt.Fprintf(stderr, "Found a server at %s\n", target)
	}
	warnSharedPort(ctx, stderr, target)
	warnHostHeader(stderr, o.hostHeader)

	p := &printer{
		stdout: cmd.OutOrStdout(),
		stderr: stderr,
		json:   jsonOut,
		bold:   styled(cmd.OutOrStdout()),
		target: target,
		now:    time.Now,
		loc:    time.Local,
	}
	level := slog.LevelWarn
	if viper.GetBool("verbose") {
		level = slog.LevelDebug
	}
	expired, err := runQuick(ctx, connector.QuickOptions{
		URL:          target,
		HostHeader:   o.hostHeader,
		Bootstrap:    o.bootstrap,
		Ports:        o.ports,
		ServerCAFile: o.serverCAFile,
		Version:      api.UserAgent(),
		Logger:       slog.New(slog.NewTextHandler(stderr, &slog.HandlerOptions{Level: level})),
		Owner:        NewOwnerTicketFunc(DefaultOwnerOptions(o.noOwner)),
		OnReady:      p.ready,
		OnReachable:  p.reachable,
		OnPause:      p.paused,
	})
	return p.ended(classify(expired, err, cmd.Root().Name(), utils.EnvPrefix+"API_TOKEN"))
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
