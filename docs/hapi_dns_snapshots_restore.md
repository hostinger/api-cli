## hapi dns snapshots restore

Restore DNS snapshot

### Synopsis

Restore DNS zone to the selected snapshot.

Use this endpoint to revert domain DNS to a previous configuration.

```
hapi dns snapshots restore <domain> <snapshot-id> [flags]
```

### Options

```
  -h, --help   help for restore
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hapi.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hapi dns snapshots](hapi_dns_snapshots.md)	 - Snapshot commands

