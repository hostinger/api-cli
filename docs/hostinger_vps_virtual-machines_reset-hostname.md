## hostinger vps virtual-machines reset-hostname

Reset hostname

### Synopsis

Reset hostname and PTR record of a specified virtual machine to default value.

Use this endpoint to restore default hostname configuration for VPS instances.

```
hostinger vps virtual-machines reset-hostname <virtual-machine-id> [flags]
```

### Examples

```
  hostinger vps virtual-machines reset-hostname 1268054
```

### Options

```
  -h, --help                    help for reset-hostname
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

* [hostinger vps virtual-machines](hostinger_vps_virtual-machines.md)	 - Virtual machine commands

