## hostinger vps snapshots get

Get snapshot

### Synopsis

Retrieve snapshot for a specified virtual machine.

Use this endpoint to view current VPS snapshot information.

```
hostinger vps snapshots get <virtual-machine-id> [flags]
```

### Examples

```
  hostinger vps snapshots get 1268054
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

* [hostinger vps snapshots](hostinger_vps_snapshots.md)	 - Snapshots commands

