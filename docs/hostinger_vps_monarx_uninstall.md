## hostinger vps monarx uninstall

Uninstall Monarx

### Synopsis

Uninstall the Monarx malware scanner on a specified virtual machine.

If Monarx is not installed, the request will still be processed without any effect.

Use this endpoint to remove malware scanner from VPS instances.

```
hostinger vps monarx uninstall <virtual-machine-id> [flags]
```

### Examples

```
  hostinger vps monarx uninstall 1268054
```

### Options

```
  -h, --help                    help for uninstall
      --wait                    Wait until the action finishes
      --wait-timeout duration   Maximum time to wait with --wait (default 10m0s)
  -y, --yes                     Skip the confirmation prompt
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

