# Hostinger API CLI

`hostinger` is the command line interface for the [Hostinger API](https://developers.hostinger.com).

> **This repository is generated.** Do not open pull requests here — every file
> is overwritten by the next generation run. The CLI is generated from the
> OpenAPI specification by
> [hostinger/public-api-generator](https://github.com/hostinger/public-api-generator)
> (see the `cli/` directory there). Report issues or contribute in that repository.

## Installation

### Homebrew (macOS & Linux)

```sh
brew install hostinger/tap/hostinger
```

Upgrade with `brew upgrade hostinger`. Shell tab-completion (bash/zsh/fish) and man pages
(`man hostinger`) are installed automatically.

### Binary download

Download the binary for your platform from the
[releases page](https://github.com/hostinger/api-cli/releases). Then enable tab-completion
with `hostinger completion install` (see [AUTOCOMPLETE.md](AUTOCOMPLETE.md)).

## Configuration

Sign in through the browser:

```sh
hostinger auth login
```

or use an API token instead: run `hostinger config set api_token` (it asks for the token
without echoing it) or set the `HOSTINGER_API_TOKEN` environment variable. A token takes precedence over
the sign-in. Generate one at [hPanel → API](https://hpanel.hostinger.com/profile/api).

`hostinger auth status` shows which credentials are in use. Settings live in
`$HOME/.hostinger.yaml` (`hostinger config list` shows them):

```yaml
api_token: <your API token>
format: json        # default output format
no_input: false     # true: never prompt
```

## Usage

```
hostinger <group> <subgroup> <verb> [args] [flags]
```

Examples:

```
hostinger billing catalog list
hostinger vps virtual-machines list
hostinger vps vm start 123          # vm is an alias for virtual-machines
hostinger dns records list example.com
```

Every command's `--help` lists its required flags separately and shows examples. Flags that take
JSON accept inline JSON, `@file.json`, or `-` for stdin; the help shows the expected shape.

### Output

- Tables on a terminal, JSON when the output is piped or redirected. Choose explicitly with
  `-o/--format table|json|yaml|tree`, or set a default with `hostinger config set format json`.
- `--jq '<expr>'` filters the JSON response (strings print raw), e.g. `--jq '.[].id'`.
- List commands with pages take `--all` to fetch every page.

### Prompts and safety

- In a terminal, missing arguments and required flags are asked for, with choices listed where
  they are known. `--no-input` (or `no_input: true`) turns prompting off.
- Commands that delete, stop, restore or purchase ask for confirmation. Pass `--yes` to skip it;
  non-interactive runs (scripts, CI) must pass `--yes`.
- `--dry-run` prints the request as a `curl` command instead of sending it.
- Commands that start an asynchronous action take `--wait` to block until it finishes.
- `--verbose` prints each HTTP request and its status to stderr.

### Exit codes

| Code | Meaning                                       |
|------|-----------------------------------------------|
| 0    | Success                                       |
| 1    | API, network or other error; aborted prompt   |
| 2    | Invalid arguments, flags or input             |
| 3    | No credentials, or authentication failed (HTTP 401/403) |
| 4    | Not found (HTTP 404)                          |
| 5    | Validation failed (HTTP 422)                  |
| 77   | `preview`: the service requires a signed-in account |
| 78   | `preview`: the Preview was taken down         |

## Deploying a website

`hostinger up` deploys the current directory (or a path) to a Web or Cloud
hosting website, asking only for what it cannot work out: which plan, which
website (or a new one with a free subdomain), and a confirmation of the detected
type (static files or a Node.js app, built on the server). It remembers the
answers in `.hostinger/site.json`, so the next `up` is one command. Inside a Git
checkout `.hostinger/` is added to `.gitignore`, since the file names the hosting
account; CI should pass `--domain` instead.

```
hostinger up                      # deploy the current directory, ask what is missing
hostinger up ./dist               # deploy a folder; the link is written in the current directory
hostinger up --type nodejs        # override the detected or remembered type
hostinger up --settings           # review the Node.js build settings again
hostinger up --source git         # deploy the pushed branch through a connected GitHub/GitLab account
hostinger link                    # point this directory at another website
hostinger unlink                  # forget it
```

Commands that act on a website, such as `hostinger hosting cdn status`, take the
website from the same link when the domain is left out, or from `--domain`,
`HOSTINGER_DOMAIN` and `default_domain` in the config.

### From scripts and agents

Without a terminal nothing is asked. Pass `--domain` for an existing website, or
`--create` for a new one (`--order-id` when the account has several plans), and
`--format json` for one JSON object on stdout; progress goes to stderr, failures
exit non-zero (API errors with the exit codes above) and carry an `error` field. Prompts are also off when `CI` or
`HOSTINGER_NONINTERACTIVE` is set, or with `--no-input`.

```
hostinger up ./dist --domain example.com --format json
hostinger up --create --type nodejs --app-type express --entry-file server.js --format json
hostinger up --dry-run --domain example.com      # the plan as JSON; changes nothing
hostinger up --no-wait --domain example.com      # start the Node.js build and return its uuid
hostinger hosting nodejs build example.com <uuid>  # poll it
```

`--dry-run` shows the resolved website, type and source, how many files would be
packed and which are skipped, and `overwritesContent` when the first deploy would
replace what the website serves (the real run then asks, or needs `--yes`). Every deploy replaces the website's current files.

## Sharing a local dev server

`hostinger preview` shares a web server running on your computer at a temporary
public https URL until you press Ctrl+C or its time runs out. Without a port it looks
for a dev server on the usual ports (5173, 3000, 4321, 8080, 8000, 5000, 4200, 8888).

```
hostinger preview                                  # find the dev server and share it
hostinger preview 3000                             # share http://localhost:3000
hostinger preview 8080 --host-header shop.test     # for a server that serves by host name
hostinger preview --url https://localhost:8443     # share another origin
hostinger preview --json                           # one JSON event per line: ready, reachable, paused, ended
```

Anyone with the URL can reach the server, including its debug consoles and dev tools,
so ports of databases, debuggers and other system services are never shared from your
computer. A dev server that checks the Host header must allow `.hostingerpreview.com`
(in Vite, `server.allowedHosts: ['.hostingerpreview.com']`); `--host-header localhost`
bypasses that check and makes its debug consoles public.

When the service asks for an owner, the Preview is linked to the account you signed in
with; when it requires one, run `hostinger auth login` or set `HOSTINGER_API_TOKEN`.
The Preview needs outbound TCP port 7844.

Command documentation lives in [docs/](docs/) and in the man pages, and `manifest.json` maps every
API operation to its command. Shell autocompletion: see [AUTOCOMPLETE.md](AUTOCOMPLETE.md).
