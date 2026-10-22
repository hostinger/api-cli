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

### Examples

```
  hostinger vps firewall delete-rule 9449049 8941182
```

### Options

```
  -h, --help   help for delete-rule
  -y, --yes    Skip the confirmation prompt
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

