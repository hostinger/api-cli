## hostinger vps backups restore

Restore backup

### Synopsis

Restore a backup for a specified virtual machine.

The system will then initiate the restore process, which may take some time depending on the size of the backup.

**All data on the virtual machine will be overwritten with the data from the backup.**

Use this endpoint to recover VPS data from backup points.

```
hostinger vps backups restore <virtual-machine-id> <backup-id> [flags]
```

### Options

```
  -h, --help   help for restore
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger vps backups](hostinger_vps_backups.md)	 - Backups commands

