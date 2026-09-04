## hostinger vps firewall sync-to-all-assigned-v-ms

Sync firewall to all assigned VMs

### Synopsis

Sync a firewall's rules to every virtual machine it's assigned to.

Firewall can lose sync with a virtual machine if the firewall has new rules added, removed or updated.

Use this endpoint to apply updated firewall rules to all VPS instances assigned to the firewall.

```
hostinger vps firewall sync-to-all-assigned-v-ms <firewall-id> [flags]
```

### Options

```
  -h, --help   help for sync-to-all-assigned-v-ms
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger vps firewall](hostinger_vps_firewall.md)	 - Firewall commands

