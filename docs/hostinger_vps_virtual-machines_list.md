## hostinger vps virtual-machines list

Get virtual machines

### Synopsis

Retrieve all available virtual machines.

Use this endpoint to view available VPS instances.

```
hostinger vps virtual-machines list [flags]
```

### Examples

```
  hostinger vps vm list
  hostinger vps vm list --jq '.[] | {id, hostname, state}'
```

### Options

```
  -h, --help   help for list
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

* [hostinger vps virtual-machines](hostinger_vps_virtual-machines.md)	 - Virtual machine commands

