## hapi vps snapshots restore

Restore snapshot

### Synopsis

Restore a specified virtual machine to a previous state using a snapshot.

Restoring from a snapshot allows users to revert the virtual machine to that state,
which is useful for system recovery, undoing changes, or testing.

Use this endpoint to revert VPS instances to previous saved states.

```
hapi vps snapshots restore <virtual-machine-id> [flags]
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

* [hapi vps snapshots](hapi_vps_snapshots.md)	 - Snapshots commands

