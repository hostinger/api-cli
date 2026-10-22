## hostinger agency-hosting metrics list-order-resource-usage

List order resource usage metrics

### Synopsis

Returns aggregated CPU, memory, and process usage for the Agency Plan order
over the selected time frame, plus the plan quotas and a per-website
breakdown. Each website is identified by uid. Suspended and deleted websites
are excluded from both the order totals and the per-website breakdown.
Values may be up to one hour stale. Disk and inode usage are on the
disk-usage-metrics endpoint.

```
hostinger agency-hosting metrics list-order-resource-usage <order_id> [flags]
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
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --dry-run         Print the request as a curl command instead of sending it
  -o, --format string   Output format: table|json|yaml|tree (default: table on a terminal, json when piped)
      --jq string       Filter the JSON response with a jq expression, e.g. '.data[].id'
      --no-input        Never prompt; fail when an argument, flag or confirmation is missing
      --verbose         Print the config file used and each HTTP request to stderr
```

### SEE ALSO

* [hostinger agency-hosting metrics](hostinger_agency-hosting_metrics.md)	 - Metrics commands

