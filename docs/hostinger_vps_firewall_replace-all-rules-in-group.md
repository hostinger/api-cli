## hostinger vps firewall replace-all-rules-in-group

Replace all firewall rules in group

### Synopsis

Replaces all firewall rules within a specified firewall group with the provided set of rules
in a single atomic operation, instead of creating or deleting rules one by one.

Any virtual machine using this firewall group will need to be synchronized after replacing rules;
pass the "sync" parameter to trigger synchronization immediately.

```
hostinger vps firewall replace-all-rules-in-group <firewall-id> [flags]
```

### Options

```
  -h, --help           help for replace-all-rules-in-group
      --rules string   The complete set of firewall rules that atomically replaces all existing rules in the group (JSON)
      --sync           Synchronize the firewall group to all its virtual machines after replacing the rules
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger vps firewall](hostinger_vps_firewall.md)	 - Firewall commands

