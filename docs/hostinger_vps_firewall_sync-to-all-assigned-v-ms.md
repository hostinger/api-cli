## hostinger vps firewall sync-to-all-assigned-v-ms

Sync firewall to all assigned VMs

### Synopsis

Sync a firewall's rules to every virtual machine it's assigned to.

Firewall can lose sync with a virtual machine if the firewall has new rules added, removed or updated.

Use this endpoint to apply updated firewall rules to all VPS instances assigned to the firewall.

```
hostinger vps firewall sync-to-all-assigned-v-ms <firewall-id> [flags]
```

### Examples

```
  hostinger vps firewall sync-to-all-assigned-v-ms 9449049
```

### Options

```
  -h, --help   help for sync-to-all-assigned-v-ms
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

