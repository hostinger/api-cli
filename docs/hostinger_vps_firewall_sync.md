## hostinger vps firewall sync

Sync firewall

### Synopsis

Sync a firewall for a specified virtual machine.

Firewall can lose sync with virtual machine if the firewall has new rules added, removed or updated.

Use this endpoint to apply updated firewall rules to VPS instances.

```
hostinger vps firewall sync <firewall-id> <virtual-machine-id> [flags]
```

### Options

```
  -h, --help   help for sync
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger vps firewall](hostinger_vps_firewall.md)	 - Firewall commands

