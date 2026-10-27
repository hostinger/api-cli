## hostinger agent-readiness scans get

Get agent readiness scan

### Synopsis

Get the latest agent readiness scan of a website.\n
Scans started with the `Run agent readiness scan` operation are asynchronous. While the status is
`pending`, the result fields are null; wait a few seconds and call this operation again until the status is
`completed` or `failed`. A `failed` scan is not retried; run the scan again.\n
A completed scan has a 0-100 score, a letter grade, a score per category, and the result of every
check with the evidence found and how to fix it. Check names are listed by the `List checks` operation.\n
Returns 404 when the domain has never been scanned; use the `Run agent readiness scan` operation first.

The domain can be given as [domain] or left out, in which case it comes from --domain, HOSTINGER_DOMAIN, the .hostinger/site.json written by `hostinger up` or `hostinger link` when that website's domain belongs to the account, default_domain in the config file, or a prompt when run in a terminal.

```
hostinger agent-readiness scans get [domain] [flags]
```

### Examples

```
  hostinger agent-readiness scans get example.com
```

### Options

```
  -h, --help   help for get
```

### Options inherited from parent commands

```
      --config string     Config file (default is $HOME/.hostinger.yaml)
      --domain string     Website domain for commands that act on a website (default: HOSTINGER_DOMAIN, .hostinger/site.json, or default_domain in the config)
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

