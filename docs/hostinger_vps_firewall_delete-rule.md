## hostinger vps firewall delete-rule

Delete firewall rule

### Synopsis

Delete a specific firewall rule from a specified firewall.

Any virtual machine that has this firewall activated will lose sync with the firewall
and will have to be synced again manually.

Use this endpoint to remove specific firewall rules.

```
hostinger vps firewall delete-rule <firewall-id> <rule-id> [flags]
```

### Options

```
  -h, --help   help for delete-rule
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger vps firewall](hostinger_vps_firewall.md)	 - Firewall commands

