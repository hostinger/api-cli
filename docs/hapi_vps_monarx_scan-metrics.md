## hapi vps monarx scan-metrics

Get scan metrics

### Synopsis

Retrieve scan metrics for the [Monarx](https://www.monarx.com/) malware scanner
installed on a specified virtual machine.

The scan metrics provide detailed information about malware scans performed
by Monarx, including number of scans, detected threats, and other relevant
statistics. This information is useful for monitoring security status of the
virtual machine and assessing effectiveness of the malware scanner.

Use this endpoint to monitor VPS security scan results and threat detection.

```
hapi vps monarx scan-metrics <virtual-machine-id> [flags]
```

### Options

```
  -h, --help   help for scan-metrics
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hapi.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hapi vps monarx](hapi_vps_monarx.md)	 - Malware scanner commands

