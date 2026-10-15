## hostinger agent-readiness scans run-website

Run website scan

### Synopsis

Start a scan of how ready a public website is for AI agents to discover, read and act on.\n
Scanning is asynchronous: this operation only starts the scan, and the returned scan is usually `pending`
with no results yet. To get the results, poll the `Get website scan` operation with the same domain every
few seconds until its status is `completed` or `failed`. The scan itself is limited to 60 seconds,
plus any time it waits in the queue. A `failed` scan is not retried; run this operation again.\n
Starting a scan for a domain that already has a pending scan from the last 5 minutes
returns that scan instead of starting a new one.

```
hostinger agent-readiness scans run-website [flags]
```

### Options

```
      --domain string   Domain name or full URL of the website to scan. Only the host is used.
  -h, --help            help for run-website
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger agent-readiness scans](hostinger_agent-readiness_scans.md)	 - Scans commands

