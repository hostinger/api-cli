## hostinger vps firewall create

Create new firewall

### Synopsis

Create a new firewall.

Use this endpoint to set up new firewall configurations for VPS security.

```
hostinger vps firewall create [flags]
```

### Examples

```
  hostinger vps firewall create --name 'My Firewall Group'
```

### Options

```
  -h, --help          help for create
      --name string   
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

