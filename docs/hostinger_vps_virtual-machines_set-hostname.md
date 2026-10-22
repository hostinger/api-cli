## hostinger vps virtual-machines set-hostname

Set hostname

### Synopsis

Set hostname for a specified virtual machine.

Changing hostname does not update PTR record automatically.
If you want your virtual machine to be reachable by a hostname,
you need to point your domain A/AAAA records to virtual machine IP as well.

Use this endpoint to configure custom hostnames for VPS instances.

```
hostinger vps virtual-machines set-hostname <virtual-machine-id> [flags]
```

### Examples

```
  hostinger vps virtual-machines set-hostname 1268054 --hostname my.server.tld
```

### Options

```
  -h, --help                    help for set-hostname
      --hostname string         
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

