## hostinger vps monarx scan-metrics

Get scan metrics

### Synopsis

Retrieve scan metrics for the Monarx (https://www.monarx.com/) malware scanner
installed on a specified virtual machine.

The scan metrics provide detailed information about malware scans performed
by Monarx, including number of scans, detected threats, and other relevant
statistics. This information is useful for monitoring security status of the
virtual machine and assessing effectiveness of the malware scanner.

Use this endpoint to monitor VPS security scan results and threat detection.

```
hostinger vps monarx scan-metrics <virtual-machine-id> [flags]
```

### Examples

```
  hostinger vps monarx scan-metrics 1268054
```

### Options

```
  -h, --help   help for scan-metrics
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

* [hostinger vps monarx](hostinger_vps_monarx.md)	 - Malware scanner commands

