## hostinger vps recovery stop

Stop recovery mode

### Synopsis

Stop recovery mode for a specified virtual machine.

If virtual machine is not in recovery mode, this operation will fail.

Use this endpoint to exit system rescue mode and return VPS to normal operation.

```
hostinger vps recovery stop <virtual-machine-id> [flags]
```

### Examples

```
  hostinger vps recovery stop 1268054
```

### Options

```
  -h, --help                    help for stop
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

* [hostinger vps recovery](hostinger_vps_recovery.md)	 - Recovery commands

