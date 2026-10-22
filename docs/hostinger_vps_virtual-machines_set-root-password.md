## hostinger vps virtual-machines set-root-password

Set root password

### Synopsis

Set root password for a specified virtual machine.

Requirements for password are same as in the recreate virtual machine
endpoint (/#tag/vps-virtual-machine/POST/api/vps/v1/virtual-machines/{virtualMachineId}/recreate).

Use this endpoint to update administrator credentials for VPS instances.

```
hostinger vps virtual-machines set-root-password <virtual-machine-id> [flags]
```

### Examples

```
  hostinger vps virtual-machines set-root-password 1268054 --password oMeNRustosIO
```

### Options

```
  -h, --help                    help for set-root-password
      --password string         Root password for the virtual machine
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

