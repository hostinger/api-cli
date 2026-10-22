## hostinger vps firewall list

Get firewall list

### Synopsis

Retrieve all available firewalls.

Use this endpoint to view existing firewall configurations.

```
hostinger vps firewall list [flags]
```

### Examples

```
  hostinger vps firewall list
```

### Options

```
      --all        Fetch every page and print all items
  -h, --help       help for list
      --page int   Page number
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

