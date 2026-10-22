## hostinger vps firewall sync

Sync firewall

### Synopsis

Deprecated: use `POST /api/vps/v1/firewall/{firewallId}/sync` instead, which syncs the firewall
to all virtual machines assigned to it.

Sync a firewall for a specified virtual machine.

Firewall can lose sync with virtual machine if the firewall has new rules added, removed or updated.

Use this endpoint to apply updated firewall rules to VPS instances.

```
hostinger vps firewall sync <firewall-id> <virtual-machine-id> [flags]
```

### Examples

```
  hostinger vps firewall sync 9449049 1268054
```

### Options

```
  -h, --help                    help for sync
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

