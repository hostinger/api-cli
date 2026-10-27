## hostinger agency-hosting metrics list-order-resource-usage

List order resource usage metrics

### Synopsis

Returns aggregated CPU, memory, and process usage for the Agency Plan order
over the selected time frame, plus the plan quotas and a per-website
breakdown. Each website is identified by uid. Suspended and deleted websites
are excluded from both the order totals and the per-website breakdown.
Values may be up to one hour stale. Disk and inode usage are on the
disk-usage-metrics endpoint.

The order can be given as [order-id] or left out, in which case it comes from --order-id, HOSTINGER_ORDER_ID, the .hostinger/site.json written by `hostinger up` or `hostinger link` (hosting orders), or a prompt when run in a terminal (the only order is used without asking).

```
hostinger agency-hosting metrics list-order-resource-usage [order-id] [flags]
```

### Examples

```
  hostinger agency-hosting metrics list-order-resource-usage 123456
```

### Options

```
  -h, --help                   help for list-order-resource-usage
      --time-frame-hours int   Length of the window in hours, ending now. Bucket size grows with the window. (one of: 1, 24, 168, 336, 720) (default 24)
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

* [hostinger agency-hosting metrics](hostinger_agency-hosting_metrics.md)	 - Metrics commands

