## hostinger dns snapshots restore

Restore DNS snapshot

### Synopsis

Restore DNS zone to the selected snapshot.

Use this endpoint to revert domain DNS to a previous configuration.

```
hostinger dns snapshots restore <domain> <snapshot-id> [flags]
```

### Examples

```
  hostinger dns snapshots restore mydomain.tld 53513053
```

### Options

```
  -h, --help   help for restore
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

* [hostinger dns snapshots](hostinger_dns_snapshots.md)	 - Snapshot commands

