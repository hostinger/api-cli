## hostinger vps firewall delete

Delete firewall

### Synopsis

Delete a specified firewall.

Any virtual machine that has this firewall activated will automatically have it deactivated.

Use this endpoint to remove unused firewall configurations.

```
hostinger vps firewall delete <firewall-id> [flags]
```

### Examples

```
  hostinger vps firewall delete 9449049
```

### Options

```
  -h, --help   help for delete
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

