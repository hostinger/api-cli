## hostinger vps firewall get

Get firewall details

### Synopsis

Retrieve firewall by its ID and rules associated with it.

Use this endpoint to view specific firewall configuration and rules.

```
hostinger vps firewall get <firewall-id> [flags]
```

### Examples

```
  hostinger vps firewall get 9449049
```

### Options

```
  -h, --help   help for get
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

