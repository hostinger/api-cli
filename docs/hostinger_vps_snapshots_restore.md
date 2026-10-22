## hostinger vps snapshots restore

Restore snapshot

### Synopsis

Restore a specified virtual machine to a previous state using a snapshot.

Restoring from a snapshot allows users to revert the virtual machine to that state,
which is useful for system recovery, undoing changes, or testing.

Use this endpoint to revert VPS instances to previous saved states.

```
hostinger vps snapshots restore <virtual-machine-id> [flags]
```

### Examples

```
  hostinger vps snapshots restore 1268054
```

### Options

```
  -h, --help                    help for restore
      --wait                    Wait until the action finishes
      --wait-timeout duration   Maximum time to wait with --wait (default 10m0s)
  -y, --yes                     Skip the confirmation prompt
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

