## hostinger vps firewall activate

Activate firewall

### Synopsis

Activate a firewall for a specified virtual machine.

Only one firewall can be active for a virtual machine at a time.

Use this endpoint to apply firewall rules to VPS instances.

```
hostinger vps firewall activate <firewall-id> <virtual-machine-id> [flags]
```

### Examples

```
  hostinger vps firewall activate 9449049 1268054
```

### Options

```
  -h, --help                    help for activate
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

* [hostinger vps firewall](hostinger_vps_firewall.md)	 - Firewall commands

