## hostinger vps virtual-machines metrics

Get metrics

### Synopsis

Retrieve historical metrics for a specified virtual machine.

It includes the following metrics:
- CPU usage
- Memory usage
- Disk usage
- Network usage
- Uptime

Use this endpoint to monitor VPS performance and resource utilization over time.

```
hostinger vps virtual-machines metrics <virtual-machine-id> [flags]
```

### Examples

```
  hostinger vps virtual-machines metrics 1268054 --date-from 2025-05-01T00:00:00Z --date-to 2025-06-01T00:00:00Z
```

### Options

```
      --date-from string   
      --date-to string     
  -h, --help               help for metrics
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

* [hostinger vps virtual-machines](hostinger_vps_virtual-machines.md)	 - Virtual machine commands

