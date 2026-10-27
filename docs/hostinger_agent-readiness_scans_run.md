## hostinger agent-readiness scans run

Run agent readiness scan

### Synopsis

Start a scan of how ready a public website is for AI agents to discover, read and act on.\n
Scanning is asynchronous: this operation only starts the scan, and the returned scan is usually `pending`
with no results yet. To get the results, poll the `Get agent readiness scan` operation with the same
domain every few seconds until its status is `completed` or `failed`. The scan itself is limited to 60 seconds,
plus any time it waits in the queue. A `failed` scan is not retried; run this operation again.\n
Starting a scan for a domain that already has a pending scan from the last 5 minutes
returns that scan instead of starting a new one.

```
hostinger agent-readiness scans run [flags]
```

### Examples

```
  hostinger agent-readiness scans run --domain example.com
```

### Options

```
      --domain string   Domain name or full URL of the website to scan. Only the host is used.
  -h, --help            help for run
```

### Options inherited from parent commands

```
      --config string     Config file (default is $HOME/.hostinger.yaml)
      --dry-run           Print the request as a curl command instead of sending it
  -o, --format string     Output format: table|json|yaml|tree (default: table on a terminal, json when piped)
      --jq string         Filter the JSON response with a jq expression, e.g. '.data[].id'
      --no-input          Never prompt; fail when an argument, flag or confirmation is missing
      --order-id string   Order for commands that act on an order (default: HOSTINGER_ORDER_ID, .hostinger/site.json for hosting, or a prompt)
      --username string   Hosting account for commands that act on an account (default: HOSTINGER_USERNAME, the website from --domain, or .hostinger/site.json)
      --verbose           Print the config file used and each HTTP request to stderr
```

### SEE ALSO

* [hostinger agent-readiness scans](hostinger_agent-readiness_scans.md)	 - Scans commands

