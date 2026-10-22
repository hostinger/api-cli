## hostinger vps virtual-machines start

Start virtual machine

### Synopsis

Start a specified virtual machine.

If the virtual machine is already running, the request will still be processed without any effect.

Use this endpoint to power on stopped VPS instances.

```
hostinger vps virtual-machines start <virtual-machine-id> [flags]
```

### Examples

```
  hostinger vps vm start 1234567
  hostinger vps vm start 1234567 --wait
```

### Options

```
  -h, --help                    help for start
      --wait                    Wait until the action finishes
      --wait-timeout duration   Maximum time to wait with --wait (default 10m0s)
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

