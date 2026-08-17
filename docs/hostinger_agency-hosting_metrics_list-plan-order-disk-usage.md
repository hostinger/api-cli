## hostinger agency-hosting metrics list-plan-order-disk-usage

List Agency Plan order disk usage metrics

### Synopsis

Returns aggregated disk and inode usage for the Agency Plan order over the
selected time frame, plus the plan quotas. Figures cover the whole order
account. Values may be up to one hour stale. CPU, memory, and process usage
are on the resource-usage-metrics endpoint.

```
hostinger agency-hosting metrics list-plan-order-disk-usage <order_id> [flags]
```

### Options

```
  -h, --help                  help for list-plan-order-disk-usage
      --time-frame-days int   Length of the window in days, ending now. Bucket size grows with the window. (one of: 1, 7, 14, 30) (default 1)
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger agency-hosting metrics](hostinger_agency-hosting_metrics.md)	 - Metrics commands

