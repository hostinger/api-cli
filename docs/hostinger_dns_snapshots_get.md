## hostinger dns snapshots get

Get DNS snapshot

### Synopsis

Retrieve particular DNS snapshot with contents of DNS zone records.

Use this endpoint to view historical DNS configurations for domains.

```
hostinger dns snapshots get <domain> <snapshot-id> [flags]
```

### Examples

```
  hostinger dns snapshots get mydomain.tld 53513053
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

* [hostinger dns snapshots](hostinger_dns_snapshots.md)	 - Snapshot commands

