## hostinger preview

Share a local dev server at a temporary public URL

### Synopsis

Share a web server running on this computer at a temporary public https URL, for
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
for any other failure.

```
hostinger preview [port] [flags]
```

### Examples

```
  hostinger preview
  hostinger preview 3000
  hostinger preview 8080 --host-header shop.test
  hostinger preview --url https://localhost:8443
  hostinger preview --json
```

### Options

```
  -h, --help[=true|false]    help for preview
      --host-header string   Host header sent to the server (default: the Preview's host name; prefer allowing .hostingerpreview.com in the dev server)
      --json[=true|false]    Print one JSON object per event on stdout
      --url string           Share this origin instead of a local port, e.g. http://192.168.1.20:8080 (http or https)
```

### Options inherited from parent commands

```
      --config string           Config file (default is $HOME/.hostinger.yaml)
      --domain string           Website domain for commands that act on a website (default: HOSTINGER_DOMAIN, .hostinger/site.json, or default_domain in the config)
      --dry-run[=true|false]    Print the request as a curl command instead of sending it
  -o, --format string           Output format: table|json|yaml|tree (default: table on a terminal, json when piped)
      --jq string               Filter the JSON response with a jq expression, e.g. '.data[].id'
      --no-input[=true|false]   Never prompt; fail when an argument, flag or confirmation is missing
      --order-id string         Order for commands that act on an order (default: HOSTINGER_ORDER_ID, .hostinger/site.json for hosting, or a prompt)
      --username string         Hosting account for commands that act on an account (default: HOSTINGER_USERNAME, the website from --domain, or .hostinger/site.json)
      --verbose[=true|false]    Print the config file used and each HTTP request to stderr
```

### SEE ALSO

* [hostinger](hostinger.md)	 - Hostinger API Command Line Interface

