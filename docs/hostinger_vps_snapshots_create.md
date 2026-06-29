## hostinger vps snapshots create

Create snapshot

### Synopsis

Create a snapshot of a specified virtual machine.

A snapshot captures the state and data of the virtual machine at a specific point in time, 
allowing users to restore the virtual machine to that state if needed. 
This operation is useful for backup purposes, system recovery, 
and testing changes without affecting the current state of the virtual machine.

**Creating new snapshot will overwrite the existing snapshot!**

Use this endpoint to capture VPS state for backup and recovery purposes.

```
hostinger vps snapshots create <virtual-machine-id> [flags]
```

### Options

```
  -h, --help   help for create
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger vps snapshots](hostinger_vps_snapshots.md)	 - Snapshots commands

